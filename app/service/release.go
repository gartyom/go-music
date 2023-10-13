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

func (rs *release_service) New(artists []*model.Artist, release_title string, cover image.Image) error {
	// Generating uuid for image
	var release_id string
	for true {
		release_id = helpers.GenerateUUID()
		_, err := rs.repo.Release.GetById(release_id)
		if err == sql.ErrNoRows {
			break
		} else if err != nil {
			return err
		}
	}

	//Insterting release with no image in DataBase
	_, err := rs.repo.Release.New(&model.Release{
		Uuid:  release_id,
		Title: release_title,
		Image: "", // image is empty
	})
	if err != nil {
		return err
	}

	//Creating image inside {UploadsDir} folder
	ipath := "/images/release/" + release_id
	f, err := os.Create(config.Conf.UploadsDir + ipath)
	defer f.Close()
	if err != nil {
		rs.repo.Release.Delete(release_id)
		return err
	}
	//Writing form image to the image we created
	err = jpeg.Encode(f, cover, nil)
	if err != nil {
		rs.repo.Release.Delete(release_id)
		return err
	}
	f.Close()

	//Updating image column of inserted release
	_, err = rs.repo.Release.UpdateImage(release_id, ipath)
	if err != nil {
		rs.repo.Release.Delete(release_id)
		return err
	}

	//Creating many-to-many release<->artist relation
	for _, a := range artists {
		_, err = rs.repo.RelaseArtist.New(&model.ReleaseArtist{
			ArtistUuid:  a.Uuid,
			ReleaseUuid: release_id,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
