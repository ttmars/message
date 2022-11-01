---
title: "HTTP2协议"
date: 2022-05-06T10:26:00+08:00
draft: true
---
### HTTP2历史

- 2009年，谷歌的SPDY协议问世，主要特性：流多路利用、请求优先级、HTTP首部压缩
- 2015年，基于SPDY的HTTP2协议正式成为标准问世

### HTTP2新增概念

- 二进制协议
- 多路复用
- 流量控制功能
- 数据流优先级
- 首部压缩
- 服务端推送

### 组合

- HTTP/1.1+http
- HTTP/1.1+https
- HTTP/2+https=h2
- HTTP/2+http=h2c

### 命令

- telnet，使用telnet模拟tcp连接和发送http请求
- nc(netcat)

### HTTP版本区别

[https://www.cnblogs.com/yjh1995/p/16368320.html](https://www.cnblogs.com/yjh1995/p/16368320.html)

- HTTP/0.9：只有一个命令GET
- HTTP/1.0：增加了请求方法、HTTP版本号、HTTP首部、响应状态码
- HTTP/1.1：强制添加host首部、支持持久连接
- HTTP/2：host首部改为:authority伪首部，可通过这个字段区别1.1和2版本；HTTP/2请求中未明确声明版本号

### HTTPS

- HTTPS在HTTP的基础上，添加了TLS加密
- TLS的前身是SSL
- HTTPS提供了消息加密、消息完整性验证、服务器身份验证

### 模拟发送HTTP请求

- telnet、nc（netcat）、curl、wget
- 浏览器、postman

### 其他

优化http/1.1

- 使用多个http连接
- 合并http请求

### 网站测试瀑布图

[https://www.webpagetest.org/](https://www.webpagetest.org/)

### 升级到HTTP/2

### 负载均衡

- HTTP负载均衡器，也称7层负载均衡器
- TCP负载均衡器，也称4层负载均衡器

### TCP队头阻塞

[https://blog.csdn.net/weixin_44260459/article/details/120797655](https://blog.csdn.net/weixin_44260459/article/details/120797655)

HTTP/2解决了http层的队头阻塞，但TCP层的队头阻塞依然存在

因为tcp要保证有序，如果前面的包丢失，尽管后边的包成功收到，仍无法向上交付，必须等丢失的包重传后才能向应用层交付

### QUIC协议

- 谷歌提出，全称：Quick UDP Internet Connections
- 基于UDP