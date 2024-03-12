-- Set timezone
SET time_zone = '+07:00';

-- Name: issue_statuses; Type: TABLE; Schema: public; Owner: -

CREATE TABLE IF NOT EXISTS `issue_statuses` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `status_type` int NOT NULL DEFAULT '1',
  `color` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '#fff',
  `created_at` datetime(6) NOT NULL,
  `updated_at` datetime(6) NOT NULL,
  `lock_version` int NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `index_issue_statuses_on_title` (`title`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
