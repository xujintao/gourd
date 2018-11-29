package model

import (
	"errors"
	"regexp"
)

// validate a username (e.g. from github)
var reUsername = regexp.MustCompile("^[a-zA-Z0-9-_.]+$")

var errUserLoginInvalid = errors.New("Invalid User Login")

// User represents a registered user.
//
// swagger:model user
type User struct {
	// the id for this user.
	ID int64 `json:"id" db:"user_id"`

	// the username for this user.
	Name string `json:"name"  db:"user_name"`

	// Token is the oauth2 token.
	Token string `json:"-"  db:"user_token"`

	// Secret is the oauth2 token secret.
	Secret string `json:"-" db:"user_secret"`

	// Expiry is the token and secret expiration timestamp.
	Expiry int64 `json:"-" db:"user_expiry"`

	// Email is the email address for this user.
	Email string `json:"email" db:"user_email"`

	// the avatar url for this user.
	Avatar string `json:"avatar_url" db:"user_avatar"`

	// Activate indicates the user is active in the system.
	Active bool `json:"active" db:"user_active"`

	// Synced is the timestamp when the user was synced with the remote system.
	Synced int64 `json:"synced" db:"user_synced"`

	// Admin indicates the user is a system administrator.
	//
	// NOTE: This is sourced from the DRONE_ADMINS environment variable and is no
	// longer persisted in the database.
	Admin bool `json:"admin,omitempty" db:"-"`

	// Hash is a unique token used to sign tokens.
	Hash string `json:"-" db:"user_hash"`

	// DEPRECATED Admin indicates the user is a system administrator.
	XAdmin bool `json:"-" db:"user_admin"`
}

// Validate validates the required fields and formats.
func (u *User) Validate() error {
	switch {
	case len(u.Name) == 0:
		return errUserLoginInvalid
	case len(u.Name) > 250:
		return errUserLoginInvalid
	case !reUsername.MatchString(u.Name):
		return errUserLoginInvalid
	default:
		return nil
	}
}
