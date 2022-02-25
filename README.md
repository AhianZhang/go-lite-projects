# 介绍

本仓库是学习 golang 时敲的一些项目集合

## im
命令行聊天工具，支持公聊、私聊、超时下线、改名。
来源：https://www.bilibili.com/video/BV1gf4y1r79E
 原作者：刘丹冰Aceld

编译命令
```shell
go bulid -o server main.go server.go user.go
go bulid -o client client.go
```
执行
```shell
./server
./client
```