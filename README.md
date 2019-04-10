# Golang实现的小说爬虫数据分析

> 通过爬取起点网的注册用户信息进行数据分析

[![Travis Status for henson/proxypool](https://travis-ci.org/henson/proxypool.svg?branch=master)](https://travis-ci.org/henson/proxypool) [![Go Report Card](https://goreportcard.com/badge/distributed/henson/proxypool)](https://goreportcard.com/report/distributed/henson/proxypool)

## 版本更新

- 2019年4月10日 v1.0
  - 采用 mysql 作为数据持久化
  - 结构简洁，适合二次开发

### 1、主要结构

proxypool：
    
   - 在运行程序之前需要下载[phantomjs](http://phantomjs.org/)
　　
   - 需要自己配置conf中的app.ini文件.

```
[server]
HTTP_ADDR       = 0.0.0.0 //这个ip不要动
HTTP_PORT       = 8080  //填写一个没有被占用的端口，默认8080
; Session expires time
SESSION_EXPIRES = 168h0m0s  //默认即可

[database]
; Either "mysql", "postgres" or "sqlite3", you can connect to TiDB with MySQL protocol
DB_TYPE  = mysql  //填写你想要存储的数据库类型
HOST     = 127.0.0.1:3306  //数据库的ip及端口
NAME     = ippool   //由自己创建的数据表用来存放ip池
USER     = root    //访问数据库的用户名
PASSWD   = 56781234 //访问数据库的密码
```
- 通过上述的修改可以直接go build，通过Get `http://localhost:8080/v2/ip` 以及  `http://localhost:8080/v2/https` 获取随机ip
- 如果编译报错，请查看[proxypool开源网址](https://github.com/henson/ProxyPool)


reptile：

- 爬虫实现，使用框架goquery
- 数据库mysql
- ip池自带检查ip有效性有一定疏漏，在ippool函数中加入异常处理机制，保证ip有效性
- 破解反爬字体使用工具[fonttools](https://github.com/fonttools/fonttools)


web：
- 实现简单的web服务端
- 提供接口输出json数据
- 简单的前端代码


### 2、代码实现
- crack：

　　存放字体破解文件

- engine：

　　相关的处理机制代码

- debugs：

　　是否开启调试模式

- 其他：

　　xorm生成数据库表

### 3、安装及使用
本项目用到的依赖库包括有：
```
github.com/go-clog/clog
github.com/go-ini/ini
github.com/go-xorm/xorm
github.com/go-xorm/core
github.com/go-sql-driver/mysql
github.com/lib/pq
github.com/Aiicy/htmlquery
github.com/PuerkitoBio/goquery
github.com/parnurzeal/gorequest
github.com/nladuo/go-phantomjs-fetcher
github.com/donnie4w/dom4g
```

下载本项目：
```
go get -u *
```

按步骤执行
1.创建mysql数据表
2.启动爬虫爬取数据
3.启动web服务端
4.启动html静态页面


### 5、容错处理

代码具备基本的容错处理，如果效果不如意可以自行添加。

### 6、诚挚的感谢

- 首先感谢您的使用，如果觉得程序还不错也能帮助您解决实际问题，不妨添个赞以鼓励本人继续努力，谢谢！
- 如果您对程序有任何建议和意见，也欢迎提交issue。
- 当然，如果您愿意贡献代码和我一起改进本程序，那再好不过了。
