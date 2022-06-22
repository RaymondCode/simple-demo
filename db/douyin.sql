/*
 Navicat Premium Data Transfer

 Source Server         : douyin
 Source Server Type    : MySQL
 Source Server Version : 80029
 Source Host           : xinzf0520.top:20607
 Source Schema         : douyin

 Target Server Type    : MySQL
 Target Server Version : 80029
 File Encoding         : 65001

 Date: 11/06/2022 17:17:38
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NULL DEFAULT NULL COMMENT '用户id',
  `video_id` int NULL DEFAULT NULL COMMENT '视频id',
  `status` tinyint NULL DEFAULT NULL COMMENT '评论状态',
  `content` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '内容',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_video_id_status`(`video_id`, `status`) USING BTREE COMMENT '获取某个视频的所有评论'
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for favorite
-- ----------------------------
DROP TABLE IF EXISTS `favorite`;
CREATE TABLE `favorite`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NULL DEFAULT NULL COMMENT '用户id',
  `video_id` int NULL DEFAULT NULL COMMENT '视频id',
  `status` tinyint NULL DEFAULT NULL COMMENT '点赞状态',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_user_id_status_video_id`(`user_id`, `status`, `video_id`) USING BTREE COMMENT '获取某个用户所有点赞视频id'
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for follow
-- ----------------------------
DROP TABLE IF EXISTS `follow`;
CREATE TABLE `follow`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NULL DEFAULT NULL COMMENT '用户id',
  `followed_user` int NULL DEFAULT NULL COMMENT '被关注者id',
  `status` tinyint NULL DEFAULT 1 COMMENT '关注状态',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_user_id_status_followed_user`(`user_id`, `status`, `followed_user`) USING BTREE COMMENT '查询关注列表',
  INDEX `idx_followed_user_status_user_id`(`followed_user`, `status`, `user_id`) USING BTREE COMMENT '查询粉丝列表'
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '姓名',
  `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '密码',
  `follow_count` int NULL DEFAULT NULL COMMENT '关注数',
  `follower_count` int NULL DEFAULT NULL COMMENT '粉丝数',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_name`(`name`) USING BTREE COMMENT '根据用户名查询信息（如注册）'
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for video
-- ----------------------------
DROP TABLE IF EXISTS `video`;
CREATE TABLE `video`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `author_id` int NULL DEFAULT NULL COMMENT '作者id',
  `play_url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '播放路径',
  `cover_url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '封面路径',
  `favorite_count` int NULL DEFAULT NULL COMMENT '喜欢数',
  `comment_count` int NULL DEFAULT NULL COMMENT '评论数',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_author_id`(`author_id`) USING BTREE COMMENT '根据用户id查询所有发布的视频'
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
