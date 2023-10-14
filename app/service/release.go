package service

import (
	"database/sql"
	"image"
	"image/jpeg"
	"os"

	"github.com/gartyom/go-music/config"
	"github.com/gartyom/go-music/helpers"
	"github.com/gartyom/go-music/model"
	"github.com/gartyom/go-music/repository"
)

type release_service struct {
	repo *repository.Repository
}

func NewReleaseService(repo *repository.Repository) releaseService {
	return &release_service{
		repo,
	}
}

func (rs *release_service) New(artists []*model.Artist, releaseTitle string, cover image.Image) (string, error) {
	// Generating uuid for image
	var releaseUuid string
	for true {
		releaseUuid = helpers.GenerateUUID()
		_, err := rs.repo.Release.GetById(releaseUuid)
		if err == sql.ErrNoRows {
			break
		} else if err != nil {
			return "", err
		}
	}

	//Insterting release with no image in DataBase
	_, err := rs.repo.Release.New(&model.Release{
		Uuid:  releaseUuid,
		Title: releaseTitle,
		Image: "", // image is empty
	})
	if err != nil {
		return "", err
	}

	//Creating image inside {UploadsDir} folder
	ipath := "/images/release/" + releaseUuid
	f, err := os.Create(config.Conf.UploadsDir + ipath)
	defer f.Close()
	if err != nil {
		rs.repo.Release.Delete(releaseUuid)
		return "", err
	}
	//Writing form image to the image we created
	err = jpeg.Encode(f, cover, nil)
	if err != nil {
		rs.repo.Release.Delete(releaseUuid)
		return "", err
	}

	//Updating image column of inserted release
	_, err = rs.repo.Release.UpdateImage(releaseUuid, ipath)
	if err != nil {
		rs.repo.Release.Delete(releaseUuid)
		return "", err
	}

	//Creating many-to-many release<->artist relation
	for _, a := range artists {
		_, err = rs.repo.RelaseArtist.New(&model.ReleaseArtist{
			ArtistUuid:  a.Uuid,
			ReleaseUuid: releaseUuid,
		})
		if err != nil {
			return "", err
		}
	}

	return releaseUuid, err
}
