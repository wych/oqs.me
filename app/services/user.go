package services

import "oqs.me/models"

type UserRecord struct {
	models.UserRecord
}

func (r *UserRecord) ModelObj() *models.UserRecord {
	return &r.UserRecord
}

func (r *UserRecord) Save() error {
	err := models.DB.Create(r.ModelObj()).Error
	if err != nil {
		return err
	}
	return nil

}
