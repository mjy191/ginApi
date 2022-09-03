## 1.基本介绍
### 1.1 项目介绍
> ginApi是一个基于gin开发的框架，集成了api签名、鉴权、全链路日志追踪，减少在项目中重复造轮子
### 1.2 链路日志追踪
通过logid可以追踪一次http请求的日志，logid均在接口中返回

日志是按照小时切割，接口中请求和返回的数据均通过中间件<span style="color:green">LoggerMiddleware.go</span>自动记录，mysql的操作sql语句也自动记录

例如：用户请求登录/api/user/login

返回如下数据

```
{
	"code": 1,
	"data": {
		"token": "PqrUdnsC0kE7VARhqoMGY5Jm4zwkfixr"
	},
	"logid": "cd79e47c610343cbb3345c8603e7dacd",
	"msg": "success"
}
```
进入的日志目录，根据返回的logid<span style="color:red">cd79e47c610343cbb3345c8603e7dacd</span>查找，一次http请求的日志
```
[root@localhost 202209]# cd logs/202209
[root@localhost 202209]# grep cd79e47c610343cbb3345c8603e7dacd 2022090312.log 
2022/09/03 12:11:42 logid[cd79e47c610343cbb3345c8603e7dacd]  url[/api/user/login?sign=f9e9eb19af3aa19f0fe83bc5b43742c9ec65dc57] ip[127.0.0.1] method[POST] post_data[] body[{"username": "zhangsan", "password": "123456", "token": "A5tvtvTZR56Vekj8K7V71I25OpVuQ8jg"}]
2022/09/03 12:11:42 logid[cd79e47c610343cbb3345c8603e7dacd]  url[/api/user/login?sign=f9e9eb19af3aa19f0fe83bc5b43742c9ec65dc57] signPre[abc123{"username": "zhangsan", "password": "123456", "token": "A5tvtvTZR56Vekj8K7V71I25OpVuQ8jg"}abc123]
2022/09/03 12:11:42 logid[cd79e47c610343cbb3345c8603e7dacd]  mysql[SELECT * FROM `user` WHERE username='zhangsan' ORDER BY `user`.`id` LIMIT 1] time[9.4602] filePath[D:/www/go/ginApi/Service/UserService.go:130]
2022/09/03 12:11:42 logid[cd79e47c610343cbb3345c8603e7dacd]  url[/api/user/login?sign=f9e9eb19af3aa19f0fe83bc5b43742c9ec65dc57] response[{"code":1,"data":{"token":"PqrUdnsC0kE7VARhqoMGY5Jm4zwkfixr"},"logid":"cd79e47c610343cbb3345c8603e7dacd","msg":"success"}]
```

### 1.3 框架架构
```
        ┌ Common                             (配置包)
        │   └── Enum                         (枚举)
        │   └── Logger                       (自定义日志)
        │   └── Tools                        (工具文件)
        ├── Config                           (配置文件)
        ├── Controller                       (控制器)
        │   ├── Admin                        (后台控制器)
        │   └── Api                          (api控制器)
        ├── logs                             (日志目录)
        ├── Middleware                       (中间件)
        │   ├── CheckSignMiddleware.go       (签名中间件)
        │   ├── CheckTokenMiddleware.go      (校验token中间件)          
        │   └── LoggerMiddleware.go          (日志中间件)
        ├── Models                           (中间件层)
        │   ├── Mysql.go                     (mysql model)          
        │   └── Redis.go                     (redis model)
        ├── Routers                          (路由)
        │   ├── AdminRouter.go               (后台路由)
        │   ├── ApiRouter.go                 (api路由)
        │   └── Router.go.go                 (路由)
        ├── Service                          (服务层)
        ├── static                           (静态资源)
        │   ├── css                          (css文件)
        │   ├── image                        (图片文件)
        │   └── js                           (js文件)
        └── template                         (路由层)
            └── Admin                        (后台模板文件)                        
```

## 2. 使用说明

```
- golang版本 >= v1.16
- IDE推荐：Goland
```

### 2.1 server项目

使用 `Goland` 等编辑工具

```bash

# 克隆项目
git clone https://github.com/mjy191/ginApi.git
# 进入gitApi文件夹
cd gitApi
# 配置代理防止无法下载依赖包
go env -W GOPROXY=https://goproxy.cn,https://goproxy.io,direct

# 使用 go mod 并安装go依赖包
go mod tidy

# 编译 
go build -o server main.go (windows编译命令为go build -o server.exe main.go )

# 运行二进制
./server (windows运行命令为 server.exe)
```