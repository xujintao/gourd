package dao

import "github.com/xujintao/gourd/apps/tpl/model"

// Cache interface
type Cache interface {
}

// DB interface
type DB interface {
	// GetUser gets a user by unique ID.
	GetUserByID(int64) (*model.User, error)

	// GetUserLogin gets a user by unique Login name.
	GetUserByName(string) (*model.User, error)

	// GetUserList gets a list of all users in the system.
	GetUserList() ([]*model.User, error)

	// GetUserCount gets a count of all users in the system.
	GetUserCount() (int, error)

	// CreateUser creates a new user account.
	CreateUser(*model.User) error

	// UpdateUser updates a user account.
	UpdateUser(*model.User) error

	// DeleteUser deletes a user account.
	DeleteUser(*model.User) error

	GetFeedList(*model.User) ([]*model.Feed, error)
	GetFeedListLatest(*model.User) ([]*model.Feed, error)

	GetRepoList(*model.User) ([]*model.Repo, error)
}
