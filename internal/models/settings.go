package models

type Settings struct {
	Workspaces map[string]Workspace `json:"workspaces"`
}
