package main

import "strings"

func processLine(line string) (string, bool) {
	line = strings.TrimSpace(line)
	switch {
	case line == "": // Skip empty lines
		return "", true
	case strings.HasPrefix(line, "--"): // Skip comments
		return "", true
	case strings.HasPrefix(line, "function"): // Prepend "local" to functions
		return "local " + line, false
	default: // No change for other lines
		return line, false
	}
}
