# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.7.16)
# Database: yehelaoren
# Generation Time: 2019-11-25 02:58:55 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table category
# ------------------------------------------------------------

DROP TABLE IF EXISTS `category`;

CREATE TABLE `category` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL,
  `description` varchar(512) DEFAULT NULL,
  `status` tinyint(1) unsigned NOT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table comment_like
# ------------------------------------------------------------

DROP TABLE IF EXISTS `comment_like`;

CREATE TABLE `comment_like` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) unsigned NOT NULL,
  `comment_id` bigint(20) unsigned NOT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_comment_id` (`comment_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table favorite
# ------------------------------------------------------------

DROP TABLE IF EXISTS `favorite`;

CREATE TABLE `favorite` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) unsigned NOT NULL,
  `topic_id` bigint(20) unsigned NOT NULL,
  `title` varchar(20) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_topic_id` (`topic_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table lunar_calendar
# ------------------------------------------------------------

DROP TABLE IF EXISTS `lunar_calendar`;

CREATE TABLE `lunar_calendar` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `day` varchar(10) DEFAULT '' COMMENT '日期20190928',
  `lunar_year` varchar(5) DEFAULT '' COMMENT '农历年',
  `lunar_month` varchar(5) DEFAULT '' COMMENT '农历月',
  `lunar_day` varchar(5) DEFAULT '' COMMENT '农历日',
  `suit` varchar(200) DEFAULT '' COMMENT '适宜',
  `taboo` varchar(50) DEFAULT '' COMMENT '禁忌',
  `jieqi` varchar(50) DEFAULT '' COMMENT '节气',
  `created_at` datetime NOT NULL COMMENT '时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_day` (`day`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='日历';



# Dump of table message
# ------------------------------------------------------------

DROP TABLE IF EXISTS `message`;

CREATE TABLE `message` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `from_id` bigint(20) unsigned NOT NULL,
  `user_id` bigint(20) unsigned NOT NULL,
  `content` text NOT NULL,
  `obj_type` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `obj_id` int(10) unsigned NOT NULL DEFAULT '0',
  `status` tinyint(3) unsigned NOT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table sign_in
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sign_in`;

CREATE TABLE `sign_in` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `days` int(10) unsigned NOT NULL DEFAULT '0',
  `experience` int(10) unsigned NOT NULL DEFAULT '0',
  `day` varchar(10) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table sys_config
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_config`;

CREATE TABLE `sys_config` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `key` varchar(128) NOT NULL,
  `value` text,
  `name` varchar(32) NOT NULL,
  `description` varchar(128) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `key` (`key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table topic
# ------------------------------------------------------------

DROP TABLE IF EXISTS `topic`;

CREATE TABLE `topic` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL,
  `category_id` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `is_top` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `is_best` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `title` varchar(128) NOT NULL DEFAULT '',
  `content` text NOT NULL,
  `view_count` int(10) unsigned NOT NULL DEFAULT '0',
  `comment_count` int(10) unsigned NOT NULL DEFAULT '0',
  `experience` int(10) unsigned NOT NULL DEFAULT '0',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `last_comment_time` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_status` (`status`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_category_id` (`category_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table topic_comment
# ------------------------------------------------------------

DROP TABLE IF EXISTS `topic_comment`;

CREATE TABLE `topic_comment` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) unsigned NOT NULL,
  `topic_id` bigint(20) unsigned NOT NULL,
  `parent_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `is_accept` tinyint(1) unsigned NOT NULL,
  `content` text NOT NULL,
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `like_count` int(10) unsigned NOT NULL DEFAULT '0',
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_topic_id` (`topic_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table topic_tag
# ------------------------------------------------------------

DROP TABLE IF EXISTS `topic_tag`;

CREATE TABLE `topic_tag` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `topic_id` bigint(20) unsigned NOT NULL,
  `tag_id` bigint(20) unsigned NOT NULL,
  `status` tinyint(1) unsigned NOT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_topic_id` (`topic_id`),
  KEY `idx_tag_id` (`tag_id`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table user
# ------------------------------------------------------------

DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `email` varchar(128) DEFAULT NULL,
  `mobile` varchar(11) DEFAULT NULL,
  `nickname` varchar(16) DEFAULT NULL,
  `avatar` varchar(200) DEFAULT NULL,
  `password` varchar(200) DEFAULT NULL,
  `status` tinyint(4) NOT NULL,
  `roles` varchar(200) DEFAULT NULL,
  `description` varchar(512) DEFAULT NULL,
  `experience` int(10) unsigned NOT NULL DEFAULT '0',
  `topic_num` int(10) unsigned NOT NULL DEFAULT '0',
  `comment_num` int(10) unsigned NOT NULL DEFAULT '0',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_email` (`email`),
  UNIQUE KEY `uniq_mobile` (`mobile`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
