package main

import (
	"github.com/mxzinke/geofabrik"
	"log"
)

func main() {
	// it's the path like the same in geofabrik download server
	filename := "europe/germany-latest.osm.pbf"

	// Getting the md5-hash
	hashBytes, err := geofabrik.OSMFileHash(filename)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Hash: %s", string(hashBytes))

	// Actually downloading:
	downloadFolderPath := "."
	filepath, err := geofabrik.OSMFile(downloadFolderPath, filename)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Downloaded to the path %s", filepath)
}
