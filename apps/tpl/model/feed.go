package model

// Feed represents an item in the user's feed or timeline.
//
// swagger:model feed
type Feed struct {
	Owner    string `json:"owner"         db:"repo_owner"`
	Name     string `json:"name"          db:"repo_name"`
	FullName string `json:"full_name"     db:"repo_full_name"`

	Number   int    `json:"number,omitempty"        db:"build_number"`
	Event    string `json:"event,omitempty"         db:"build_event"`
	Status   string `json:"status,omitempty"        db:"build_status"`
	Created  int64  `json:"created_at,omitempty"    db:"build_created"`
	Started  int64  `json:"started_at,omitempty"    db:"build_started"`
	Finished int64  `json:"finished_at,omitempty"   db:"build_finished"`
	Commit   string `json:"commit,omitempty"        db:"build_commit"`
	Branch   string `json:"branch,omitempty"        db:"build_branch"`
	Ref      string `json:"ref,omitempty"           db:"build_ref"`
	Refspec  string `json:"refspec,omitempty"       db:"build_refspec"`
	Remote   string `json:"remote,omitempty"        db:"build_remote"`
	Title    string `json:"title,omitempty"         db:"build_title"`
	Message  string `json:"message,omitempty"       db:"build_message"`
	Author   string `json:"author,omitempty"        db:"build_author"`
	Avatar   string `json:"author_avatar,omitempty" db:"build_avatar"`
	Email    string `json:"author_email,omitempty"  db:"build_email"`
}
