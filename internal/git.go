package internal

import (
	"fmt"
	"os/exec"
	"strings"
)

func CreateGitBranch(repo, issueKey, summary string) error {

	newBranchName := fmt.Sprintf("%s-%s", issueKey, slugify(summary))
	cmd := exec.Command("git", "checkout", "-b", newBranchName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to create new branch: %w, output: %s", err, string(output))
	}

	// Push the new branch to the remote repository
	cmd = exec.Command("git", "push", "-u", "origin", newBranchName)
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to push new branch: %w", err)
	}

	return nil
}

func slugify(s string) string {
	// very naive example:
	return strings.ReplaceAll(strings.ToLower(s), " ", "-")
}
