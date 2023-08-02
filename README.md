# Momotok-Server

## 抖音项目服务端

具体功能内容参考飞书说明文档

工程无其他依赖，直接编译运行即可

```shell
go build && ./Momotok-Server
```

### 功能说明

接口功能正在开发中

* 用户登录数据保存在MySQL数据库中
* 视频上传后会保存到本地 public 目录中，访问时用 127.0.0.1:8080/static/video_name 即可
* 
### 目前进度
- [x] basic apis
  - [x] controller.Feed
  - [x] controller.UserInfo
  - [x] controller.Register
  - [x] controller.Login
  - [x] controller.Publish
  - [x] controller.PublishList
- [ ] extra apis - I
  - [ ] apiRouter.POST("/favorite/action/", controller.FavoriteAction)
  - [ ] apiRouter.GET("/favorite/list/", controller.FavoriteList)
  - [ ] apiRouter.POST("/comment/action/", controller.CommentAction)
  - [ ] apiRouter.GET("/comment/list/", controller.CommentList)
- [ ] extra apis - II
  -[ ] apiRouter.POST("/relation/action/", controller.RelationAction)
  -[ ] apiRouter.GET("/relation/follow/list/", controller.FollowList)
  -[ ] apiRouter.GET("/relation/follower/list/", controller.FollowerList)
  -[ ] apiRouter.GET("/relation/friend/list/", controller.FriendList)
  -[ ] apiRouter.GET("/message/chat/", controller.MessageChat)
  -[ ] apiRouter.POST("/message/action/", controller.MessageAction) 

### 测试

test 目录下为不同场景的功能测试case，可用于验证功能实现正确性

其中 common.go 中的 _serverAddr_ 为服务部署的地址，默认为本机地址，可以根据实际情况修改

测试数据写在 demo_data.go 中，用于列表接口的 mock 测试

### 建库说明
````mysql
create table user
(
    id         int auto_increment
        primary key,
    username   varchar(50)                         not null,
    ip         varchar(15)                         null,
    password   varchar(60)                         null,
    created_at timestamp default CURRENT_TIMESTAMP not null,
    constraint name
        unique (username)
)ENGINE = InnoDB;

create table video
(
    id              int auto_increment
        primary key,
    author_id       int                                 null,
    play_url        linestring                          null,
    cover_url       linestring                          null,
    favourite_count int                                 null,
    comment_count   int                                 null,
    title           varchar(72)                         null,
    publish_time     timestamp default CURRENT_TIMESTAMP null,
    FOREIGN KEY (author_id) REFERENCES user(id)
)ENGINE = InnoDB;

create table likes
(
    id       int auto_increment
        primary key,
    video_id int                                 null,
    liked_at timestamp default CURRENT_TIMESTAMP not null,
    user_id  int                                 null,
    FOREIGN KEY (user_id) REFERENCES user(id),
    FOREIGN KEY (video_id) REFERENCES video(id)
)ENGINE = InnoDB;
````