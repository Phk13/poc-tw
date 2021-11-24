package models

/* LoginResponse contains the token return during login*/
type LoginResponse struct {
	Token string `json:"token,omitempty"`
}
