package event

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

//align all formats to send to outputs
type MessageNotify struct {
	MsgType  string   `json:"msg_type"`
	Notifier []string `json:"notifier"`
	Content  string   `json:"content"`
}
