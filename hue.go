package main

import (
	"os"

	"github.com/jsncmgs1/hue_cli/lib/room"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	app := kingpin.New("hue", "Phillips Hue CLI")
	room.ConfigureLightCommand(app)
	room.ConfigureRoomsLightCommand(app)
	kingpin.MustParse(app.Parse(os.Args[1:]))
}
