package room_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/jsncmgs1/hue_cli/lib/room"
	"gopkg.in/alecthomas/kingpin.v2"
)

func TestConfigureRoomsLightCommand(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"1": {"name": "Kitchen"}, "2": {"name": "Office"}}`)
	}))

	u, _ := url.Parse(ts.URL)
	oldHueUrl := os.Getenv("HUE_URL")
	os.Setenv("HUE_URL", fmt.Sprintf("%s/", u.String()))
	defer os.Setenv("HUE_URL", oldHueUrl)

	app := kingpin.New("hue", "testing app")
	room.ConfigureRoomsLightCommand(app)

	cmd := app.GetCommand("kitchen")
	off := cmd.GetCommand("off")
	on := cmd.GetCommand("on")

	if cmd == nil {
		t.Error("Expected Kitchen command to exist")
	}

	if on == nil {
		t.Error("Expected Kitchen sub command on to exist")
	}

	if off == nil {
		t.Error("Expected Kitchen sub command off to exist")
	}

	cmd = app.GetCommand("office")
	off = cmd.GetCommand("off")
	on = cmd.GetCommand("on")

	if cmd == nil {
		t.Error("Expected office command to exist")
	}

	if on == nil {
		t.Error("Expected Office sub command on to exist")
	}

	if off == nil {
		t.Error("Expected Office sub command off to exist")
	}
}
