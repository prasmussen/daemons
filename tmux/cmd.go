package tmux

import (
    "os/exec"
)

func run(args ...string) error {
    if len(args) == 0 {
        return nil
    }

    if len(args) == 1 {
        return exec.Command(args[0]).Run()
    }

    return exec.Command(args[0], args[1:]...).Run()
}
