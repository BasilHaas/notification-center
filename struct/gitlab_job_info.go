package gitlab_info

type Job struct {
	ObjectKind         string      `json:"object_kind"`
	Ref                string      `json:"ref"`
	Tag                bool        `json:"tag"`
	BeforeSha          string      `json:"before_sha"`
	Sha                string      `json:"sha"`
	BuildID            string      `json:"build_id"`
	BuildName          string      `json:"build_name"`
	BuildStage         string      `json:"build_stage"`
	BuildStatus        string      `json:"build_status"`
	BuildStartedAt     interface{} `json:"build_started_at"`
	BuildFinishedAt    interface{} `json:"build_finished_at"`
	BuildDuration      interface{} `json:"build_duration"`
	BuildAllowFailure  bool        `json:"build_allow_failure"`
	BuildFailureReason string      `json:"build_failure_reason"`
	PipelineID         int         `json:"pipeline_id"`
	ProjectID          int         `json:"project_id"`
	ProjectName        string      `json:"project_name"`
	User               struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Email     string `json:"email"`
		AvatarURL string `json:"avatar_url"`
	} `json:"user"`
	Commit struct {
		ID          int         `json:"id"`
		Sha         string      `json:"sha"`
		Message     string      `json:"message"`
		AuthorName  string      `json:"author_name"`
		AuthorEmail string      `json:"author_email"`
		Status      string      `json:"status"`
		Duration    interface{} `json:"duration"`
		StartedAt   interface{} `json:"started_at"`
		FinishedAt  interface{} `json:"finished_at"`
	} `json:"commit"`
	Repository struct {
		Name            string `json:"name"`
		Description     string `json:"description"`
		Homepage        string `json:"homepage"`
		GitSSHURL       string `json:"git_ssh_url"`
		GitHTTPURL      string `json:"git_http_url"`
		VisibilityLevel int    `json:"visibility_level"`
	} `json:"repository"`
	Runner struct {
		Active      bool   `json:"active"`
		IsShared    bool   `json:"is_shared"`
		ID          int    `json:"id"`
		Description string `json:"description"`
	} `json:"runner"`
}
