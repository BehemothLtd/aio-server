-- Set timezone
SET time_zone = '+07:00';

-- Name: user_groups_permissions; Type: TABLE; Schema: public; Owner: -

CREATE TABLE IF NOT EXISTS `user_groups_permissions` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_group_id` bigint NOT NULL,
  `permission_id` bigint NOT NULL,
  `active` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` datetime(6) NOT NULL,
  `updated_at` datetime(6) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `index_user_groups_permissions_on_user_group_id` (`user_group_id`)
) ENGINE=InnoDB AUTO_INCREMENT=92 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
