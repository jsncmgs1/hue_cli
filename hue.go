package main

import (
  "os"
  "fmt"
  "io/ioutil"
  "gopkg.in/alecthomas/kingpin.v2"
  "encoding/json"
  "strconv"
  "strings"
  "hue_cli/lib/utils"
  httpClient "hue_cli/lib/client"
)

type LightCommand struct {
}

type RoomLightCommand struct {
  LightGroup *Group
}

type Group struct {
  Name string
  Id string
}

var url string = os.Getenv("HUE_URL")

func getLightsUrl() string {
  return fmt.Sprintf("%slights", url)
}

func getGroupsUrl() string {
  return fmt.Sprintf("%sgroups", url)
}

func groupsActionUrl(groupId string) string {
  return fmt.Sprintf("%sgroups/%s/action", url, groupId)
}

var client = &httpClient.Client{}

func (light *LightCommand) run(c *kingpin.ParseContext) error {
  resp, _ := client.Get(getLightsUrl())
  defer resp.Body.Close()
  bodyBytes, _ := ioutil.ReadAll(resp.Body)
  utils.PrettyPrint(bodyBytes)
  return nil
}

func (lights *RoomLightCommand) run(c *kingpin.ParseContext) error {
  state := os.Args[2] == "on"
  url := groupsActionUrl(lights.LightGroup.Id)
  var jsonStr = fmt.Sprintf(`{"on":%t}`, state)

  client.Put(url, jsonStr)
  return nil
}

func configureLightCommand(app *kingpin.Application) {
  c:= &LightCommand{}
  app.Command("lights", "Get light info").Action(c.run)
}

func configureRoomsLightCommand(app *kingpin.Application) {
  var result map[string]interface{}
  resp, _ := client.Get(getGroupsUrl())
  defer resp.Body.Close()
  bodyBytes, _ := ioutil.ReadAll(resp.Body)
  json.Unmarshal(bodyBytes, &result)

  for i := 0; i < len(result); i++ {
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
