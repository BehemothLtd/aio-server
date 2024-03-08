-- Set timezone
SET time_zone = '+07:00';

-- Name: snippets_collections; Type: TABLE; Schema: public; Owner: -

CREATE TABLE IF NOT EXISTS `snippets_collections` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `snippet_id` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `collection_id` bigint NOT NULL,
  `created_at` datetime(6) NOT NULL,
  `updated_at` datetime(6) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `index_snippets_collections_on_snippet_id_and_collection_id` (`snippet_id`,`collection_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
