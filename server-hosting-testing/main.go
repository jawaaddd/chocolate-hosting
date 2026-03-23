package main

import (
	"encoding/json"
	"net/http"

	"fmt"
)

type VersionInfo struct {
	JavaVersion string `json:"id"`
	ServerURL   string `json:"url"`
}

type MCServerVersions struct {
	Versions []VersionInfo `json:"versions"`
}

func main() {
	const manifestJSONUrl string = "https://piston-meta.mojang.com/mc/game/version_manifest.json"

	var versions MCServerVersions

	resp, err := http.Get(manifestJSONUrl)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&versions)

}
