# 数据库设计与技术选型

[TOC]

## MySQL数据库字段

> MySQL包含5个表，分别是User、Video、Comment、Favorite、Follow，分别记录用户信息、视频信息、评论信息、点赞信息、关注信息

## 数据库表

### User

| 字段             | 类型       | 备注   |
| -------------- | -------- | ---- |
| id             | int      |      |
| name           | varchar  | 姓名   |
| follow_count   | Int      | 关注数  |
| follower_count | Int      | 粉丝数  |
| create_time    | datetime | 创建时间 |

### Video

| 字段             | 类型       | 备注   |
| -------------- | -------- | ---- |
| id             | int      |      |
| author_id      | Int      | 姓名   |
| play_url       | varcher  | 播放地址 |
| cover_url      | varcher  | 封面地址 |
| favorite_count | Int      | 点赞数  |
| comment_count  | int      | 评论数  |
| created_at     | datetime | 创建时间 |

### Comment

| 字段         | 类型       | 备注   |
| ---------- | -------- | ---- |
| id         | int      |      |
| user_id    | Int      | 用户id |
| video_id   | int      | 视频id |
| content    | text     | 评论内容 |
| created_at | datetime | 创建时间 |

### Favorite

| 字段         | 类型       | 备注        |
| ---------- | -------- | --------- |
| id         | int      |           |
| user_id    | Int      | 用户id      |
| video_id   | int      | 视频id      |
| status     | tinyint  | 是否点赞（1/0） |
| created_at | datetime | 创建时间      |
| updated_at | datetime | 修改时间      |

### Follow

| 字段            | 类型       | 备注     |
| ------------- | -------- | ------ |
| id            | int      |        |
| user_id       | Int      | 用户id   |
| followed_user | int      | 被关注者id |
| status        | tinyint  | 关注状态   |
| created_at    | datetime | 创建时间   |
| updated_at    | datetime | 修改时间|

## MongoDB数据库字段

> MongoDB数据库主要是用于记录用户之间的关系，即关注和被关注，考虑到在大量用户的场景下，使用MySQL记录和查询用户关系的开销比较大且不够方便，所以考虑使用MongoDB来存放用户之间的关系
>
> 查询时，直接取出一个用户的关注和被关注列表

### 关注列表

```json
{
    "user_id": "1",
    "follow_list": [
        {
            "user_id": "2",
            "name": "2号用户"
        },
        {
            "user_id": "3",
            "name": "3号用户"
        }
    ]
}
```

### 被关注列表

```json
{
    "user_id": "1",
    "follower_list": [
        {
            "user_id": "2",
            "name": "2号用户"
        },
        {
            "user_id": "3",
            "name": "3号用户"
        }
    ]
}
```

## Redis数据库字段（MongoDB数据库的另一个备选方案）

> 上述使用MongoDB所完成的用户关系记录也可以考虑使用Redis来完成；Redis和方案相比Mongo方案，查询速度应该是更快的，但是相比之下，查看关注列表和被关注列表可能不是特别常用的操作，这些数据长期占用Redis的内存可能有些浪费。
>
> 由此也可能有一个更扩展的做法，先从MongoDB读取列表，然后存放在Redis中，分批次加载给用户（一般场景可能是这样的：用户通常不会查看关注列表，若要查看关注列表，则可能会查看多个关注列表）

### follow库

```json
{
    "[user_id]": [
        {
            "user_id": "2",
            "name": "2号用户"
        },
        {
            "user_id": "3",
            "name": "3号用户"
        }
    ]
}
```

### follower库

```json
{
    "[user_id]": [
        {
            "user_id": "2",
            "name": "2号用户"
        },
        {
            "user_id": "3",
            "name": "3号用户"
        }
    ]
}
```

## 接口字段

|接口|依赖的表|
|:---:|:---:|
|douyin/feed|user, video|
|douyin/user/register|user|
|douyin/user/login|user|
|douyin/user|user|
|douyin/publish/action|user, video|
|douyin/publish/list|user, video|
|douyin/favorite/action|video|
|douyin/favorite/list|video, user|
|douyin/comment/action|comment, user|
|douyin/comment/list|comment, user|
|douyin/relation/action||
|douyin/relation/follow/list||
|douyin/relation/follower/list||
