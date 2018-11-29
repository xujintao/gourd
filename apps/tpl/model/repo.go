package model

// Repo represents a repository.
//
// swagger:model repo
type Repo struct {
	ID          int64       `json:"id,omitempty"             db:"repo_id"`
	UserID      int64       `json:"-"                        db:"repo_user_id"`
	Owner       string      `json:"owner"                    db:"repo_owner"`
	Name        string      `json:"name"                     db:"repo_name"`
	FullName    string      `json:"full_name"                db:"repo_full_name"`
	Avatar      string      `json:"avatar_url,omitempty"     db:"repo_avatar"`
	Link        string      `json:"link_url,omitempty"       db:"repo_link"`
	Kind        string      `json:"scm,omitempty"            db:"repo_scm"`
	Clone       string      `json:"clone_url,omitempty"      db:"repo_clone"`
	Branch      string      `json:"default_branch,omitempty" db:"repo_branch"`
	Timeout     int64       `json:"timeout,omitempty"        db:"repo_timeout"`
	Visibility  string      `json:"visibility"               db:"repo_visibility"`
	IsPrivate   bool        `json:"private"                  db:"repo_private"`
	IsTrusted   bool        `json:"trusted"                  db:"repo_trusted"`
	IsStarred   bool        `json:"starred,omitempty"        db:"-"`
	IsGated     bool        `json:"gated"                    db:"repo_gated"`
	IsActive    bool        `json:"active"                   db:"repo_active"`
	AllowPull   bool        `json:"allow_pr"                 db:"repo_allow_pr"`
	AllowPush   bool        `json:"allow_push"               db:"repo_allow_push"`
	AllowDeploy bool        `json:"allow_deploys"            db:"repo_allow_deploys"`
	AllowTag    bool        `json:"allow_tags"               db:"repo_allow_tags"`
	Counter     int         `json:"last_build"               db:"repo_counter"`
	Config      string      `json:"config_file"              db:"repo_config_path"`
	Hash        string      `json:"-"                        db:"repo_hash"`
	Permission  *Permission `json:"-"                        db:"-"`
}

// func (r *Repo) ResetVisibility() {
// 	r.Visibility = VisibilityPublic
// 	if r.IsPrivate {
// 		r.Visibility = VisibilityPrivate
// 	}
// }
