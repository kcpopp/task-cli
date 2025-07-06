# task-cli

A simple CLI tool to create Jira tasks under an epic and corresponding GitHub branches. Uses your existing Jira and GitHub authentication configured in environment variables or local tools.

---

## Prerequisites

- Go 1.20+ installed ([install Go](https://go.dev/doc/install))
- Git installed
- You should be already authenticated with:
  - **GitHub:** via environment variable `GITHUB_TOKEN` or `GH_TOKEN`, or `gh auth login`
  - **Jira:** environment variables `JIRA_BASE_URL`, `JIRA_USERNAME`, `JIRA_API_TOKEN`

---

## Build and Install

```bash
git clone https://github.com/yourusername/task-cli.git
cd task-cli
make build
make install
```

## Environment Variables

Before using `task-cli`, export these environment variables in your shell:

```bash
export GITHUB_TOKEN="your_github_token_here"          # or GH_TOKEN
export JIRA_BASE_URL="https://yourcompany.atlassian.net"
export JIRA_USERNAME="your_jira_email@example.com"
export JIRA_API_TOKEN="your_jira_api_token_here"
```

## Set configuration

You can set default Jira epic and GitHub repo to avoid typing them every time:

```bash
task-cli config --default-epic="EPIC" --default-repo="REPO" --default-project="PROJECT" --jira-base-url="https://yourcompany.atlassian.net"
```

## Create Task

```bash
task-cli new --epic="EPIC" --task="Add auth to checkout" --repo="REPO" --project="PROJECT"
```

or use the defaults:

```bash
task-cli new --task="Fix login bug"
```

# Troubleshooting

```bash
task-cli --help
task-cli new --help
task-cli config --help
```