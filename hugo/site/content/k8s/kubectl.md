---
title: "kubectl"
date: 2022-06-05T15:27:00+08:00
draft: true
---

### get

```bash
# 详细输出
kubectl get pods -o wide

# 输出yaml格式
kubectl get svc nginx -o yaml

# 输出json格式
kubectl get svc nginx -o json

# 实时输出
kubectl get pods -w
```

### config

```bash
# 切换命名空间
kubectl config set-context --current --namespace default

# 查看当前上下文
kubectl config get-contexts

# 切换上下文
kubectl config use-context kubernetes-admin@kubernetes

# 查看集群信息
kubectl cluster-info

# 查看集群资源使用率，依赖metrics-server插件（https://blog.csdn.net/liuyanwuyu/article/details/119793631）
kubectl top nodes
```

### explain

```bash
# 递归获取资源定义
kubectl explain pods --recursive

# 获取资源模板
kubectl create deploy demo --image=busybox:1.28 --dry-run=client -o yaml
```

### run

```bash
# 运行单个Pod
kubectl run busybox --image=busybox:1.28 -- sleep 99999

# 进入容器
kubectl exec -it busybox -- sh
```
### attach

```bash
# 将本地终端附着到容器
[root@master nginx]# kubectl attach nginx-68b6878866-428tz 
If you don't see a command prompt, try pressing enter.
172.16.219.64 - - [03/Jun/2022:07:43:44 +0000] "GET / HTTP/1.1" 200  615 "-" "curl/7.29.0" "-"
```

### port-forward

```bash
# 转发服务端口
kubectl port-forward svc/nginx --address 0.0.0.0 9999:80

# 转发POD端口
kubectl port-forward pods/nginx-68b6878866-428tz --address 0.0.0.0 9999:8080
```

### describe

```bash
# 获取资源状态
kubectl describe svc nginx
```

### logs

```bash
# 实时展示最近10行日志
kubectl logs ignite-0 --tail=10 -f
```

### cp

```bash
# 将主机文件复制到容器
kubectl cp localfile <pod_name>:<pod_filepath> -c containerName

# 将容器中的文件复制到主机
kubectl cp <pod_name>:<pod_filepath> localfile -c containerName
```

### tools

1. [kube-shell](https://github.com/cloudnativelabs/kube-shell)

与 Kubernetes CLI 一起使用的集成 shell

```bash
yum install -y python3
pip3 install kube-shell
```

2. [squash](https://github.com/solo-io/squash)

微服务调试器

3. [JQ](https://jqplay.org/)

json查询工具

4. [kubespy](https://github.com/pulumi/kubespy)

实时监控k8s资源变化情况

5. [kompose](https://github.com/kubernetes/kompose)

轻松将docker-compose项目迁移到k8s

6. [kubeval](https://github.com/instrumenta/kubeval)

k8s资源验证器

7. other

click、kubed-sh、stern