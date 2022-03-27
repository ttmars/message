---
title: "First"
date: 2022-03-27T11:22:17+08:00
draft: true
---
### 一.基本概念
1.使用存储设备的一般方法

	创建分区	
	格式化，即创建文件系统（创建inode、block）
	挂载

2.硬盘的接口方式

	PATA接口：即IDE
	SATA接口：多用于个人桌面
	服务器常用：SCSI、SAS、FC-AL

3.分区表类型

MBR与GPT的区别，参考：
[https://www.dujin.org/11274.html](https://www.dujin.org/11274.html)
[https://www.dazhuanlan.com/2020/01/02/5e0dafbd204e4/](https://www.dazhuanlan.com/2020/01/02/5e0dafbd204e4/)

4.文件系统类型

常见文件系统，参考：
[https://www.cnblogs.com/eddie1127/p/11260402.html](https://www.cnblogs.com/eddie1127/p/11260402.html)
[https://blog.csdn.net/liushengxi_root/article/details/87439737](https://blog.csdn.net/liushengxi_root/article/details/87439737)
centos7默认文件系统类型：xfs
win7默认文件系统类型：NTFS
### 二.常用命令
1.查看磁盘和分区

	[root@localhost ~]# ls /dev/sda
	sda   sda1  sda2
	新建分区后，用partprobe刷新磁盘

2.查看已经挂载的分区

	[root@localhost ~]# df -Th|grep -v tmp
	文件系统                类型      容量  已用  可用 已用% 挂载点
	/dev/mapper/centos-root xfs        17G  1.6G   16G   10% /
	/dev/sda1               xfs      1014M  138M  877M   14% /boot

3.查看磁盘的分区、挂载信息

	[root@localhost ~]# lsblk
	NAME            MAJ:MIN RM  SIZE RO TYPE MOUNTPOINT
	sda               8:0    0   20G  0 disk 
	├─sda1            8:1    0    1G  0 part /boot
	└─sda2            8:2    0   19G  0 part 
	  ├─centos-root 253:0    0   17G  0 lvm  /
	  └─centos-swap 253:1    0    2G  0 lvm  [SWAP]
	sr0              11:0    1  4.4G  0 rom

4.查询磁盘分区详情

	[root@localhost ~]# fdisk -l
	磁盘 /dev/sda：21.5 GB, 21474836480 字节，41943040 个扇区
	Units = 扇区 of 1 * 512 = 512 bytes
	扇区大小(逻辑/物理)：512 字节 / 512 字节
	I/O 大小(最小/最佳)：512 字节 / 512 字节
	磁盘标签类型：dos
	磁盘标识符：0x0009c530
	
	   设备 Boot      Start         End      Blocks   Id  System
	/dev/sda1   *        2048     2099199     1048576   83  Linux
	/dev/sda2         2099200    41943039    19921920   8e  Linux LVM
	
	磁盘 /dev/mapper/centos-root：18.2 GB, 18249416704 字节，35643392 个扇区
	Units = 扇区 of 1 * 512 = 512 bytes
	扇区大小(逻辑/物理)：512 字节 / 512 字节
	I/O 大小(最小/最佳)：512 字节 / 512 字节
	
	
	磁盘 /dev/mapper/centos-swap：2147 MB, 2147483648 字节，4194304 个扇区
	Units = 扇区 of 1 * 512 = 512 bytes
	扇区大小(逻辑/物理)：512 字节 / 512 字节
	I/O 大小(最小/最佳)：512 字节 / 512 字节
### 二.磁盘分区、格式化、挂载
参考：[https://www.cnblogs.com/duzhaoqi/p/7395619.html](https://www.cnblogs.com/duzhaoqi/p/7395619.html)
1.MBR分区工具fdisk

	fdisk /dev/sdb
	o:创建一个空的DOS类型分区表
	通知内核重新读取分区表：partprobe /dev/sdb 

2.GPT分区工具gdisk
	
	gdisk /dev/sdb

3.通用分区工具parted

	parted /dev/sdb

4.格式化

	mkfs.xfs /dev/sdb1

5.挂载

	mount /dev/sdb1 /root/usb

### 三.LVM逻辑卷管理

	物理卷PV（Physical Volume）：可以是硬盘、硬盘分区
	卷组VG（Volume Group）：有一个或多个物理卷组成，类似非LVM系统中的物理硬盘
	逻辑卷LV（Logical Volume）：从卷组中“切出”的一块空间，可以动态扩展，类似非LVM系统中的硬盘分区
	物理区域PE（Physical Extent）：物理卷的基本单位，默认4M，同一个卷组中的物理卷的PE大小一致
	逻辑区域LE（Logical Extent）：同一个卷组中，LE与PE大小一致，且一一对应

### 四.RAID磁盘阵列
参考：[https://www.cnblogs.com/Wolf-Dreams/p/10693115.html](https://www.cnblogs.com/Wolf-Dreams/p/10693115.html)
### 五.磁盘限额
### 六.磁盘扩容
参考：[https://blog.51cto.com/daisywei/1897487](https://blog.51cto.com/daisywei/1897487)

1.fdisk创建一个新分区

	[root@localhost ~]# fdisk /dev/sda
	看不到新分区？执行一下：partprobe

2.使用新分区创建一个物理卷PV

	[root@localhost ~]# pvcreate /dev/sda3

3.查看卷信息

	物理卷pv：pvdisplay
	逻辑卷组vg：vgdisplay
	逻辑卷lv：lvdisplay

4.将物理卷扩展到逻辑卷组中

	[root@localhost ~]# vgextend centos /dev/sda3

5.扩容逻辑卷

	[root@localhost ~]# lvextend -L +5GB -n /dev/centos/root

6.使扩容生效

	[root@localhost ~]# xfs_growfs /dev/centos/root
