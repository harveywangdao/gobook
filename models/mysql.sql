mysql -u root -p
create database gobook;
show databases;
use gobook;

CREATE TABLE IF NOT EXISTS `article_tb` (
    `id` bigint AUTO_INCREMENT NOT NULL PRIMARY KEY COMMENT 'index',
    `user_id` bigint NOT NULL DEFAULT 0  COMMENT 'userId',
    `title` varchar(100) NOT NULL DEFAULT ''  COMMENT 'title',
    `text` varchar(10000) NOT NULL DEFAULT ''  COMMENT 'text',
    `tags` varchar(100) NOT NULL DEFAULT ''  COMMENT 'tags',
    `createTime` datetime NOT NULL COMMENT 'createTime'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `user_tb` (
    `id` bigint AUTO_INCREMENT NOT NULL PRIMARY KEY COMMENT 'index',
    `email` varchar(100) NOT NULL DEFAULT ''  UNIQUE COMMENT 'email',
    `token` varchar(100) NOT NULL DEFAULT ''  COMMENT 'token',
    `createTime` datetime NOT NULL COMMENT 'createTime'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

exit