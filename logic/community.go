package logic

import (
	"errors"
	"goweb/dao/mysql"
	"goweb/models"
)

func GetCommunityList() ([]*models.Community, error) {
	return mysql.GetCommunityList()
}

func GetCommunityDetailByID(id int64) (*models.Community, error) {
	community, err := mysql.GetCommunityDetailByID(id)
	if errors.Is(err, mysql.ErrorCommunityNotExist) {
		return nil, ErrCommunityNotExist
	}
	return community, err
}
