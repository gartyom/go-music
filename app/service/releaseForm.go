package service

import (
	"archive/zip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	ffmpeg_go "github.com/u2takey/ffmpeg-go"
)

type release_form struct {
}

func NewReleaseFormService() releaseFormService {
	return &release_form{}
}

func parseReleaseForm(r *http.Request) (multipart.File, error) {
	archive, _, err := r.FormFile("archive")
	return archive, err
}

func getArchiveSize(archive multipart.File) (int64, error) {
	fileSize, err := archive.Seek(0, 2)
	if err != nil {
		return fileSize, err
	}
	_, err = archive.Seek(0, 0)
	if err != nil {
		return fileSize, err
	}

	return fileSize, nil
}

func getUnzipper(r *http.Request) (*zip.Reader, error) {
	a, err := parseReleaseForm(r)
	if err != nil {
		return nil, err
	}

	s, err := getArchiveSize(a)
	if err != nil {
		return nil, err
	}

	reader, err := zip.NewReader(a, s)
	if err != nil {
		return nil, err
	}

	return reader, nil
}

func checkFileType(file *zip.File, flacMarkerBuffer *[]byte) error {
	reader, err := file.Open()
	defer reader.Close()
	if err != nil {
		return err
	}

	_, err = reader.Read(*flacMarkerBuffer)
	if err != nil {
		return err
	}

	// if string(*flacMarkerBuffer) != "fLaC" {
	// 	return errors.New("Error: Only flac files are supported")
	// }

	return nil
}

type tags struct {
	Title       string `json:"TITLE"`
	Album       string `json:"ALBUM"`
	AlbumArtist string `json:"album_artist"`
	Artist      string `jsong:"ARTIST"`
	Track       string `json:"track"`
}

type format struct {
	Tags tags `json:"tags"`
}

type Meta struct {
	Format format `json:"format"`
}

func extractmetadata(reader io.Reader) (*Meta, error) {
	metaStr, err := ffmpeg_go.ProbeReader(reader)

	if err != nil {
		return nil, err
	}

	fmt.Println(metaStr)
	var metadata Meta
	err = json.Unmarshal([]byte(metaStr), &metadata)
	if err != nil {
		return nil, err
	}
	return &metadata, nil
}

func checkmetadata(file *zip.File) (*Meta, error) {

	reader, err := file.Open()
	defer reader.Close()
	if err != nil {
		return nil, err
	}

	s, err := extractmetadata(reader)
	if err != nil {
		return nil, err
	}

	if s.Format.Tags.Album != "" && s.Format.Tags.Title != "" && s.Format.Tags.Artist != "" {
		return s, err
	} else {
		return nil, errors.New("No metadata in song(s)")
	}

}

func checkArchiveFiles(zipReader *zip.Reader) (*[]Meta, error) {
	var releaseMeta []Meta
	flacMarkerBuffer := make([]byte, 4)
	for _, file := range zipReader.File {
		if file.FileInfo().IsDir() {
			continue
		}

		err := checkFileType(file, &flacMarkerBuffer)
		if err != nil {
			return nil, err
		}

		m, err := checkmetadata(file)
		if err != nil {
			return nil, err
		}
		releaseMeta = append(releaseMeta, *m)
	}
	return &releaseMeta, nil
}

func (service *release_form) Deconstruct(r *http.Request) (*[]Meta, *zip.Reader, error) {

	unzipper, err := getUnzipper(r)
	if err != nil {
		return nil, nil, err
	}

	releaseMeta, err := checkArchiveFiles(unzipper)
	if err != nil {
		return nil, nil, err
	}

	return releaseMeta, unzipper, nil
}

// func SaveReleaseLocally(zipReader *zip.Reader) error {
// 	for _, file := range zipReader.File {
// 		if file.FileInfo().IsDir() {
// 			continue
// 		}

// 		reader, err := file.Open()
// 		if err != nil {
// 			return nil
// 		}
// 	}

// }
