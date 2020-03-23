package gitlab_info

type MergeComment struct {
	ObjectKind string `json:"object_kind"`
	User       struct {
		Name      string `json:"name"`
		Username  string `json:"username"`
		AvatarURL string `json:"avatar_url"`
	} `json:"user"`
	ProjectID int `json:"project_id"`
	Project   struct {
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
		Homepage          string      `json:"homepage"`
		URL               string      `json:"url"`
		SSHURL            string      `json:"ssh_url"`
		HTTPURL           string      `json:"http_url"`
	} `json:"project"`
	Repository struct {
		Name        string `json:"name"`
		URL         string `json:"url"`
		Description string `json:"description"`
		Homepage    string `json:"homepage"`
	} `json:"repository"`
	ObjectAttributes struct {
		ID           int         `json:"id"`
		Note         string      `json:"note"`
		NoteableType string      `json:"noteable_type"`
		AuthorID     int         `json:"author_id"`
		CreatedAt    string      `json:"created_at"`
		UpdatedAt    string      `json:"updated_at"`
		ProjectID    int         `json:"project_id"`
		Attachment   interface{} `json:"attachment"`
		LineCode     interface{} `json:"line_code"`
		CommitID     string      `json:"commit_id"`
		NoteableID   int         `json:"noteable_id"`
		System       bool        `json:"system"`
		StDiff       interface{} `json:"st_diff"`
		URL          string      `json:"url"`
	} `json:"object_attributes"`
	MergeRequest struct {
		ID              int    `json:"id"`
		TargetBranch    string `json:"target_branch"`
		SourceBranch    string `json:"source_branch"`
		SourceProjectID int    `json:"source_project_id"`
		AuthorID        int    `json:"author_id"`
		AssigneeID      int    `json:"assignee_id"`
		Title           string `json:"title"`
		CreatedAt       string `json:"created_at"`
		UpdatedAt       string `json:"updated_at"`
		MilestoneID     int    `json:"milestone_id"`
		State           string `json:"state"`
		MergeStatus     string `json:"merge_status"`
		TargetProjectID int    `json:"target_project_id"`
		Iid             int    `json:"iid"`
		Description     string `json:"description"`
		Position        int    `json:"position"`
		Source          struct {
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
			Homepage          string      `json:"homepage"`
			URL               string      `json:"url"`
			SSHURL            string      `json:"ssh_url"`
			HTTPURL           string      `json:"http_url"`
		} `json:"source"`
		Target struct {
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
			Homepage          string      `json:"homepage"`
			URL               string      `json:"url"`
			SSHURL            string      `json:"ssh_url"`
			HTTPURL           string      `json:"http_url"`
		} `json:"target"`
		LastCommit struct {
			ID        string `json:"id"`
			Message   string `json:"message"`
			Timestamp string `json:"timestamp"`
			URL       string `json:"url"`
			Author    struct {
				Name  string `json:"name"`
				Email string `json:"email"`
			} `json:"author"`
		} `json:"last_commit"`
		WorkInProgress bool `json:"work_in_progress"`
		Assignee       struct {
			Name      string `json:"name"`
			Username  string `json:"username"`
			AvatarURL string `json:"avatar_url"`
		} `json:"assignee"`
	} `json:"merge_request"`
}
