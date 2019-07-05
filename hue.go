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

type KitchenLightCommand struct {
  All bool
}

func (lights *KitchenLightCommand) run(c *kingpin.ParseContext) error {
  state := os.Args[2] == "on"

  for i:= 0; i < len(kitchenLights); i++ {
    url:= fmt.Sprintf("%slights/%d/state", url, kitchenLights[i])
    var jsonStr = []byte(fmt.Sprintf(`{"on":%t}`, state))
    req.Put(url, bytes.NewBuffer(jsonStr))
  }

  return nil
}

func configureKitchenLightCommand(app *kingpin.Application) {
  c := &KitchenLightCommand{}
  kitchenCommand  := app.Command("kitchen", "Control Kitchen Lights")
  kitchenCommand.Command("off", "turn kitchen lights off").Action(c.run)
  kitchenCommand.Command("on", "turn kitchen lights on").Action(c.run)
}

func configureCommands(app *kingpin.Application) {
  configureLightCommand(app)
  configureKitchenLightCommand(app)
}

func main() {
  app := kingpin.New("hue", "Phillips Hue CLI")
  configureCommands(app)
  kingpin.MustParse(app.Parse(os.Args[1:]))
}
