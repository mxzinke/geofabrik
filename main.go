package geofabrik

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

const (
	baseURL = "http://download.geofabrik.de/"
)

// OSMFileHash returns the md5 hash of the file, fetched from GeoFabrik Server.
func OSMFileHash(fileName string) ([]byte, error) {
	downloadURL := baseURL + fileName

	resp, err := http.Get(downloadURL + ".md5")
	if err != nil {
		return []byte{}, fmt.Errorf("Fetching the md5-hash from GeoFabrik server: %v ", err)
	}
	defer resp.Body.Close()

	// Only take first 32 bytes, because it contains 32 chars + the file name in there
	hash, err := ioutil.ReadAll(io.LimitReader(resp.Body, 32))
	if err != nil {
		return []byte{}, err
	}

	return hash, nil
}

// OSMFile will download the given GeoFabrik filename, to the baseFolder.
// It will return a string of the path, where the file could be found.
func OSMFile(baseFolder string, fileName string) (string, error) {
	downloadURL := baseURL + fileName
	filePath := path.Join(baseFolder, fileName)

	if err := downloadFile(filePath, downloadURL); err != nil {
		return "", fmt.Errorf("Downloading the OSM File from GeoFabrik Server: %v ", err)
	}

	return filePath, nil
}

// downloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func downloadFile(filepath string, url string) error {
	// Create the file, but give it a tmp file extension, this means we won't overwrite a
	// file until it's downloaded, but we'll remove the tmp extension once downloaded.
	out, err := os.Create(filepath + ".tmp")
	if err != nil {
		return err
	}

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		out.Close()
		return err
	}
	defer resp.Body.Close()

	if _, err = io.Copy(out, resp.Body); err != nil {
		out.Close()
		return err
	}

	// Close the file without defer so it can happen before Rename()
	out.Close()

	if err = os.Rename(filepath+".tmp", filepath); err != nil {
		return err
	}
	return nil
}
