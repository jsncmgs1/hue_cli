# Phillips HUE CLI

This is really just an excuse for me to learn Go in my off time. I would not
depend on this library for anything serious. Proceed with caution.
I have not run this on Windows or Linux.

## Setup

This project uses GOMODULES. No need to clone to your GOPATH.

1. Make sure your `GOBIN` is correctly setup.
2. Follow https://developers.meethue.com/develop/get-started-2/ and get your UserId and the bridge IP.
3. `export HUE_CLI=http://<BRIDGE IP>/api/<YOUR IP>/` in your shell config.
4. `$ go build -o hue`
5. `$ go install hue.go`


## Useage

```bash
$ hue
```

Shows you available rooms set up on your bridge. Pretend you have a kitchen and office:

```bash
$ hue
usage: hue [<flags>] <command> [<args> ...]

Phillips Hue CLI

Flags:
  --help  Show context-sensitive help (also try --help-long and --help-man).

Commands:
  help [<command>...]
    Show help.

  lights
    Get light info

  office off
    Turn office lights on

  office on
    Turn office lights off

  kitchen off
    Turn kitchen lights on

  kitchen on
    Turn kitchen lights off
```

Run `hue lights` to get the JSON for all of your lights.
Run `hue office on` to turn your office lights on.
Run `hue kitchen off` to turn your kitchen lights off.

## Future things

Individual light control, scene control, dunno what else.
