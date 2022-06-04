---
title: "k8s存储与配置"
date: 2022-06-04T21:46:00+08:00
draft: true
---
### 一、本地存储卷

#### emptyDir

1. 将主机的临时目录映射到容器里
2. 与pod同生命周期，pod销毁，存储卷同时销毁

```yaml
# 基于磁盘的临时存储
volumes:
  - name: cache-volume
    emptyDir: {}
  
# 基于内存的临时存储
volumes:
  - name: cache-volume
    emptyDir:
      medium: Memory
```

#### hostPath

1. 将主机上的指定目录映射到容器里
2. 永久保存

```yaml
volumes:
  - name: cache-volume
    hostPath:
      path: /data
```

### 二、网络存储卷

1. 可以跨节点共享存储
2. 支持多种存储方案，如NFS

```yaml
volumes:
  - name: cache-volume
    nfs:
      path: /data
      server: 192.168.66.100
```

### 三、持久存储卷

#### 静态PV                                            

```yaml
# NFS服务==>持久卷（PersistentVolume）==>持久卷声明（PersistentVolumeClaim）==>Pod引用
volumes:
  - name: cache-volume
    persistentVolumeClaim:
      claimName: <pvcName>
```

#### 动态PV

```yaml
# NFS服务==>存储分配器（provisioner）==>持久卷类（StorageClass）==>持久卷声明（PersistentVolumeClaim）==>Pod引用（自动绑定PV）
volumes:
  - name: cache-volume
    persistentVolumeClaim:
      claimName: <scName>
```

### 四、配置存储卷

#### ConfigMap

1. 模板示例

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: example-configmap
data:
  name: valar
  age: 25
  hobby: |
    game=LOL
    movie=Titanic
```

2. 通过环境变量引用

```yaml
# 依次引用
env:
  - name: NAME
    valueFrom: 
      configMapKeyRef:
        name: example-configmap
        key: name
  - name: AGE
    valueFrom: 
      configMapKeyRef:
        name: example-configmap
        key: age
  - name: HOBBY
    valueFrom: 
      configMapKeyRef:
        name: example-configmap
        key: hobby
# 一次引用所有
envFrom:
  - configMapRef:
      name: example-configmap
```

3. 通过存储卷引用

```yaml
# 会将configMap中的键值对映射到容器的指定目录，键->文件名，值->文件内容
volumes:
  - name: volume-config
    configMap:
      name: example-configmap
```

#### Secret

1. OpaqueSecret，对数据进行base64编码

```bash
加密：echo "hello" | base64
解密：echo "aGVsbG8K" | base64 --decode
```

2. ImagePullSecret，用于存储私有仓库的认证信息

3. ServiceAccountSecret，用于存储访问KubernetesAPI的认证信息

#### DownwardAPI

向Pod中运行的容器暴露Pod自身的信息