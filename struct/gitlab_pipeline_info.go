package gitlab_info

import "time"

type Pipeline struct {
	ObjectKind       string `json:"object_kind"`
	ObjectAttributes struct {
		ID         int      `json:"id"`
		Ref        string   `json:"ref"`
		Tag        bool     `json:"tag"`
		Sha        string   `json:"sha"`
		BeforeSha  string   `json:"before_sha"`
		Source     string   `json:"source"`
		Status     string   `json:"status"`
		Stages     []string `json:"stages"`
		CreatedAt  string   `json:"created_at"`
		FinishedAt string   `json:"finished_at"`
		Duration   int      `json:"duration"`
		Variables  []struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"variables"`
	} `json:"object_attributes"`
	MergeRequest struct {
		ID              int    `json:"id"`
		Iid             int    `json:"iid"`
		Title           string `json:"title"`
		SourceBranch    string `json:"source_branch"`
		SourceProjectID int    `json:"source_project_id"`
		TargetBranch    string `json:"target_branch"`
		TargetProjectID int    `json:"target_project_id"`
		State           string `json:"state"`
		MergeStatus     string `json:"merge_status"`
		URL             string `json:"url"`
	} `json:"merge_request"`
	User struct {
		Name      string `json:"name"`
		Username  string `json:"username"`
		AvatarURL string `json:"avatar_url"`
		Email     string `json:"email"`
	} `json:"user"`
	Project struct {
		ID                int         `json:"id"`
		Name              string      `json:"name"`
		Description       string      `json:"description"`
		WebURL            string      `json:"web_url"`
		AvatarURL         interface{} `json:"avatar_url"`
		GitSSHURL         string      `json:"git_ssh_url"`
		GitHTTPURL        string      `json:"git_http_url"`
		Namespace         string      `json:"namespace"`
		VisibilityLevel   int         `json:"visibility_level"`
		PathWithNamespace string      `json:"path_with_namespace"`
		DefaultBranch     string      `json:"default_branch"`
	} `json:"project"`
	Commit struct {
		ID        string    `json:"id"`
		Message   string    `json:"message"`
		Timestamp time.Time `json:"timestamp"`
		URL       string    `json:"url"`
		Author    struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		} `json:"author"`
	} `json:"commit"`
	Builds []struct {
		ID           int         `json:"id"`
		Stage        string      `json:"stage"`
		Name         string      `json:"name"`
		Status       string      `json:"status"`
		CreatedAt    string      `json:"created_at"`
		StartedAt    interface{} `json:"started_at"`
		FinishedAt   interface{} `json:"finished_at"`
		When         string      `json:"when"`
		Manual       bool        `json:"manual"`
		AllowFailure bool        `json:"allow_failure"`
		User         struct {
			Name      string `json:"name"`
			Username  string `json:"username"`
			AvatarURL string `json:"avatar_url"`
		} `json:"user"`
		Runner        interface{} `json:"runner"`
		ArtifactsFile struct {
			Filename interface{} `json:"filename"`
			Size     interface{} `json:"size"`
		} `json:"artifacts_file"`
	} `json:"builds"`
}
