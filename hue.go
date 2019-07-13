package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/jsncmgs1/hue_cli/lib/utils"
	"gopkg.in/alecthomas/kingpin.v2"
)

type LightCommand struct{}

type RoomLightCommand struct {
	LightGroup *Group
}

type Group struct {
	Name string
	ID   string
}

var url = os.Getenv("HUE_URL")

func groupsActionUrl(groupId string) string {
	return fmt.Sprintf("%sgroups/%s/action", url, groupId)
}

var client = &http.Client{}

func (light *LightCommand) run(c *kingpin.ParseContext) error {
	url := fmt.Sprintf("%slights", url)
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	utils.PrettyPrintJSON(bodyBytes)
	return nil
}

func putRequest(url string, data io.Reader) {
	req, err := http.NewRequest(http.MethodPut, url, data)
	if err != nil {
		log.Fatal(err)
	}
	_, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
}
func (lights *RoomLightCommand) run(c *kingpin.ParseContext) error {
	state := os.Args[2] == "on"
	url := fmt.Sprintf("%sgroups/%s/action", url, lights.LightGroup.ID)
	jsonStr := fmt.Sprintf(`{"on":%t}`, state)
	putRequest(url, strings.NewReader(jsonStr))
	return nil
}

func configureLightCommand(app *kingpin.Application) {
	c := &LightCommand{}
	app.Command("lights", "Get light info").Action(c.run)
}

func configureRoomsLightCommand(app *kingpin.Application) {
	var result map[string]interface{}
	url := fmt.Sprintf("%sgroups", url)
	resp, _ := client.Get(url)
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bodyBytes, &result)

	for i := 0; i < len(result); i++ {
		c := &RoomLightCommand{}
		id := strconv.Itoa(i + 1)
		first := result[id].(map[string]interface{})
		name := strings.ToLower(first["name"].(string))
		c.LightGroup = &Group{Name: name, ID: id}

		on := fmt.Sprintf("Turn %s lights on", name)
		off := fmt.Sprintf("Turn %s lights off", name)

		roomCommand := app.Command(name, "")
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
