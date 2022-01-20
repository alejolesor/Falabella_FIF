CREATE TABLE IF NOT EXISTS `BeerItem`
(
    `id` bigint NOT NULL AUTO_INCREMENT,
    `name` varchar(50) NOT NULL,
    `brewery` varchar(60) DEFAULT NULL,
    `country` varchar(50) DEFAULT NULL,
    `price` float DEFAULT NULL,
    `currency` varchar(50) DEFAULT NULL,
    PRIMARY KEY (`id`)
)
