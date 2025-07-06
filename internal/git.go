package internal

import (
    "context"
    "fmt"
    "os"

    "github.com/google/go-github/v53/github"
    "golang.org/x/oauth2"
)

func newGitHubClient(token string) *github.Client {
    ctx := context.Background()
    ts := oauth2.StaticTokenSource(
        &oauth2.Token{AccessToken: token},
    )
    tc := oauth2.NewClient(ctx, ts)
    client := github.NewClient(tc)
    return client
}

// CreateGitBranch creates a git branch in the given repo named after the issue key and task summary
func CreateGitBranch(repo, issueKey, summary string) error {
    token := os.Getenv("GITHUB_TOKEN")
    if token == "" {
        token = os.Getenv("GH_TOKEN")
    }
    if token == "" {
        return fmt.Errorf("GITHUB_TOKEN or GH_TOKEN environment variable required")
    }

    client := newGitHubClient(token)
    ctx := context.Background()

    // get default branch to branch off from
    repository, _, err := client.Repositories.Get(ctx, "your-github-org-or-user", repo)
    if err != nil {
        return fmt.Errorf("failed to get repo info: %w", err)
    }

    defaultBranch := repository.GetDefaultBranch()

    // get ref for the default branch
    ref, _, err := client.Git.GetRef(ctx, "your-github-org-or-user", repo, "refs/heads/"+defaultBranch)
    if err != nil {
        return fmt.Errorf("failed to get ref of default branch: %w", err)
    }

    newBranchName := fmt.Sprintf("%s-%s", issueKey, slugify(summary))

    // create new ref (branch)
    newRef := &github.Reference{
        Ref:    github.String("refs/heads/" + newBranchName),
        Object: &github.GitObject{SHA: ref.Object.SHA},
    }

    _, _, err = client.Git.CreateRef(ctx, "your-github-org-or-user", repo, newRef)
    if err != nil {
        return fmt.Errorf("failed to create git ref: %w", err)
    }

    return nil
}

// slugify is a helper to make branch names safe (implement as you want)
func slugify(s string) string {
    // very naive example:
    return strings.ReplaceAll(strings.ToLower(s), " ", "-")
}
