-- Set timezone
SET time_zone = '+07:00';

-- Name: pins; Type: TABLE; Schema: public; Owner: -

CREATE TABLE IF NOT EXISTS `pins` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `parent_type` bigint NOT NULL,
  `parent_id` bigint NOT NULL,
  `created_at` datetime(6) NOT NULL,
  `updated_at` datetime(6) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `index_on_parent` (`user_id`,`parent_id`,`parent_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
