-- Set timezone
SET time_zone = '+07:00';

-- Name: project_assignees; Type: TABLE; Schema: public; Owner: -

CREATE TABLE IF NOT EXISTS `project_assignees` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `project_id` bigint NOT NULL,
  `user_id` bigint NOT NULL,
  `development_role_id` int NOT NULL DEFAULT '1',
  `active` tinyint(1) NOT NULL DEFAULT '1',
  `created_at` datetime(6) NOT NULL,
  `updated_at` datetime(6) NOT NULL,
  `join_date` date DEFAULT NULL,
  `leave_date` date DEFAULT NULL,
  `lock_version` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_pas_on_project_user_role` (`project_id`,`user_id`,`development_role_id`),
  KEY `index_project_assignees_on_project_id` (`project_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;