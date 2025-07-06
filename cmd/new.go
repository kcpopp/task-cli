package cmd

import (
	"fmt"
	"task-cli/internal"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var epic string
var task string
var repo string

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create new Jira task and Git branch",
	Run: func(cmd *cobra.Command, args []string) {
		// Get flags first
		epicFlag, _ := cmd.Flags().GetString("epic")
		repoFlag, _ := cmd.Flags().GetString("repo")
		taskFlag, _ := cmd.Flags().GetString("task")
		projectFlag, _ := cmd.Flags().GetString("project")
		baseURLFlag, _ := cmd.Flags().GetString("jira-base-url")

		// Fallback to config if flags empty
		epicToUse := epicFlag
		if epicToUse == "" {
			epicToUse = viper.GetString("default_epic")
		}
		repoToUse := repoFlag
		if repoToUse == "" {
			repoToUse = viper.GetString("default_repo")
		}
		projectToUse := projectFlag
		if projectToUse == "" {
			projectToUse = viper.GetString("default_project")
		}

		jiraBaseUrlToUse := baseURLFlag
		if jiraBaseUrlToUse == "" {
			jiraBaseUrlToUse = viper.GetString("default_jira_base_url")
		}

		if taskFlag == "" {
			fmt.Println("Task description (--task) is required")
			return
		}

		issue, err := internal.CreateJiraTask(jiraBaseUrlToUse, projectToUse, epicToUse, taskFlag)
		if err != nil {
			fmt.Println("Failed to create Jira task:", err)
			return
		}

		err = internal.CreateGitBranch(repoToUse, issue.Key, taskFlag)
		if err != nil {
			fmt.Println("Git error:", err)
			return
		}

		fmt.Printf("✅ Created Jira task %s and branch\n", issue.Key)
	},
}

func init() {
	newCmd.Flags().StringVar(&epic, "epic", "", "Epic name")
	newCmd.Flags().StringVar(&task, "task", "", "Task description (required)")
	newCmd.Flags().StringVar(&repo, "repo", "", "Repository name")
	rootCmd.AddCommand(newCmd)
}
