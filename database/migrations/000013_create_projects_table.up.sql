-- Set timezone
SET time_zone = '+07:00';

-- Name: projects; Type: TABLE; Schema: public; Owner: -

CREATE TABLE IF NOT EXISTS `projects` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `code` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `project_type` int NOT NULL DEFAULT '1',
  `client_id` bigint DEFAULT NULL,
  `jira_id` bigint DEFAULT NULL,
  `sprint_duration` int DEFAULT NULL,
  `description` longtext COLLATE utf8mb4_unicode_ci,
  `current_sprint_id` bigint DEFAULT NULL,
  `project_priority` int NOT NULL DEFAULT '2',
  `setting` json DEFAULT NULL,
  `state` bigint NOT NULL DEFAULT '1',
  `actived_at` datetime(6) DEFAULT NULL,
  `inactived_at` datetime(6) DEFAULT NULL,
  `started_at` datetime(6) DEFAULT NULL,
  `ended_at` datetime(6) DEFAULT NULL,
  `created_at` datetime(6) NOT NULL,
  `updated_at` datetime(6) NOT NULL,
  `lock_version` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `index_projects_on_code` (`code`),
  UNIQUE KEY `index_projects_on_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;