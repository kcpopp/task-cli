package cmd

import (
	"fmt"
	"task-cli/internal"

	"github.com/andygrunwald/go-jira"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var epic string
var task string
var taskDescription string

var repo string
var fromBranch string
var branchPrefix string

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create new Jira task and Git branch",
	Run: func(cmd *cobra.Command, args []string) {
		// Get flags first
		epicFlag, _ := cmd.Flags().GetString("epic")
		repoFlag, _ := cmd.Flags().GetString("repo")
		taskFlag, _ := cmd.Flags().GetString("task")
		taskDescription, _ := cmd.Flags().GetString("task-description")
		projectFlag, _ := cmd.Flags().GetString("project")
		baseURLFlag, _ := cmd.Flags().GetString("jira-base-url")
		fromBranch, _ := cmd.Flags().GetString("from-branch")
		branchPrefix, _ := cmd.Flags().GetString("branch-prefix")

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

		fromBranchToUse := fromBranch
		if fromBranchToUse == "" {
			fromBranchToUse = viper.GetString("default_branch_from_branch")
			if fromBranchToUse == "" || fromBranchToUse == "CURRENT" {
				fromBranchToUse = internal.GetCurrentGitBranch()
			}

		}
		branchPrefixToUse := branchPrefix
		if branchPrefixToUse == "" {
			branchPrefixToUse = "feat"

		}
		if taskFlag == "" {
			fmt.Println("Task description (--task) is required")
			return
		}

		issue := &jira.Issue{}
		jiraClient, err := internal.GetJiraClient(jiraBaseUrlToUse)

		currentBranchIssueKey := internal.GetIssueKeyFromBranchName(fromBranchToUse)

		if err != nil {
			fmt.Println("Failed to get current branch:", err)
			return
		}
		if fromBranchToUse == "develop" || currentBranchIssueKey == "" {
			issue, err = internal.CreateTask(jiraClient, projectToUse, epicToUse, taskFlag, taskDescription)

		} else {
			issue, err = internal.CreateSubTask(jiraClient, currentBranchIssueKey, projectToUse, epicToUse, taskFlag, taskDescription)

		}

		if err != nil {
			fmt.Println("Failed to create Jira client:", err)
			return
		}
		if err != nil {
			fmt.Println("Failed to create Jira task:", err)
			return
		}

		branchName, err := internal.CreateGitBranch(repoToUse, branchPrefixToUse, issue.Key, taskFlag)
		if err != nil {
			fmt.Println("Git error:", err)
			return
		}

		fmt.Printf("✅ Created Jira task %s and branch %s\n", issue.Key, branchName)
	},
}

func init() {
	newCmd.Flags().StringVar(&epic, "epic", "", "Epic name")
	newCmd.Flags().StringVar(&task, "task", "", "Task title (required)")
	newCmd.Flags().StringVar(&taskDescription, "task-description", "", "Task description")
	newCmd.Flags().StringVar(&repo, "repo", "", "Github repository name.")
	newCmd.Flags().StringVar(&fromBranch, "from-branch", "", "Branch from which to branch out of. If not provided, the current branch will be used. If the current branch is develop, the Jira issue will be created as a TASK. Otherwise, it will be created as a SUBTASK of the task associated with the current branch. If the current branch is not linked to a JIRA task, the CLI will branch out of develop and create a task.")
	newCmd.Flags().StringVar(&branchPrefix, "branch-prefix", "", "Semantic branch prefix. Defaults to feat.")

	rootCmd.AddCommand(newCmd)
}
