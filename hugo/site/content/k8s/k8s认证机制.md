---
title: "k8s认证机制"
date: 2022-06-05T13:11:00+08:00
draft: true
---

### APIServer三重认证

1. 身份认证：决定用户能不能访问
2. 授权：决定用户能访问的资源
3. 准入控制器：决定用户访问资源的规范、准则

### 身份认证

1. 常规用户认证
2. ServiceAccount认证，集群内部Pod使用

### RBAC授权

![](https://raw.githubusercontent.com/va-len-tine/image/master/k8s/k8s-1.png)

### 准入控制器