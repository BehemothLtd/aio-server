-- Set timezone
SET time_zone = '+07:00';
-- Name: snippets; Type: TABLE; Schema: public; Owner: -
CREATE TABLE IF NOT EXISTS `snippets` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `content` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `user_id` bigint NOT NULL,
  `slug` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `snippet_type` int NOT NULL DEFAULT '0',
  `favorites_count` int NOT NULL DEFAULT '0',
  `lock_version` int DEFAULT NULL,
  `created_at` datetime(6) NOT NULL,
  `updated_at` datetime(6) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=51 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;