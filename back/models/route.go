package models

import (
	"github.com/2418071565/GoTicket/dto"
	"github.com/2418071565/GoTicket/storage/db"
)

type Route struct{}

const routeQuerySQL = `
SELECT
    accessible_routes.route_id as route_id,
    CONCAT(trains.train_type, train_numbers.code) as code,
    accessible_routes.station_ids as station_ids,
    accessible_routes.station_distances as station_distances,
    accessible_routes.price_pk as price_pk,
    train_numbers.available_seats as available_seats,
    trains.avg_speed as avg_speed,
    trains.seats as seats
FROM
    train_numbers
    INNER JOIN
    trains
    ON train_numbers.train_id = trains.id
    INNER JOIN 
    (
        SELECT 
            rs.route_id as route_id,
            GROUP_CONCAT(station_id ORDER BY distance_from_start) AS station_ids,
            GROUP_CONCAT(distance_from_start ORDER BY distance_from_start) AS station_distances,
            r.price_pk as price_pk
        FROM
            route_station AS rs
        INNER JOIN
            routes AS r
            ON r.id = rs.route_id
        WHERE
            rs.route_id IN (
                SELECT s1.route_id
                FROM 
                    (SELECT route_id FROM route_station WHERE station_id = ?) AS s1
                INNER JOIN 
                    (SELECT route_id FROM route_station WHERE station_id = ?) AS s2
                ON s1.route_id = s2.route_id
            ) 
        GROUP BY rs.route_id
    ) AS accessible_routes
    ON train_numbers.route_id = accessible_routes.route_id;
WHERE
    train_numbers.status = 1;
`

// func (Route) GetRouteById(route_id uint16) (*dto.Route, error) {
// 	route = &dto.Route{ID: route_id}
// 	if err != db.DB.Table("routes").Select()
// }

func (Route) GetRoutesByStationId(start uint16, end uint16) ([]dto.AvailableRoute, error) {
	routes := make([]dto.AvailableRoute, 0)
	if err := db.DB.Exec(routeQuerySQL, start, end).Scan(&routes).Error; err != nil {
		return nil, err
	}
	return routes, nil
}
