-- Set timezone
SET time_zone = '+07:00';

-- Name: working_timelogs; Type: TABLE; Schema: public; Owner: -

CREATE TABLE IF NOT EXISTS `working_timelogs` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `project_id` bigint NOT NULL,
  `issue_id` bigint NOT NULL,
  `minutes` int NOT NULL,
  `description` text COLLATE utf8mb4_unicode_ci,
  `logged_at` date NOT NULL,
  `created_at` datetime(6) NOT NULL,
  `updated_at` datetime(6) NOT NULL,
  `lock_version` int DEFAULT '0', 
  PRIMARY KEY (`id`),
  KEY `index_working_timelogs_on_logged_at` (`logged_at`),
  KEY `index_working_timelogs_on_project_id_and_issue_id` (`project_id`,`issue_id`),
  KEY `index_working_timelogs_on_project_id_and_logged_at` (`project_id`,`logged_at`),
  KEY `index_working_timelogs_on_project_id` (`project_id`),
  KEY `index_working_timelogs_on_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;