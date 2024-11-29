package repositories

import (
	"errors"

	"github.com/Bruno07/tasks-api/internal/infra/db"
	"github.com/Bruno07/tasks-api/internal/models"
)

type IProfileRepository interface {
	Find(profileId int64) (models.Profile, error)
}

type ProfileRepository struct{}

type ProfileMockRepository struct{}

func (pr *ProfileRepository) Find(profileId int64) (models.Profile, error) {

	profile := models.Profile{}

	result := db.GetInstance().Model(profile).Find(map[string]interface{}{})

	return profile, result.Error

}

func (pr *ProfileMockRepository) Find(profileId int64) (models.Profile, error) {

	var err error
	var profile = models.Profile{}
	
	profiles := map[int64]string{
		1: "Technician",
		2: "Manager",
	}

	if profiles[profileId] == "" {
		err = errors.New("Invalid profile!")
	} else {
		profile = models.Profile{
			ID: profileId,
			Name: profiles[profileId],
		}
	}


	return profile, err

}
