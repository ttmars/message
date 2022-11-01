### 文件

```go
// 链接文件
func Link(oldname, newname string) error	// 创建硬链接
func Symlink(oldname, newname string) error	// 创建软连接（符号链接）
func Readlink(name string) (string, error)	// 读取符号链接指向的文件名

// 管道文件
func Pipe() (r *File, w *File, err error)	// 创建管道

// 文件读写
func ReadFile(name string) ([]byte, error)						// 读取整个文件到内存
func WriteFile(name string, data []byte, perm FileMode) error	  // 将数据保存到文件

// 文件删除
func Remove(name string) error		// 删除文件或目录
func RemoveAll(path string) error	// 递归删除文件、目录，相当于rm -rf，慎用！

// 文件权限修改
func Getgid() int		// 获取gid
func Getegid() int		// 获取有效gid
func Getuid() int		// 获取uid
func Geteuid() int		// 获取有效uid

func Chown(name string, uid, gid int) error		// 更改文件所有者，更改符号链接指向的文件
func Lchown(name string, uid, gid int) error	// 更改文件所有者，更改符号链接本身

// 文件信息修改
func Rename(oldpath, newpath string) error							// 重命名文件
func Truncate(name string, size int64) error						// 更改文件大小
func Chtimes(name string, atime time.Time, mtime time.Time) error	  // 更改文件时间
```

### 目录

```go
func Chdir(dir string) error {}				// 更改文件权限
func Getwd() (dir string, err error)		// 获取当前目录

// 创建目录
func Mkdir(name string, perm FileMode) error		// 创建目录
func MkdirAll(path string, perm FileMode) error		// 创建目录，包括父目录
func MkdirTemp(dir, pattern string) (string, error)	// 创建临时目录

// 默认目录
func TempDir() string					// 返回操作系统默认的临时目录
func UserCacheDir() (string, error)		// 返回操作系统默认的缓存目录
func UserConfigDir() (string, error)	// 返回操作系统默认的配置目录
func UserHomeDir() (string, error)		// 返回当前用户的家目录
```

### 系统

```go
func Exit(code int)							// 退出程序
func Hostname() (name string, err error)	  // 获取主机名
func Executable() (string, error)			 // 输出当前运行程序的完整路径
func Getpagesize() int						// 获取内存页大小，4096

func Getpid() int		// 获取进程ID
func Getppid() int		// 获取父进程ID

func DirFS(dir string) fs.FS
func Getgroups() ([]int, error)
```

### 环境变量

```go
func Setenv(key, value string) error	// 设置环境变量
func Unsetenv(key string) error			// 删除环境变量
func Clearenv()						// 删除所有环境变量
func Environ() []string				// 输出所有环境变量
func Getenv(key string) string					// 查找环境变量
func LookupEnv(key string) (string, bool)		// 查找环境变量，可以区分值为空或不存在两种情况

// 变量名扩展，https://pkg.go.dev/os@go1.19.1#ExpandEnv
func Expand(s string, mapping func(string) string) string
func ExpandEnv(s string) string
```

### DirEntry

```go
func ReadDir(name string) ([]DirEntry, error)		// 读取目录
```

### File

https://pkg.go.dev/os@go1.19.1#File
