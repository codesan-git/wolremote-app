package commands

import (
	"fmt"
	"strings"
	"wolremote/src/checker"
	"wolremote/src/config"
	"wolremote/src/ssh"

	"github.com/spf13/cobra"
)

var CheckCmd = &cobra.Command{
	Use:   "check",
	Short: "Check if the remote program is running",
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

		// Check if the remote program is running
		running, err := checker.CheckIfRunning(client, cfg.RemoteProgram.ProgramName)
		if err != nil {
			fmt.Println("Error checking if program is running:", err)
			return
		}

		// Output the result based on whether the program is running
		if running {
			// Command to get detailed process info, similar to ps aux
			checkCmd := fmt.Sprintf("ps aux | grep '%s' | grep -v grep", cfg.RemoteProgram.ProgramName)
			if cfg.RemoteProgram.SudoUser {
				checkCmd = "sudo " + checkCmd
			}

			// Run the command to get detailed process info
			output, err := client.RunCommand(checkCmd)
			if err != nil {
				fmt.Println("Failed to retrieve process details:", err)
				return
			}

			// Parse the output to format it neatly into a table
			// We will print the output in a human-readable table format

			lines := strings.Split(output, "\n")
			if len(lines) < 2 {
				fmt.Println("Program not found!")
				return
			}

			// Output the result as a table
			fmt.Println("Program is running! Here are the details:")
			fmt.Println("+----------------------------------------------------------------------------------------------------------+")
			fmt.Println("|  USER  |  PID  |  CPU%  |  RAM%  |  VSZ(KB)  |  RSS(KB)  |  TTY  | STAT |            COMMAND             |")
			fmt.Println("|--------+-------+--------+--------+-----------+-----------+-------+------+--------------------------------|")

			// Process each line of the ps output and format it into the table
			for _, line := range lines {
				// Skip empty lines
				if line == "" {
					continue
				}

				// Split the line into columns
				cols := strings.Fields(line)

				// Extract relevant columns
				user := cols[0]
				pid := cols[1]
				cpu := cols[2]
				mem := cols[3]
				vsz := cols[4]
				rss := cols[5]
				tty := cols[6]
				stat := cols[7]
				cmd := strings.Join(cols[10:], " ") // COMMAND can contain spaces

				// Print the formatted line in table format
				fmt.Printf("| %-6s | %-5s | %-6s | %-6s | %-9s | %-9s | %-5s | %-4s | %-30s |\n", user, pid, cpu, mem, vsz, rss, tty, stat, cmd)
			}

			fmt.Println("+----------------------------------------------------------------------------------------------------------+")
		} else {
			fmt.Printf("Program %s is NOT running on the remote machine.\n", cfg.RemoteProgram.ProgramName)
		}
	},
}

func init() {
	// Add the --remote flag
	CheckCmd.Flags().String("remote", "", "Path to the remote PC config file")
	CheckCmd.MarkFlagRequired("remote")

	// Add CheckCmd to the root command
	RootCmd.AddCommand(CheckCmd)
}
