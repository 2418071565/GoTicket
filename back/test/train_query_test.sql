DROP TABLE IF EXISTS `train_numbers`;
DROP TABLE IF EXISTS `trains`;
DROP TABLE IF EXISTS `route_station`;
DROP TABLE IF EXISTS `stations`;
DROP TABLE IF EXISTS `routes`;

CREATE TABLE IF NOT EXISTS `stations` (
  `id` smallint unsigned AUTO_INCREMENT NOT NULL,
  `name` varchar(32) NOT NULL,
  `postion` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `station_index` (`name`),
  KEY `postion_index` (`postion`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO
    stations (
        name,
        postion
    ) VALUES
        ('宁波站', '宁波'),
        ('杭州东站', '杭州'),
        ('上海虹桥站', '上海'),
        ('南京南站', '南京');

CREATE TABLE IF NOT EXISTS `routes` (
  `id` smallint unsigned AUTO_INCREMENT NOT NULL,
  `price_pk` decimal(9, 2)  NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO
    routes (
        price_pk
    ) VALUES
        (10),
        (20);


CREATE TABLE IF NOT EXISTS `route_station` (
  `route_id` smallint unsigned NOT NULL,
  `station_id` smallint unsigned NOT NULL,
  `distance_from_start` double NOT NULL,
  PRIMARY KEY (`route_id`,`station_id`),
  FOREIGN KEY (`route_id`) REFERENCES routes(`id`),
  FOREIGN KEY (`station_id`) REFERENCES stations(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO 
    route_station (
        station_id,
        route_id, 
        distance_from_start
    ) VALUES
        (1, 1, 0),
        (2, 1, 23.1),
        (4, 1, 46.3),
        (1, 2, 0),
        (2, 2, 10.2),
        (3, 2, 50.43),
        (4, 2, 127);


CREATE TABLE IF NOT EXISTS `trains` (
  `id` smallint unsigned AUTO_INCREMENT NOT NULL,
  `train_type` enum('G','D','K') NOT NULL,
  `max_capacity` smallint unsigned NOT NULL,
  `seats` json NOT NULL,
  `avg_speed` double NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO
    trains (
        train_type,
        max_capacity,
        seats,
        avg_speed
    ) VALUES
        ('G', 1000, '[[10,10,10,10],[10,10,10,10]]', 270),
        ('G', 1000, '[[10,10,10,10],[10,10,10,10]]', 270),
        ('K', 1000, '[[10,10,10,10],[10,10,10,10]]', 120),
        ('D', 1000, '[[10,10,10,10],[10,10,10,10]]', 180);


CREATE TABLE IF NOT EXISTS `train_numbers` (
  `id` int unsigned AUTO_INCREMENT NOT NULL,
  `code` varchar(4) NOT NULL,
  `status` enum('Online','Offline') NOT NULL,
  `available_seats` smallint unsigned NOT NULL,
  `start_time` datetime NOT NULL,
  `train_id` smallint unsigned NOT NULL,
  `route_id` smallint unsigned NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`train_id`) REFERENCES trains(`id`),
  FOREIGN KEY (`route_id`) REFERENCES routes(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO
    train_numbers (
        status,
        code,
        available_seats,
        start_time,
        train_id,
        route_id
    ) VALUES
        (1, '2134',10, '2024-12-06 14:30', 1, 1),
        (2, '1111',10, '2024-12-05 14:30', 1, 1),
        (1, '4525',10, '2024-12-06 15:21', 2, 1),
        (1, '975',10, '2024-12-07 20:51', 2, 2),
        (2, '6421',10, '2024-12-07 16:31', 3, 2);


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
                    (SELECT route_id FROM route_station WHERE station_id = 1) AS s1
                INNER JOIN 
                    (SELECT route_id FROM route_station WHERE station_id = 3) AS s2
                ON s1.route_id = s2.route_id
            ) 
        GROUP BY rs.route_id
    ) AS accessible_routes
    ON train_numbers.route_id = accessible_routes.route_id
WHERE
    train_numbers.status = 1;