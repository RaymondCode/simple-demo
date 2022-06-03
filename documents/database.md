# 数据库设计与技术选型

[TOC]

## MySQL数据库字段

> MySQL包含3个表，分别是User表、Video表、Comment表，分别记录用户信息、视频信息、评论信息

### User表

|字段|类型|备注|
|:---:|:---:|:---:|
|id|Integer|Video id|
|name|varchar||
|follow_count|int|关注数，要不要用字符串？|
|follower_count|int|粉丝数|
|is_follow|book|是否关注？不清楚什么用|

### Video表

|字段|类型|备注|
|:---:|:---:|:---:|
|id|Integer|Video id|
|user_id|Integer|foreign key|
|play_url|varchar||
|cover_url|varchar||
|favorite_count|int|点赞数|
|comment_count|int|评论数|
|is_favorite|bool||
|title|varchar||

### Comment表

|字段|类型|备注|
|:---:|:---:|:---:|
|id|Integer||
|user_id|Integer|foreign key|
|content|text||
|create_date|date|auto create|

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
