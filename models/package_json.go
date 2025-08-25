package package_json

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

type PackageJson struct {
	Name    string            `json:"name"`
	Scripts map[string]string `json:"scripts"`
}

func (p *PackageJson) setScripts(file []byte) {
	err := json.Unmarshal(file, &p)
	if err != nil {
		log.Fatal("Failed to unmarshal JSON from file.")
	}
}

type PackageJsonFiles []PackageJson

func (p *PackageJsonFiles) appendPackageJson(pkg PackageJson) {
	*p = append(*p, pkg)
}

func (p *PackageJsonFiles) GetAllScripts() map[string]string {
	scripts := make(map[string]string)

	for _, pkg := range *p {
		for key, script := range pkg.Scripts {
			scripts[pkg.Name+" - "+key] = script
		}
	}

	return scripts
}

func (p *PackageJsonFiles) ReadDirectoryContent(path string) {
	var pkgJson PackageJson

	file, err := os.ReadFile(path + "package.json")
	if err == nil {
		pkgJson.setScripts(file)
		p.appendPackageJson(pkgJson)
	}
}

func (p *PackageJsonFiles) ReadDirectoryContentRecursive(path string) {
	p.ReadDirectoryContent(path)

	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatalln("failed to read directory", err)
	}

	for _, entry := range entries {
		if entry.IsDir() && !strings.HasPrefix(entry.Name(), ".") {
			p.ReadDirectoryContentRecursive(path + entry.Name() + "/")
		}
	}
}
