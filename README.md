daemons
=======

Start configured applications inside a tmux session.
The configuration file consists of a session name and
one or more windows with a list of commands that will
execute inside that window.

## Usage
    daemons config.json   Start daemons defined in config
    daemons --example     Print example config

## Example config
    {
      "name": "daemons",
      "windows": [
        {
          "name": "hello world",
          "cmds": [
            "cd /tmp",
            "echo hello world"
          ]
        },
        {
          "name": "uptime",
          "cmds": [
            "uptime"
          ]
        }
      ]
    }
