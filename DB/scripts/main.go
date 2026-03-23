package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"context"

	"github.com/jackc/pgx/v5"
)

const MANIFEST_URL string = "https://piston-meta.mojang.com/mc/game/version_manifest_v2.json"

type VersionManifest struct {
	Versions []Version `json:"versions"`
}

type Version struct {
	ID          string `json:"id"`
	MetaDataURL string `json:"url"`
	VersionType string `json:"type"`
	MetaData    VersionMetaData
}

func main() {
	resp, err := http.Get(MANIFEST_URL)

	if err != nil {
		fmt.Println("Error getting the version_manifest JSON:", err.Error())
		return
	}

	defer resp.Body.Close()

	var manifest VersionManifest

	err = json.NewDecoder(resp.Body).Decode(&manifest)

	if err != nil {
		fmt.Println("Error decoding the manifest into object:", err.Error())
		return
	}

	type CleanData struct {
		ID          string
		VersionType string
		DownloadURL string
		Hash        string
		JavaVersion int
	}

	conn, err := pgx.Connect(context.Background(), "postgres://postgres:password@localhost:5432/chocolatehosting")
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer conn.Close(context.Background())

	for i := 0; i < len(manifest.Versions); i++ {
		getDownloadURL(manifest.Versions[i].MetaDataURL, &manifest.Versions[i].MetaData)
		var version CleanData = CleanData{
			ID:          manifest.Versions[i].ID,
			VersionType: manifest.Versions[i].VersionType,
			DownloadURL: manifest.Versions[i].MetaData.DownloadInfo.Server.URL,
			Hash:        manifest.Versions[i].MetaData.DownloadInfo.Server.Hash,
			JavaVersion: manifest.Versions[i].MetaData.JavaVersion.Version,
		}

		fmt.Println(version)

		_, err = conn.Exec(context.Background(),
			`INSERT INTO game_versions (id, version_type, java_version, download_url, server_hash)
             VALUES ($1, $2, $3, $4, $5)
             ON CONFLICT (id) DO NOTHING`,
			version.ID,
			version.VersionType,
			version.JavaVersion,
			version.DownloadURL,
			version.Hash,
		)

		if err != nil {
			fmt.Println("Error inserting", version.ID, ":", err)
		} else {
			fmt.Println("Inserted", version.ID)
		}

	}

}

type VersionMetaData struct {
	DownloadInfo struct {
		Server struct {
			URL  string `json:"url"`
			Hash string `json:"sha1"`
		} `json:"server"`
	} `json:"downloads"`
	JavaVersion struct {
		Version int `json:"majorVersion"`
	} `json:"JavaVersion"`
}

func getDownloadURL(versionJSON string, versionObj *VersionMetaData) {
	resp, err := http.Get(versionJSON)

	if err != nil {
		fmt.Println("Error getting a game version's metadata:", err.Error())
		return
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(versionObj)

	if err != nil {
		fmt.Println("Error decoding version metadata:", err.Error())
		return
	}
}
