# Phillips HUE CLI

This is really just an excuse for me to learn Go in my off time. I would not
depend on this library for anything serious. Proceed with caution. 
I have not run this on Windows or Linux.

## Setup

1. Make sure your `GOBIN` is correctly setup.
2. Follow https://developers.meethue.com/develop/get-started-2/ and get your UserId and the bridge IP.
3. `export HUE_CLI=http://<BRIDGE IP>/api/<YOUR IP>/` in your shell config.
2. `$go install .`


## Useage

```bash
$ hue_cli
```

Shows you available rooms set up on your bridge. Pretend you have a kitchen and office:

```bash
$ hue_cli
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

Run `hue_cli lights` to get the JSON for all of your lights.
Run `hue_cli office on` to turn your office lights on.
Run `hue_cli kitchen off` to turn your office kitchen on.
