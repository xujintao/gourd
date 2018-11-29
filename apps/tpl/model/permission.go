package model

// Permission defines a repository permission for an individual user.
type Permission struct {
	UserID int64  `json:"-"      db:"perm_user_id"`
	RepoID int64  `json:"-"      db:"perm_repo_id"`
	Repo   string `json:"-"      db:"-"`
	Pull   bool   `json:"pull"   db:"perm_pull"`
	Push   bool   `json:"push"   db:"perm_push"`
	Admin  bool   `json:"admin"  db:"perm_admin"`
	Synced int64  `json:"synced" db:"perm_synced"`
}
