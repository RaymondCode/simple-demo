/*
 Navicat Premium Data Transfer

 Source Server Type    : MySQL
 Source Server Version : 100703
 Source Host           : fitenne.com:3306
 Source Schema         : dousheng

 Target Server Type    : MySQL
 Target Server Version : 100703
 File Encoding         : 65001

 Date: 21/05/2022 18:28:14
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for videos
-- ----------------------------
DROP TABLE IF EXISTS `videos`;
CREATE TABLE `videos`  (
   `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '短视频id',
   `author_id` bigint NOT NULL COMMENT '作者id',
   `play_url` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '短视频url',
   `cover_url` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '封面url',
   `favorite_count` bigint NOT NULL DEFAULT 0 COMMENT '点赞数',
   `comment_count` bigint NOT NULL DEFAULT 0 COMMENT '评论数',
   `created_at` bigint NOT NULL COMMENT '投递时间',
   `deleted_at` datetime(3) default NULL COMMENT '删除标记位',
   PRIMARY KEY (`id`) USING BTREE,
   #created_at聚簇索引
   UNIQUE INDEX `create_time_index`(`created_at`, `deleted_at`, `id`, `author_id`, `play_url`, `cover_url`, `favorite_count`, `comment_count`) USING BTREE,
   #deleted_at普通索引
   INDEX `idx_videos_deleted_at`(`deleted_at`) USING BTREE,
   #author_id普通索引
   INDEX `fk_videos_author`(`author_id`) USING BTREE,
   #users.id->videos.author_id外键约束
   CONSTRAINT `fk_videos_author` FOREIGN KEY (`author_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
