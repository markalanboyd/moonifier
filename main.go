package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/markalanboyd/moonifier/tokens"
)

func main() {
	fmt.Println("Opening test.lua...")

	file, err := os.Open("test.lua")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	fmt.Println("Scanning test.lua...")

	scanner := bufio.NewScanner(file)
	var builder strings.Builder
	tokenSet := tokens.GetTokenSet()
	for scanner.Scan() {
		var skip bool
		line := scanner.Text()
		line, skip = processLine(line)
		if skip {
			continue
		}
		line = tokenProcessor(line, &tokenSet)
		builder.WriteString(line + " ")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	minifiedString := builder.String()

	fmt.Println("Writing to minified.lua...")

	outFile, err := os.Create("minified.lua")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer outFile.Close()

	writer := bufio.NewWriter(outFile)

	writer.WriteString(minifiedString)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing to file:", err)
	}

	fmt.Println("Success!")
}
