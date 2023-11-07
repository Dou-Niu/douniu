CREATE TABLE `user`
(
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `nickname`         varchar(32) NULL DEFAULT '',
    `password`         varchar(32)  NOT NULL DEFAULT '',
    `phone`            varchar(11)  NOT NULL DEFAULT '',
    `avatar`           varchar(255) NOT NULL DEFAULT '',
    `background_image` varchar(255) NOT NULL DEFAULT '',
    `signature`        varchar(255) NOT NULL DEFAULT '',
    `create_time`      datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time`      datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `delete_time`      datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `id` (`id`),
    UNIQUE KEY `phone` (`phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci

