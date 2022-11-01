### 获取当前路径

```go
filepath.Abs(".")
```

### 路径截取

```go
path := "/aa/bb/cc/dd.js"
filepath.Dir(path)		// \aa\bb\cc 截取前部分
filepath.Base(path)		// dd.js 截取后部分
filepath.Ext(path)		// .js 截取扩展名
filepath.Split(path)	// /aa/bb/cc/ dd.js
```

### 路径拼接

```go
// 跨平台兼容
filepath.Join("aa", "bb")	//linux:aa/bb windowns:aa\bb
```

### 路径匹配

```go
filepath.Match("/home/catch/*", "/home/catch/foo")		// true nil
```

### 路径斜杠转换

```go
filepath.FromSlash("/aa/bb/cc/dd.js")		// \aa\bb\cc\dd.js
filepath.ToSlash("\\aa\\bb\\cc\\dd.js")		// /aa/bb/cc/dd.js
```

### 路径遍历并回调

```go
// Walk调用了os.Lstat返回更详细的文件信息，因此效率比WalkDir要低
func Walk(root string, fn WalkFunc) error {}
func WalkDir(root string, fn fs.WalkDirFunc) error {}
```