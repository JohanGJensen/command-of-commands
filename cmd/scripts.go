package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

type ScriptContent struct {
	Scripts map[string]string `json:"scripts"`
}

func findScriptsInFile(file []byte) ScriptContent {
	var content ScriptContent

	err := json.Unmarshal(file, &content)
	if err != nil {
		fmt.Print("Failed to marshal JSON from file.")
		os.Exit(1)
	}

	return content
}

func scripts() (result string) {
	file, err := os.ReadFile("package.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "No file named 'package.json' in this directory: '%s'\n", err)
		os.Exit(1)
	}

	scripts := findScriptsInFile(file).Scripts

	keys := make([]string, 0, len(scripts))
	values := make([]string, 0, len(scripts))
	for k, v := range scripts {
		keys = append(keys, k)
		values = append(values, v)
	}

	prompt := promptui.Select{
		Label: "Select script command:",
		Items: keys,
	}
	index, _, errPrompt := prompt.Run()
	if errPrompt != nil {
		fmt.Printf("Prompt failed %v\n", errPrompt)
		return
	}

	return values[index]
}

var scriptsCmd = &cobra.Command{
	Use:   "scripts",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		command := scripts()
		/**
		 * Passing a command as a single string like this is not a safe pattern.
		 * Risk of shell injection.
		 * Reconsider this approach.
		 */
		err := exec.Command("bash", "-c", command).Run()
		if err != nil {
			fmt.Println("could not run command: ", err)
		}
	},
}
