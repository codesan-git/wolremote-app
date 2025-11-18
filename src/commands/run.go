package commands

import (
	"fmt"
	"path/filepath"
	"time" // Import time package for delay
	"wolremote/src/checker"
	"wolremote/src/config"
	"wolremote/src/ssh"

	"github.com/spf13/cobra"
)

var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the remote program in the background",
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

		// Prepare the command to run the remote program in the background using nohup and disown
		remoteCmd := fmt.Sprintf("nohup bash -c 'cd %s && %s' > /dev/null 2>&1 & disown",
			filepath.Clean(cfg.RemoteProgram.FolderPath), cfg.RemoteProgram.ProgramName)
		if cfg.RemoteProgram.SudoUser {
			remoteCmd = "sudo " + remoteCmd
		}

		// Run the command without capturing output
		_, err = client.RunCommand(remoteCmd)
		if err != nil {
			fmt.Println("Failed to run remote program:", err)
			return
		}

		// Output confirmation
		fmt.Println("Execute program in background:", cfg.RemoteProgram.ProgramName)

		// Adding a small delay to allow the program to start
		time.Sleep(2 * time.Second) // Delay of 2 seconds

		// Check if the program is running after the delay
		running, err := checker.CheckIfRunning(client, cfg.RemoteProgram.ProgramName)
		if err != nil {
			fmt.Println("Error checking if program is running:", err)
			return
		}

		// Output based on whether the program is running
		if running {
			fmt.Println("[INFO] The program is running!")
		} else {
			fmt.Println("[INFO] The program is not running.")
		}
	},
}

func init() {
	// Add the --remote flag
	RunCmd.Flags().String("remote", "", "Path to the remote PC config file")
	RunCmd.MarkFlagRequired("remote")

	// Add RunCmd to the root command
	RootCmd.AddCommand(RunCmd)
}
