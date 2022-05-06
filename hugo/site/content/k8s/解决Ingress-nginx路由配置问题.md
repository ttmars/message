---
title: "解决Ingress-nginx路由配置问题"
date: 2022-05-06T09:33:00+08:00
draft: true
---
### Ingress-nginx路由配置

官方文档：[https://kubernetes.github.io/ingress-nginx/examples/rewrite/](https://kubernetes.github.io/ingress-nginx/examples/rewrite/)

### 静态文件404问题

[https://www.qikqiak.com/post/url-rewrite-on-ingress-nginx/](https://www.qikqiak.com/post/url-rewrite-on-ingress-nginx/)

[https://stackoverflow.com/questions/50871970/kubernetes-ingress-nginx-loading-resources-404](https://stackoverflow.com/questions/50871970/kubernetes-ingress-nginx-loading-resources-404)

### 配置示例

```yaml
kind: Ingress
apiVersion: extensions/v1beta1
metadata:
  name: **
  namespace: **
  annotations:
    kubernetes.io/ingress.class: nginx
    nginx.ingress.kubernetes.io/rewrite-target: /$1		# $1为正则匹配占位符
    nginx.ingress.kubernetes.io/configuration-snippet: |
      rewrite ^/css/(.*)$ /app/css/$1 redirect;  		# 给静态文件目录添加app前缀
spec:
  rules:
    - host: **
      http:
        paths:
          - path: /app/(.*)		# 添加app路径，以路径区别后端服务
            pathType: Prefix
            backend:
              serviceName: **
              servicePort: 80
```

### 注意事项

后端对静态资源的引用，要删除前导斜杠

```html
如:
<script src=/js/example.js></script>
删除前导斜杠:    
<script src=js/example.js></script>
```

