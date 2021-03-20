package interfaces

// A BrowserType is created when Serpent connects to a browser instance, either through browserType.launch([options]) or
// browserType.connect(params).
// An example of using a Browser to create a Page:
// See ChromiumBrowser, FirefoxBrowser and WebKitBrowser for browser-specific features. Note that
// browserType.connect(params) and browserType.launch([options]) always return a specific browser instance, based on the browser
// being connected to or launched.
type BrowserType interface {
	// In case this browser is obtained using browserType.launch([options]), closes the browser and all of its pages (if any were
	// opened).
	// In case this browser is obtained using browserType.connect(params), clears all created contexts belonging to this browser
	// and disconnects from the browser server.
	// The Browser object itself is considered to be disposed and cannot be used anymore.
	Close() error
	// Returns an array of all open browser contexts. In a newly created browser, this will return zero browser contexts.
	// Returns the browser version.
	Version() string
}
