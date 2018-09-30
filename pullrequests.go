package bitbucket

import (
	"encoding/json"
	"os"
	"fmt"

	"github.com/k0kubun/pp"
)

type PullRequests struct {
	c *Client
}

func (p *PullRequests) Create(po *PullRequestsOptions) (interface{}, error) {
	data := p.buildPullRequestBody(po)
	urlStr := p.c.requestUrl("/repositories/%s/%s/pullrequests/", po.Owner, po.RepoSlug)
	return p.c.execute("POST", urlStr, data)
}

func (p *PullRequests) Update(po *PullRequestsOptions) (interface{}, error) {
	data := p.buildPullRequestBody(po)
	urlStr := GetApiBaseURL() + "/repositories/" + po.Owner + "/" + po.RepoSlug + "/pullrequests/" + po.ID
	return p.c.execute("PUT", urlStr, data)
}

func (p *PullRequests) Gets(po *PullRequestsOptions) (interface{}, error) {
	urlStr := GetApiBaseURL() + "/repositories/" + po.Owner + "/" + po.RepoSlug + "/pullrequests/"
	return p.c.execute("GET", urlStr, "")
}

func (p *PullRequests) Get(po *PullRequestsOptions) (interface{}, error) {
	urlStr := GetApiBaseURL() + "/repositories/" + po.Owner + "/" + po.RepoSlug + "/pullrequests/" + po.ID
	return p.c.execute("GET", urlStr, "")
}

func (p *PullRequests) Activities(po *PullRequestsOptions) (interface{}, error) {
	urlStr := GetApiBaseURL() + "/repositories/" + po.Owner + "/" + po.RepoSlug + "/pullrequests/activity"
	return p.c.execute("GET", urlStr, "")
}

func (p *PullRequests) Activity(po *PullRequestsOptions) (interface{}, error) {
	urlStr := GetApiBaseURL() + "/repositories/" + po.Owner + "/" + po.RepoSlug + "/pullrequests/" + po.ID + "/activity"
	return p.c.execute("GET", urlStr, "")
}

func (p *PullRequests) Commits(po *PullRequestsOptions) (interface{}, error) {
	urlStr := GetApiBaseURL() + "/repositories/" + po.Owner + "/" + po.RepoSlug + "/pullrequests/" + po.ID + "/commits"
	return p.c.execute("GET", urlStr, "")
}

func (p *PullRequests) Patch(po *PullRequestsOptions) (interface{}, error) {
	urlStr := GetApiBaseURL() + "/repositories/" + po.Owner + "/" + po.RepoSlug + "/pullrequests/" + po.ID + "/patch"
	return p.c.execute("GET", urlStr, "")
}

func (p *PullRequests) Diff(po *PullRequestsOptions) (interface{}, error) {
	urlStr := GetApiBaseURL() + "/repositories/" + po.Owner + "/" + po.RepoSlug + "/pullrequests/" + po.ID + "/diff"
	return p.c.execute("GET", urlStr, "")
}

func (p *PullRequests) Merge(po *PullRequestsOptions) (interface{}, error) {
	data := p.buildPullRequestBody(po)
	urlStr := GetApiBaseURL() + "/repositories/" + po.Owner + "/" + po.RepoSlug + "/pullrequests/" + po.ID + "/merge"
	return p.c.execute("POST", urlStr, data)
}

func (p *PullRequests) Decline(po *PullRequestsOptions) (interface{}, error) {
	data := p.buildPullRequestBody(po)
	urlStr := GetApiBaseURL() + "/repositories/" + po.Owner + "/" + po.RepoSlug + "/pullrequests/" + po.ID + "/decline"
	return p.c.execute("POST", urlStr, data)
}

func (p *PullRequests) GetComments(po *PullRequestsOptions) (interface{}, error) {
	urlStr := GetApiBaseURL() + "/repositories/" + po.Owner + "/" + po.RepoSlug + "/pullrequests/" + po.ID + "/comments/"
	return p.c.execute("GET", urlStr, "")
}

func (p *PullRequests) GetComment(po *PullRequestsOptions) (interface{}, error) {
	urlStr := GetApiBaseURL() + "/repositories/" + po.Owner + "/" + po.RepoSlug + "/pullrequests/" + po.ID + "/comments/" + po.CommentID
	return p.c.execute("GET", urlStr, "")
}

func (p *PullRequests) AddComment(po *PullRequestsOptions) (interface{}, error) {
	data := p.buildPullRequestComment(po)
	urlStr := GetApiBaseURL() + "/repositories/" + po.Owner + "/" + po.RepoSlug + "/pullrequests/" + po.ID + "/comments/"
	return p.c.execute("POST", urlStr, data)
}

func (p *PullRequests) buildPullRequestComment(po *PullRequestsOptions) string {
	rawMessasge := strings.Replace(po.Message, `"`, `\"`, -1)
	rawMessasge = strings.Replace(rawMessasge, `_`, `\\_`, -1)
	rawMessasge = strings.Replace(rawMessasge, `+`, `\\+`, -1)
	rawMessasge = strings.Replace(rawMessasge, "\n", `\n\n`, -1)
	return fmt.Sprintf("{\"content\":{\"raw\":\"%s\"}}", rawMessasge)
}

func (p *PullRequests) buildPullRequestBody(po *PullRequestsOptions) string {

	body := map[string]interface{}{}
	body["source"] = map[string]interface{}{}
	body["destination"] = map[string]interface{}{}
	body["reviewers"] = []map[string]string{}
	body["title"] = ""
	body["description"] = ""
	body["message"] = ""
	body["close_source_branch"] = false

	if n := len(po.Reviewers); n > 0 {
		body["reviewers"] = make([]map[string]string, n)
		for i, user := range po.Reviewers {
			body["reviewers"].([]map[string]string)[i] = map[string]string{"username": user}
		}
	}

	if po.SourceBranch != "" {
		body["source"].(map[string]interface{})["branch"] = map[string]string{"name": po.SourceBranch}
	}

	if po.SourceRepository != "" {
		body["source"].(map[string]interface{})["repository"] = map[string]interface{}{"full_name": po.SourceRepository}
	}

	if po.DestinationBranch != "" {
		body["destination"].(map[string]interface{})["branch"] = map[string]interface{}{"name": po.DestinationBranch}
	}

	if po.DestinationCommit != "" {
		body["destination"].(map[string]interface{})["commit"] = map[string]interface{}{"hash": po.DestinationCommit}
	}

	if po.Title != "" {
		body["title"] = po.Title
	}

	if po.Description != "" {
		body["description"] = po.Description
	}

	if po.Message != "" {
		body["message"] = po.Message
	}

	if po.CloseSourceBranch == true || po.CloseSourceBranch == false {
		body["close_source_branch"] = po.CloseSourceBranch
	}

	data, err := json.Marshal(body)
	if err != nil {
		pp.Println(err)
		os.Exit(9)
	}

	return string(data)
}
