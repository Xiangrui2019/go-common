set NAMES utf8mb4;
set FOREIGN_KEY_CHECKS = 0;

CREATE TABLE `t_config_map` (
	`id` int(11) NOT NULL AUTO_INCREMENT,
	`name` varchar(255) NOT NULL,
	`description` varchar(255) DEFAULT NULL,
	PRIMARY KEY ( `id` )
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `t_config_entry` (
	`id` int(11) NOT NULL AUTO_INCREMENT,
	`key` varchar(255) NOT NULL,
	`value` varchar(255) NOT NULL,
	`config_map_id` int(11) NOT NULL,
	PRIMARY KEY ( `id` )
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

set FOREIGN_KEY_CHECKS = 1;

