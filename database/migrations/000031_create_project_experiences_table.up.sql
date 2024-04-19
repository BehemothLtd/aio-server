-- Set timezone
SET time_zone = '+07:00';

-- Name: project_experiences; Type: TABLE; Schema: public; Owner: -

CREATE TABLE IF NOT EXISTS `project_experiences` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `description` text COLLATE utf8mb4_unicode_ci,
  `user_id` bigint NOT NULL,
  `project_id` bigint NOT NULL,
  `created_at` datetime(6) NOT NULL,
  `updated_at` datetime(6) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `index_project_experiences_on_project_id` (`project_id`),
  KEY `index_project_experiences_on_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
