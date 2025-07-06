package internal

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/andygrunwald/go-jira"
	"github.com/trivago/tgo/tcontainer"
)

func GetJiraClient(baseURL string) (*jira.Client, error) {
	username := os.Getenv("JIRA_USERNAME")
	token := os.Getenv("JIRA_API_TOKEN")
	if username == "" || token == "" {
		return nil, fmt.Errorf("please set JIRA_USERNAME, and JIRA_API_TOKEN environment variables")
	}

	tp := jira.BasicAuthTransport{
		Username: username,
		Password: token,
	}

	client, err := jira.NewClient(tp.Client(), baseURL)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func CreateSubTask(client *jira.Client, parentIssueKey, project, epic, summary string, description string) (*jira.Issue, error) {

	issueReq := jira.Issue{
		Fields: &jira.IssueFields{
			Summary:     summary,
			Description: description,
			Project: jira.Project{
				Key: project,
			},
			Type: jira.IssueType{
				Name: "Sub-task",
			},
			Parent: &jira.Parent{
				Key: parentIssueKey,
			},
		},
	}

	createdIssue, resp, err := client.Issue.Create(&issueReq)
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to create Jira task: %v, body: %s", err, string(body))
	}

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("failed to create Jira task, status: %s, response: %s", resp.Status, string(body))
	}

	return createdIssue, nil
}
func CreateTask(client *jira.Client, project, epic, summary string, description string) (*jira.Issue, error) {
	issueReq := jira.Issue{
		Fields: &jira.IssueFields{
			Summary:     summary,
			Description: description,
			Project: jira.Project{
				Key: project,
			},
			Type: jira.IssueType{
				Name: "Task",
			},
			Unknowns: tcontainer.MarshalMap{
				"customfield_10000": epic,
			},
		},
	}

	createdIssue, resp, err := client.Issue.Create(&issueReq)
	if err != nil {
		return nil, fmt.Errorf("failed to create Jira task: %v", err)
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("failed to create Jira task, status: %s, response: %s", resp.Status, string(body))
	}

	return createdIssue, nil
}
