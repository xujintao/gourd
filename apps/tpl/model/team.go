package model

// Team represents a team or organization in the remote version control system.
//
// swagger:model user
type Team struct {
	// Name is the username for this team.
	Name string `json:"name"`

	// the avatar url for this team.
	Avatar string `json:"avatar_url"`
}
