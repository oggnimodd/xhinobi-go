package helpers

import (
	"fmt"
	"os/exec"
)

// GetTreeOutput executes the 'tree' command with ignored patterns and returns its output.
// It returns an empty string if 'tree' is not found.
func GetTreeOutput() string {
	_, err := exec.LookPath("tree")
	if err != nil {
		fmt.Println("Warning: 'tree' command not found. Skipping tree generation.")
		return ""
	}

	cmd := exec.Command("tree", "-I", "node_modules|dist|vendor|*.log|tmp|images|go.sum|*.lock")
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Warning: 'tree' command finished with an error: %v\n", err)
	}

	header := "--- FOLDER TREE ---\n"
	footer := "\n--- FILE CONTENT ---\n\n"
	return header + string(output) + footer
}