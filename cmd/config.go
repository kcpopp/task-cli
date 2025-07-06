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
		if defaultBranchFromBranch, _ := cmd.Flags().GetString("default-branch-from-branch"); defaultBranchFromBranch != "" {
			viper.Set("default_branch_from_branch", defaultBranchFromBranch)
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
	configCmd.Flags().String("jira-base-url", "", "Jira base URL.")
	configCmd.Flags().String("default-epic", "", "Default Jira epi.")
	configCmd.Flags().String("default-repo", "", "Default GitHub repository.")
	configCmd.Flags().String("default-project", "", "Default Jira project code.")
	configCmd.Flags().String("default-branch-from-branch", "", "Default branch to branch out of. Seto to CURRENT to always branch out of the current branch, or a branch name such as develop or master. Defaults to CURRENT.")

	rootCmd.AddCommand(configCmd)
}
