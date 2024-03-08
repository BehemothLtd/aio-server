SET time_zone = '+07:00';

-- Name: devices; Type: TABLE; Schema: public; Owner: -

CREATE TABLE IF NOT EXISTS `devices` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint DEFAULT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `code` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `description` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `state` int NOT NULL DEFAULT '1',
  `device_type_id` int NOT NULL,
  `seller` text COLLATE utf8mb4_unicode_ci,
  `buy_at` datetime(6) DEFAULT NULL,
  `created_at` datetime(6) NOT NULL,
  `updated_at` datetime(6) NOT NULL,
  `lock_version` int NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `index_devices_on_code` (`code`),
  KEY `index_devices_on_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;