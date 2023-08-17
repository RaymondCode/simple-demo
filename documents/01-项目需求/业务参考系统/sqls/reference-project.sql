-- 使用数据库
USE douyin_simple;

-- 插入 users 表数据
INSERT INTO users (id, name, password) VALUES
(1, 'user_1', 'password_1'),
(2, 'user_2', 'password_2'),
(3, 'user_3', 'password_3'),
(4, 'user_4', 'password_4'),
(5, 'user_5', 'password_5');

-- 插入 videos 表数据
INSERT INTO videos (id, author_id, play_url, cover_url, title, publish_time) VALUES
(1, 1, 'play_url_1', 'cover_url_1', 'title_1', '2022-01-01 10:00:00'),
(2, 1, 'play_url_2', 'cover_url_2', 'title_2', '2022-02-01 14:30:00'),
(3, 2, 'play_url_3', 'cover_url_3', 'title_3', '2022-03-12 09:15:00'),
(4, 3, 'play_url_4', 'cover_url_4', 'title_4', '2022-04-25 18:45:00'),
(5, 4, 'play_url_5', 'cover_url_5', 'title_5', '2022-05-09 21:20:00');

-- 插入 follows 表数据
INSERT INTO follows (user_id, follow_user_id) VALUES
(1, 2),
(1, 3),
(2, 4),
(3, 1),
(4, 5);

-- 插入 comments 表数据
INSERT INTO comments (id, user_id, video_id, content, create_date) VALUES
(1, 2, 1, 'comment_1', '2022-01-01 12:30:00'),
(2, 3, 1, 'comment_2', '2022-01-02 09:45:00'),
(3, 1, 3, 'comment_3', '2022-03-14 13:10:00'),
(4, 5, 2, 'comment_4', '2022-02-02 16:20:00');

-- 插入 likes 表数据
INSERT INTO likes (id, user_id, video_id) VALUES
(1, 1, 2),
(2, 2, 1),
(3, 3, 3),
(4, 4, 5),
(5, 5, 4);
