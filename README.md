# task-cli

A simple CLI tool to create Jira tasks under an epic and corresponding GitHub branches. Uses your existing GitHub authentication and branches out of the branch you are currently on unless specified otherwise. Once configured, creating a ticket and branch can be as simple as running:

```bash
task-cli new --task="implement-oauth"
```

---

## Prerequisites

- Go 1.20+ installed ([install Go](https://go.dev/doc/install))
- Git installed
- You should be already authenticated with:
  - **GitHub:** via SSH
  - **Jira:** environment variables `JIRA_USERNAME`, `JIRA_API_TOKEN`

---

## Build and Install

```bash
git clone https://github.com/kcpopp/task-cli.git
cd task-cli
make install
```

Make sure the tool is added to path.

For zsh run:
```zsh
echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.zshrc
source ~/.zshrc
```

For bash run:
```bash
echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.bashrc
source ~/.bashrc
```

## Environment Variables

Before using `task-cli`, export these environment variables in your terminal:

```bash
export JIRA_USERNAME="your_jira_email@example.com"
export JIRA_API_TOKEN="your_jira_api_token_here"
```

## Set configuration

You can set defaults such as a Jira epic and GitHub repo to avoid typing them every time:

```bash
task-cli config --default-epic="EPIC" --default-repo="REPO" --default-project="PROJECT" --jira-base-url="https://yourcompany.atlassian.net" --default-branch-from-branch="CURRENT"
```

Run: 
```bash
task-cli config --help
```
 for a full list of options.

## Create Task Examples

Specifying epic, repo and Jira project

```bash
task-cli new --epic="EPIC" --task="Add auth to checkout" --task-description="My description" --repo="REPO" --project="PROJECT"
```

Using the defaults:

```bash
task-cli new --task="Fix login bug"
```

From a specific branch:

```bash
task-cli new --task="Fix login bug --from-branch="feat/my-main-branch""
```

Run: 
```bash
task-cli new --help
```
 for a full list of options.

# Troubleshooting

```bash
task-cli --help
task-cli new --help
task-cli config --help
```
