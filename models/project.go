package models

type Project struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	ClientID string `json:"client_id"`
	Status   string `json:"status"`
}
