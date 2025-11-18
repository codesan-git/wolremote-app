package commands

import (
	"fmt"
	"wolremote/src/config"
	"wolremote/src/ssh"

	"github.com/spf13/cobra"
)

var OffCmd = &cobra.Command{
	Use:   "off",
	Short: "Shutdown remote PC",
	Run: func(cmd *cobra.Command, args []string) {
		// Get the config file path from the --remote flag
		cfgFile, _ := cmd.Flags().GetString("remote")

		// Load the configuration
		cfg, err := config.LoadConfig(cfgFile)
		if err != nil {
			fmt.Println("Error loading config:", err)
			return
		}

		// Connect via SSH
		client, err := ssh.NewClient(cfg.SSH)
		if err != nil {
			fmt.Println("SSH connection failed:", err)
			return
		}
		defer client.Close()

		// Prepare the shutdown command with sudo password
		// Here we echo the password to sudo using -S
		remoteCmd := fmt.Sprintf("echo '%s' | sudo -S shutdown 0", cfg.SSH.Password)

		// Run shutdown command
		output, err := client.RunCommand(remoteCmd)
		if err != nil {
			fmt.Println("Failed to shutdown remote PC:", err)
			fmt.Println("Error output:", output)
			return
		}

		fmt.Println("Remote PC shutdown triggered for", cfgFile)
		fmt.Println("Output:", output)
	},
}

func init() {
	// Add the --remote flag
	OffCmd.Flags().String("remote", "", "Path to the remote PC config file")
	OffCmd.MarkFlagRequired("remote")

	// Add OffCmd to the root command
	RootCmd.AddCommand(OffCmd)
}
