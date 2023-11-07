CREATE TABLE `video`
(
    `id`          bigint                                      NOT NULL AUTO_INCREMENT,
    `user_id`     bigint                                      NOT NULL,
    `title`       varchar(255) COLLATE utf8mb4_unicode_520_ci NOT NULL DEFAULT '',
    `partition`   bigint                                      NOT NULL DEFAULT '0',
    `play_url`    varchar(255) COLLATE utf8mb4_unicode_520_ci NOT NULL DEFAULT '',
    `cover_url`   varchar(255) COLLATE utf8mb4_unicode_520_ci NOT NULL DEFAULT '',
    `create_time` datetime                                    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` datetime                                    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `delete_time` datetime                                             DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `id` (`id`),
    KEY           `idx_user_id_create_time` (`user_id`,`create_time`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci

