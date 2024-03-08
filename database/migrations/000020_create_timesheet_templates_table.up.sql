SET time_zone = '+07:00';

-- Name: timesheet_templates; Type: TABLE; Schema: public; Owner: -

CREATE TABLE IF NOT EXISTS `timesheet_templates` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `user_id` bigint NOT NULL,
  `settings` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `lock_version` int DEFAULT NULL,
  `created_at` datetime(6) NOT NULL,
  `updated_at` datetime(6) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `index_timesheet_templates_on_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;