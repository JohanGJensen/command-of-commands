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

func runSelectPrompt(pkgSlice models.PackageJsonSlice) string {
	scripts := pkgSlice.GetAllScripts()
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

func runConfirmSelectPrompt(cmd string) string {
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

func getCmdFromPackageJSON(path string, recursive bool) (result string) {
	var pkgSlice models.PackageJsonSlice

	if recursive {
		readDirectoryContentRecursive(path, &pkgSlice)
	} else {
		readDirectoryContent(path, &pkgSlice)
	}

	log.Println(pkgSlice)
	selectedCmd := runSelectPrompt(pkgSlice)
	confirmedCmd := runConfirmSelectPrompt(selectedCmd)

	return confirmedCmd
}

func readDirectoryContent(path string, pkgSlice *models.PackageJsonSlice) {
	var pkgJson models.PackageJson

	file, err := os.ReadFile(path + "package.json")
	if err == nil {
		pkgJson.SetScripts(file)
		pkgSlice.AppendPackageJson(pkgJson)
	}
}

func readDirectoryContentRecursive(path string, pkgSlice *models.PackageJsonSlice) {
	readDirectoryContent(path, pkgSlice)

	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatalln("failed to read directory", err)
	}

	for _, entry := range entries {
		if entry.IsDir() && !strings.HasPrefix(entry.Name(), ".") {
			readDirectoryContentRecursive(path+entry.Name()+"/", pkgSlice)
		}
	}
}

var scriptsCmd = &cobra.Command{
	Use:   "scripts",
	Short: "Lists all script commands from package.json file",
	Long:  "Lists all script commands from package.json file",
	Run: func(cmd *cobra.Command, args []string) {
		recursive, err := cmd.Flags().GetBool("recursive")
		if err != nil {
			log.Fatalln("could not read recursive flag properly")
		}

		command := getCmdFromPackageJSON("./", recursive)
		/**
		 * Passing a command as a single string like this is not a safe pattern.
		 * Risk of shell injection.
		 * Reconsider this approach.
		 */
		runErr := exec.Command("bash", "-c", command).Run()
		if runErr != nil {
			log.Fatalln("could not run command: ", runErr)
		}

		os.Exit(1)
	},
}
