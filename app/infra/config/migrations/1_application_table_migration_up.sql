set NAMES utf8mb4;
set FOREIGN_KEY_CHECKS = 0;

CREATE TABLE `t_applications` (
	`id` int(11) NOT NULL AUTO_INCREMENT,
	`name` varchar(255) NOT NULL,
	`description` varchar(255) DEFAULT NULL,
	PRIMARY KEY ( `id` )
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

set FOREIGN_KEY_CHECKS = 1;

