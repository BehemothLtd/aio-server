-- Set timezone
SET time_zone = '+07:00';

-- Name: attendances; Type: TABLE; Schema: public; Owner: -

CREATE TABLE IF NOT EXISTS `attendances` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `checkin_at` datetime(6) DEFAULT NULL,
  `checkout_at` datetime(6) DEFAULT NULL,
  `created_user_id` bigint NOT NULL,
  `created_at` datetime(6) NOT NULL,
  `updated_at` datetime(6) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `index_attendances_on_checkin_at` (`checkin_at`),
  KEY `index_attendances_on_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;