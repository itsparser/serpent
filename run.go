package serpent

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/workfoxes/serpent/src/exception"
)

func getDefaultCacheDirectory() (string, error) {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return "", exception.UnableToLocateRootDirectory() // fmt.Errorf("could not get user home directory: %v", err)
	}
	switch runtime.GOOS {
	case "windows":
		return filepath.Join(userHomeDir, "AppData", "Local"), nil
	case "darwin":
		return filepath.Join(userHomeDir, "Library", "Caches"), nil
	case "linux":
		return filepath.Join(userHomeDir, ".cache"), nil
	}
	return "", exception.UnableToLocateDriver()
}

func getDriverName() string {
	switch runtime.GOOS {
	case "windows":
		return "serpent-cli.exe"
	case "darwin":
		return "serpent-cli.sh"
	case "linux":
		return "serpent-cli.sh"
	}
	panic("Not supported OS!")
}

// RunOption is a custom option to run the Serpant
// This will be used when serpent is executed as lib
type RunOption struct {
	DriverDirectory string
	ChromeVersion   string
}

type serpentDriver struct {
	driverDirectory,
	driverBinaryLocation,
	version string
}

func getNewDriver(options *RunOption) (*serpentDriver, error) {
	baseDriverDirectory := ""
	if options != nil {
		baseDriverDirectory = options.DriverDirectory
	}
	if baseDriverDirectory == "" {
		var err error
		baseDriverDirectory, err = getDefaultCacheDirectory()
		if err != nil {
			return nil, fmt.Errorf("could not get default cache directory: %v", err)
		}
	}
	driverDirectory := filepath.Join(options.DriverDirectory, "serpent-drive", Version)
	driverBinaryLocation := filepath.Join(options.DriverDirectory, getDriverName())
	return &serpentDriver{
		driverBinaryLocation: driverBinaryLocation,
		driverDirectory:      driverDirectory,
		version:              Version,
	}, nil
}

// Init function will get the Serpent connection for starting the automation testing
func Init(options *RunOption) {
	// driver, err := getNewDriver(options)

}

// Setup will setup the enviroment with all the dependent driver for execution in in serpant
func Setup() {

}
