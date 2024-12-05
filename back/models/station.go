package models

import (
	"github.com/2418071565/GoTicket/dto"
	"github.com/2418071565/GoTicket/storage/db"
)

type Station struct{}

func (Station) GetStations(postion string) ([]uint16, error) {
	sts := make([]uint16, 0)
	if err := db.DB.Table("stations").
		Select("stations.id").
		Where("stations.postion = ? OR stations.name = ?", postion, postion).
		Scan(&sts).Error; err != nil {
		return nil, err
	}
	return sts, nil
}

func (Station) GetStationsByIds(station_ids []uint16) ([]dto.Station, error) {
	sts := make([]dto.Station, 0)
	if err := db.DB.Table("stations").
		Select("stations.name, stations.postion").
		Where("stations.id IN ?", station_ids).
		Scan(&sts).Error; err != nil {
		return nil, err
	}
	return sts, nil
}
