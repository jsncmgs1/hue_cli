package room_test

import (
	"testing"

	"github.com/jsncmgs1/hue_cli/lib/room"
	"gopkg.in/alecthomas/kingpin.v2"
)

func TestBuildRoomCommadn(t *testing.T) {
	app := kingpin.New("hue", "testing app")
	room.BuildRoomCommand(app)

	if true {
		t.Error("make me pass")
	}
}
