package models

type Workspace struct {
	Context  string            `json:"context"`
	Commands map[string]string `json:"commands"`
}
