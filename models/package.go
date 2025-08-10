package models

import (
	"encoding/json"
	"log"
)

type PackageJson struct {
	Scripts map[string]string `json:"scripts"`
}

func (p *PackageJson) SetScripts(file []byte) {
	err := json.Unmarshal(file, &p)
	if err != nil {
		log.Fatal("Failed to marshal JSON from file.")
	}
}
