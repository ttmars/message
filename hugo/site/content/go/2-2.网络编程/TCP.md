---
title: "TCP协议"
date: 2022-05-06T10:26:00+08:00
draft: true
---
[https://liuyehcf.github.io/2019/10/28/TCP-IP%E8%AF%A6%E8%A7%A3-%E8%AF%BB%E4%B9%A6%E7%AC%94%E8%AE%B0/](https://liuyehcf.github.io/2019/10/28/TCP-IP%E8%AF%A6%E8%A7%A3-%E8%AF%BB%E4%B9%A6%E7%AC%94%E8%AE%B0/)

### 术语

- ARQ（Automatic repeat request）：自动重复请求
- 面向连接、可靠的字节流、全双工
- 一个TCP连接由一个套接字（四元组）标识，四元组：由IP头中的源ip和目的ip、tcp头中的源端口和目的端口组成

### 抓包

#### wireshark

常用过滤规则：[https://blog.csdn.net/liu_yanxiaobai/article/details/124943492](https://blog.csdn.net/liu_yanxiaobai/article/details/124943492)

#### tcpdump

### 命令

#### telnet

telnet可以建立tcp连接，一般用于端口探测

[https://blog.csdn.net/weixin_45322291/article/details/116780620](https://blog.csdn.net/weixin_45322291/article/details/116780620)

#### netstat

```shell
[root@iZ8vb8qjajxkxytobaq81hZ ~]# netstat -h
-n, --numeric            don't resolve names						# 显示ip地址而不是域名
-l, --listening          display listening server sockets			 # 只显示监听套接字
-a, --all                display all sockets (default: connected)	  # 显示所有套接字
{-t|--tcp} {-u|--udp}											  # 显示TCP/UDP连接
-p, --programs           display PID/Program name for sockets		  # 显示对应的应用程序
```

常用组合

```shell
# 显示所有tcp监听套接字
netstat -nltp

# 显示所有tcp套接字
netstat -natp

# 同理udp
netstat -nlup
netstat -naup
```

### TCP头部组成

- 16位源端口+16位目的端口
- 32位序列号：代表每个分组的第一个字节在整个数据流中的字节偏移，ISN初始序列号是随机的
- 32位确认号：表示该序列号的字节已成功接收，期待接收下一个字节，即序列号+1（只有标志位ACK开启才生效）
- 4位头部长度+4位保留+8位标志号+16位窗口大小
- 16位TCP校验和+16紧急指针（只有标志位URG开启才生效）
- 选项

<img src="https://raw.githubusercontent.com/ttmars/image/master/network/tcp2.jpg" alt="image"  />

#### 头部长度

tcp头部占20字节，选项部分最多占40字节，总长度最大为60字节
为什么总长度最大为60字节？头部长度用4位表示，最大值为15，且以4字节为单位，故15*4=60

#### 标志位

- CWR：拥塞控制
- ECE：拥塞控制
- URG：当URG开启时，16位紧急指针字段才会生效
- ACK：确认位
- PSH：表示发送端缓存为空，即PSH数据包发送完后，发送端没有其他数据包需要传输
- PST：重置报文段，用于错误连接、立刻终止连接等等
- SYN：用于建立连接
- FIN：用于断开连接

#### 头部选项

| 种类 | 长度 | 名称                            | 描述                                          |
| ---- | ---- | ------------------------------- | --------------------------------------------- |
| 2    | 4    | MSS（Maximum segment size）     | 最大段大小，默认值：536，一般值：1500-40=1460 |
| 3    | 3    | WSOPT（window scale operation） | 窗口缩放因子（窗口左移量）                    |
| 4    | 2    | SACK（SACK permitted）          | 选择确认，发送者支持SACK选项                  |
| 1    | 1    | NOP（No Operation）             | 无操作，用作填充                              |

### TCP连接状态转换过程

#### TIME_WAIT状态

- tcp连接处于time_wait状态时，需要等待2MSL时长
- Linux：net.ipv4.tcp_fin_timeout；Windows：TcpTimedWaitDelay；取值范围：30~300s
- 主动关闭的一端才会进入time_wait状态
- 影响：在2MSL时长内，端口处于占用状态，无法再次使用！

<img src="https://raw.githubusercontent.com/ttmars/image/master/network/tcp3.jpg" alt="image"  />

<img src="https://raw.githubusercontent.com/ttmars/image/master/network/tcp1.jpg" alt="image"  />

### TCP三次握手

<img src="https://raw.githubusercontent.com/ttmars/image/master/network/tcp.jpg" alt="image"  />

### TCP服务器选项

1. 如何设置Local Address和Local Address？
2. 监听套接字和连接套接字的区别？

```shell
[root@iZ8vb8qjajxkxytobaq81hZ ~]# netstat -nltp
Active Internet connections (only servers)
Proto Recv-Q Send-Q Local Address           Foreign Address         State       PID/Program name    
tcp        0      0 127.0.0.1:8080          0.0.0.0:*               LISTEN      8340/./message      
tcp        0      0 0.0.0.0:80              0.0.0.0:*               LISTEN      1212/nginx: master  
tcp        0      0 0.0.0.0:22              0.0.0.0:*               LISTEN      1079/sshd           
tcp        0      0 127.0.0.1:8001          0.0.0.0:*               LISTEN      5698/v2ray          
tcp6       0      0 :::3400                 :::*                    LISTEN      12761/node          
tcp6       0      0 :::3800                 :::*                    LISTEN      1586/./musicAPI
```

- 一台主机可能有多个IP地址
- :::3400为IPv6格式的全零通配符，IPv4格式表现为0.0.0.0:3400
- 监听套接字：套接字处于LISTEN状态，可以接收SYN连接请求
- 连接套接字：套接字处于ESTABLISHED状态，可以传输数据

### TCP连接队列

### *TCP重传

了解即可，两种方式：

1. 超时重传：基于计时器
2. 快速重传：基于接收端的反馈信息来引发重传

### *TCP流量控制

了解即可，发送方比接收方快，强迫发送方降低速率

### *TCP拥塞控制

了解即可，保护发送方和接收方之间的网络设备，如路由器

### TCP保活机制

保活功能不是TCP规范的一部分，原因：

1. 在出现短暂的网络错误时，保活机制会使一个正常的连接断开
2. 占用不必要的带宽
3. 按流量计费的情况下，花费更多的成本

保活功能一般由服务器使用，探测客户端是否奔溃或离开，从而决定是否为客户端绑定资源

保活功能默认关闭，TCP连接的任何一端都可以自定义开启

