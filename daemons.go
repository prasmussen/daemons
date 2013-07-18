package main

import (
    "fmt"
    "os"
    "encoding/json"
    "./tmux"
)

func parseConfig(fname string) (*tmux.Session, error) {
    f, err := os.Open(fname)
    if err != nil {
        return nil, err
    }
    defer f.Close()

    var session *tmux.Session
    if err := json.NewDecoder(f).Decode(&session); err != nil {
        return nil, err
    }

    return session, nil
}

func printDefaultConfig() {
    session := &tmux.Session{
        Name: "daemons",
        Windows: []tmux.Window{
            {
                Name: "hello world",
                Cmds: []string{"cd /tmp", "echo hello world"},
            },
            {
                Name: "uptime",
                Cmds: []string{"uptime"},
            },
        },
    }
    data, _ := json.MarshalIndent(session, "", "  ")
    fmt.Println(string(data))
}

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage:")
        fmt.Printf("%s config.json\tStart daemons defined in config\n", os.Args[0])
        fmt.Printf("%s --example\tPrint example config\n", os.Args[0])
        return
    }

    if os.Args[1] == "--example" {
        printDefaultConfig()
        return
    }

    session, err := parseConfig(os.Args[1])
    if err != nil {
        fmt.Println(err)
        return
    }

    err = session.Create()
    if err != nil {
        fmt.Println(err)
        return
    }
}
