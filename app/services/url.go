package services

import (
	"errors"

	"oqs.me/models"
	"oqs.me/utils"
)

type URLRecord struct {
	ShortURL utils.Base62Int
	RealURL  string
}

func (u *URLRecord) Create(id uint) error {
	base62n := utils.Base10to62(id)
	challenge := utils.GenChallenge()
	u.ShortURL = base62n + challenge
	err := models.DB.Create(&models.OQSRecord{
		ID:        id,
		Challenge: string(challenge),
		RealURL:   u.RealURL,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *URLRecord) Fetch() error {
	base62ID := u.ShortURL[:len(u.ShortURL)-2]
	challenge := u.ShortURL[len(u.ShortURL)-2:]
	id, err := utils.Base62to10(base62ID)
	if err != nil {
		return err
	}
	var record models.OQSRecord
	models.DB.Where("id = ?", id).First(&record)
	if string(challenge) != record.Challenge {
		return errors.New("challenge check false")
	}
	u.RealURL = record.RealURL
	return nil
}

func (u *URLRecord) CreateAndCache(id uint) error {
	err := u.Create(id)
	if err != nil {
		return err
	}
	err = makeRecordCache(string(u.ShortURL), u.RealURL)
	if err != nil {
		return err
	}
	return nil
}

func (u *URLRecord) FetchAndCache() error {
	err := u.Fetch()
	if err != nil {
		return err
	}
	err = makeRecordCache(string(u.ShortURL), u.RealURL)
	if err != nil {
		return err
	}
	return nil
}

func (u *URLRecord) Valid() bool {
	maxid := MaxOQSID()
	base62ID := u.ShortURL[:len(u.ShortURL)-2]
	return utils.Base10to62(maxid).BiggerThan(base62ID)
}

func (u *URLRecord) GetCache() bool {
	s, err := readRecordCache(string(u.ShortURL))
	if err != nil {
		return false
	}
	u.RealURL = s
	return true
}
