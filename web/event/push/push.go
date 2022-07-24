package push

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    pushRequest, err := UnmarshalPushRequest(bytes)
//    bytes, err = pushRequest.Marshal()

type PushRequest struct {
	ObjectKind        string      `json:"object_kind"`
	EventName         string      `json:"event_name"`
	Before            string      `json:"before"`
	After             string      `json:"after"`
	Ref               string      `json:"ref"`
	CheckoutSHA       string      `json:"checkout_sha"`
	Message           interface{} `json:"message"`
	UserID            int64       `json:"user_id"`
	UserName          string      `json:"user_name"`
	UserUsername      string      `json:"user_username"`
	UserEmail         interface{} `json:"user_email"`
	UserAvatar        string      `json:"user_avatar"`
	ProjectID         int64       `json:"project_id"`
	Project           Project     `json:"project"`
	Commits           []Commit    `json:"commits"`
	TotalCommitsCount int64       `json:"total_commits_count"`
	PushOptions       PushOptions `json:"push_options"`
	Repository        Repository  `json:"repository"`
}

type Commit struct {
	ID        string        `json:"id"`
	Message   string        `json:"message"`
	Title     string        `json:"title"`
	Timestamp string        `json:"timestamp"`
	URL       string        `json:"url"`
	Author    Author        `json:"author"`
	Added     []string      `json:"added"`
	Modified  []interface{} `json:"modified"`
	Removed   []interface{} `json:"removed"`
}

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
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

type PushOptions struct {
}

type Repository struct {
	Name            string `json:"name"`
	URL             string `json:"url"`
	Description     string `json:"description"`
	Homepage        string `json:"homepage"`
	GitHTTPURL      string `json:"git_http_url"`
	GitSSHURL       string `json:"git_ssh_url"`
	VisibilityLevel int64  `json:"visibility_level"`
}
