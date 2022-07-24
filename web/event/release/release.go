// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    release, err := UnmarshalRelease(bytes)
//    bytes, err = release.Marshal()

package release

type ReleaseRequest struct {
	ID          int64          `json:"id"`
	CreatedAt   string         `json:"created_at"`
	Description string         `json:"description"`
	Name        string         `json:"name"`
	ReleasedAt  string         `json:"released_at"`
	Tag         string         `json:"tag"`
	ObjectKind  string         `json:"object_kind"`
	Project     ReleaseProject `json:"project"`
	URL         string         `json:"url"`
	Action      string         `json:"action"`
	Assets      Assets         `json:"assets"`
	Commit      Commit         `json:"commit"`
}

type Assets struct {
	Count   int64         `json:"count"`
	Links   []interface{} `json:"links"`
	Sources []Source      `json:"sources"`
}

type Source struct {
	Format string `json:"format"`
	URL    string `json:"url"`
}

type Commit struct {
	ID        string        `json:"id"`
	Message   string        `json:"message"`
	Title     string        `json:"title"`
	Timestamp string        `json:"timestamp"`
	URL       string        `json:"url"`
	Author    ReleaseAuthor `json:"author"`
}

type ReleaseAuthor struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ReleaseProject struct {
	ID                int64       `json:"id"`
	Name              string      `json:"name"`
	Description       interface{} `json:"description"`
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
