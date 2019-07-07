package main

import (
  "fmt"
  "os"
  "net/http"
  "io/ioutil"
  "gopkg.in/alecthomas/kingpin.v2"
  "github.com/imroc/req"
  "bytes"
  "utils"
  "sync"
  "encoding/json"
  "strconv"
  "strings"
)

var url string = os.Getenv("HUE_URL")
var kitchenLights = [4]int{1,2,4}

type LightCommand struct {
  All bool
}

func (light *LightCommand) run(c *kingpin.ParseContext) error {
  url := fmt.Sprintf("%slights", url)
  resp, _ := http.Get(url)
  defer resp.Body.Close()
  bodyBytes, _ := ioutil.ReadAll(resp.Body)
  utils.PrettyPrint(bodyBytes)
  return nil
}

func configureLightCommand(app *kingpin.Application) {
  c:= &LightCommand{}
  app.Command("lights", "Get light info").Action(c.run)
}

type RoomLightCommand struct {
  All bool
}

func (lights *RoomLightCommand) run(c *kingpin.ParseContext) error {
  var wg sync.WaitGroup
  wg.Add(len(kitchenLights))
  state := os.Args[2] == "on"

  for i:= 0; i < len(kitchenLights); i++ {
    url:= fmt.Sprintf("%slights/%d/state", url, kitchenLights[i])
    var jsonStr = []byte(fmt.Sprintf(`{"on":%t}`, state))
    go func() {
      defer wg.Done()
      req.Put(url, bytes.NewBuffer(jsonStr))
    }()
  }

  wg.Wait()
  return nil
}

type Room struct {
  name string
  lightIds []string
}

func configureRoomLightCommand(app *kingpin.Application) {
  var result map[string]interface{}

  c := &RoomLightCommand{}
  groupsUrl := fmt.Sprintf("%sgroups", url)

  resp, _ := http.Get(groupsUrl)
  bodyBytes, _ := ioutil.ReadAll(resp.Body)
  json.Unmarshal(bodyBytes, &result)

  for i := 0; i < len(result); i++ {
    first:= result[strconv.Itoa(i+1)].(map[string]interface{})
    name := strings.ToLower(first["name"].(string))
    lights := first["lights"].([]interface{})
    length := len(lights)
    lightsArr := make([]string, length)

    for i := 0; i < length; i++ {
      t, _ := lights[i].(string)
      lightsArr[i] = t
    }

    room := &Room{lightIds: lightsArr, name: name}
    on := fmt.Sprintf("Turn %s lights on", name)
    off := fmt.Sprintf("Turn %s lights off", name)

    roomCommand  := app.Command(name, "")
    roomCommand.Command("off", on).Action(c.run)
    roomCommand.Command("on", off).Action(c.run)
  }
}

func configureCommands(app *kingpin.Application) {
  configureLightCommand(app)
  configureRoomLightCommand(app)
}

func main() {
  app := kingpin.New("hue", "Phillips Hue CLI")
  configureCommands(app)
  kingpin.MustParse(app.Parse(os.Args[1:]))
}
