CREATE TABLE IF NOT EXISTS `BeerItem`
(
    `id` bigint NOT NULL AUTO_INCREMENT,
    `name` varchar(100) NOT NULL,
    `brewery` varchar(100) DEFAULT NULL,
    `country` varchar(100) DEFAULT NULL,
    `price` float DEFAULT NULL,
    `currency` varchar(100) DEFAULT NULL,
    PRIMARY KEY (`id`)
)
