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

// CGIExtensionRequest is the structure that matches your JSON input.
type CGIExtensionRequest struct {
	// Define the fields according to your JSON structure
	Field1 string `json:"field1"`
	Field2 int    `json:"field2"`
	// Add more fields as needed
}

func main() {
	// Define and parse command-line flags
	extVerify := flag.Bool("ext_verify", false, "")
	flag.Parse()

	// Create the Verify struct
	verify := Verify{
		Alias:        "cgi-upshot",
		Description:  "upshot cgi extension for blockless runtime",
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
						fmt.Println("Parsed request:", request)
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

	// Proceed with the main loop of your program
	// For example, executing a Python script
	cmd := exec.Command("python3", "main.py")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error running script:", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", stdoutStderr)
}
