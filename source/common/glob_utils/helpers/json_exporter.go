package helpers

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// ExportJSON exports data as formatted JSON to build/debug folder
// filename: if empty, auto-generates with timestamp
func ExportJSON(v any, filename string) error {
	// Create build/debug directory if not exists
	buildDir := filepath.Join("build", "debug")
	if err := os.MkdirAll(buildDir, 0755); err != nil {
		return fmt.Errorf("failed to create build/debug directory: %w", err)
	}

	// Generate filename with timestamp if empty
	if filename == "" {
		filename = fmt.Sprintf("export_%s.json", time.Now().Format("20060102_150405"))
	}

	// Marshal data with indentation
	jsonBytes, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	// Write to file
	filePath := filepath.Join(buildDir, filename)
	if err := os.WriteFile(filePath, jsonBytes, 0644); err != nil {
		return fmt.Errorf("failed to write JSON file: %w", err)
	}

	return nil
}

// ExportJSONDebug is convenience wrapper that logs to console + exports to file
func ExportJSONDebug(v any, filename string) {
	// Print to console
	// Export to file
	if err := ExportJSON(v, filename); err != nil {
		fmt.Printf("Warning: Failed to export JSON: %v\n", err)
	} else {
		filePath := filepath.Join("build", "debug", filename)
		if filename == "" {
			filePath = filepath.Join("build", "debug", fmt.Sprintf("export_%s.json", time.Now().Format("20060102_150405")))
		}
		fmt.Printf("✓ JSON exported to: %s\n", filePath)
	}
}
