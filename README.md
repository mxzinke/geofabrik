# Geofabrik

A small Golang package to access data of [Geofabrik.de (Download Server)](https://download.geofabrik.de) for OSM Data faster.

## Example

You can use the md5-hash to check if the file has changed and to prevent useless requests, for downloading the lastest version.

```golang
	// it's the path like the same in geofabrik download server
	filename := "europe/germany-latest.osm.pbf"

	// Getting the md5-hash
	hashBytes, err := geofabrik.OSMFileHash(filename)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Hash: %s", string(hashBytes))

	// Actually downloading:
	downloadFolderPath := "./downloading-to-here"
	filepath, err := geofabrik.OSMFile(downloadFolderPath, filename)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Downloaded to the path %s", filepath)
```

The file is beeing downloaded to the path, returned by the `OSMFile(...)` function.

