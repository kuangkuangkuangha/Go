CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_german2_ci NOT NULL,
  `password` varchar(255) COLLATE utf8mb4_german2_ci NOT NULL,
  `avatar` varchar(255) COLLATE utf8mb4_german2_ci NOT NULL,
  `role` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_german2_ci;

CREATE TABLE `team` (
  `id` int NOT NULL AUTO_INCREMENT,
  `teamname` varchar(255) COLLATE utf8mb4_german2_ci NOT NULL,
  `logo` varchar(255) COLLATE utf8mb4_german2_ci NOT NULL,
  `member` varchar(255) COLLATE utf8mb4_german2_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_german2_ci;

CREATE TABLE `players` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_german2_ci NOT NULL,
  `picture` varchar(255) COLLATE utf8mb4_german2_ci NOT NULL,
  `team_belong_id` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_german2_ci;

CREATE TABLE `game` (
  `id` int NOT NULL AUTO_INCREMENT,
  `creat_at` timestamp(4) NULL DEFAULT CURRENT_TIMESTAMP(4) ON UPDATE CURRENT_TIMESTAMP(4),
  `name` varchar(255) COLLATE utf8mb4_german2_ci DEFAULT NULL,
  `data` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_german2_ci DEFAULT NULL,
  `place` varchar(255) COLLATE utf8mb4_german2_ci DEFAULT NULL,
  `team_aid` int DEFAULT NULL,
  `Team_bid` int DEFAULT NULL,
  `count` int DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_german2_ci;

CREATE TABLE `user_game` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL,
  `game_id` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_german2_ci;

CREATE TABLE `team_players` (
  `id` int NOT NULL AUTO_INCREMENT,
  `team_id` int NOT NULL,
  `playerid` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_german2_ci;
