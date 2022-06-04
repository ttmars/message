---
title: "K8S服务"
date: 2022-06-04T17:49:00+08:00
draft: true
---

### 五种服务类型

向外发布服务

- ClusterIP
- NodePort
- LoadBalancer

向内发布服务

- 无头服务
- ExternalName

### 两种服务发现方式

- 环境变量

- DNS服务器

### ClusterIP

1. 通过kube-proxy组件实现负载均衡
2. 自动分配集群IP

```yaml
# https://kubernetes.io/docs/concepts/services-networking/service/
apiVersion: v1
kind: Service
metadata:
  name: service-simple-service
spec:
  selector:
    app: service-simple-app
  ports:
    - protocol: TCP
      port: 80
      targetPort: 8080
```

### NodePort

1. 基于ClusterIP模式
2. 端口范围：30000-32767

```yaml
# https://kubernetes.io/docs/concepts/services-networking/service/#nodeport
apiVersion: v1
kind: Service
metadata:
  name: service-node-port-service
spec:
  type: NodePort
  selector:
    app: MyApp
  ports:
    - port: 80         # By default and for convenience, the `targetPort` is set to the same value as the `port` field.
      targetPort: 80   # Optional field.
      nodePort: 30007  # Optional field. By default and for convenience, the Kubernetes control plane will allocate a port from a range (default: 30000-32767)
 
```

### LoadBalancer

1. 基于ClusterIP+NodePort
2. 公有云才支持LoadBalancer

```yaml
# https://kubernetes.io/docs/concepts/services-networking/service/#loadbalancer
apiVersion: v1
kind: Service
metadata:
  name: service-load-balancer-service
spec:
  selector:
    app: MyApp
  ports:
    - protocol: TCP
      port: 80
      targetPort: 9376
  type: LoadBalancer
 
```

在私有集群中部署负载均衡器[metallb](https://metallb.universe.tf/installation/)，通过helm安装，修改values.yaml如下

```yaml
configInline:
  address-pools:
    - name: default
      protocol: layer2
      addresses:
      - 192.168.66.100-192.168.66.101
```

### 无头服务

1. 不会分配集群IP
2. 不经过kube-proxy实现负载均衡
3. 普通服务将DNS解析为集群IP，无头服务将DNS解析为Pod地址
4. 只能在POD中访问，集群节点和外部节点无法直接访问
5. 主要在StatefulSet中使用

```yaml
apiVersion: v1
kind: Service
metadata:
  name: headless-service-headless-service
spec:
  clusterIP: None  # This marks this service out as a headless service
  selector:
    app: headless-service-app
  ports:
    - protocol: TCP
      port: 80
      targetPort: 3000
```

最新版busybox中的nslookup命令不能正常解析DNS，应指定版本1.28

1. 运行busybox

```bash
kubectl run busybox --image=busybox:1.28 -- sleep 999999
```

2. 进入容器

```bash
kubectl exec -it busybox -- sh
```

```shell
/ # nslookup nginx
Server:    10.96.0.10
Address 1: 10.96.0.10 kube-dns.kube-system.svc.cluster.local

Name:      nginx
Address 1: 10.100.203.13 nginx.test.svc.cluster.local						# 解析为集群IP
/ # 
/ # 
/ # nslookup nginx-headless
Server:    10.96.0.10
Address 1: 10.96.0.10 kube-dns.kube-system.svc.cluster.local

Name:      nginx-headless
Address 1: 172.16.166.157 172-16-166-157.nginx.test.svc.cluster.local		# 解析为POD地址
```

### ExternalName

1. 将外部服务引入集群
2. 没有选择器、端口定义

```yaml
# https://kubernetes.io/docs/concepts/services-networking/service/#externalname
apiVersion: v1
kind: Service
metadata:
  name: service-external-name-service
spec:
  type: ExternalName
  externalName: my.database.example.com
 
```

### 其他方式

1. 自定义Endpoints
2. ExternalIP

### 服务访问方式

#### 创建一个nginx服务

```bash
[root@master nginx (⎈ |admin@kubernetes:test)]# kubectl get svc/nginx 
NAME    TYPE           CLUSTER-IP      EXTERNAL-IP      PORT(S)        AGE
nginx   LoadBalancer   10.100.203.13   192.168.66.100   80:31006/TCP   25m
[root@master nginx (⎈ |admin@kubernetes:test)]# kubectl describe svc/nginx 
Name:                     nginx
Namespace:                test
Labels:                   app.kubernetes.io/instance=nginx
                          app.kubernetes.io/managed-by=Helm
                          app.kubernetes.io/name=nginx
                          helm.sh/chart=nginx-12.0.0
Annotations:              meta.helm.sh/release-name: nginx
                          meta.helm.sh/release-namespace: test
Selector:                 app.kubernetes.io/instance=nginx,app.kubernetes.io/name=nginx
Type:                     LoadBalancer
IP Family Policy:         SingleStack
IP Families:              IPv4
IP:                       10.100.203.13
IPs:                      10.100.203.13
LoadBalancer Ingress:     192.168.66.100
Port:                     http  80/TCP
TargetPort:               http/TCP
NodePort:                 http  31006/TCP
Endpoints:                172.16.166.157:8080
Session Affinity:         None
External Traffic Policy:  Cluster
Events:
  Type    Reason        Age   From                Message
  ----    ------        ----  ----                -------
  Normal  IPAllocated   25m   metallb-controller  Assigned IP ["192.168.66.100"]
  Normal  nodeAssigned  25m   metallb-speaker     announcing from node "master"

```

#### 外部节点访问

```
任意节点IP+NodePort访问
http://192.168.66.100:31006/
http://192.168.66.101:31006/

负载均衡器EXTERNAL-IP+NodePort访问
http://192.168.66.100:31006/
```

#### 集群节点访问

```
集群IP+Port访问
curl 10.100.203.13:80

Endpoints访问
curl 172.16.166.157:8080
```

#### 集群Pod访问

```
集群IP+Port访问
wget -O- 10.100.203.13:80

Endpoints访问
wget -O- 172.16.166.157:8080

服务名访问
wget -O- nginx:80

完整DNS访问
wget -O- nginx.test.svc.cluster.local:80
```

### Ingress

[https://kubernetes.github.io/ingress-nginx/](https://kubernetes.github.io/ingress-nginx/)













