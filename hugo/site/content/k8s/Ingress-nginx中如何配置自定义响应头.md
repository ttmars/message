---
title: "Ingress-nginx中如何配置自定义响应头"
date: 2022-05-06T10:26:00+08:00
draft: true
---
官方文档：[https://kubernetes.github.io/ingress-nginx/](https://kubernetes.github.io/ingress-nginx/)

### 方式一

通过Ingress注释配置，add_header指令会映射到nginx配置文件中的localtion模块中，只对单个服务生效

[https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/annotations/#configuration-snippet](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/annotations/#configuration-snippet)

示例：

```yaml
apiVersion: extensions/v1betal
kind: Ingress
metadata:
  annotations:
    kubernetes.io/ingress.class: nginx
    nginx.ingress.kubernetes.io/configuration-snippet: |
      add_header "**" "**";
spec:
  rules:
  **
```

### 方式二

将配置存储在configmap中，在Ingress中间接引用，此方式会将add_header指令映射到nginx配置文件中的http模块中，对所有服务生效

[https://kubernetes.github.io/ingress-nginx/examples/customization/custom-headers/](https://kubernetes.github.io/ingress-nginx/examples/customization/custom-headers/)

示例：

自定义custom-headers文件，存储响应头

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: custom-headers
  namespace: ingress-nginx
data:
  X-Different-Name: "true"
  X-Using-Nginx-Controller: "true"
```

在Ingress-nginx默认配置文件中，间接引用custom-headers

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: ingress-nginx-controller
  namespace: ingress-nginx
data:
  add-headers: "ingress-nginx/custom-headers"
```

### 注意事项

1. 方式二修改配置文件后，需要重启nginx-ingress-controller，配置才能生效！
2. 由于nginx指令add_header的继承性问题，两种方式不能同时配置，否则方式二会失效！请参考：[https://juejin.cn/post/6844904041126838286](https://juejin.cn/post/6844904041126838286)