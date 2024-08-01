-- Set timezone
SET time_zone = '+07:00';

-- Name: messages; Type: TABLE; Schema: public; Owner: -

CREATE TABLE IF NOT EXISTS `messages` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `leave_day_request_id` bigint NOT NULL,
  `content` text COLLATE utf8mb4_unicode_ci,
  `timestamp` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` datetime(6) NOT NULL,
  `updated_at` datetime(6) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `index_on_leave_day_request_id` (`leave_day_request_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
