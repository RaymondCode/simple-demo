-- 创建数据库
CREATE DATABASE IF NOT EXISTS douyin-simple;

-- 使用数据库
USE douyin-simple;

-- 创建 Users 表
CREATE TABLE IF NOT EXISTS Users (
  id BIGINT PRIMARY KEY COMMENT '用户ID',
  name VARCHAR(255) NOT NULL COMMENT '用户名',
  follow_count BIGINT DEFAULT 0 COMMENT '关注数量',
  follower_count BIGINT DEFAULT 0 COMMENT '粉丝数量',
  is_follow BOOLEAN NOT NULL DEFAULT FALSE COMMENT '是否关注',
  avatar VARCHAR(255) COMMENT '头像',
  background_image VARCHAR(255) COMMENT '背景图片',
  signature TEXT COMMENT '个性签名',
  total_favorited BIGINT DEFAULT 0 COMMENT '总喜欢数',
  work_count BIGINT DEFAULT 0 COMMENT '作品数量',
  favorite_count BIGINT DEFAULT 0 COMMENT '收藏数量'
);

-- 创建 Videos 表
CREATE TABLE IF NOT EXISTS Videos (
  id BIGINT PRIMARY KEY COMMENT '视频ID',
  author_id BIGINT NOT NULL COMMENT '作者ID',
  play_url VARCHAR(255) NOT NULL COMMENT '播放链接',
  cover_url VARCHAR(255) NOT NULL COMMENT '封面链接',
  favorite_count BIGINT DEFAULT 0 COMMENT '喜欢数量',
  comment_count BIGINT DEFAULT 0 COMMENT '评论数量',
  is_favorite BOOLEAN NOT NULL DEFAULT FALSE COMMENT '是否喜欢',
  title VARCHAR(255) NOT NULL COMMENT '标题',
  FOREIGN KEY (author_id) REFERENCES Users(id)
);

-- 创建 Follows 表
CREATE TABLE IF NOT EXISTS Follows (
  user_id BIGINT NOT NULL,
  follow_user_id BIGINT NOT NULL,
  PRIMARY KEY (user_id, follow_user_id),
  FOREIGN KEY (user_id) REFERENCES Users(id) ON DELETE CASCADE,
  FOREIGN KEY (follow_user_id) REFERENCES Users(id) ON DELETE CASCADE
);

-- 创建 Fans 表
CREATE TABLE IF NOT EXISTS Fans (
  user_id BIGINT NOT NULL,
  follower_user_id BIGINT NOT NULL,
  PRIMARY KEY (user_id, follower_user_id),
  FOREIGN KEY (user_id) REFERENCES Users(id) ON DELETE CASCADE,
  FOREIGN KEY (follower_user_id) REFERENCES Users(id) ON DELETE CASCADE
);

-- 创建 Comments 表
CREATE TABLE IF NOT EXISTS Comments (
  id BIGINT PRIMARY KEY COMMENT '评论ID',
  user_id BIGINT NOT NULL COMMENT '用户ID',
  content TEXT NOT NULL COMMENT '评论内容',
  create_date DATE NOT NULL COMMENT '创建日期',
  FOREIGN KEY (user_id) REFERENCES Users(id) ON DELETE CASCADE
);

-- 创建 Messages 表
CREATE TABLE IF NOT EXISTS Messages (
  id BIGINT PRIMARY KEY COMMENT '消息ID',
  to_user_id BIGINT NOT NULL COMMENT '接收用户ID',
  from_user_id BIGINT NOT NULL COMMENT '发送用户ID',
  content TEXT NOT NULL COMMENT '消息内容',
  create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  FOREIGN KEY (to_user_id) REFERENCES Users(id) ON DELETE CASCADE,
  FOREIGN KEY (from_user_id) REFERENCES Users(id) ON DELETE CASCADE
);
