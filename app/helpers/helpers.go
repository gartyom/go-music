package helpers

import (
	"archive/zip"
	"errors"
	"image"
	"image/jpeg"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

func GenerateUUID() string {
	id := uuid.New().String()
	id = strings.Replace(id, "-", "", -1)[3:14]
	return id
}

type ReleaseForm struct {
	Artist  string
	Title   string
	Image   image.Image
	Archive multipart.File
	Err     error
}

func ParseReleaseForm(r *http.Request) *ReleaseForm {

	var artistErr error
	var titleErr error

	artist_name := r.FormValue("artist")
	if artist_name == "" {
		artistErr = errors.New("Error: artist name is missing")
	}

	release_title := r.FormValue("release")
	if release_title == "" {
		titleErr = errors.New("Error: release title is missing")
	}

	cover, _, coverErr := r.FormFile("release_cover")

	archive, _, archiveErr := r.FormFile("release_songs")

	errs := errors.Join(artistErr, titleErr, coverErr, archiveErr)

	if errs != nil {
		return &ReleaseForm{
			Err: errs,
		}
	}

	image, imageErr := jpeg.Decode(cover)

	errs = errors.Join(errs, imageErr)

	return &ReleaseForm{
		Artist:  artist_name,
		Title:   release_title,
		Image:   image,
		Archive: archive,
		Err:     errs,
	}
}

func GetArchiveSize(archive multipart.File) (int64, error) {

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

func CheckArchiveFiles(zipReader *zip.Reader) error {
	signatureBuffer := make([]byte, 3)
	for _, file := range zipReader.File {
		if file.FileInfo().IsDir() {
			continue
		}

		reader, err := file.Open()
		if err != nil {
			return err
		}

		_, err = reader.Read(signatureBuffer)
		if err != nil {
			return err
		}

		if !((signatureBuffer[0] == 0x49 && signatureBuffer[1] == 0x44 && signatureBuffer[2] == 0x33) || (signatureBuffer[0] == 0xFF && signatureBuffer[1] == 0xFB)) {
			return errors.New("Error: Song must be a mp3 file")
		}
	}
	return nil
}

// NOT IMPLEMENTED YET
// GOAL: extract metadata from files (artist, album)
func ExtractID3Metadata(zipReader *zip.Reader) error {
	for _, file := range zipReader.File {
		if file.FileInfo().IsDir() {
			continue
		}

		reader, err := file.Open()
		if err != nil {
			return err
		}

		fs := file.FileInfo().Size()
		metadataBuffer := make([]byte, fs)
		_, err = reader.Read(metadataBuffer)
		if err != nil {
			return err
		}
	}
	return nil
}
