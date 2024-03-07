-- Set timezone
SET time_zone = '+07:00';

-- Name: project_issue_statuses; Type: TABLE; Schema: public; Owner: -

CREATE TABLE IF NOT EXISTS `project_issue_statuses` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `project_id` bigint NOT NULL,
  `issue_status_id` bigint NOT NULL,
  `position` int NOT NULL DEFAULT '1',
  `created_at` datetime(6) NOT NULL,
  `updated_at` datetime(6) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `index_project_issue_statuses_on_project_id_and_issue_status_id` (`project_id`,`issue_status_id`)
) ENGINE=InnoDB AUTO_INCREMENT=351 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;