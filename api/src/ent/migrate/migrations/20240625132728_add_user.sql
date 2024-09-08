-- Create "users" table
CREATE TABLE `users` (`id` bigint NOT NULL AUTO_INCREMENT, `uid` varchar(255) NOT NULL, `last_name` varchar(255) NOT NULL, `first_name` varchar(255) NOT NULL, `email` varchar(255) NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `deleted_at` datetime NULL, PRIMARY KEY (`id`), UNIQUE INDEX `uid` (`uid`), UNIQUE INDEX `email` (`email`), INDEX `user_deleted_at` (`deleted_at`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;