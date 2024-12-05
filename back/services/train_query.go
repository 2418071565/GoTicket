package services

import (
	"strconv"
	"strings"
	"time"

	"github.com/2418071565/GoTicket/dto"
	"github.com/2418071565/GoTicket/models"
)

// 获取通过所有经过 start 和 end 的路线
func TrainsQuery(start string, end string) ([]dto.AvailableRoute, error) {
	start_ids, err := models.Station{}.GetStations(start)
	if err != nil {
		return nil, err
	}
	end_ids, err := models.Station{}.GetStations(end)
	if err != nil {
		return nil, err
	}
	result := make([]dto.AvailableRoute, 0)
	for _, start_id := range start_ids {
		for _, end_id := range end_ids {
			routes, err := models.Route{}.GetRoutesByStationId(start_id, end_id)
			if err != nil {
				return nil, err
			}
			result = append(result, routes...)
		}
	}
	return result, nil
}

// 过滤掉所有不符合用户偏好的车次
func TrainsFilter(routes []dto.AvailableRoute, pref *dto.Preferences) ([]dto.AvailableRoute, error) {
	result := make([]dto.AvailableRoute, 0)

	for i := 0; i < len(routes); i++ {
		station_ids_str := routes[i].Station_ids.(string)
		station_ids := strings.Split(station_ids_str, ",")
		ids := make([]uint16, 0)
		for _, id_str := range station_ids {
			id, _ := strconv.Atoi(id_str)
			ids = append(ids, uint16(id))
		}

		station_distances_str := routes[i].Station_distances.(string)
		station_distances := strings.Split(station_distances_str, ",")
		distances := make([]float64, 0)
		for _, distance_str := range station_distances {
			distance, _ := strconv.ParseFloat(distance_str, 64)
			distances = append(distances, distance)
		}

		routes[i].Station_ids = ids
		routes[i].Station_distances = distances

		routes[i].Station_expected_time = make([]time.Duration, 1)
		for _, 
	}
}
