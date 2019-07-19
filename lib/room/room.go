package room

import (
	"fmt"
	"os"
	"strings"

	hueclient "github.com/jsncmgs1/hue_cli/lib/client"
	"github.com/jsncmgs1/hue_cli/lib/utils"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	client = hueclient.New()
)

type LightCommand struct {
	JSON bool
}

type RoomLightCommand struct {
	LightGroup
}

type LightGroup struct {
	Name string
	ID   string
}

func hueURL(path string) string {
	return fmt.Sprintf("%s%s", os.Getenv("HUE_URL"), path)
}

func (light *LightCommand) run(c *kingpin.ParseContext) error {
	resp := client.GetJSON(
		hueURL("lights"),
		make(map[string]map[string]string),
	)
	lights := make(map[string]string)

	for id := range resp.Body {
		lights[id] = resp.Body[id]["name"]
	}

	if resp.Error != nil {
		return resp.Error
	}

	if light.JSON == true {
		fmt.Println(resp.Body)
	} else {
		utils.PrintSortedMap(lights)
	}
	return nil
}

func buildRoomCommand(app *kingpin.Application, result hueclient.JSONResponse, id string) {
	first := result.Body[id]
	name := strings.ToLower(first["name"])
	c := &RoomLightCommand{LightGroup{Name: name, ID: id}}

	on := fmt.Sprintf("Turn %s lights on", name)
	off := fmt.Sprintf("Turn %s lights off", name)

	roomCommand := app.Command(name, "")
	roomCommand.Command("off", on).Action(c.run)
	roomCommand.Command("on", off).Action(c.run)
}

func (lights *RoomLightCommand) run(c *kingpin.ParseContext) error {
	state := os.Args[2] == "on"
	roomURL := hueURL(fmt.Sprintf("groups/%s/action", lights.LightGroup.ID))
	jsonStr := fmt.Sprintf(`{"on":%t}`, state)
	client.Put(roomURL, strings.NewReader(jsonStr))
	return nil
}

func ConfigureRoomsLightCommand(app *kingpin.Application) error {
	response := client.GetJSON(
		hueURL("groups"),
		make(map[string]map[string]string),
	)

	if response.Error != nil {
		return response.Error
	}

	for id := range response.Body {
		buildRoomCommand(app, response, id)
	}

	return nil
}

func ConfigureLightCommand(app *kingpin.Application) {
	lc := &LightCommand{}
	c := app.Command("lights", "Get light info")
	c.Flag("json", "Returns JSON data for all lights").BoolVar(&lc.JSON)
	c.Action(lc.run)
}
