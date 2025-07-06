package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Set or view configuration",
	Run: func(cmd *cobra.Command, args []string) {
		// Read flags
		if defaultJiraBaseURL, _ := cmd.Flags().GetString("jira-base-url"); defaultJiraBaseURL != "" {
			viper.Set("default_jira_base_url", defaultJiraBaseURL)
		}
		if defaultEpic, _ := cmd.Flags().GetString("default-epic"); defaultEpic != "" {
			viper.Set("default_epic", defaultEpic)
		}
		if defaultRepo, _ := cmd.Flags().GetString("default-repo"); defaultRepo != "" {
			viper.Set("default_repo", defaultRepo)
		}
		if defaultProject, _ := cmd.Flags().GetString("default-project"); defaultProject != "" {
			viper.Set("default_project", defaultProject)
		}
		// Save config file to $HOME/.task-cli.yaml
		configPath := os.ExpandEnv("$HOME/.task-cli.yaml")
		err := viper.WriteConfigAs(configPath)
		if err != nil {
			// If config doesn't exist, create it
			err = viper.SafeWriteConfigAs(configPath)
			if err != nil {
				fmt.Println("Error writing config:", err)
				os.Exit(1)
			}
		}

		fmt.Println("Config updated at", configPath)
	},
}

func init() {
	configCmd.Flags().String("jira-base-url", "", "Jira base URL")
	configCmd.Flags().String("default-epic", "", "Default Jira epic")
	configCmd.Flags().String("default-repo", "", "Default GitHub repo")
	configCmd.Flags().String("default-project", "", "Default Jira project")
	rootCmd.AddCommand(configCmd)
}
