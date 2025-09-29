package main

import (
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"os"
)

func main() {
	err := KillServer("server.pid")
	if err != nil {
		fmt.Println("ERROR:", err)
		if errors.Is(err, fs.ErrNotExist) {
			fmt.Println("not found")
		}
		for e := err; e != nil; e = errors.Unwrap(e) {
			fmt.Printf("> %s\n", err)
		}
	}
}

// Tries to acquire a resource, check for error, defer release
func KillServer(pidFile string) error {
	file, err := os.Open(pidFile)
	if err != nil {
		return err
	}
	// Defer happens when the function exits, no matter what (panic)
	// Defer works at the function level
	// If multipler defers, they are executed in reverse order (LIFO)
	defer func() {
		if err := file.Close(); err != nil {
			slog.Warn("close", "file", pidFile, "error", err)
		}
	}()

	var pid int 
	if  _, err := fmt.Fscanf(file, "%d", &pid); err != nil {
		return fmt.Errorf("%q - bad pid: %w", pidFile, err)
	} 

	slog.Info("killing", "pid", pid)

	if err := os.Remove(pidFile); err != nil {
		// we are not failing, only warning
		slog.Warn("delete", "file", pidFile, "error", err)
	}

	return nil
}