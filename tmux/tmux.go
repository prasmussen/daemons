package tmux

import (
    "fmt"
)

type Session struct {
    Name string `json:"name"`
    Windows []Window `json:"windows"`
}

func (self *Session) Create() error {
    if err := run("tmux", "new-session", "-s", self.Name, "-d"); err != nil {
        return err
    }

    err := self.createWindows()

    // Kill window 0, daemons is 1 indexed
    windowZero := fmt.Sprintf("%s:0", self.Name)
    run("tmux", "kill-window", "-t", windowZero)

    return err
}

func (self *Session) createWindows() error {
    for i, w := range self.Windows {
        if err := w.Create(self.Name, i + 1); err != nil {
            return err
        }
    }

    return nil
}

type Window struct {
    Name string `json:"name"`
    Cmds []string `json:"cmds"`
}

func (self *Window) Create(session string, index int) error {
    target := fmt.Sprintf("%s:%d", session, index)
    if err := run("tmux", "new-window", "-k", "-t", target, "-n", self.Name); err != nil {
        return err
    }

    cmdTarget := fmt.Sprintf("%s:%s", session, self.Name)

    for _, cmd := range self.Cmds {
        if err := self.SendKeys(cmdTarget, cmd); err != nil {
            return err
        }
    }
    return nil
}

func (self *Window) SendKeys(target, cmd string) error {
    return run("tmux", "send-keys", "-t", target, cmd, "C-m")
}
