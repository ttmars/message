在HBuilderX.exe中打开终端，通过hugo将md文档编译为html
PS D:\project\goProject\message\hugo\site> 
hugo --buildDrafts --baseUrl="/public/"

### 模板渲染
https://blog.csdn.net/weixin_43761212/article/details/122301027

### Centos安装Chrome浏览器
https://blog.taliove.com/centos-headless-chrome

```shell
[root@ message]# cat /etc/yum.repos.d/google-chrome.repo 
[google-chrome]
name=google-chrome
baseurl=http://dl.google.com/linux/chrome/rpm/stable/$basearch
enabled=1
gpgcheck=0
gpgkey=https://dl-ssl.google.com/linux/linux_signing_key.pub

[root@ message]# yum install -y chromium
```

