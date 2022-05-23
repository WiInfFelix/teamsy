package organisations

import (
	"gorm.io/gorm"
	"teamsy/internal/pkg/db"
)

type Organisation struct {
	gorm.Model

	OrganisationName string
	Email            string
}

func (org Organisation) Save() (uint, error) {

	orgDB := Organisation{
		Model:            gorm.Model{},
		OrganisationName: org.OrganisationName,
		Email:            org.Email,
	}

	result := db.Db.Create(&orgDB)

	if result.Error != nil {
		return 0, nil
	}

	return orgDB.ID, nil
}

func GetAll() ([]*Organisation, error) {
	var orgs []*Organisation

	result := db.Db.Find(&orgs)

	return orgs, result.Error
}
