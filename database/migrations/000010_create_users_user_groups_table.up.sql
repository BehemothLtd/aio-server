-- Set timezone
SET time_zone = '+07:00';

-- Name: users_user_groups; Type: TABLE; Schema: public; Owner: -

CREATE TABLE IF NOT EXISTS `users_user_groups` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `user_group_id` bigint NOT NULL,
  `created_at` datetime(6) NOT NULL,
  `updated_at` datetime(6) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `index_users_user_groups_on_user_id_and_user_group_id` (`user_id`,`user_group_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;