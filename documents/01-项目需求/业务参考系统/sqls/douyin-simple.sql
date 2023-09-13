-- 创建数据库
CREATE DATABASE IF NOT EXISTS douyin_simple;

-- 使用数据库
USE douyin_simple;

-- 创建 users 表
CREATE TABLE IF NOT EXISTS users (
  id BIGINT PRIMARY KEY COMMENT '用户ID',
  name VARCHAR(255) NOT NULL COMMENT '用户名',
  password VARCHAR(255) NOT NULL COMMENT '密码'
);

-- 创建 videos 表
CREATE TABLE IF NOT EXISTS videos (
  id BIGINT PRIMARY KEY COMMENT '视频ID',
  author_id BIGINT NOT NULL COMMENT '作者ID',
  play_url VARCHAR(255) NOT NULL COMMENT '播放链接',
  cover_url VARCHAR(255) NOT NULL COMMENT '封面链接',
  title VARCHAR(255) NOT NULL COMMENT '标题',
  publish_time DATETIME NOT NULL COMMENT '发布时间戳',
  FOREIGN KEY (author_id) REFERENCES users(id)
);

-- 创建 follows 表
CREATE TABLE IF NOT EXISTS follows (
  user_id BIGINT NOT NULL COMMENT '用户ID',
  follow_user_id BIGINT NOT NULL COMMENT '被关注的用户ID',
  PRIMARY KEY (user_id, follow_user_id),
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  FOREIGN KEY (follow_user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- 创建 comments 表
CREATE TABLE IF NOT EXISTS comments (
  id BIGINT PRIMARY KEY COMMENT '评论ID',
  user_id BIGINT NOT NULL COMMENT '用户ID',
  video_id BIGINT NOT NULL COMMENT '视频ID',
  content TEXT NOT NULL COMMENT '评论内容',
  create_date DATETIME NOT NULL COMMENT '创建日期',
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  FOREIGN KEY (video_id) REFERENCES videos(id) ON DELETE CASCADE
);

-- 创建 likes 表
CREATE TABLE IF NOT EXISTS likes (
  id BIGINT NOT NULL COMMENT '主键ID',
  user_id BIGINT NOT NULL COMMENT '点赞者ID',
  video_id BIGINT NOT NULL COMMENT '视频ID',
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  FOREIGN KEY (video_id) REFERENCES videos(id) ON DELETE CASCADE
);