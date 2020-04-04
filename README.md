# award
A small but complete lottery system

### Quick Start

在程序运行前，你需要先安装 redis 和 mysql

#### 1、安装 redis

redis 的安装可以参考：https://www.runoob.com/redis/redis-install.html

安装完成后，启动 redis，默认会在 127.0.0.1:6379 端口监听

#### 2、安装并部署 mysql

mysql 的安装可以参考：https://www.runoob.com/mysql/mysql-install.html

安装完成后，需要创建相应库表

```
create database award;
use award;

CREATE TABLE `award_user_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `award_name` varchar(100) NOT NULL,
  `user_name` varchar(100) NOT NULL,
  `award_time` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9139 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

```

#### 3、更改 conf/config.toml

修改奖品发放的开始、结束时间，奖品总数、mysql 的账号密码

dsn 格式为：用户名:密码@/数据库名

```
[award]
startTime="2020-04-04 19:00:00"  ## 奖品发放开始时间
endTime="2020-04-04 20:00:00"   ## 奖品发放结束时间


[awardMap]
A=20
B=200
C=500

[mysql]
dsn="root:diuge123456@/award"

[redis]
network="tcp"
ip="127.0.0.1"
port=6379
```

#### 4、RUN

```bash
git clone https://github.com/lubanproj/award
go build -v
./award
```

程序会在 8080 端口监听 http 请求，打开浏览器，访问

http://localhost:8080/draw?username=lubanproj

即可看到中奖情况
