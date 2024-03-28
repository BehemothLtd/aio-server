SET time_zone = '+07:00';

-- Name: project_sprints; Type: TABLE; Schema: public; Owner: -
CREATE TABLE IF NOT EXISTS `project_sprints` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `project_id` bigint NOT NULL,
  `start_date` date DEFAULT NULL,
  `end_date` date DEFAULT NULL,
  `archived` tinyint(1) DEFAULT '0',
  `created_at` datetime(6) NOT NULL,
  `updated_at` datetime(6) NOT NULL,
  `lock_version` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `index_project_sprints_on_title_and_project_id` (`title`,`project_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;