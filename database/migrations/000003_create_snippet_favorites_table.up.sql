-- Set timezone
SET time_zone = '+07:00';
-- Name: snippet_favorites; Type: TABLE; Schema: public; Owner: -

CREATE TABLE `snippets_favorites` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `snippet_id` bigint NOT NULL,
  `user_id` bigint NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `index_snippets_favorites_on_snippet_id_and_user_id` (`snippet_id`,`user_id`),
  KEY `index_snippets_favorites_on_snippet_id` (`snippet_id`),
  KEY `index_snippets_favorites_on_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
