SET time_zone = '+07:00';

-- Name: issues; Type: TABLE; Schema: public; Owner: -

CREATE TABLE IF NOT EXISTS `issues` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `project_id` bigint NOT NULL,
  `jira_id` bigint DEFAULT NULL,
  `issue_type` int NOT NULL DEFAULT '1',
  `parent_id` bigint DEFAULT NULL,
  `title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `description` text COLLATE utf8mb4_unicode_ci,
  `code` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `priority` int NOT NULL DEFAULT '3',
  `issue_status_id` int DEFAULT '1',
  `position` int DEFAULT '1',
  `project_sprint_id` bigint DEFAULT NULL,
  `start_date` date DEFAULT NULL,
  `deadline` date DEFAULT NULL,
  `archived` tinyint(1) DEFAULT '0',
  `creator_id` bigint NOT NULL,
  `lock_version` int DEFAULT '0',
  `data` json DEFAULT NULL,
  `created_at` datetime(6) NOT NULL,
  `updated_at` datetime(6) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `index_issues_on_project_id_and_code` (`project_id`,`code`),
  KEY `index_issues_on_deadline` (`deadline`),
  KEY `index_issues_on_project_id_and_deadline` (`project_id`,`deadline`),
  KEY `index_issues_on_project_id` (`project_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1501 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;