package serpent

import "github.com/workfoxes/serpent/src/engine/interfaces"

// Commit will have the commit if for the head
// date Date of the release date of current version
// Version will have the current version of the application
var (
	Commit  string
	Date    string
	Version string
)

// Navigator will have all the User Device info on which test case need to be executed
type Navigator struct {
	UserAgent          string `json:"userAgent"`
	DeviceScaleFactor  int    `json:"deviceScaleFactor"`
	IsMobile           bool   `json:"isMobile"`
	HasTouch           bool   `json:"hasTouch"`
	DefaultBrowserType string `json:"defaultBrowserType"`
}

// Serpent will have respective Serpent instance
type Serpent struct {
	Chromium interfaces.BrowserType
	Firefox  interfaces.BrowserType
	WebKit   interfaces.BrowserType
	Devices  map[string]*Navigator
}