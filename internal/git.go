package internal

import (
	"fmt"
	"os/exec"
	"strings"
)

func GetCurrentGitBranch() string {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	currentBranchOutput, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("failed to get current branch: %w, output: %s", err, string(currentBranchOutput))
		return ""
	}
	currentBranch := strings.TrimSpace(string(currentBranchOutput))
	return currentBranch
}

func CreateGitBranch(repo, branchPrefix string, issueKey, summary string) (string, error) {

	newBranchName := makeBranchName(branchPrefix, issueKey, summary)
	cmd := exec.Command("git", "checkout", "-b", newBranchName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to create new branch: %w, output: %s", err, string(output))
	}

	cmd = exec.Command("git", "push", "-u", "origin", newBranchName)
	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("failed to push new branch: %w", err)
	}

	return newBranchName, nil
}

func makeBranchName(branchPrefix, issueKey string, summary string) string {
	return fmt.Sprintf("%s/%s-%s", branchPrefix, issueKey, slugify(summary))
}

func GetIssueKeyFromBranchName(branchName string) string {
	nameNoPrefix := strings.Split(branchName, "/")
	if len(nameNoPrefix) < 2 {
		return ""
	}
	parts := strings.Split(nameNoPrefix[1], "-")

	if len(parts) < 2 {
		return ""
	}
	return strings.Join(parts[:2], "-")
}

func slugify(s string) string {
	// very naive example:
	return strings.ReplaceAll(strings.ToLower(s), " ", "-")
}
