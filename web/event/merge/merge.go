package merge

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    pushRequest, err := UnmarshalPushRequest(bytes)
//    bytes, err = pushRequest.Marshal()

type MergeRequest struct {
	ObjectKind       string           `json:"object_kind"`
	EventType        string           `json:"event_type"`
	User             User             `json:"user"`
	Project          Project          `json:"project"`
	ObjectAttributes ObjectAttributes `json:"object_attributes"`
	Labels           []interface{}    `json:"labels"`
	Changes          Changes          `json:"changes"`
	Repository       Repository       `json:"repository"`
	Assignees        []User           `json:"assignees"`
}

type User struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	AvatarURL string `json:"avatar_url"`
	Email     string `json:"email"`
}

type Changes struct {
	MergeStatus MergeStatus `json:"merge_status"`
}

type MergeStatus struct {
	Previous string `json:"previous"`
	Current  string `json:"current"`
}

type ObjectAttributes struct {
	AssigneeID                  int64         `json:"assignee_id"`
	AuthorID                    int64         `json:"author_id"`
	CreatedAt                   string        `json:"created_at"`
	Description                 string        `json:"description"`
	HeadPipelineID              interface{}   `json:"head_pipeline_id"`
	ID                          int64         `json:"id"`
	Iid                         int64         `json:"iid"`
	LastEditedAt                interface{}   `json:"last_edited_at"`
	LastEditedByID              interface{}   `json:"last_edited_by_id"`
	MergeCommitSHA              interface{}   `json:"merge_commit_sha"`
	MergeError                  interface{}   `json:"merge_error"`
	MergeParams                 MergeParams   `json:"merge_params"`
	MergeStatus                 string        `json:"merge_status"`
	MergeUserID                 interface{}   `json:"merge_user_id"`
	MergeWhenPipelineSucceeds   bool          `json:"merge_when_pipeline_succeeds"`
	MilestoneID                 interface{}   `json:"milestone_id"`
	SourceBranch                string        `json:"source_branch"`
	SourceProjectID             int64         `json:"source_project_id"`
	StateID                     int64         `json:"state_id"`
	TargetBranch                string        `json:"target_branch"`
	TargetProjectID             int64         `json:"target_project_id"`
	TimeEstimate                int64         `json:"time_estimate"`
	Title                       string        `json:"title"`
	UpdatedAt                   string        `json:"updated_at"`
	UpdatedByID                 interface{}   `json:"updated_by_id"`
	URL                         string        `json:"url"`
	Source                      Project       `json:"source"`
	Target                      Project       `json:"target"`
	LastCommit                  LastCommit    `json:"last_commit"`
	WorkInProgress              bool          `json:"work_in_progress"`
	TotalTimeSpent              int64         `json:"total_time_spent"`
	TimeChange                  int64         `json:"time_change"`
	HumanTotalTimeSpent         interface{}   `json:"human_total_time_spent"`
	HumanTimeChange             interface{}   `json:"human_time_change"`
	HumanTimeEstimate           interface{}   `json:"human_time_estimate"`
	AssigneeIDS                 []int64       `json:"assignee_ids"`
	Labels                      []interface{} `json:"labels"`
	State                       string        `json:"state"`
	BlockingDiscussionsResolved bool          `json:"blocking_discussions_resolved"`
	FirstContribution           bool          `json:"first_contribution"`
	Action                      string        `json:"action"`
}

type LastCommit struct {
	ID        string `json:"id"`
	Message   string `json:"message"`
	Title     string `json:"title"`
	Timestamp string `json:"timestamp"`
	URL       string `json:"url"`
	Author    Author `json:"author"`
}

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type MergeParams struct {
	ForceRemoveSourceBranch string `json:"force_remove_source_branch"`
}

type Project struct {
	ID                int64       `json:"id"`
	Name              string      `json:"name"`
	Description       string      `json:"description"`
	WebURL            string      `json:"web_url"`
	AvatarURL         interface{} `json:"avatar_url"`
	GitSSHURL         string      `json:"git_ssh_url"`
	GitHTTPURL        string      `json:"git_http_url"`
	Namespace         string      `json:"namespace"`
	VisibilityLevel   int64       `json:"visibility_level"`
	PathWithNamespace string      `json:"path_with_namespace"`
	DefaultBranch     string      `json:"default_branch"`
	CiConfigPath      string      `json:"ci_config_path"`
	Homepage          string      `json:"homepage"`
	URL               string      `json:"url"`
	SSHURL            string      `json:"ssh_url"`
	HTTPURL           string      `json:"http_url"`
}

type Repository struct {
	Name        string `json:"name"`
	URL         string `json:"url"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
}
