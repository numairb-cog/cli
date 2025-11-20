package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/numairb-cog/cli/go-cli/cmd"
)

func main() {
	args := os.Args[1:]
	
	newArgs := append([]string{os.Args[0], "exec"}, args...)
	os.Args = newArgs
	
	removedSwitches := map[string]bool{
		"--always-spawn":        true,
		"--ignore-existing":     true,
		"--shell-auto-fallback": true,
		"--npm":                 true,
		"--node-arg":            true,
		"-n":                    true,
	}
	
	filteredArgs := []string{os.Args[0], "exec"}
	sawRemovedFlags := false
	
	for i := 0; i < len(args); i++ {
		arg := args[i]
		
		if strings.HasPrefix(arg, "--") {
			key := strings.Split(arg, "=")[0]
			if removedSwitches[key] {
				fmt.Fprintf(os.Stderr, "npx: the %s argument has been removed.\n", key)
				sawRemovedFlags = true
				if !strings.Contains(arg, "=") && i+1 < len(args) && !strings.HasPrefix(args[i+1], "-") {
					i++ // Skip the value too
				}
				continue
			}
		}
		
		if arg == "-p" || strings.HasPrefix(arg, "-p=") {
			arg = strings.Replace(arg, "-p", "--package", 1)
		} else if arg == "--shell" || strings.HasPrefix(arg, "--shell=") {
			arg = strings.Replace(arg, "--shell", "--script-shell", 1)
		} else if arg == "--no-install" {
			arg = "--yes=false"
		}
		
		filteredArgs = append(filteredArgs, arg)
	}
	
	if sawRemovedFlags {
		fmt.Fprintln(os.Stderr, "See `npm help exec` for more information")
	}
	
	os.Args = filteredArgs
	
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
