-- Set timezone
SET time_zone = '+07:00';

-- Name: attachments; Type: TABLE; Schema: public; Owner: -

CREATE TABLE IF NOT EXISTS `attachments` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `attachment_blob_id` bigint DEFAULT NULL,
  `owner_id` bigint DEFAULT NULL,
  `owner_type` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` datetime(6) NOT NULL,
  `updated_at` datetime(6) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `index_attachments_on_attachment_blob_id` (`attachment_blob_id`),
  KEY `index_attachments_on_owner_id` (`owner_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;