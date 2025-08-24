package models

import (
	"encoding/json"
	"log"
)

type PackageJson struct {
	Name    string            `json:"name"`
	Scripts map[string]string `json:"scripts"`
}

func (p *PackageJson) SetScripts(file []byte) {
	err := json.Unmarshal(file, &p)
	if err != nil {
		log.Fatal("Failed to unmarshal JSON from file.")
	}
}

type PackageJsonSlice []PackageJson

func (p *PackageJsonSlice) AppendPackageJson(pkg PackageJson) {
	*p = append(*p, pkg)
}

func (p *PackageJsonSlice) GetAllScripts() map[string]string {
	scripts := make(map[string]string)

	for _, pkg := range *p {
		for key, script := range pkg.Scripts {
			scripts[pkg.Name+" - "+key] = script
		}
	}

	return scripts
}
