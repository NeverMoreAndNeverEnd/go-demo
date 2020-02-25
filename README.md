# go-demo
go crud demo

目的

本项目采用了一系列Golang中比较流行的组件，可以以本项目为基础快速搭建Restful Web API

特色

本项目已经整合了许多开发API所必要的组件：

Gin: 轻量级Web框架，自称路由速度是golang最快的

GORM: ORM工具。本项目需要配合Mysql使用

Gin-Session: Gin框架提供的Session操作工具

Go-Redis: Golang Redis客户端

godotenv: 开发环境下的环境变量工具，方便使用环境变量

Gin-Cors: Gin框架提供的跨域中间件

自行实现了国际化i18n的一些基本功能

本项目是使用基于cookie实现的session来保存登录状态的，如果需要可以自行修改为token验证

本项目已经预先创建了一系列文件夹划分出下列模块:

api文件夹就是MVC框架的controller，负责协调各部件完成任务

model文件夹负责存储数据库模型和数据库操作相关的代码

service负责处理比较复杂的业务，把业务代码模型化可以有效提高业务代码的质量（比如用户注册，充值，下单等）

serializer储存通用的json模型，把model得到的数据库模型转换成api需要的json对象

cache负责redis缓存相关的代码

auth权限控制文件夹

util一些通用的小工具

conf放一些静态存放的配置文件，其中locales内放置翻译相关的配置文件
