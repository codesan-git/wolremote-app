package ssh

import (
	"fmt"
	"wolremote/src/config"

	"github.com/bugfan/gossh"
)

// Client wraps a gossh.Client
type Client struct {
	cfg    config.SSHConfig
	client *gossh.Client
}

// NewClient creates and connects an SSH client
func NewClient(cfg config.SSHConfig) (*Client, error) {
	auth := gossh.Password(cfg.Password)

	cli, err := gossh.New(cfg.User, cfg.Host, uint(cfg.Port), auth)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to %s:%d: %w", cfg.Host, cfg.Port, err)
	}

	return &Client{
		cfg:    cfg,
		client: cli,
	}, nil
}

// RunCommand runs a simple command on the remote PC
func (c *Client) RunCommand(cmd string) (string, error) {
	out, err := c.client.Run(cmd)
	if err != nil {
		return "", fmt.Errorf("error running command '%s': %w", cmd, err)
	}
	return string(out), nil
}

// Close closes the SSH connection
func (c *Client) Close() error {
	return c.client.Close()
}
