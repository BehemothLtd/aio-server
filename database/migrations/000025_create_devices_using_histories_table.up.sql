SET time_zone = '+07:00';

-- Name: devices_using_histories; Type: TABLE; Schema: public; Owner: -

CREATE TABLE IF NOT EXISTS `devices_using_histories` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint DEFAULT NULL,
  `device_id` bigint NOT NULL,
  `state` int NOT NULL,
  `created_at` datetime(6) NOT NULL,
  `updated_at` datetime(6) NOT NULL,
  `lock_version` int NOT NULL,
  PRIMARY KEY (`id`),
  KEY `index_devices_using_histories_on_user_id_and_device_id` (`user_id`,`device_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;