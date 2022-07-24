package gitlab_api

import (
	"encoding/json"
	"errors"
	"fmt"
	"gitrabbit/conf"
	"io/ioutil"
	"net/http"
)

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    merge, err := UnmarshalMerge(bytes)
//    bytes, err = merge.Marshal()

//tag and release handles deliver parameters to GitLab as individual field
type GitLab struct {
	ProjectID int64
}

type MergeResponse struct {
	ID                          int64                `json:"id"`
	Iid                         int64                `json:"iid"`
	ProjectID                   int64                `json:"project_id"`
	Title                       string               `json:"title"`
	Description                 string               `json:"description"`
	State                       string               `json:"state"`
	CreatedAt                   string               `json:"created_at"`
	UpdatedAt                   string               `json:"updated_at"`
	MergedBy                    interface{}          `json:"merged_by"`
	MergeUser                   interface{}          `json:"merge_user"`
	MergedAt                    interface{}          `json:"merged_at"`
	ClosedBy                    *Assignee            `json:"closed_by"`
	ClosedAt                    *string              `json:"closed_at"`
	TargetBranch                string               `json:"target_branch"`
	SourceBranch                string               `json:"source_branch"`
	UserNotesCount              int64                `json:"user_notes_count"`
	Upvotes                     int64                `json:"upvotes"`
	Downvotes                   int64                `json:"downvotes"`
	Author                      Assignee             `json:"author"`
	Assignees                   []Assignee           `json:"assignees"`
	Assignee                    Assignee             `json:"assignee"`
	Reviewers                   []Assignee           `json:"reviewers"`
	SourceProjectID             int64                `json:"source_project_id"`
	TargetProjectID             int64                `json:"target_project_id"`
	Labels                      []interface{}        `json:"labels"`
	Draft                       bool                 `json:"draft"`
	WorkInProgress              bool                 `json:"work_in_progress"`
	Milestone                   interface{}          `json:"milestone"`
	MergeWhenPipelineSucceeds   bool                 `json:"merge_when_pipeline_succeeds"`
	MergeStatus                 string               `json:"merge_status"`
	SHA                         string               `json:"sha"`
	MergeCommitSHA              interface{}          `json:"merge_commit_sha"`
	SquashCommitSHA             interface{}          `json:"squash_commit_sha"`
	DiscussionLocked            interface{}          `json:"discussion_locked"`
	ShouldRemoveSourceBranch    interface{}          `json:"should_remove_source_branch"`
	ForceRemoveSourceBranch     bool                 `json:"force_remove_source_branch"`
	Reference                   string               `json:"reference"`
	References                  References           `json:"references"`
	WebURL                      string               `json:"web_url"`
	TimeStats                   TimeStats            `json:"time_stats"`
	Squash                      bool                 `json:"squash"`
	TaskCompletionStatus        TaskCompletionStatus `json:"task_completion_status"`
	HasConflicts                bool                 `json:"has_conflicts"`
	BlockingDiscussionsResolved bool                 `json:"blocking_discussions_resolved"`
	ApprovalsBeforeMerge        interface{}          `json:"approvals_before_merge"`
}

type Assignee struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Name      string `json:"name"`
	State     string `json:"state"`
	AvatarURL string `json:"avatar_url"`
	WebURL    string `json:"web_url"`
}

type References struct {
	Short    string `json:"short"`
	Relative string `json:"relative"`
	Full     string `json:"full"`
}

type TaskCompletionStatus struct {
	Count          int64 `json:"count"`
	CompletedCount int64 `json:"completed_count"`
}

type TimeStats struct {
	TimeEstimate        int64       `json:"time_estimate"`
	TotalTimeSpent      int64       `json:"total_time_spent"`
	HumanTimeEstimate   interface{} `json:"human_time_estimate"`
	HumanTotalTimeSpent interface{} `json:"human_total_time_spent"`
}

var getMergeApi = "https://gitlab.com/api/v4/projects/%d/merge_requests?state=merged&per_page=1"

//获取最后一个MR的审批人列表
//obtains assignee and other info list in last Merge Request
func (gitlab *GitLab) GetNotifyUser() []string {
	notifyUser := make([]string, 0)

	commit, err := gitlab.GetLastCommit()
	//send missing info without obtaining notifier from gitlab
	//panic is missing with this method
	if err != nil {
		return notifyUser
	}
	//get assignees at last merge request commit
	for _, user := range (*commit).Assignees {
		notifyUser = append(notifyUser, user.Name)
	}
	//get reviewers at last merge request commit
	for _, user := range (*commit).Reviewers {
		notifyUser = append(notifyUser, user.Name)
	}

	return notifyUser
}

//Obatin last commit info in gitlab api
func (gitlab *GitLab) GetLastCommit() (*MergeResponse, error) {

	client := &http.Client{}
	//ProjectID needs to grap and gen the correct url
	url := fmt.Sprintf(getMergeApi, gitlab.ProjectID)
	//主动去GITLAB 开放的HTTP API获取内容
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("PRIVATE-TOKEN", conf.Config.Gitlab.AccessToken)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var m []MergeResponse
	err = json.Unmarshal(body, &m)
	if err != nil {
		return nil, err
	}

	if len(m) < 1 {
		return nil, errors.New("Merge Request List is empty, so could not obtain valid info")
	}

	return &m[0], nil
}
