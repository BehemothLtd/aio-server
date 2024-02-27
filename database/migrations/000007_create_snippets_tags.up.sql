-- Set timezone
SET time_zone = '+07:00';

-- Name: tags; Type: TABLE; Schema: public; Owner: -

CREATE TABLE IF NOT EXISTS `snippets_tags` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `snippet_id` bigint NOT NULL,
  `tag_id` bigint NOT NULL,
  `lock_version` int DEFAULT NULL,
  `created_at` datetime(6) NOT NULL,
  `updated_at` datetime(6) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `index_snippets_tags_on_snippet_id_and_tag_id` (`snippet_id`,`tag_id`),
  KEY `index_snippets_tags_on_snippet_id` (`snippet_id`),
  KEY `index_snippets_tags_on_tag_id` (`tag_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;