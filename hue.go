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
  "encoding/json"
  "strconv"
  "strings"
)


type LightCommand struct {
  All bool
}

type RoomLightCommand struct {
  All bool
  LightGroup *Group
}

type Group struct {
  Name string
  Id string
}

type Client struct {
}

func(client *Client) Get(url string) (*http.Response, error) {
  resp, err := http.Get(url)
  return resp, err
}

var client = &Client{}

var url string = os.Getenv("HUE_URL")

func getLightsUrl() string {
  return fmt.Sprintf("%slights", url)
}

func getGroupsUrl() string {
  return fmt.Sprintf("%sgroups", url)
}

func (light *LightCommand) run(c *kingpin.ParseContext) error {
  resp, _ := client.Get(getLightsUrl())
  defer resp.Body.Close()
  bodyBytes, _ := ioutil.ReadAll(resp.Body)
  utils.PrettyPrint(bodyBytes)
  return nil
}

func configureLightCommand(app *kingpin.Application) {
  c:= &LightCommand{}
  app.Command("lights", "Get light info").Action(c.run)
}

func (lights *RoomLightCommand) run(c *kingpin.ParseContext) error {
  state := os.Args[2] == "on"
  url:= fmt.Sprintf("%sgroups/%s/action", url, lights.LightGroup.Id)
  fmt.Println(url)
  var jsonStr = []byte(fmt.Sprintf(`{"on":%t}`, state))
  req.Put(url, bytes.NewBuffer(jsonStr))
  return nil
}

func configureRoomsLightCommand(app *kingpin.Application) {
  var result map[string]interface{}
  resp, _ := client.Get(getGroupsUrl())
  defer resp.Body.Close()
  bodyBytes, _ := ioutil.ReadAll(resp.Body)
  json.Unmarshal(bodyBytes, &result)

  for i := 0; i < len(result); i++ {
    fmt.Println("*******************")
    c := &RoomLightCommand{}
    id := strconv.Itoa(i+1)
    first:= result[id].(map[string]interface{})
    name := strings.ToLower(first["name"].(string))
    c.LightGroup = &Group{Name: name, Id: id}

    on := fmt.Sprintf("Turn %s lights on", name)
    off := fmt.Sprintf("Turn %s lights off", name)

    roomCommand  := app.Command(name, "")
    roomCommand.Command("off", on).Action(c.run)
    roomCommand.Command("on", off).Action(c.run)
  }
}

func configureCommands(app *kingpin.Application) {
  configureLightCommand(app)
  configureRoomsLightCommand(app)
}

func main() {
  app := kingpin.New("hue", "Phillips Hue CLI")
  configureCommands(app)
  kingpin.MustParse(app.Parse(os.Args[1:]))
}
