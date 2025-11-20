package main

import (
	"fmt"
	"os/exec"
	"runtime"

	"notifier/internal/infrastructure/api"

	_ "notifier/docs" // swagger docs
)

// @title Notifier API
// @version 1.0
// @description This is a high performance notification system API
// @host localhost:8080
// @BasePath /
func main() {
	// Create a new Gin router
	router := api.NewRouter()

	// Start server in a goroutine
	go func() {
		if err := router.Run(":8080"); err != nil {
			panic(err)
		}
	}()

	// Auto-open browser
	openBrowser("http://localhost:8080/swagger/index.html")

	// Keep main goroutine alive
	select {}
}

// openBrowser opens the default browser to the specified URL
func openBrowser(url string) {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)

	if err := exec.Command(cmd, args...).Start(); err != nil {
		fmt.Printf("Failed to open browser: %v\n", err)
		fmt.Printf("Please open manually: %s\n", url)
	}
}
