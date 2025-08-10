package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"command-of-commands/models"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

type CommandObject struct {
	Key   string
	Value string
}

func runSelectPrompt(pkgJson models.PackageJson) string {
	scripts := pkgJson.Scripts
	keys := make([]string, 0, len(scripts))
	values := make([]string, 0, len(scripts))

	for k, v := range scripts {
		keys = append(keys, k)
		values = append(values, v)
	}

	selectPrompt := promptui.Select{
		Label: "Select script command:",
		Items: keys,
	}
	index, _, err := selectPrompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}

	return values[index]
}

func confirmSelectPrompt(cmd string) string {
	confirmPrompt := promptui.Prompt{
		Label: fmt.Sprintf("Are you sure you want to run: %s (y/n)", cmd),
		Validate: func(input string) error {
			if input != "y" && input != "n" {
				return fmt.Errorf("please enter y or n")
			}
			return nil
		},
	}
	r, err := confirmPrompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}

	if strings.ToLower(r) == "y" {
		log.Println("Confirmed!")
	} else {
		log.Println("Cancelled.")
		os.Exit(1)
	}

	return cmd
}

func getCmdFromPackageJSON() (result string) {
	file, err := os.ReadFile("package.json")
	if err != nil {
		log.Fatalf("No file named 'package.json' in this directory: '%s'\n", err)
	}

	var pkgJson models.PackageJson
	pkgJson.SetScripts(file)

	selectedCmd := runSelectPrompt(pkgJson)
	confirmedCmd := confirmSelectPrompt(selectedCmd)

	return confirmedCmd
}

var scriptsCmd = &cobra.Command{
	Use:   "scripts",
	Short: "Lists all script commands from package.json file",
	Long:  "Lists all script commands from package.json file",
	Run: func(cmd *cobra.Command, args []string) {
		command := getCmdFromPackageJSON()
		/**
		 * Passing a command as a single string like this is not a safe pattern.
		 * Risk of shell injection.
		 * Reconsider this approach.
		 */
		err := exec.Command("bash", "-c", command).Run()
		if err != nil {
			log.Fatalln("could not run command: ", err)
		}

		os.Exit(1)
	},
}
