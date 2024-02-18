-- Set timezone
SET time_zone = '+07:00';

-- Name: user_groups; Type: TABLE; Schema: public; Owner: -

CREATE TABLE IF NOT EXISTS `user_groups` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `title` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
    `created_at` datetime(6) NOT NULL,
    `updated_at` datetime(6) NOT NULL,
    `lock_version` int DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `index_user_groups_on_title` (`title`)
) ENGINE = InnoDB AUTO_INCREMENT = 5 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;
