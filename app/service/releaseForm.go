package service

import (
	"archive/zip"
	"encoding/json"
	"errors"
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

	var metadata Meta
	err = json.Unmarshal([]byte(metaStr), &metadata)
	if err != nil {
		return nil, err
	}
	return &metadata, nil
}

func checkmetadata(file *zip.File) error {

	reader, err := file.Open()
	defer reader.Close()
	if err != nil {
		return err
	}

	s, err := extractmetadata(reader)
	if err != nil {
		return err
	}

	if s.Format.Tags.Album != "" && s.Format.Tags.AlbumArtist != "" && s.Format.Tags.Title != "" && s.Format.Tags.Artist != "" {
		return nil
	} else {
		return errors.New("No metadata in song(s)")
	}

}

func checkArchiveFiles(zipReader *zip.Reader) error {
	flacMarkerBuffer := make([]byte, 4)
	for _, file := range zipReader.File {
		if file.FileInfo().IsDir() {
			continue
		}

		err := checkFileType(file, &flacMarkerBuffer)
		if err != nil {
			return err
		}

		err = checkmetadata(file)
		if err != nil {
			return err
		}

	}
	return nil
}

func (service *release_form) Deconstruct(r *http.Request) (*[]Meta, error) {

	unzipper, err := getUnzipper(r)
	if err != nil {
		return nil, err
	}

	err = checkArchiveFiles(unzipper)
	if err != nil {
		return nil, err
	}

	var a []Meta

	for _, file := range unzipper.File {
		rc, err := file.Open()
		if err != nil {
			return nil, err
		}
		m, err := extractmetadata(rc)
		if err != nil {
			return nil, err
		}
		a = append(a, *m)
	}

	return &a, nil
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
