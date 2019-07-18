package room

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	hueclient "github.com/jsncmgs1/hue_cli/lib/client"
	"github.com/jsncmgs1/hue_cli/lib/utils"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	client = hueclient.New()
)

type LightCommand struct{}

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
	resp, err := client.Get(hueURL("lights"))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(utils.PrettyPrintJSON(bodyBytes))
	return nil
}

func buildRoomCommand(app *kingpin.Application, result map[string]map[string]string, id string) {
	first := result[id]
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
	result := make(map[string]map[string]string)
	roomURL := hueURL("groups")

	resp, err := client.Get(roomURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bodyBytes, &result)

	for id := range result {
		buildRoomCommand(app, result, id)
	}
	return nil
}

func ConfigureLightCommand(app *kingpin.Application) {
	c := &LightCommand{}
	app.Command("lights", "Get light info").Action(c.run)
}
