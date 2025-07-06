package internal

import (
    "fmt"
    "os"
    "github.com/andygrunwald/go-jira"
)

func CreateJiraTask(project, epic, summary string) (*jira.Issue, error) {
    baseURL := os.Getenv("JIRA_BASE_URL")
    username := os.Getenv("JIRA_USERNAME")
    token := os.Getenv("JIRA_API_TOKEN")
    if baseURL == "" || username == "" || token == "" {
        return nil, fmt.Errorf("please set JIRA_BASE_URL, JIRA_USERNAME, and JIRA_API_TOKEN environment variables")
    }

    tp := jira.BasicAuthTransport{
        Username: username,
        Password: token,
    }

    client, err := jira.NewClient(tp.Client(), baseURL)
    if err != nil {
        return nil, err
    }

    issueReq := jira.Issue{
        Fields: &jira.IssueFields{
            Summary:     summary,
            Description: summary,
            Project: jira.Project{
                Key: project, // assuming epic is project key; adjust if needed
            },
            Type: jira.IssueType{
                Name: "Task",
            },
			"customfield_10000": epic,
        },
    }

    createdIssue, resp, err := client.Issue.Create(&issueReq)
    if err != nil {
        return nil, fmt.Errorf("jira create error: %v, response: %v", err, resp)
    }

    return createdIssue, nil
}
