package remote

import (
	"fmt"
	"wolremote/src/config"
	"wolremote/src/ssh"
)

// Controller manages a remote program via SSH
type Controller struct {
	SSHClient *ssh.Client
	Config    config.RemoteProgramConfig
	Password  string
}

// NewController creates a new remote program controller
func NewController(client *ssh.Client, cfg config.RemoteProgramConfig, password string) *Controller {
	return &Controller{
		SSHClient: client,
		Config:    cfg,
		Password:  password,
	}
}

// StartProgram starts the remote program
func (c *Controller) StartProgram() (string, error) {
	cmd := fmt.Sprintf("cd %s && nohup %s > /dev/null 2>&1 &",
		c.Config.FolderPath, c.Config.ProgramName)

	if c.Config.SudoUser {
		cmd = fmt.Sprintf("sudo -S -p '' %s", cmd)
		cmd = fmt.Sprintf("echo %q | %s", c.Password, cmd)
	}

	return c.SSHClient.RunCommand(cmd)
}

// StopProgram stops the remote program
func (c *Controller) StopProgram() (string, error) {
	cmd := fmt.Sprintf("pkill -f %s", c.Config.ProgramName)

	if c.Config.SudoUser {
		cmd = fmt.Sprintf("sudo -S -p '' %s", cmd)
		cmd = fmt.Sprintf("echo %q | %s", c.Password, cmd)
	}

	return c.SSHClient.RunCommand(cmd)
}
