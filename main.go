package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// Verify structure for JSON output
type Verify struct {
	Alias        string `json:"alias"`
	Description  string `json:"description"`
	IsCGI        bool   `json:"is_cgi"`
}

// CGIExtensionRequest
type CGIExtensionRequest struct {
	Arguments []string `json:"arguments"`
}

func main() {
	// Define and parse command-line flags
	extVerify := flag.Bool("ext_verify", false, "")
	flag.Parse()

	// Create the Verify struct
	verify := Verify{
		Alias:        "cgi-allora-infer",
		Description:  "allora cgi extension for blockless runtime",
		IsCGI:        true,
	}

	// Check if the --ext_verify flag is set
	if *extVerify {
		jsonData, err := json.Marshal(verify)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Println(string(jsonData))
		os.Exit(0)
	}

	// Read and parse JSON from stdin
	scanner := bufio.NewScanner(os.Stdin)
	var request CGIExtensionRequest
	var inputProcessed bool

	// Read the first line for the length of the data
	if scanner.Scan() {
		lengthStr := strings.TrimSpace(scanner.Text())
		if lengthStr != "" {
			length, err := strconv.Atoi(lengthStr)
			if err != nil {
				fmt.Println("Invalid length value:", err)
			} else if scanner.Scan() {
				jsonData := scanner.Text()
				if len(jsonData) == length {
					if err := json.Unmarshal([]byte(jsonData), &request); err != nil {
						fmt.Println("Error parsing JSON:", err)
					} else {
						inputProcessed = true
					}
				} else {
					fmt.Println("Data length mismatch")
				}
			}
		}
	}

	if !inputProcessed {
		fmt.Println("No valid input provided, continuing with the main loop.")
	}

	scriptPath := "/tmp/runtime/extensions/main.py"
	if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
		scriptPath = "/app/main.py"
		if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
			scriptPath = "/app/runtime/extensions/main.py"
		}
	}

	// Execute the Python script with arguments
	cmdArgs := append([]string{ scriptPath }, request.Arguments...)
	cmd := exec.Command("python3", cmdArgs...)
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error running script:", err)
		os.Exit(1)
	}

	fmt.Printf("%s", stdoutStderr)
}
