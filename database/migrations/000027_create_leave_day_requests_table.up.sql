-- Set timezone
SET time_zone = '+07:00';

-- Name: leave_day_requests; Type: TABLE; Schema: public; Owner: -

CREATE TABLE IF NOT EXISTS `leave_day_requests` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `approver_id` bigint DEFAULT NULL,
  `from` datetime(6) NOT NULL,
  `to` datetime(6) NOT NULL,
  `time_off` float NOT NULL,
  `reason` text COLLATE utf8mb4_unicode_ci,
  `request_type` int NOT NULL,
  `request_state` int DEFAULT '1',
  `lock_version` int DEFAULT NULL,
  `created_at` datetime(6) NOT NULL,
  `updated_at` datetime(6) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_leave_day_time_range` (`user_id`,`from`,`to`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
