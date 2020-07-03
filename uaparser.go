package uaparser

import "github.com/mssola/user_agent"

type UserAgent struct {
	Mobile       bool   `json:"mobile,omitempty"`
	Bot          bool   `json:"bot,omitempty"`
	Mozilla      string `json:"mozilla,omitempty"`
	Platform     string `json:"platform,omitempty"`
	Localization string `json:"localization,omitempty"`

	OS struct {
		Name     string `json:"name,omitempty"`
		FullName string `json:"full_name,omitempty"`
		Version  string `json:"version,omitempty"`
	} `json:"os,omitempty"`
	Engine struct {
		Name    string `json:"name,omitempty"`
		Version string `json:"version,omitempty"`
	} `json:"engine,omitempty"`
	Browser struct {
		Name    string `json:"name,omitempty"`
		Version string `json:"version,omitempty"`
	} `json:"browser,omitempty"`
}

func Parse(uaString string) UserAgent {
	ua := user_agent.New(uaString)
	res := UserAgent{
		Mobile:  ua.Mobile(),
		Bot:     ua.Bot(),
		Mozilla: ua.Mozilla(),
	}
	res.Browser.Name, res.Browser.Version = ua.Browser()
	res.Engine.Name, res.Engine.Version = ua.Engine()
	osInfo := ua.OSInfo()
	res.OS.Name = osInfo.Name
	res.OS.FullName = osInfo.FullName
	res.OS.Version = osInfo.Version
	return res
}
