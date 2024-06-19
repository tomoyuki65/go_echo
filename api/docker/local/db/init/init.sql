-- 存在しない場合のみ通常DBを作成
CREATE DATABASE IF NOT EXISTS `echo-db` CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;

-- 存在しない場合のみテストDBを作成
CREATE DATABASE IF NOT EXISTS `echo-db-testing` CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;

-- 存在しない場合のみユーザーを作成
CREATE USER IF NOT EXISTS 'echo-user'@'%' IDENTIFIED BY 'echo-password';

-- DBに権限付与
GRANT ALL PRIVILEGES ON `echo-db`.* TO 'echo-user'@'%';
GRANT ALL PRIVILEGES ON `echo-db-testing`.* TO 'echo-user'@'%';

-- 権限の変更反映
FLUSH PRIVILEGES;
