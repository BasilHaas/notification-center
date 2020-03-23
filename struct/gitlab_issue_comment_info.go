package gitlab_info

import "time"

type IssueComment struct {
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
	Issue struct {
		ID          int           `json:"id"`
		Title       string        `json:"title"`
		AssigneeIds []interface{} `json:"assignee_ids"`
		AssigneeID  interface{}   `json:"assignee_id"`
		AuthorID    int           `json:"author_id"`
		ProjectID   int           `json:"project_id"`
		CreatedAt   string        `json:"created_at"`
		UpdatedAt   string        `json:"updated_at"`
		Position    int           `json:"position"`
		BranchName  interface{}   `json:"branch_name"`
		Description string        `json:"description"`
		MilestoneID interface{}   `json:"milestone_id"`
		State       string        `json:"state"`
		Iid         int           `json:"iid"`
		Labels      []struct {
			ID          int         `json:"id"`
			Title       string      `json:"title"`
			Color       string      `json:"color"`
			ProjectID   interface{} `json:"project_id"`
			CreatedAt   time.Time   `json:"created_at"`
			UpdatedAt   time.Time   `json:"updated_at"`
			Template    bool        `json:"template"`
			Description interface{} `json:"description"`
			Type        string      `json:"type"`
			GroupID     int         `json:"group_id"`
		} `json:"labels"`
	} `json:"issue"`
}
