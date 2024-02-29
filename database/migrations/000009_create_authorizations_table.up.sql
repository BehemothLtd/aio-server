-- Set timezone
SET time_zone = '+07:00';

-- Name: authorizations; Type: TABLE; Schema: public; Owner: -

CREATE TABLE IF NOT EXISTS `authorizations` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_group_id` bigint NOT NULL,
  `permission_id` int NOT NULL,
  `created_at` datetime(6) NOT NULL,
  `updated_at` datetime(6) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `index_authorizations_on_user_group_id_and_permission_id` (`user_group_id`,`permission_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;