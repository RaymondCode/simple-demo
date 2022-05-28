# simple-demo

## 抖音项目服务端

### 点赞功能更新

使用Redis实现，数据暂时没有同步到MySQL的video_favorite，Redis我放的是服务器上了，不用改配置了

#### 点赞操作

在Redis中维护了两个set

+ 以userId为key, value为此用户点赞的视频id集合；
+ 以videoId为key，value为点赞此视频的用户id集合

#### 点赞列表

从Redis中查询用户点赞的视频id集合，再对MySQL进行批量查询

#### 其他更新

* 抽取配置文件config.yaml, 配置信息启动时加载到config目录下的对应的struct，包括MySQL、Redis、JWT的配置       [利用Viper实现，参考链接](https://www.topgoer.cn/docs/goday/goday-1crg2dneqeek8)
* 使用JWT生成token，在gin中添加了JWT中间件（middleware目录）      [gin中使用jwt](https://www.cnblogs.com/supery007/p/13724121.html)
* initialize 目录下进行一些初始化，包括Config、MySQL、Redis、Router
* global 目录下存放一些公共变量，包括Config、MySQL的连接、Redis的连接
* utils 目录下抽取一些工具方法
  * **`utils.GetUserId(c *gin.Context)` 可以获取token中的userId**
  * **`utils.IsFavorite(userId, videoId) ` 判断用户是否点赞了该视频**
* service下存放业务操作，repository下直接操作数据库
* model下存放模型与表对应，model.response下存放返回的response结构
* 为了测试方便, 我在demo_data.go的根据model下的模型，定义了user1,user2,以及videos数据，登录接口直接利用user1.Id生成token并返回，用户信息接口也是直接返回user1的信息，视频流接口直接返回videos
* 数据库SQL我也放上来了，之前文档有更新，表结构有改动