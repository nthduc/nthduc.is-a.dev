package libs

import (
	"html/template"
)

type GitHubData struct {
	User     User
	Projects []Project
}

type User struct {
	Name      string
	AvatarURL template.URL
}
