/**
Go standard librabry provides tools for outputing logs
	1. log package for free-form output
	2. log/slog package for structured output(you log key-value pairs like json logs)
*/

package main

import (
	"log"
	"log/slog"
	"os"
)

func main() {
	log.Println("This is a normal log message.")
	log.Printf("This is a formatted log: %d + %d = %d", 2, 3, 2+3)
	//log.Fatal("This stops the program!") // Logs and then exits with code

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil)) // Set up a JSON logger writing to stdout

	logger.Info("User logged in", "user", "alice", "role", "admin")
	logger.Error("Something went wrong", "error", "db timeout")
}
