

+++

title = "Linux.文件"
description = "Linux.文件"
tags = ["it","os","Linux"]

+++



# Linux.文件

## 通配符

通配符：通用匹配符号，shell 内建的符号

作用：用于操作多个相似（有简单规律）的文件

常用：

- `*`：匹配一或多个字符
- `?`：匹配一个字符
- `[xyz]`：匹配指定字符之中任意一个字符
- `[!xyz], [^xyz]`：匹配指定字符之外任意一个字符
- `[a-z]`：匹配指定范围之中任意一个字符
- `[!a-z], [^a-z]`：匹配指定范围之外任意一个字符

## 路径

- 绝对路径：以根目录 `/` 开头的路径
- 相对路径
  - 当前目录：`.` 或 `./`
  - 前一目录：`..` 或 `../`

## 操作

一切皆文件。在linux中，设备、磁盘、文件系统 等一切都可以用文件表示

- 空文件：
- 普通文件：
- 设备文件：
- 文件空洞：文件从开头到结尾的扇区所占用的磁盘空间中，未存储任何数据的部分被称为空洞。比如给linux创建一个虚拟机的磁盘空间1T，但是实际只存放1M的数据，其它都是空洞，用`du`查看将是1M，`ls -lh`查看将是1T



### ls

ls：list

作用：列出目录内容

语法：`ls [OPTION]... [FILE]...`

参数（常用）

- `-l`：`ls -l` 可缩写为 `ll`
  - 长格式显示文件：文件类型和权限，文件个数，创建用户，创建用户所属用户组，文件大小，最后修改时间，文件名
  - 文件类型：文件为-，目录为d
- `-a`：显示隐藏文件。文件名以"."开头即隐藏文件
- `-h`：文件大小将使用单位m、g、t等（根据文件大小自判定的），一般配合`-l`使用
- `-t`：按照时间顺序显示
- `-R`：递归显示所有文件

```shell
$ ls
anaconda-ks.cfg  Desktop  Documents  Downloads  initial-setup-ks.cfg  it  Music  Pictures  Public  Templates  Videos

$ ls -a /
.  ..  bin  boot  dev  etc  home  lib  lib64  media  mnt  opt  proc  root  run  sbin  srv  sys  tmp  usr  var

$ ls -lh /root /mnt
/mnt:
总用量 0
drwxrwx---. 1 root vboxsf 0 8月  22 15:52 share

/root:
总用量 8.0K
-rw-------. 1 root root 1.3K 8月  22 14:37 anaconda-ks.cfg
drwxr-xr-x. 2 root root    6 8月  22 15:42 Desktop
drwxr-xr-x. 2 root root    6 8月  22 15:42 Documents
drwxr-xr-x. 2 root root    6 8月  22 15:42 Downloads
-rw-r--r--. 1 root root 1.6K 8月  22 14:43 initial-setup-ks.cfg
drwxr-xr-x. 3 root root   16 8月  22 16:09 it
drwxr-xr-x. 2 root root    6 8月  22 15:42 Music
drwxr-xr-x. 3 root root   24 8月  26 22:20 Pictures
drwxr-xr-x. 2 root root    6 8月  22 15:42 Public
drwxr-xr-x. 2 root root    6 8月  22 15:42 Templates
drwxr-xr-x. 2 root root    6 8月  22 15:42 Videos
```



### tree

作用：以树状格式列出目录的内容  

```shell
# 非 linux 自带，需安装
$ yum install tree -y
```

语法：`ls [OPTION]... [FILE]...`

参数（常用）

- `-L level [-R]`： 目录树的最大显示深度

```shell
$ tree ~/Pictures/
/root/Pictures/
└── Wallpapers
    └── 91362440_p0.png
$ tree ~/Pictures/ -L 1
/root/Pictures/
└── Wallpapers
```



### cp

`cp`：copy

作用：复制文件

语法

- `cp [OPTION]... SOURCE... DIRECTORY`
- `cp [OPTION]... [-T] SOURCE DEST`
- `cp [OPTION]... -t DIRECTORY SOURCE...`

参数（常用）

- `-r, -R, --recursive`：递归复制目录
- `-v, --verbose`：解释正在做什么。即打印当前正在复制的内容（进度）
- `--preserve[=ATTR_LIST]`：保留指定的属性，默认有 mode、ownership、timestamps，如果可能的话还有其它他属性 context、links、xattr、all
- `-p`：等价于 `--preserve=mode,ownership,timestamps`。即保留 `--preserve`默认的属性
- `-a, --archive`：等价于 `-dR --preserve=all`。即保留所有属性
- `-d`：等价于 `--no-dereference --preserve=links`。即保留 links 属性
- `-P, --no-dereference`：永不遵循 SOURCE 中的符号链接（symbolic links）

```shell
$ mkdir -p /a/b/c
$ cp -r /a a.bk
```



### mv

`mv`：move

作用：移动（重命名）文件

语法

- `mv [OPTION]... SOURCE... DIRECTORY`
- `mv [OPTION]... [-T] SOURCE DEST`
- `mv [OPTION]... -t DIRECTORY SOURCE...`

```shell
# 移动到相同目录下即等于改名
$ mv a.bk a
```



### rm

`rm`：remove

作用：删除文件或目录

语法：`rm [OPTION]... [FILE]...`

参数（常用）

- `-f, --force`：忽略不存在的文件和参数，不作提示。rm 操作大多都有各种提示和询问，使用 `-f` 将直接删除不做询问
- `-r, -R, --recursive`：递归地删除目录及其内容。配合 `-f` 使用将不再提示询问，将直接进行递归删除

```shell
$ mkdir -p a/b/c
$ touch a/b/c/d.md

$ rm a/b/c/d.md
rm：是否删除普通文件 'a/b/c/d.md'？
$ rm -r a/b/c
rm：是否删除目录 'a/b/c'？
$ rm -r a
rm：是否进入目录'a'?

$ rm -rf a
```



### touch

作用：修改文件（或目录）的时间戳（属性）。如果文件不存在则新建文件

语法：`touch [OPTION]... FILE...`

```shell
$ touch a.md
```



### dd

作用：转换和复制文件

语法

- `dd [OPERAND]...`
-  `dd OPTION`

参数（常用）

- `bs=BYTES`：一次读或写的块大小（bytes），默认 512
- `ibs=BYTES`：一次写的块大小（bytes），默认 512
- `obs=BYTES`：一次读的块大小（bytes），默认 512
- `count=N`：复制块的数量
- `if=FILE`：输入文件，从文件读代替从 stdin
- `of=FILE`：输出文件，写文件代替写入 stdout
- `seek=N`：在输出开始时跳过 N 个 obs 大小的块。跳过部分即成为**文件空洞**
- `skip=N`：在输入开始时跳过 N 个 obs 大小的块

```shell
# 从 /dev/zero 可以读取无穷多个0，常用于测试 dd 命令
dd if=/dev/zero of=test.tar.gz bs=1M count=10
```



### du

作用：估计文件空间使用情况。默认计算磁盘使用情况（不包含文件空洞的大小）

语法

- `du [OPTION]... [FILE]...`
- `du [OPTION]... --files0-from=F`

参数（常用）

- `--apparent-size`：打印表面大小（即包含文件空洞的大小），而不是磁盘使用情况。虽然表面大小通常较小，但也可能由于漏洞（“稀疏”）文件、内部碎片、间接块等更大
- `-B, --block-size=SIZE`：打印前按 SIZE 缩放尺寸。即指定单位，如 1M、1G，默认 1K
- `-b, --bytes`：等价于`--apparent-size --block-size=1`，即以 Byte 为单位

```shell
$ du -b Pictures/
5745716 Pictures/Wallpapers
5745740 Pictures/

$ dd if=/dev/zero of=test.tar.gz bs=1M count=10 seek=5
$ ll test.tar.gz -h
-rw-r--r--. 1 root root 15M 9月  21 00:49 test.tar.gz
$ du test.tar.gz --B=1M --apparent-size
15      test.tar.gz
$ du test.tar.gz --B=1M
10      test.tar.gz
```



### cd

cd：change directory。cd 是 bash 的内置命令

作用：切换当前工作目录

语法：`cd [DIRECTORY]`

```shell
# 切换到指定目录
$ cd /root
# 切换到上两层目录
$ cd ../..
# 切换到上一次使用的目录。可以实现两个目录来回切换的效果
$ cd -
# 切换到家目录
$ cd ~
$ cd
```



### pwd

pwd：print working directory

作用：打印当前工作目录名

语法：`pwd [OPTION]...`

```shell
$ pwd
/
```



### mkdir

mkdir：make directories

语法：`mkdir [OPTION]... DIRECTORY...`

参数（常用）

- `-p, --parents`：根据需要创建父目录

```shell
$ mkdir a a/b
$ mkdir -p c/d/e
```



### rmdir

作用：删除空目录

语法：`rmdir [OPTION]... DIRECTORY...`

参数（常用）

- `-p, --parents`：删除DIRECTORY及其祖先

```shell
$ rmdir -p a
```



## 打包

- Linux的备份压缩
  - 最早的Linux备份介质是磁带，使用的命令是tar，即打包
  - 可以对打包后的磁带文件进行压缩储存，压缩的命令是gzip和bzip2
  - 经常使用的扩展名是.tar.gz、tar.bz2、tgz
- /etc一般是保存配置文件的目录，属于重点备份的文件，以etc为例进行打包和压缩
- 一般采用双扩展名以方便知道是那种方式打包的

### tar命令

作用：一个归档工具。即压缩打包工具

语法

- `tar [OPTION] `

参数

- `-c, --create`：创建一个新的归档文件，参数提供要归档的文件的名称，目录是递归归档的，除非给出了`--no-recursive`选项。即**打包**。
- `-x, --extract, --get`：从归档文件中提取文件， 参数是可选的，给出时，它们指定要提取的存档成员的名称。即**解包**。
- `-f, --file=ARCHIVE`：使用归档文件，或者设备归档
- `-z, --gzip, --gunzip, --ungzip`：使用 gzip 过滤归档文件。即使用 **gzip 压缩或解压缩**。
- `-j, --bzip2`：使用 bzip2 过滤归档文件。即使用 **bzip2 压缩或解压缩**。
- `-d, --directory=DIR`：在执行任何操作之前，请更改为DIR。 这个选项是顺序敏感的，即它影响后面的所有选项。即比如解包时指定保存路径，默认保存到当前工作目录
- `-v, --verbose`：详细列出所处理的文件

```shell
# 打包
$ tar -cf a.tar.gz a
# 解包
$ tar -xf a.tar.gz -C ./
```



## 权限

### chmod

作用：修改文件 mod bits

语法：

- `chmod [OPTION]... MODE[,MODE]... FILE...`
- `chmod [OPTION]... OCTAL-MODE FILE...`
- `chmod [OPTION]... --reference=RFILE FILE...`

```shell
# 添加执行权限，a表示所有，x表示执行
$ chmod a+x demo.exe
```

## 文本

### cat

- cat：文本内容显示到终端

### head

- head：查看文件开头。默认显示10行，-5参数显示5行

### tail

- tail：查看文件结尾。默认显示10行，-5参数显示5行
  - 常用参数-f 文件内容更新后，显示信息同步更新，ctrl+c停止

### wc

- wc：统计文件内容信息。-l参数显示文件行数

### more

- more：将分页显示
- less more

### 编辑

#### vi

- vi是多模式文本编辑器
- 多模式产生的原因
- 四种模式，通过模式的切换，就可以无需鼠标仅使用键盘进行各种各样的文本操作
  - 正常模式(Normal-mode)：进入编辑器界面时的初始模式，显示文本内容，有光标，光标可以移动。该模式下所有键盘输入的按键都是对编辑器所下的命令。
  - 插入模式(Insert-mode)：
    - Normal模式下，使用快捷键 i 进入insert模式，可以输入文本内容
    - 
  - 命令模式(Command-mode)
  - 可视模式(Visual-mode)
- `vi`：进入编辑器，默认进入的是vim的版本，是原始vi编辑器的扩展，是一个同vi向上兼容的文本编辑器
- `vim`：进入vim编辑器，或者`vim <filename>`以vim编辑器打开某个文件
  - esc回到Normal模式，有光标
  - 正常模式
    - h、j、k、l：左、下、上、右，移动光标。在图形界面或远程终端上，如果有 左、下、上、右（箭头）方向键，效果是一样的，但是如果是字符终端，可能会出现乱码
    - 复制
      - 一行：Normal模式下，按yy可以复制一行内容
      - 多行：Normal模式下，按3yy即可复制当前光标所在行开始的3行
      - 光标到行尾：y$
    - 剪切：dd、d$，与y类似
    - 粘贴：按p键，可以在光标所在行的下一行粘贴，继续按p键可以粘贴多次
    - 撤销：u键
    - 重做：把撤销的内容恢复，ctrl+R
    - 删除单个字符：x
    - 替换单个字符：按r键，在输入新的字符
    - 显示行数：`:set nu`
    - 移动到指定行：5G移动到第5行，g移动到第一行，G移动到最后一行
    - ^，或者说shift+6：移动光标到一行开头
    - $，或者或shift+4：移动光标到一行末尾
  - insert
    - Normal模式下，使用快捷键 i 进入insert模式，光标位置不变
    - 大写的`I`，进入插入模式，光标将从，光标将移动到光标所在行的开头
    - 小写a：光标右移一位，
    - 大写A：光标行最右
    - o：光标到下一行，并且是新开空行
    - O：光标到上一行，并且是新开空行
  - 命令模式：
    - `:`：进入命令模式，窗口末尾显示`:`时即表示处于命令模式了，可以键入命令
      - 按esc退出命令模式
      - `w <filename>`：保存文件到目标，如`w /tmp/a.txt`，如果是通过文件名打开的vim编辑器，只需要`w`命令即可保存修改
      - `q`：退出vim编辑器
      - `wq`、`wq <filename>`：连用，即保存并退出
      - `q!`：不保存并退出
      - `!<命令>`：如果想要临时执行一些系统命令，查看系统命令的结果，可以通过该命令来进行，如`!ip addr`查看ip地址，按`回车`即可重新回到编辑器界面
    - `/字符串`：查找指定字符串，按n跳到下一个，shift+n跳到上一个
    - `:s/旧串/新串`：替换光标所在的行的匹配的字符
    - `:%s/旧串/新串`：替换全文匹配的字符第一个字符
    - `:%s/旧串/新串/g`：替换全文匹配的字符所有
    - `:set nu`：显示行数。只单次生效，需要修改vim配置以长期生效
      - `vim /etc/vimrc`：移动到最后一行插入内容`set nu`，保存退出即可
    - `:set nonu`：不显示行数
  - Visual：可视模式 
    - v：字符可视模式
    - V：行可视模式
    - ctrl+v：块可视模式（类似于列模式）
      - `d`：删除选中块
      - `ctrl+i`：进入插入模式，输入要插入的内容，按2次esc，可以批量给选中的每一行前面添加相同内容



## 文件系统

Linux支持多种文件系统，常见的有

- ext4（centos6默认）
- xfs（centos7默认）
- NTFS （需安装额外软件ntfs-3g，有版权的，windows用）

### ext4

结构

- **超级块**：会记录文件数，有副本
- **超级块副本**：恢复数据用
- **inode**：i 节点，记录每一个文件的信息，权限、编号等。文件名与编号不在同一inode，而是记录在其父目录的inode中
  - `ls`查看的是inode中的文件信息，`ls -i`可以查到每一个文件的inode
  - vim写文件会改变inode，vim编辑时，会在家目录下复制一份临时文件进行修改，然后保存复制到源目录，并删除原来的文件（应该是）
  - `rm testfile`是从父inode删除文件名，所以文件再大都是秒删
  - `ln afile bfile`将bfile指向afile并也加入到父inode，且不会占用额外空间
- **datablock**：数据块，存放文件的数据内容，挂在inode上的，如果一个数据块不够就接着往后挂，链式。默认创建的数据块为4k，即使只写了1个字符也是4k，所以存储大量小文件会很费磁盘，所以网络上有一些专门用来存储小文件的文件系统
  - `du`统计的是数据块的个数用来计算大小
  - `echo >`写入文件只会改变数据块
- **软连接**：亦称符号链接，类似快捷方式。`ln -s afile cfile`链接两个文件，`ls -li afile cfile`查看两个文件链接信息，其实cfile就记录了目标文件afile的路径，链接文件的权限修改对其自身是无意义的，对其权限修改将在目标文件上得到反馈。可以跨分区（跨文件系统）
- facl：文件访问控制，getfacl afile查看文件权限，setfacl -m u:user1:r afile：u表示为用户分配权限，g表示用户组，r表示读权限，m改成x即可收回对应用户（组）权限

- 磁盘配额的使用：给多个用户之间磁盘使用做限制



### mkfs

- `mkfs`：make filesystem，创建文件系统。实际上是一个综合的指令，它会去调用正确的文件系统格式化工具软件。常说的“格式化”其实就是“make filesystem”，创建的其实是 xfs 文件系统， 因此使用的是 mkfs.xfs 这个指令

- `mkfs.xfs [-b bsize] [-d parms] [-i parms] [-l parms] [-L label] [-f] \ [-r parms]`：加单位则为是Bytes值，可以用 k,m,g,t,p （小写）等来解释，s指的是 sector 个数
  - `-b`：block 容量，可由 512 到 64k，不过最大容量限制为 Linux 的 4k 
  - `-d`：data section 的相关参数值
    - `agcount=数值`：设置需要几个储存群组的意思（AG），通常与 CPU 有关
    - `agsize=数值`：每个 AG 设置为多少容量的意思，通常 agcount/agsize 只选一个设置即可
    - `file`：指的是“格式化的设备是个文件而不是个设备”的意思！（例如虚拟磁盘）
    - `size=数值`：data section 的容量，亦即你可以不将全部的设备容量用完的意思
    - `su=数值`：当有 RAID 时，那个 stripe 数值的意思，与下面的 sw 搭配使用
    - `sw=数值`：当有 RAID 时，用于储存数据的磁盘数量（须扣除备份碟与备用碟）
    - `sunit=数值`：与 su 相当，不过单位使用的是“几个 sector（512Bytes大小）”的意思
    - `swidth=数值`：就是 su*sw 的数值，但是以“几个 sector（512Bytes大小）”来设置
  - `-f`：如果设备内已经有文件系统，则需要使用这个 -f 来强制格式化才行
  - `-i`：与 inode 有较相关的设置，主要的设置值有
    - `size=数值`：最小256Bytes，最大2k，一般保留 256 就足够使用了
    - `internal=[0&#124;1]`：log 设备是否为内置？默认为 1 内置，如果要用外部设备，使用下面设置
    - `logdev=device`：log 设备为后面接的那个设备上头的意思，需设置 internal=0 才可
    - `size=数值`：指定这块登录区的容量，通常最小得要有 512 个 block，大约 2M 以上才行
  - `-L`：后面接这个文件系统的标头名称 Label name 的意思
- `-r`：指定 realtime section 的相关设置值，常见的有： 
  - `extsize=数值`：就是那个重要的 extent 数值，一般不须设置，但有 RAID 时，最好设置与 swidth的数值相同较佳！最小为 4K 最大为 1G 。

```sh
# 不带参数将使用默认值
mkfs.xfs /dev/sdb3
blkid /dev/sdb*
```

因为 xfs 可以使用多个数据流来读写系统，以增加速度，因此那个 agcount 可以跟 CPU 的核心数来做搭配！举例来说，如果我的服务器仅有一颗 4 核心，但是有启动 Intel 超线程功能，则系统会仿真 出 8 颗 CPU 时，那个 agcount 就可以设置为 8 喔

```sh
# 找出系统的 CPU 数，并据以设置 agcount 数值
grep 'processor' /proc/cpuinfo
mkfs.xfs -f -d agcount=2 /dev/sdb4
```



### 分区

如果是虚拟机，可以直接在virtualbox上给其添加一块硬盘进行练习（比如叫sdc），可能需要关机才能添加

### 挂载

如果是虚拟机，可以直接在virtualbox上给其添加一块硬盘进行练习（比如叫sdc），可能需要关机才能添加

- 磁盘的分区与挂载：
  - 常用命令
    - mkfs：使用分区，输入mkfs.可以看到有很多不同后缀，都是指不同的文件系统，比如mkfd.ext4 /dev/sdc2即可格式化为ext4的文件系统。但是文件操作是文件系统之上的操作，无法直接操作，需要将其挂载到某个目录，对目录进行操作，
    - `mount`：mount -t auto自动检测文件系统，或者直接mount /dev/sdc2 /mnt/sdc2也会自动检测，将/dev/sd2挂载到/mnt/sd2，但是是临时挂载，vim /etc/fstab进行修改，dev/sdc1 /mnt/sdc1 ext4 defaults 0 0，即磁盘目录 挂载目录 文件系统指定 权限(defauls表示可读写) 磁盘配额相关参数1  磁盘配额相关参数2
      - `-t`：指定档案系统的型态，通常不必指定，`mount` 会自动选择正确的型态，或者 `mount -t auto` 也会自动检测文件系统类型。
    - `mount -t proc proc /proc`：把proc这个虚拟文件系统挂载到/proc目录，mount的标准用法是 `mount -t type device dir `，但是内核代码中，proc filesystem根本没有处理 dev_name这个参数，所以传什么都没有影响，只是影响mount命令的输出内容，好的实践应该将设备名定义为 nodev、none 等，或者就叫 proc 亦可。
    - parted：如果磁盘大于2T，不要用fdisk进行分区，而是parted，parted /dev/sdd

  - 用户磁盘配额：限制用户对磁盘的使用，比如创建文件数量（即限制i节点数）、数据块数量

    - xfs文件系统的用户磁盘配额quota，修改步骤如下

      - mkfs.xfs /dev/sdb1：创建分区，如果分区已经存在，为了防止这是一个误操作，会提示使用-f参数强制覆盖，mkfs.xfs -f /dev/sdb1
      - mkdir /mnt/disk1
      - mount -o uquota,gquota /dev/sdb1 /mnt/disk1，-o开启磁盘配额，uquota表示支持用户磁盘配额，gquota表示支持用户组磁盘配额
      - chmod 1777 /mnt/disk1：赋予1777权限
      - xfs_quota -x -c 'report -ugibh'/mnt/disk1：有参数时直接非交互配置，xfs_quota可以直接进入交互模式，但是一般非交互即可，
        - -c表示命令，report -ugibh，report表示报告（查看）磁盘配额，-u表示用户磁盘配额，g表示组磁盘配额，i表示节点，b表示块，h可以更人性化显示

      - xfs_quota -x -c 'limit -u isoft=5 ihard=10 user1' /mnt/disk1：root是无限制的，不要对root进行磁盘配额，没意义。这里对user1进行磁盘配额，limit表示限制磁盘配额，限制用户磁盘配额加-u，限制组磁盘配额加-g；isoft软限制i节点，ihard将硬限制，软限制比硬限制的配置的值更小，达到软限制之后，会提示用户在某一个宽限的时间条件内可以用超过软限制的值，硬限制则绝对不能超过限制的值；数据块限制即bsoft、bhard

  - 常见配置文件

    - letc/fstab
- 交换分区（虚拟内存）的查看与创建

  - free查看mem和swap，前面提到过
  - 增加交换分区的大小，使用硬盘分区扩充swap
    - mkswap：如mkswap /dev/sdd1将标记上swap
    - swapon：swapon /dev/sdd1打开swap，通过free可以看到swap被扩充了，swapoff /dev/sdd1关闭swap
  - 使用文件制作交换分区：可以直接创建一个比如10G的文件，或者创建带有空洞的文件，使其在swap的使用过程中逐渐扩大也可以
    - dd if=/dev/zero bs=4M count=1024 of=/swapfile：创建文件
    - mkswap /swapfile：即可为文件打上swap标记使其成为swap的空间
    - chmod 600 /swapfile：为了安全起见，一般修改为600权限
    - swapon、swapoff
  - 同样swap设置也是临时的，vi /etc/fstab。/swapfile swap swap defaults 0 0，即 swap文件或分区 挂载到swap目录（这是一个虚拟目录，因为swap不需要用文件目录来进行操作，挂载到这个虚拟目录即可） 文件系统格式（也是swap），第一个0表示做dump备份时要不要备份该硬盘（分区），但是现在一般都是tar进行备份，所以设置为0即可，第二个0表示开机的时候进行磁盘的自检的顺序问题，是针对之前的ext2、ext3的文件系统的设置，但是现在已经不需要了，如果发现写入是不完整的自动会对那个分区进行检查，所以也是0即可
    - 如果写错了东西，发现重启启动不起来了，通过grap进入到单一用户模式，来去修改/etc/fstab