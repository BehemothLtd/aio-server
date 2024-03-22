-- Set timezone
SET time_zone = '+07:00';

-- Name: issue_assignees; Type: TABLE; Schema: public; Owner: -

CREATE TABLE IF NOT EXISTS `issue_assignees` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `issue_id` bigint NOT NULL,
  `user_id` bigint NOT NULL,
  `development_role_id` int NOT NULL DEFAULT '1',
  `created_at` datetime(6) NOT NULL,
  `updated_at` datetime(6) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `index_issue_assignees_uniqueness` (`issue_id`,`user_id`,`development_role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;