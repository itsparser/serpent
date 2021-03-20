package serpent

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"testing"
)

func TestSetup(t *testing.T) {
	// tests := []struct {
	// 	name string
	// }{
	// 	// TODO: Add test cases.
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		Setup()
	// 	})
	// }
	var path = ""
	switch runtime.GOOS {
	case "windows":
		path = ""
	case "darwin":
		path = `/Applications/Google\ Chrome.app/Contents/MacOS/Google\ Chrome`
	case "linux":
		path = "/usr/bin/google-chrome"
	}
	cmd, err := exec.Command(path, "--version").Output()
	if err != nil {
		log.Fatal(err)
		fmt.Printf("The date is %s\n", err)
	}
	log.Fatal(cmd)
	fmt.Printf("The date is %s\n", cmd)
}
