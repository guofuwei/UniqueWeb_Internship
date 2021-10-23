# **Linux**基本命令

## Linux系统文件介绍

### 所有的目录

- **/bin**：
  bin 是 Binaries (二进制文件) 的缩写, 这个目录存放着最经常使用的命令。

- **/boot：**
  这里存放的是启动 Linux 时使用的一些核心文件，包括一些连接文件以及镜像文件。

- **/dev ：**
  dev 是 Device(设备) 的缩写, 该目录下存放的是 Linux 的外部设备，在 Linux 中访问设备的方式和访问文件的方式是相同的。

- **/etc：**
  etc 是 Etcetera(等等) 的缩写,这个目录用来存放所有的系统管理所需要的配置文件和子目录。

- **/home**：
  用户的主目录，在 Linux 中，每个用户都有一个自己的目录，一般该目录名是以用户的账号命名的，如上图中的 alice、bob 和 eve。

- **/lib**：
  lib 是 Library(库) 的缩写这个目录里存放着系统最基本的动态连接共享库，其作用类似于 Windows 里的 DLL 文件。几乎所有的应用程序都需要用到这些共享库。

- **/lost+found**：
  这个目录一般情况下是空的，当系统非法关机后，这里就存放了一些文件。

- **/media**：
  linux 系统会自动识别一些设备，例如U盘、光驱等等，当识别后，Linux 会把识别的设备挂载到这个目录下。

- **/mnt**：
  系统提供该目录是为了让用户临时挂载别的文件系统的，我们可以将光驱挂载在 /mnt/ 上，然后进入该目录就可以查看光驱里的内容了。

- **/opt**：
  opt 是 optional(可选) 的缩写，这是给主机额外安装软件所摆放的目录。比如你安装一个ORACLE数据库则就可以放到这个目录下。默认是空的。

- **/proc**：
  proc 是 Processes(进程) 的缩写，/proc 是一种伪文件系统（也即虚拟文件系统），存储的是当前内核运行状态的一系列特殊文件，这个目录是一个虚拟的目录，它是系统内存的映射，我们可以通过直接访问这个目录来获取系统信息。
  这个目录的内容不在硬盘上而是在内存里，我们也可以直接修改里面的某些文件，比如可以通过下面的命令来屏蔽主机的ping命令，使别人无法ping你的机器：

  ```
  echo 1 > /proc/sys/net/ipv4/icmp_echo_ignore_all
  ```

- **/root**：
  该目录为系统管理员，也称作超级权限者的用户主目录。

- **/sbin**：
  s 就是 Super User 的意思，是 Superuser Binaries (超级用户的二进制文件) 的缩写，这里存放的是系统管理员使用的系统管理程序。

- **/selinux**：
   这个目录是 Redhat/CentOS 所特有的目录，Selinux 是一个安全机制，类似于 windows 的防火墙，但是这套机制比较复杂，这个目录就是存放selinux相关的文件的。

- **/srv**：
   该目录存放一些服务启动之后需要提取的数据。

- **/sys**：

  这是 Linux2.6 内核的一个很大的变化。该目录下安装了 2.6 内核中新出现的一个文件系统 sysfs 。

  sysfs 文件系统集成了下面3种文件系统的信息：针对进程信息的 proc 文件系统、针对设备的 devfs 文件系统以及针对伪终端的 devpts 文件系统。

  该文件系统是内核设备树的一个直观反映。

  当一个内核对象被创建的时候，对应的文件和目录也在内核对象子系统中被创建。

- **/tmp**：
  tmp 是 temporary(临时) 的缩写这个目录是用来存放一些临时文件的。

- **/usr**：
   usr 是 unix shared resources(共享资源) 的缩写，这是一个非常重要的目录，用户的很多应用程序和文件都放在这个目录下，类似于 windows 下的 program files 目录。

- **/usr/bin：**
  系统用户使用的应用程序。

- **/usr/sbin：**
  超级用户使用的比较高级的管理程序和系统守护程序。

- **/usr/src：**
  内核源代码默认的放置目录。

- **/var**：
  var 是 variable(变量) 的缩写，这个目录中存放着在不断扩充着的东西，我们习惯将那些经常被修改的目录放在这个目录下。包括各种日志文件。

- **/run**：
  是一个临时文件系统，存储系统启动以来的信息。当系统重启时，这个目录下的文件应该被删掉或清除。如果你的系统上有 /var/run 目录，应该让它指向 run。

### 常用的重要目录

**/etc**： 上边也提到了，这个是系统中的配置文件，如果你更改了该目录下的某个文件可能会导致系统不能启动。

**/bin, /sbin, /usr/bin, /usr/sbin**: 这是系统预设的执行文件的放置目录，比如 ls 就是在 /bin/ls 目录下的。

值得提出的是，/bin, /usr/bin 是给系统用户使用的指令（除root外的通用户），而/sbin, /usr/sbin 则是给 root 使用的指令。

**/var**： 这是一个非常重要的目录，系统上跑了很多程序，那么每个程序都会有相应的日志产生，而这些日志就被记录到这个目录下，具体在 /var/log 目录下，另外 mail 的预设放置也是在这里。

-------------------------

------------------

## Linux文件基本属性

- chown (change ownerp) ： 修改所属用户与组。
- chmod (change mode) ： 修改用户的权限。

## 文件权限

![image-20211012201949888](C:\Users\26254\AppData\Roaming\Typora\typora-user-images\image-20211012201949888.png)

```shell
Linux 中第一个字符代表这个文件是目录、文件或链接文件
当为 d 则是目录
当为 - 则是文件；
若是 l 则表示为链接文档(link file)；
若是 b 则表示为装置文件里面的可供储存的接口设备(可随机存取装置)；
若是 c 则表示为装置文件里面的串行端口设备，例如键盘、鼠标(一次性读取装置)。
```

```shell
接下来的三个为一组均为rwx的参数组合：r(可读),w(可写),x(可执行),-(无权限)
```

![img](https://www.runoob.com/wp-content/uploads/2014/06/file-llls22.jpg)

**第一个rwx字符：对于一个文件来说，其有一个特定的所有者，他所拥有的权限是第一个rwx字符**

**第二个rwx字符：文件的所有者的同组成员**

**第三个rwx字符：除所有者和所有者的同组成员的所有其他成员**

**一般文件权限对root用户不起作用**

**一般来说，所有者权限>=所有者同组成员权限>=其他用户权限**



## 更改文件属性

### 1.chgrp更改文件属组

```shell
# 语法：
chgrp [-R] 属组名 文件名
-R：递归更改文件属组，就是在更改某个目录文件的属组时，如果加上-R的参数，那么该目录下的所有文件的属组都会更改。
```

### 2.chown：更改文件属主，也可以同时更改文件属组

```shell
# 语法
chown [–R] 属主名 文件名
chown [-R] 属主名:组名 文件名
eg: chown [-R] root test.py
eg: chown [-R] root:guofuwei test2.py
```

### 3.chmod：更改文件9个属性

- r:4
- w:2
- x:1

````shell
chmod 777 test.log
chmod [-R] test
````

#### 符号类型改变文件权限

```shell
# 通用的格式
chmod u/g/o/a +/-/= r/w/x 文件或目录
eg:chmod g +w test.go
```

## Linux文件与目录管理

### 处理目录的常见命令

- ls（英文全拼：list files）: 列出目录及文件名
  - -a ：全部的文件，连同隐藏文件( 开头为 . 的文件) 一起列出来(常用)
  - -d ：仅列出目录本身，而不是列出目录内的文件数据(常用)
  - -l ：长数据串列出，包含文件的属性与权限等等数据；(常用)
- cd（英文全拼：change directory）：切换目录
- pwd（英文全拼：print work directory）：显示目前的目录
- mkdir（英文全拼：make directory）：创建一个新的目录
  - -m ：配置文件的权限喔！直接配置，不需要看默认权限 (umask) 的脸色～
  - -p ：帮助你直接将所需要的目录(包含上一级目录)递归创建起来！
- rmdir（英文全拼：remove directory）：删除一个**空的目录**
  - **-p ：**从该目录起，一次删除**多级空目录**
- cp（英文全拼：copy file）: 复制文件或目录
  - **-r：**递归持续复制，用於目录的复制行为；(常用)
  - **-p：**连同文件的属性一起复制过去，而非使用默认属性(备份常用)；
- rm（英文全拼：remove）: 删除文件或目录
  - -f ：就是 force 的意思，忽略不存在的文件，不会出现警告信息；
  - -i ：互动模式，在删除前会询问使用者是否动作
  - **-r ：递归删除啊！**最常用在目录的删除了！这是非常危险的选项！！！
- mv（英文全拼：move file）: 移动文件与目录，或修改文s件与目录的名称
  - -f ：force 强制的意思，如果目标文件已经存在，不会询问而直接覆盖；
  - -i ：若目标文件 (destination) 已经存在时，就会询问是否覆盖！
  - -u ：若目标文件已经存在，且 source 比较新，才会升级 (update)

### Linux 文件内容查看

- cat 由第一行开始显示文件内容
- tac 从最后一行开始显示，可以看出 tac 是 cat 的倒着写！
- nl  显示的时候，顺道输出行号！
- more 一页一页的显示文件内容
  - 空白键 (space)：代表向下翻一页；
  - Enter     ：代表向下翻『一行』；
  - /字串     ：代表在这个显示的内容当中，向下搜寻『字串』这个关键字；
  - :f      ：立刻显示出档名以及目前显示的行数；
  - q       ：代表立刻离开 more ，不再显示该文件内容。
  - b 或 [ctrl]-b ：代表往回翻页，不过这动作只对文件有用，对管线无用
- less 与 more 类似，但是比 more 更好的是，他可以往前翻页！
- **head 只看头几行**
  - -n ：后面接数字，代表显示几行的意思
- **tail 只看尾巴几行**
  - -n ：后面接数字，代表显示几行的意思
  - -f ：表示持续侦测后面所接的档名，要等到按下[ctrl]-c才会结束tail的侦测

### 软连接和硬连接

```shell
# 通过ln命令来创建
ln -s f1 f3 //创建f1的软连接文件f3
ln f1 f2 //创建f1的硬连接文件f2(类似于windows的快捷方式)
示例:
ubuntu@LAPTOP-2MJNJKV9:~$ ln f1 f2
ubuntu@LAPTOP-2MJNJKV9:~$ ln -s f1 f3
ubuntu@LAPTOP-2MJNJKV9:~$ ls -l
total 0
-rw-r--r-- 2 ubuntu ubuntu 0 Oct 15 22:06 f1
-rw-r--r-- 2 ubuntu ubuntu 0 Oct 15 22:06 f2
lrwxrwxrwx 1 ubuntu ubuntu 2 Oct 15 22:06 f3 -> f1
当删除f1后:
ubuntu@LAPTOP-2MJNJKV9:~$ rm f1
ubuntu@LAPTOP-2MJNJKV9:~$ cat f2
ubuntu@LAPTOP-2MJNJKV9:~$ cat f3
cat: f3: No such file or directory
```



## Linux 用户和用户组管理

### 一.Linux系统用户账号的管理

#### 1.添加新的用户

```shell
useradd 选项 用户
useradd -d /home/sam sam
useradd -s /bin/bash -g group -G adm,root gem
```

- 选项:

  - -c comment 指定一段注释性描述。
  - -d 目录 指定用户主目录，如果此目录不存在，则同时使用-m选项，可以创建主目录。
  - -g 用户组 指定用户所属的用户组。
  - -G 用户组，用户组 指定用户所属的附加组。
  - -s Shell文件 指定用户的登录Shell。
  - -u 用户号 指定用户的用户号，如果同时有-o选项，则可以重复使用其他用户的标识号。

- 用户名:

  指定新账号的登录名。

#### 2.删除新的用户

```shell
userdel 选型 用户名
常用的选项是 -r，它的作用是把用户的主目录一起删除。
eg:
userdel -r sam
此命令删除用户sam在系统文件中（主要是/etc/passwd, /etc/shadow, /etc/group等）的记录，同时删除用户的主目录。
```

#### 3.修改用户

* 修改用户账号就是根据实际情况更改用户的有关属性，如用户号、主目录、用户组、登录Shell等。

```shell
usermod 选项 用户名
```

选项和useradd基本相同

添加 -l 新用户名 选项

```shel
usermod -s /bin/dash -d /home/test -g dev -l tom sam
```

#### 4.用户口令（密码）的管理

```shell
修改密码：
passwd
# 超级用户可以为自己和其他用户指定口令，普通用户只能用它修改自己的口令。
选项：
-l 锁定口令，即禁用账号。
-u 口令解锁。
-d 使账号无口令。
-f 强迫用户下次登录时修改口令。
```



### 二、Linux系统用户组的管理

#### 1.增加一个用户组

```shell
groupadd 选项 用户组
# 选项：
-g GID 指定新用户组的组标识号（GID）。
-o 一般与-g选项同时使用，表示新用户组的GID可以与系统已有用户组的GID相同。
eg:
groupadd sudoer
groupadd -g 101 common
```

#### 2.删除一个用户组

```shell
groupdel 用户组
eg:
groupdel sudoer
```

#### 3.修改一个用户组

```shell
groupmod 选项 用户组
选项：
-g GID 为用户组指定新的组标识号。
-o 与-g选项同时使用，用户组的新GID可以与系统已有用户组的GID相同。
-n新用户组 将用户组的名字改为新名字
eg:
groupmod -g 102 -n group3 group2//修改group2GID为102，重命名为group3
```

#### 4.切换用户组

```shell
用newgrp进行切换
eg：
newgrp root//切换到root的用户组
前提条件是root用户组确实是该用户的主组或附加组
```

### 三、与用户账号有关的系统文件

#### 1.用户的密码保存文件/etc/passwd

```shell
# ubuntu@VM-4-7-ubuntu:~$ cat /etc/passwd

root:x:0:0:root:/root:/bin/bash
daemon:x:1:1:daemon:/usr/sbin:/usr/sbin/nologin
bin:x:2:2:bin:/bin:/usr/sbin/nologin
sys:x:3:3:sys:/dev:/usr/sbin/nologin
sync:x:4:65534:sync:/bin:/bin/sync
games:x:5:60:games:/usr/games:/usr/sbin/nologin
man:x:6:12:man:/var/cache/man:/usr/sbin/nologin
lp:x:7:7:lp:/var/spool/lpd:/usr/sbin/nologin
mail:x:8:8:mail:/var/mail:/usr/sbin/nologin
news:x:9:9:news:/var/spool/news:/usr/sbin/nologin
uucp:x:10:10:uucp:/var/spool/uucp:/usr/sbin/nologin
proxy:x:13:13:proxy:/bin:/usr/sbin/nologin
www-data:x:33:33:www-data:/var/www:/usr/sbin/nologin
backup:x:34:34:backup:/var/backups:/usr/sbin/nologin
list:x:38:38:Mailing List Manager:/var/list:/usr/sbin/nologin
irc:x:39:39:ircd:/var/run/ircd:/usr/sbin/nologin
gnats:x:41:41:Gnats Bug-Reporting System (admin):/var/lib/gnats:/usr/sbin/nologin
nobody:x:65534:65534:nobody:/nonexistent:/usr/sbin/nologin
systemd-network:x:100:102:systemd Network Management,,,:/run/systemd:/usr/sbin/nologin
systemd-resolve:x:101:103:systemd Resolver,,,:/run/systemd:/usr/sbin/nologin
systemd-timesync:x:102:104:systemd Time Synchronization,,,:/run/systemd:/usr/sbin/nologin
messagebus:x:103:106::/nonexistent:/usr/sbin/nologin
syslog:x:104:110::/home/syslog:/usr/sbin/nologin
_apt:x:105:65534::/nonexistent:/usr/sbin/nologin
tss:x:106:111:TPM software stack,,,:/var/lib/tpm:/bin/false
uuidd:x:107:112::/run/uuidd:/usr/sbin/nologin
tcpdump:x:108:113::/nonexistent:/usr/sbin/nologin
landscape:x:109:115::/var/lib/landscape:/usr/sbin/nologin
pollinate:x:110:1::/var/cache/pollinate:/bin/false
sshd:x:111:65534::/run/sshd:/usr/sbin/nologin
systemd-coredump:x:999:999:systemd Core Dumper:/:/usr/sbin/nologin
ubuntu:x:1000:1000:ubuntu:/home/ubuntu:/bin/bash
lxd:x:998:100::/var/snap/lxd/common/lxd:/bin/false
ntp:x:112:117::/nonexistent:/usr/sbin/nologin
_rpc:x:113:65534::/run/rpcbind:/usr/sbin/nologin
statd:x:114:65534::/var/lib/nfs:/usr/sbin/nologin
bind:x:115:118::/var/cache/bind:/usr/sbin/nologin
lighthouse:x:1001:1001::/home/lighthouse:/bin/bash
mysql:x:116:120:MySQL Server,,,:/nonexistent:/bin/false
redis:x:117:121::/var/lib/redis:/usr/sbin/nologin
celery:x:1002:1002::/home/celery:/bin/sh

每行七段所代表的意义：
用户名:口令:用户标识号:组标识号:注释性描述:主目录:登录Shell
celery:x:1002:1002::/home/celery:/bin/sh

# 补充说明：
1.通常用户标识号的取值范围是0～65 535。0是超级用户root的标识号，1～99由系统保留，作为管理账号，普通用户的标识号从100开始。在Linux系统中，这个界限是500。
2.登录Shell为空无法登录的伪用户：
	常见的伪用户
    bin 拥有可执行的用户命令文件 
    sys 拥有系统文件 
    adm 拥有帐户文件 
    uucp UUCP使用 
    lp lp或lpd子系统使用 
    nobody NFS使用
    audit, cron, mail, usenet它们也都各自为相关的进程和文件所需要。
```

#### 2./etc/shadow文件于/etc/passwd每一行的数据相对应

```shell
# cat /etc/shadow

daemon:*:18375:0:99999:7:::
bin:*:18375:0:99999:7:::
sys:*:18375:0:99999:7:::
sync:*:18375:0:99999:7:::
games:*:18375:0:99999:7:::
man:*:18375:0:99999:7:::

# 格式为：
登录名:加密口令:最后一次修改时间:最小时间间隔:最大时间间隔:警告时间:不活动时间:失效时间:标志
```

1. "登录名"是与/etc/passwd文件中的登录名相一致的用户账号
2. "口令"字段存放的是加密后的用户口令字，长度为13个字符。如果为空，则对应用户没有口令，登录时不需要口令；如果含有不属于集合 { ./0-9A-Za-z }中的字符，则对应的用户不能登录。
3. "最后一次修改时间"表示的是从某个时刻起，到用户最后一次修改口令时的天数。时间起点对不同的系统可能不一样。例如在SCO Linux 中，这个时间起点是1970年1月1日。
4. "最小时间间隔"指的是两次修改口令之间所需的最小天数。
5. "最大时间间隔"指的是口令保持有效的最大天数。
6. "警告时间"字段表示的是从系统开始警告用户到用户密码正式失效之间的天数。
7. "不活动时间"表示的是用户没有登录活动但账号仍能保持有效的最大天数。
8. "失效时间"字段给出的是一个绝对的天数，如果使用了这个字段，那么就给出相应账号的生存期。期满后，该账号就不再是一个合法的账号，也就不能再用来登录了。



#### 3、用户组的所有信息都存放在/etc/group文件中

注意：一个用户可以属于多个组，/etc/passwd中记录的是主组，其他的组为附加组，可以用newgrp从主组切换到附加组

```shell
# 示例：
ubuntu@VM-4-7-ubuntu:~$ cat /etc/group

root:x:0:
daemon:x:1:
bin:x:2:
sys:x:3:
adm:x:4:syslog,ubuntu
tty:x:5:
disk:x:6:
lp:x:7:
mail:x:8:
news:x:9:
uucp:x:10:
man:x:12:
proxy:x:13:
kmem:x:15:
dialout:x:20:
fax:x:21:
voice:x:22:


# 其中的组成为：
组名:口令:组标识号:组内用户列表
"组名"是用户组的名称，由字母或数字构成。与/etc/passwd中的登录名一样，组名不应重复。
"口令"字段存放的是用户组加密后的口令字。一般Linux 系统的用户组都没有口令，即这个字段一般为空，或者是*。
"组标识号"与用户标识号类似，也是一个整数，被系统内部用来标识组。
"组内用户列表"是属于这个组的所有用户的列表，不同用户之间用逗号(,)分隔。这个用户组可能是用户的主组，也可能是附加组。
```



#### 4.批量添加用户

##### （1）先编辑一个文本用户文件

```shell
# 每一列按照/etc/passwd密码文件的格式书写
# user.txt示例：
user001::600:100:user:/home/user001:/bin/bash
user002::601:100:user:/home/user002:/bin/bash
user003::602:100:user:/home/user003:/bin/bash
user004::603:100:user:/home/user004:/bin/bash
user005::604:100:user:/home/user005:/bin/bash
user006::605:100:user:/home/user006:/bin/bash
```

##### （2）以root身份执行命令 `/usr/sbin/newusers`，从刚创建的用户文件`user.txt`中导入数据，创建用户

```shell
# newusers < user.txt
```

##### （3）执行命令/usr/sbin/pwunconv

将 `/etc/shadow` 产生的 `shadow` 密码解码，然后回写到 `/etc/passwd` 中，并将`/etc/shadow`的`shadow`密码栏删掉。这是为了方便下一步的密码转换工作，即先取消 `shadow password` 功能。

```shell
# pwunconv
```

##### （4）编辑每个用户的密码对照文件

```shell
# 格式为：
用户名:密码
# passwd.txt示例：
user001:123456
user002:123456
user003:123456
user004:123456
user005:123456
user006:123456
```

##### （5）以 root 身份执行命令 `/usr/sbin/chpasswd`

创建用户密码，`chpasswd` 会将经过 `/usr/bin/passwd` 命令编码过的密码写入 `/etc/passwd` 的密码栏

```shell
# chpasswd < passwd.txt
```

##### （6）确定密码经编码写入/etc/passwd的密码栏后

执行命令 `/usr/sbin/pwconv` 将密码编码为 `shadow password`，并将结果写入 `/etc/shadow`

```shell
# pwconv
```



## Linux 磁盘管理

Linux 磁盘管理常用三个命令为 **df**、**du** 和 **fdisk**。

- **df**（英文全称：disk full）：列出文件系统的整体磁盘使用量
- **du**（英文全称：disk used）：检查磁盘空间使用量
- **fdisk**：用于磁盘分区

### df

```shell
df [-ahikHTm] [目录或文件名]
```

- 选项与参数：
  - -a ：列出所有的文件系统，包括系统特有的 /proc 等文件系统；
  - -k ：以 KBytes 的容量显示各文件系统；
  - -m ：以 MBytes 的容量显示各文件系统；
  - -h ：以人们较易阅读的 GBytes, MBytes, KBytes 等格式自行显示；
  - -H ：以 M=1000K 取代 M=1024K 的进位方式；
  - -T ：显示文件系统类型, 连同该 partition 的 filesystem 名称 (例如 ext3) 也列出；
  - -i ：不用硬盘容量，而以 inode 的数量来显示

```shell
ubuntu@VM-4-7-ubuntu:~$ df -h

Filesystem      Size  Used Avail Use% Mounted on
udev            1.9G     0  1.9G   0% /dev
tmpfs           394M  924K  393M   1% /run
/dev/vda2        79G   11G   65G  15% /
tmpfs           2.0G   56K  2.0G   1% /dev/shm
tmpfs           5.0M     0  5.0M   0% /run/lock
tmpfs           2.0G     0  2.0G   0% /sys/fs/cgroup
tmpfs           394M     0  394M   0% /run/user/1000
overlay          79G   11G   65G  15% /var/lib/docker/overlay2/64f37e25c4edcd4f2d88bba952c633bfc8b6fd39691ea2d5c393d8058f3b82fa/merged
overlay          79G   11G   65G  15% /var/lib/docker/overlay2/56a07d5b1f929fad3a0d162395122c8c5511a8cc4932f7910e6b3739427ced48/merged
```



### du

与 df 命令不同的是 Linux du 命令是对文件和目录磁盘使用的空间的查看

```shell
du [-ahskm] 文件或目录名称
```

选项与参数：

- -a ：列出所有的文件与目录容量，因为默认仅统计目录底下的文件量而已。
- -h ：以人们较易读的容量格式 (G/M) 显示；
- -s ：列出总量而已，而不列出每个各别的目录占用容量；
- -S ：不包括子目录下的总计，与 -s 有点差别。
- -k ：以 KBytes 列出容量显示；
- -m ：以 MBytes 列出容量显示；

```shell
ubuntu@VM-4-7-ubuntu:~$ du -h

20K	./blogsite/blogsitev1/myblog2/mytoken/__pycache__
8.0K	./blogsite/blogsitev1/myblog2/mytoken/migrations/__pycache__
4.0K	./blogsite/blogsitev1/myblog2/mytoken/migrations
24K	./blogsite/blogsitev1/myblog2/mytoken
28K	./blogsite/blogsitev1/myblog2/user/__pycache__
16K	./blogsite/blogsitev1/myblog2/user/migrations/__pycache__
12K	./blogsite/blogsitev1/myblog2/user/migrations
36K	./blogsite/blogsitev1/myblog2/user
24K	./blogsite/blogsitev1/myblog2
744K	./blogsite/blogsitev1/client/static/js
120K	./blogsite/blogsitev1/client/static/css
48M	./blogsite/blogsitev1/client/static/images
4.0K	./blogsite/blogsitev1/client/static
8.0K	./blogsite/blogsitev1/client/.idea/inspectionProfiles
20K	./blogsite/blogsitev1/client/.idea
144K	./blogsite/blogsitev1/client/templates
12K	./blogsite/blogsitev1/client/__pycache__
16K	./blogsite/blogsitev1/client
24K	./blogsite/blogsitev1
4.0K	./blogsite
8.0K	./.ssh
104K	.
............

ubuntu@VM-4-7-ubuntu:~$ du -sh
263M	.
```



### fdisk

fdisk 是 Linux 的磁盘分区表操作工具

```shell
fdisk [-l] 装置名称
```

选项与参数：

- -l ：输出后面接的装置所有的分区内容。若仅有 fdisk -l 时， 则系统将会把整个系统内能够搜寻到的装置的分区均列出来。

##### 实例 1

列出所有分区信息

```shell
# [root@AY120919111755c246621 tmp]# fdisk -l

Disk /dev/xvda: 21.5 GB, 21474836480 bytes
255 heads, 63 sectors/track, 2610 cylinders
Units = cylinders of 16065 * 512 = 8225280 bytes
Sector size (logical/physical): 512 bytes / 512 bytes
I/O size (minimum/optimal): 512 bytes / 512 bytes
Disk identifier: 0x00000000

    Device Boot      Start         End      Blocks   Id  System
/dev/xvda1   *           1        2550    20480000   83  Linux
/dev/xvda2            2550        2611      490496   82  Linux swap / Solaris

Disk /dev/xvdb: 21.5 GB, 21474836480 bytes
255 heads, 63 sectors/track, 2610 cylinders
Units = cylinders of 16065 * 512 = 8225280 bytes
Sector size (logical/physical): 512 bytes / 512 bytes
I/O size (minimum/optimal): 512 bytes / 512 bytes
Disk identifier: 0x56f40944

    Device Boot      Start         End      Blocks   Id  System
/dev/xvdb2               1        2610    20964793+  83  Linux
```

##### 实例 2

找出你系统中的根目录所在磁盘，并查阅该硬盘内的相关信息

```shell
# [root@www ~]# df /            <==注意：重点在找出磁盘文件名而已
Filesystem           1K-blocks      Used Available Use% Mounted on
/dev/hdc2              9920624   3823168   5585388  41% /

# [root@www ~]# fdisk /dev/hdc  <==仔细看，不要加上数字喔！
The number of cylinders for this disk is set to 5005.
There is nothing wrong with that, but this is larger than 1024,
and could in certain setups cause problems with:
1) software that runs at boot time (e.g., old versions of LILO)
2) booting and partitioning software from other OSs
   (e.g., DOS FDISK, OS/2 FDISK)

Command (m for help):     <==等待你的输入！
```

输入 m 后，就会看到底下这些命令介绍

```
Command (m for help): m   <== 输入 m 后，就会看到底下这些命令介绍
Command action
   a   toggle a bootable flag
   b   edit bsd disklabel
   c   toggle the dos compatibility flag
   d   delete a partition            <==删除一个partition
   l   list known partition types
   m   print this menu
   n   add a new partition           <==新增一个partition
   o   create a new empty DOS partition table
   p   print the partition table     <==在屏幕上显示分割表
   q   quit without saving changes   <==不储存离开fdisk程序
   s   create a new empty Sun disklabel
   t   change a partition's system id
   u   change display/entry units
   v   verify the partition table
   w   write table to disk and exit  <==将刚刚的动作写入分割表
   x   extra functionality (experts only)
```

离开 fdisk 时按下 `q`，那么所有的动作都不会生效！相反的， 按下`w`就是动作生效的意思。

```
Command (m for help): p  <== 这里可以输出目前磁盘的状态

Disk /dev/hdc: 41.1 GB, 41174138880 bytes        <==这个磁盘的文件名与容量
255 heads, 63 sectors/track, 5005 cylinders      <==磁头、扇区与磁柱大小
Units = cylinders of 16065 * 512 = 8225280 bytes <==每个磁柱的大小

   Device Boot      Start         End      Blocks   Id  System
/dev/hdc1   *           1          13      104391   83  Linux
/dev/hdc2              14        1288    10241437+  83  Linux
/dev/hdc3            1289        1925     5116702+  83  Linux
/dev/hdc4            1926        5005    24740100    5  Extended
/dev/hdc5            1926        2052     1020096   82  Linux swap / Solaris
# 装置文件名 启动区否 开始磁柱    结束磁柱  1K大小容量 磁盘分区槽内的系统

Command (m for help): q
```

想要不储存离开吗？按下 q 就对了！不要随便按 w 啊！

使用 `p` 可以列出目前这颗磁盘的分割表信息，这个信息的上半部在显示整体磁盘的状态。



### 磁盘格式化

```shell
mkfs [-t 文件系统格式] 装置文件名
```

选项与参数：

- -t ：可以接文件系统格式，例如 ext3, ext2, vfat 等(系统有支持才会生效)

#### 实例1：

查看mkfs支持的文件格式

```shell
[root@www ~]# mkfs[tab][tab]
mkfs         mkfs.cramfs  mkfs.ext2    mkfs.ext3    mkfs.msdos   mkfs.vfat
# 注意在mkfs后面按下两次[Tab]
```

#### 实例 2:

将分区 /dev/hdc6（可指定你自己的分区） 格式化为 ext3 文件系统：

```shell
[root@www ~]# mkfs -t ext3 /dev/hdc6
mke2fs 1.39 (29-May-2006)
Filesystem label=                <==这里指的是分割槽的名称(label)
OS type: Linux
Block size=4096 (log=2)          <==block 的大小配置为 4K 
Fragment size=4096 (log=2)
251392 inodes, 502023 blocks     <==由此配置决定的inode/block数量
25101 blocks (5.00%) reserved for the super user
First data block=0
Maximum filesystem blocks=515899392
16 block groups
32768 blocks per group, 32768 fragments per group
15712 inodes per group
Superblock backups stored on blocks:
        32768, 98304, 163840, 229376, 294912

Writing inode tables: done
Creating journal (8192 blocks): done <==有日志记录
Writing superblocks and filesystem accounting information: done

This filesystem will be automatically checked every 34 mounts or
180 days, whichever comes first.  Use tune2fs -c or -i to override.
# 这样就创建起来我们所需要的 Ext3 文件系统了！简单明了！
```



### 磁盘检验

fsck（file system check）用来检查和维护不一致的文件系统。

使用场景：若系统掉电或磁盘发生问题，可利用fsck命令对文件系统进行检查。

语法：

```shell
fsck [-t 文件系统] [-ACay] 装置名称
```

选项与参数：

- -t : 给定档案系统的型式，若在 /etc/fstab 中已有定义或 kernel 本身已支援的则不需加上此参数
- -s : 依序一个一个地执行 fsck 的指令来检查
- -A : 对/etc/fstab 中所有列出来的 分区（partition）做检查
- -C : 显示完整的检查进度
- -d : 打印出 e2fsck 的 debug 结果
- -p : 同时有 -A 条件时，同时有多个 fsck 的检查一起执行
- -R : 同时有 -A 条件时，省略 / 不检查
- -V : 详细显示模式
- -a : 如果检查有错则自动修复
- -r : 如果检查有错则由使用者回答是否修复
- -y : 选项指定检测每个文件是自动输入yes，在不确定那些是不正常的时候，可以执行 # fsck -y 全部检查修复。

#### 实例 1

查看系统有多少文件系统支持的 fsck 命令：

```
[root@www ~]# fsck[tab][tab]
fsck         fsck.cramfs  fsck.ext2    fsck.ext3    fsck.msdos   fsck.vfat
```

#### 实例 2

强制检测 /dev/hdc6 分区:

```
[root@www ~]# fsck -C -f -t ext3 /dev/hdc6 
fsck 1.39 (29-May-2006)
e2fsck 1.39 (29-May-2006)
Pass 1: Checking inodes, blocks, and sizes
Pass 2: Checking directory structure
Pass 3: Checking directory connectivity
Pass 4: Checking reference counts
Pass 5: Checking group summary information
vbird_logical: 11/251968 files (9.1% non-contiguous), 36926/1004046 blocks
```

如果没有加上 -f 的选项，则由于这个文件系统不曾出现问题，检查的经过非常快速！若加上 -f 强制检查，才会一项一项的显示过程。



## 磁盘挂载与取消挂载

Linux 的磁盘挂载使用 `mount` 命令，卸载使用 `umount` 命令

磁盘挂载语法：

```shell
mount [-t 文件系统] [-L Label名] [-o 额外选项] [-n]  装置文件名  挂载点
```

```shell
[root@www ~]# mkdir /mnt/hdc6
[root@www ~]# mount /dev/hdc6 /mnt/hdc6
[root@www ~]# df
Filesystem           1K-blocks      Used Available Use% Mounted on
.....中间省略.....
/dev/hdc6              1976312     42072   1833836   3% /mnt/hdc6
```

磁盘卸载命令 `umount` 语法：

```
umount [-fn] 装置文件名或挂载点
```

选项与参数：

- -f ：强制卸除！可用在类似网络文件系统 (NFS) 无法读取到的情况下；
- -n ：不升级 /etc/mtab 情况下卸除。

卸载/dev/hdc6

```shell
[root@www ~]# umount /dev/hdc6     
```



# Linux awk 命令

```shell
awk [选项参数] 'script' var=value file(s)
或
awk [选项参数] -f scriptfile var=value file(s)
```



# Linux grep命令

Linux grep 命令用于查找文件里符合条件的字符串。

grep 指令用于查找内容包含指定的范本样式的文件，如果发现某文件的内容符合所指定的范本样式，预设 grep 指令会把含有范本样式的那一列显示出来。**若不指定任何文件名称，或是所给予的文件名为 -，则 grep 指令会从标准输入设备读取数据。**

语法：

```shell
grep [-abcEFGhHilLnqrsvVwxy][-A<显示行数>][-B<显示列数>][-C<显示列数>][-d<进行动作>][-e<范本样式>][-f<范本文件>][--help][范本样式][文件或目录...]
```

* **-a 或 --text** : 不要忽略二进制的数据
* **-b 或 --byte-offset** : 在显示符合样式的那一行之前，标示出该行第一个字符的编号
* **-i 或 --ignore-case** : 忽略字符大小写的差别
* **-d <动作> 或 --directories=<动作>或者/r** : 当指定要查找的是目录而非文件时，必须使用这项参数，否则grep指令将回报信息并停止动作。

- **-v 或 --invert-match** : 显示不包含匹配文本的所有行。





# Linux sed命令

*Linux sed 命令是利用脚本来处理文本文件*

语法

```shell
sed [-hnV][-e<script>][-f<script文件>][文本文件]
```

**参数说明**：

- -e<script>或--expression=<script> 以选项中指定的script来处理输入的文本文件。
- -f<script文件>或--file=<script文件> 以选项中指定的script文件来处理输入的文本文件。
- -h或--help 显示帮助。
- **-n或--quiet或--silent 仅显示script处理后的结果。**
- -V或--version 显示版本信息。、
- **-i直接修改文件内容**

**动作说明**：

- a ：新增， a 的后面可以接字串，而这些字串会在新的一行出现(目前的下一行)～
- i ：插入， i 的后面可以接字串，而这些字串会在**新的一行出现**(目前的上一行)；
- d ：删除，因为是删除啊，所以 d 后面通常不接任何咚咚；
- c ：取代， c 的后面可以接字串，这些字串可以取代 n1,n2 之间的行！
- s ：取代，可以直接进行取代的工作哩！通常这个 s 的动作可以搭配正规表示法！例如 1,20s/old/new/g 就是啦！
- p ：打印，亦即将某个选择的数据印出。通常 p 会与参数 sed -n 一起运行～

**实例**

### 以行为单位的新增和删除

```shell
#  -a/-i -d
sed -e '4a newLine' testfile 
# 在testfile文件的第四行后添加一行，并将结果输出到标准输出
nl /etc/passwd | sed '3,$d' 
# 要删除第 3 到最后一行
sed '2a Drink tea or ......\
> drink beer ?'
# 在第二行后面加入两行字
注意： 每一行之间都必须要以反斜杠『 \ 』来进行新行的添加喔！
```

### 以行为单位的替换与显示

 ```shell
 #  -c -p
 
 nl /etc/passwd | sed '2,5c No 2-5 number'
 # 将第2-5行的内容取代成为『No 2-5 number』
 nl /etc/passwd | sed -n '2,10p'
 # 仅将2-10行显示出来
 ```

### 数据的搜寻并显示和删除

```shell
nl /etc/passwd | sed -n '/root/p'
# 寻找root相关的字段并显示
nl /etc/passwd | sed -n '/root/d'
# 寻找root相关的字段并删除

```

### 数据的搜寻并执行命令

**搜索/etc/passwd,找到root对应的行，执行后面花括号中的一组命令，每个命令之间用分号分隔，这里把bash替换为blueshell，再输出这行：**

```shell
nl /etc/passwd | sed -n '/root/{s/bash/blueshell/;p;q}'  
```

### 数据的搜寻并替换

```shell
sed 's/要被取代的字串/新的字串/g'
```

```shell
ubuntu@VM-4-7-ubuntu:~/tmp$ ifconfig eth0
eth0: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
        inet 10.0.4.7  netmask 255.255.252.0  broadcast 10.0.7.255
        inet6 fe80::5054:ff:fedb:9583  prefixlen 64  scopeid 0x20<link>
        ether 52:54:00:db:95:83  txqueuelen 1000  (Ethernet)
        RX packets 16289303  bytes 3823275781 (3.8 GB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 15376981  bytes 4165050715 (4.1 GB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0
# 获取eth0的ip地址，去掉其他的字符
ubuntu@VM-4-7-ubuntu:~/tmp$ ifconfig eth0 | grep 'inet '|sed 's/inet//g'|sed 's/netmask.*$//g'
10.0.4.7
```

### 多点编辑

一条sed命令，删除/etc/passwd第三行到末尾的数据，并把bash替换为blueshell

```
nl /etc/passwd | sed -e '3,$d' -e 's/bash/blueshell/'
1  root:x:0:0:root:/root:/bin/blueshell
2  daemon:x:1:1:daemon:/usr/sbin:/bin/sh
```

-e表示多点编辑，第一个编辑命令删除/etc/passwd第三行到末尾的数据，第二条命令搜索bash替换为blueshell

### 直接修改文件内容(危险动作)

```shell
# ubuntu@VM-4-7-ubuntu:~/tmp$ cat tmp1 
apple banana pear.
I don't like you.
see you again.
python is the best language around the world

# ubuntu@VM-4-7-ubuntu:~/tmp$ sed -i 's/\.$/\!/g' tmp1 
# ubuntu@VM-4-7-ubuntu:~/tmp$ cat tmp1 
apple banana pear!
I don't like you!
see you again!
python is the best language around the world!

# ubuntu@VM-4-7-ubuntu:~/tmp$ sed '$a this is the last line' tmp1 
apple banana pear!
I don't like you!
see you again!
python is the best language around the world!
this is the last line
# 注意：$表示最后一行
```













# Shell命令

## 第一个shell程序

```shell
#该文件名为test.sh
#!/bin/bash
echo "hello world!"
```

其中`#!/bin/bash`告诉系统用`/bin/bash`来执行这个文件

* 执行方法1：

  `chmod +x test.sh`

  `./test.sh`//告诉系统在该文件夹下执行文件

* 执行方法2：

  `chmod +x test.sh`

  `/bin/bash test.sh`



## Shell变量

### 变量的使用

#### 定义变量

```shell
runoob
guofuwwei
_var
# 都是合法变量
runoob='tom'
```

#### 使用变量

* 注意：记得在前面加**$**符号

* 可以用`{}`来将变量括起来方便区分

#### 只读变量

* 可以用`readonly`来申明只读变量
* 只读变量无法被删除

#### 删除变量

* 使用`unset`来删除变量，变量删除后就不能使用了

#### 变量类型

运行shell时，会同时存在三种变量：

- **1) 局部变量** 局部变量在脚本或命令中定义，仅在当前shell实例中有效，其他shell启动的程序不能访问局部变量。
- **2) 环境变量** 所有的程序，包括shell启动的程序，都能访问环境变量，有些程序需要环境变量来保证其正常运行。必要的时候shell脚本也可以定义环境变量。
- **3) shell变量** shell变量是由shell程序设置的特殊变量。shell变量中有一部分是环境变量，有一部分是局部变量，这些变量保证了shell的正常运行



### shell字符串

**字符串可以用单引号**，**双引号**

#### 单引号

* 字符串中的变量是无效的
* 单引号字串中不能出现单独一个的单引号（对单引号使用转义符后也不行），但可成对出现，作为字符串拼接使用。

```shell
str='this is a  string'
```

#### 双引号

- 双引号里可以有变量
- 双引号里可以出现转义字符

```shell
#!/bin/bash
name="runoob"
echo -e "your name is \"{$name}\"\n"
```

#### 获取字符串的长度

```shell
string="abcd"
echo ${#string} #输出 4
```

#### 提取子字符串

```shell
string="runoob is a great site"
echo ${string:1:4} # 输出 unoo
```

#### 查找子字符串

```shell
string="runoob is a great site"
echo `expr index "$string" io`  # 输出 4
```



### shell数组

bash支持一维数组，不支持多维数组

#### 定义数组

```shell
数组名=(值1 值2 ... 值n)
eg:
array_name=(value0 value1 value2 value3)
array_name=(
value0
value1
value2
value3
)s
```

#### 读取数组

```shell
${数组名[下标]}
# 用@可以获取数组的全部元素
eg:
echo ${arraynum[@]}
```

#### 获取数组的长度

* 与字符串的相同

  ```shell
  # 取得数组元素的个数
  length=${#array_name[@]}
  # 或者
  length=${#array_name[*]}
  # 取得数组单个元素的长度
  lengthn=${#array_name[n]}
  ```



### shell注释

* 单行直接用#进行注释
* 临时注释可以写一个函数
* 多行注释可以用一下的方式

```shell
:<<EOF
注释内容...
注释内容...
注释内容...
EOF
# 其中EOF可以换为其他字符
```



## Shell传递参数

 **向shell脚本中传递参数的时候用$n来表示，n表示0，1，2，3.....**

```shell
#!/bin/bash
echo "当前脚本的进程ID号为$$"
echo "下面是shell脚本的传参实例"
echo "一共传进来$#参数"
echo "全部参数是$*"
echo "这是第一个参数：$0"
echo "这是第二个参数：$1"
echo "这是第三个参数：$2"
echo "当前脚本退出的状态为$?"

# ubuntu@VM-4-7-ubuntu:~/shell_script$ ./test2.sh 10 20 30
下面是shell脚本的传参实例
这是第一个参数：./test2.sh
这是第二个参数：10
这是第三个参数：20
#  注意：./test2.sh会作为第一个参数传入shell脚本中
```

| 参数处理 | 说明                                                         |
| :------- | :----------------------------------------------------------- |
| $#       | 传递到脚本的参数个数                                         |
| $*       | 以一个单字符串显示所有向脚本传递的参数。 如"$*"用「"」括起来的情况、以"$1 $2 … $n"的形式输出所有参数。 |
| $$       | 脚本运行的当前进程ID号                                       |
| $!       | 后台运行的最后一个进程的ID号                                 |
| $@       | 与$*相同，但是使用时加引号，并在引号中返回每个参数。 如"$@"用「"」括起来的情况、以"$1" "$2" … "$n" 的形式输出所有参数。 |
| $-       | 显示Shell使用的当前选项，与[set命令](https://www.runoob.com/linux/linux-comm-set.html)功能相同。 |
| $?       | 显示最后命令的退出状态。0表示没有错误，其他任何值表明有错误。 |



## Shell 基本运算符

Shell 和其他编程语言一样，支持多种运算符，包括：

- 算数运算符
- 关系运算符
- 布尔运算符
- 字符串运算符
- 文件测试运算符

一般用`expr`进行运算操作

```shell
#!/bin/bash
val=`expr 1 + 1`
echo "1+1=${val}"

# 注意1+1不能连起来写，1和+必须用空格隔开
```

### 算术运算符

下表列出了常用的算术运算符，假定变量 a 为 10，变量 b 为 20：

| 运算符 | 说明                                          | 举例                           |
| :----- | :-------------------------------------------- | :----------------------------- |
| +      | 加法                                          | \`expr $a + $b` 结果为 30。    |
| -      | 减法                                          | \`expr $a - $b` 结果为 -10。   |
| *      | 乘法                                          | \`expr $a \* $b` 结果为  200。 |
| /      | 除法                                          | \`expr $b / $a` 结果为 2。     |
| %      | 取余                                          | \`expr $b % $a` 结果为 0。     |
| =      | 赋值                                          | a=$b 将把变量 b 的值赋给 a。   |
| ==     | 相等。用于比较两个数字，相同则返回 true。     | **[ $a == $b ] 返回 false**。  |
| !=     | 不相等。用于比较两个数字，不相同则返回 true。 | **[ $a != $b ] 返回 true。**   |

**注意：**

* 条件表达式要放在方括号之间，并且要有空格，例如: **[$a==$b]** 是错误的，必须写成 **[ $a == $b ]**。

- 乘号(*)前边必须加反斜杠(\)才能实现乘法运算

```shell
#!/bin/bash
echo "当前脚本的PID为$$"
echo "请输入两个参数a,b，我们将比较计算它们的和和比较它们的大小"
a=$1
b=$2
flag=0
echo "输入的参数为$@"
sum=`expr $a + $b`
c=`expr $a - $b`
echo "a,b的和为$sum"
if [ $a == $b ]
then
        echo "a等于b"
fi
if [ $c -gt $flag ]
then
        echo "a大于b"
fi
if [ $c -lt $flag ]
then
        echo "a小于b"
fi

if [ $? == 0 ]
then 
	status="正常"
fi
if [ $? != 0 ]
then 
	status="异常"
fi
echo "shell脚本运行情况为$status"
```

### 关系运算符

关系运算符只支持数字，不支持字符串，除非字符串的值是数字。

下表列出了常用的关系运算符，假定变量 a 为 10，变量 b 为 20：

| 运算符 | 说明                                                  | 举例                       |
| :----- | :---------------------------------------------------- | :------------------------- |
| -eq    | 检测两个数是否相等，相等返回 true。                   | [ $a -eq $b ] 返回 false。 |
| -ne    | 检测两个数是否不相等，不相等返回 true。               | [ $a -ne $b ] 返回 true。  |
| -gt    | 检测左边的数是否大于右边的，如果是，则返回 true。     | [ $a -gt $b ] 返回 false。 |
| -lt    | 检测左边的数是否小于右边的，如果是，则返回 true。     | [ $a -lt $b ] 返回 true。  |
| -ge    | 检测左边的数是否大于等于右边的，如果是，则返回 true。 | [ $a -ge $b ] 返回 false。 |
| -le    | 检测左边的数是否小于等于右边的，如果是，则返回 true。 | [ $a -le $b ] 返回 true。  |

### 布尔运算符

下表列出了常用的布尔运算符，假定变量 a 为 10，变量 b 为 20：

| 运算符 | 说明                                                | 举例                                     |
| :----- | :-------------------------------------------------- | :--------------------------------------- |
| !      | 非运算，表达式为 true 则返回 false，否则返回 true。 | [ ! false ] 返回 true。                  |
| -o     | 或运算，有一个表达式为 true 则返回 true。           | [ $a -lt 20 -o $b -gt 100 ] 返回 true。  |
| -a     | 与运算，两个表达式都为 true 才返回 true。           | [ $a -lt 20 -a $b -gt 100 ] 返回 false。 |

### 逻辑运算符

以下介绍 Shell 的逻辑运算符，假定变量 a 为 10，变量 b 为 20:

| 运算符 | 说明       | 举例                                       |
| :----- | :--------- | :----------------------------------------- |
| &&     | 逻辑的 AND | [[ $a -lt 100 && $b -gt 100 ]] 返回 false  |
| \|\|   | 逻辑的 OR  | [[ $a -lt 100 \|\| $b -gt 100 ]] 返回 true |

```shell
#!/bin/bash
echo "The shell pid is $$"
echo "please input number"
a=10
b=$1
if [ $a -gt $b ]
then 
	echo "a>b"
fi
if [ $a -lt $b ]
then 
	echo "a<b"
fi
if [ $a == $b ]
then
	echo "a=b"
fi
echo "the shell status is $?"
```

### 字符串运算符

下表列出了常用的字符串运算符，假定变量 a 为 "abc"，变量 b 为 "efg"：

| 运算符 | 说明                                         | 举例                     |
| :----- | :------------------------------------------- | :----------------------- |
| =      | 检测两个字符串是否相等，相等返回 true。      | [ $a = $b ] 返回 false。 |
| !=     | 检测两个字符串是否不相等，不相等返回 true。  | [ $a != $b ] 返回 true。 |
| -z     | 检测字符串长度是否为0，为0返回 true。        | [ -z $a ] 返回 false。   |
| -n     | 检测字符串长度是否不为 0，不为 0 返回 true。 | [ -n "$a" ] 返回 true。  |
| $      | 检测字符串是否为空，不为空返回 true。        | [ $a ] 返回 true。       |

```shell
#!/bin/bash
a="abc"
b="efg"

if [ $a = $b ]
then
   echo "$a = $b : a 等于 b"
else
   echo "$a = $b: a 不等于 b"
fi
```

### 文件测试运算符

文件测试运算符用于检测 Unix 文件的各种属性。

属性检测描述如下：

| 操作符      | 说明                                                         | 举例                      |
| :---------- | :----------------------------------------------------------- | :------------------------ |
| -b file     | 检测文件是否是块设备文件，如果是，则返回 true。              | [ -b $file ] 返回 false。 |
| -c file     | 检测文件是否是字符设备文件，如果是，则返回 true。            | [ -c $file ] 返回 false。 |
| **-d file** | 检测文件是否是目录，如果是，则返回 true。                    | [ -d $file ] 返回 false。 |
| **-f file** | 检测文件是否是普通文件（既不是目录，也不是设备文件），如果是，则返回 true。 | [ -f $file ] 返回 true。  |
| -g file     | 检测文件是否设置了 SGID 位，如果是，则返回 true。            | [ -g $file ] 返回 false。 |
| -k file     | 检测文件是否设置了粘着位(Sticky Bit)，如果是，则返回 true。  | [ -k $file ] 返回 false。 |
| -p file     | 检测文件是否是有名管道，如果是，则返回 true。                | [ -p $file ] 返回 false。 |
| -u file     | 检测文件是否设置了 SUID 位，如果是，则返回 true。            | [ -u $file ] 返回 false。 |
| **-r file** | 检测文件是否可读，如果是，则返回 true。                      | [ -r $file ] 返回 true。  |
| **-w file** | 检测文件是否可写，如果是，则返回 true。                      | [ -w $file ] 返回 true。  |
| **-x file** | 检测文件是否可执行，如果是，则返回 true。                    | [ -x $file ] 返回 true。  |
| **-s file** | 检测文件是否为空（文件大小是否大于0），不为空返回 true。     | [ -s $file ] 返回 true。  |
| **-e file** | 检测文件（包括目录）是否存在，如果是，则返回 true。          | [ -e $file ] 返回 true。  |

其他检查符：

- **-S**: 判断某文件是否 socket。
- **-L**: 检测文件是否存在并且是一个符号链接。

```shell
#!/bin/bash
file="/home/ubuntu/shell_script/test.sh"
if [ -e $file ]
then 
	echo "$file文件存在"
fi
if [ -s $file ]
then 
	echo "文件不为空"
fi
if [ -r $file ]
then
	echo "文件可读"
fi
if [ -w $file ]
then 
	echo "文件可写"
fi
```



## Shell echo命令

* echo会自动进行换行

* 显示普通字符串：双引号可以省略
* 可以使用转义字符`\`
  * `echo \"this is a test\"`
* 使用-e开启转义：比如转义\c:不换行
* 用>将结果重定向到文件
* **使用单引号原样输出字符串，不进行转义和取变量**
* 显示执行命令
  * echo \`date`

## Shell printf 命令

语法：

```
printf  format-string  [arguments...]
```

**参数说明：**

- **format-string:** 为格式控制字符串
- **arguments:** 为参数列表

```shell
#!/bin/bash
printf "%-15s %-15s %-10s\n" 班级 学号 姓名
printf "%-15s %-15s %-10s\n" 电信2001班 U202013894 郭富伟
printf "%-15s %-15s %-10s\n" 高三22班 U40521647 gfw
```

**%s %c %d %f** 都是格式替代符，**％s** 输出一个字符串，**％d** 整型输出，**％c** 输出一个字符，**％f** 输出实数，以小数形式输出。

**%-10s** 指一个宽度为 10 个字符（**-** 表示左对齐，没有则表示右对齐），任何字符都会被显示在 10 个字符宽的字符内，如果不足则自动以空格填充，超过也会将内容全部显示出来。

**%-4.2f** 指格式化为小数，其中 **.2** 指保留2位小数。

### printf 的转义序列

| 序列   | 说明                                                         |
| :----- | :----------------------------------------------------------- |
| \a     | 警告字符，通常为ASCII的BEL字符                               |
| \b     | 后退                                                         |
| \c     | 抑制（不显示）输出结果中任何结尾的换行字符（只在%b格式指示符控制下的参数字符串中有效），而且，任何留在参数里的字符、任何接下来的参数以及任何留在格式字符串中的字符，都被忽略 |
| \f     | 换页（formfeed）                                             |
| **\n** | 换行                                                         |
| **\r** | 回车（Carriage return）                                      |
| **\t** | 水平制表符                                                   |
| \v     | 垂直制表符                                                   |
| \\     | 一个字面上的反斜杠字符                                       |
| \ddd   | 表示1到3位数八进制值的字符。仅在格式字符串中有效             |
| \0ddd  | 表示1到3位的八进制值字符                                     |





## Shell test 命令

### 数值测试

| 参数 | 说明           |
| :--- | :------------- |
| -eq  | 等于则为真     |
| -ne  | 不等于则为真   |
| -gt  | 大于则为真     |
| -ge  | 大于等于则为真 |
| -lt  | 小于则为真     |
| -le  | 小于等于则为真 |

```shell
#!/bin/bash
num1=100
num2=100
if test $[num1] -eq $[num2]
then
    echo '两个数相等！'
else
    echo '两个数不相等！'
fi

# 注意：[]可以进行最基本的算数计算
eg:
#!bin/bash
a=10
b=20
ret=$[a+b]
echo "a+b=$ret"
```



### 字符串测试

| 参数      | 说明                     |
| :-------- | :----------------------- |
| =         | 等于则为真               |
| !=        | 不相等则为真             |
| -z 字符串 | 字符串的长度为零则为真   |
| -n 字符串 | 字符串的长度不为零则为真 |

```shell
#!/bin/bash
str1="runoob1"
str2="runoob2"
if test $str1 = $str2
then
	echo "str1==str2"
else
	echo "str1!=str2"
fi
```



### 文件测试

| 参数      | 说明                                 |
| :-------- | :----------------------------------- |
| -e 文件名 | 如果文件存在则为真                   |
| -r 文件名 | 如果文件存在且可读则为真             |
| -w 文件名 | 如果文件存在且可写则为真             |
| -x 文件名 | 如果文件存在且可执行则为真           |
| -s 文件名 | 如果文件存在且至少有一个字符则为真   |
| -d 文件名 | 如果文件存在且为目录则为真           |
| -f 文件名 | 如果文件存在且为普通文件则为真       |
| -c 文件名 | 如果文件存在且为字符型特殊文件则为真 |
| -b 文件名 | 如果文件存在且为块特殊文件则为真     |

```shell
#!/bin/bash
cd /bin
if test -e ./bash
then
    echo '文件已存在!'
else
    echo '文件不存在!'
fi
```

* 注意：Shell 还提供了与( -a )、或( -o )、非( ! )三个逻辑操作符用于将测试条件连接起来，其优先级为： **!** 最高， **-a** 次之， **-o** 最低。



## Shell 流程控制

### if else

```shell
if condition
then 
	command1
	command2
	command3
elif condition2
then 
	command3
else
	command4
fi
```

```shell
#!/bin/bash
echo 请输入两个参数
a=$1
b=$2
if [ $a == $b ]
then
	echo a==b
elif [ $a < $b ]
then 
	echo a<b
else
	echo a>b
fi
```

### for 循环

```shell
for var in item1 item2 ... itemN
do
    command1
    command2
    ...
    commandN
done
```

```shell
#!/bin/bash
for str in This is a string
do
	echo $str
done
# 会循环输出字符串的每一个字符
```

### while 语句

```shell
while condition
do 
	command
done
```

```shell
#!/bin/bash
int=1
while(( $int<=5 ))
do
    echo $int
    let "int++"
done
```

```shell
echo '按下 <CTRL-D> 退出'
echo -n '输入你最喜欢的网站名: '
while read FILM
do
    echo "是的！$FILM 是一个好网站"
done
```



### until 循环

```shell
until condition
do
    command
done
```

```shell
#!/bin/bash
echo "当前脚本PID为$$"
val=10
until [ $val -lt 1 ]
do 
	echo "The value is:$val"
	val=`expr $val - 1`
done
```

### case ... esac

```shell
case 值 in
模式1)
    command1
    command2
    ...
    commandN
    ;;
模式2）
    command1
    command2
    ...
    commandN
    ;;
esac
# ;;表示break
# esac 表示结束
```

```shell
    #!/bin/bash
    echo "请输入一个整数（1-4）"
    read num
    case $num in 
    1) echo "你选择了1"
    ;;
    2) echo "你选择了2"
    ;;
    3) echo "你选择了3"
    ;;
    4) echo "你选择了4"
    ;;
    esac
```

### 跳出循环

break和continue用于跳出循环

```shell
#!/bin/bash
while :
do
	echo -n "请输入一个1-5之内的数:"
	read num
	case $num in
	1|2|3|4|5) echo "你输入的数字为$num"
	;;
	*) echo "你输入的数字不是1-5之内的，游戏结束"
		break
	;;
	esac
done
```



## Shell 函数

shell脚本的格式为

```shell
[ function ] funname [()]

{

    action;

    [return int;]

}
```

说明：

- 1、可以带function fun() 定义，也可以直接fun() 定义,不带任何参数。
- 2、参数返回，可以显示加：return 返回，如果不加，将以最后一条命令运行结果，作为返回值。 return后跟数值n(0-255



```shell
#!/bin/bash
function fun(){
	echo "这是我的第一个脚本函数"
	echo -n "请输入一个数:"
	read a
	return $a
}
fun
echo "您刚才输入的是$?"

# 可以用$?读取函数的返回值
```



### 函数参数

```shell
#!/bin/bash
myprint(){
	echo "当前进程ID为$$"
	echo "第一个参数为$1"
	echo "第二个参数为$2"
	echo "第三个参数为$3"
	echo "第四个参数为$4"	
}
myprint 10 20 30 40 
```

* 注意，$10 不能获取第十个参数，获取第十个参数需要${10}。当n>=10时，需要使用${n}来获取参数

| 参数处理 | 说明                                                         |
| :------- | :----------------------------------------------------------- |
| $#       | 传递到脚本或函数的参数个数                                   |
| $*       | 以一个单字符串显示所有向脚本传递的参数                       |
| $$       | 脚本运行的当前进程ID号                                       |
| $!       | 后台运行的最后一个进程的ID号                                 |
| $@       | 与$*相同，但是使用时加引号，并在引号中返回每个参数。         |
| $-       | 显示Shell使用的当前选项，与set命令功能相同。                 |
| $?       | 显示最后命令的退出状态。0表示没有错误，其他任何值表明有错误。 |



## Shell 输入/输出重定向

| 命令            | 说明                                               |
| :-------------- | :------------------------------------------------- |
| command > file  | 将输出重定向到 file。                              |
| command < file  | 将输入重定向到 file。                              |
| command >> file | 将输出以追加的方式重定向到 file。                  |
| n > file        | 将文件描述符为 n 的文件重定向到 file。             |
| n >> file       | 将文件描述符为 n 的文件以追加的方式重定向到 file。 |
| n >& m          | 将输出文件 m 和 n 合并。                           |
| n <& m          | 将输入文件 m 和 n 合并。                           |
| << tag          | 将开始标记 tag 和结束标记 tag 之间的内容作为输入。 |

*需要注意的是文件描述符 0 通常是标准输入（STDIN），1 是标准输出（STDOUT），2 是标准错误输出（STDERR）。*

### 输出重定向

```shell
command1 > file1
eg:
echo "Hello world" > hello.txt
```

### 输入重定向

```shell
command1 < file1
```

### Here Document

基本格式为

```
command << delimiter
    document
delimiter
```

```shell
#!/bin/bash
# author:菜鸟教程
# url:www.runoob.com

cat << EOF
欢迎来到
菜鸟教程
www.runoob.com
EOF

# 结果为：
欢迎来到
菜鸟教程
www.runoob.com
```

### /dev/null 文件

/dev/null 是一个特殊的文件，写入到它的内容都会被丢弃；如果尝试从该文件读取内容，那么什么也读不到。但是 /dev/null 文件非常有用，将命令的输出重定向到它，会起到"禁止输出"的效果

```shell
echo "this msg will be ignored" > /dev/null
```



## Shell 文件包含

语法格式：

```shell
. filename   # 注意点号(.)和文件名中间有一空格

或

source filename
```

```shell
# test1.sh
#!/bin/bash
url:www.google.com

# demo6.sh
#!/bin/bash
source ./test1.sh
echo "谷歌搜索的地址是:$url"
```







# vim/vi命令

vim键盘图

![img](https://www.runoob.com/wp-content/uploads/2015/10/vi-vim-cheat-sheet-sch.gif)

## vi/vim 的使用

基本上 vi/vim 共分为三种模式，分别是**命令模式（Command mode）**，**输入模式（Insert mode）**和**底线命令模式（Last line mode）**。 这三种模式的作用分别是：

### 命令模式：

用户刚刚启动 vi/vim，便进入了命令模式。

此状态下敲击键盘动作会被Vim识别为命令，而非输入字符。比如我们此时按下i，并不会输入一个字符，i被当作了一个命令。

以下是常用的几个命令：

- **i** 切换到输入模式，以输入字符。
- **x** 删除当前光标所在处的字符。
- **:** 切换到底线命令模式，以在最底一行输入命令。

若想要编辑文本：启动Vim，进入了命令模式，按下i，切换到输入模式。

命令模式只有一些最基本的命令，因此仍要依靠底线命令模式输入更多命令。

### 输入模式

在命令模式下按下i就进入了输入模式。

在输入模式中，可以使用以下按键：

- **字符按键以及Shift组合**，输入字符
- **ENTER**，回车键，换行
- **BACK SPACE**，退格键，删除光标前一个字符
- **DEL**，删除键，删除光标后一个字符
- **方向键**，在文本中移动光标
- **HOME**/**END**，移动光标到行首/行尾
- **Page Up**/**Page Down**，上/下翻页
- **Insert**，切换光标为输入/替换模式，光标将变成竖线/下划线
- **ESC**，退出输入模式，切换到命令模式

### 底线命令模式

在命令模式下按下:（英文冒号）就进入了底线命令模式。

底线命令模式可以输入单个或多个字符的命令，可用的命令非常多。

在底线命令模式中，基本的命令有（已经省略了冒号）：

- q 退出程序
- w 保存文件

按ESC键可随时退出底线命令模式。

## vi/vim 按键说明

除了上面简易范例的 i, Esc, :wq 之外，其实 vim 还有非常多的按键可以使用。

### 第一部分：一般模式可用的光标移动、复制粘贴、搜索替换等

| 移动光标的方法                                               |                                                              |
| :----------------------------------------------------------- | ------------------------------------------------------------ |
| h 或 向左箭头键(←)                                           | 光标向左移动一个字符                                         |
| j 或 向下箭头键(↓)                                           | 光标向下移动一个字符                                         |
| k 或 向上箭头键(↑)                                           | 光标向上移动一个字符                                         |
| l 或 向右箭头键(→)                                           | 光标向右移动一个字符                                         |
| 如果你将右手放在键盘上的话，你会发现 hjkl 是排列在一起的，因此可以使用这四个按钮来移动光标。 如果想要进行多次移动的话，例如向下移动 30 行，可以使用 "30j" 或 "30↓" 的组合按键， 亦即加上想要进行的次数(数字)后，按下动作即可！ |                                                              |
| [Ctrl] + [f]                                                 | 屏幕『向下』移动一页，相当于 [Page Down]按键 (常用)          |
| [Ctrl] + [b]                                                 | 屏幕『向上』移动一页，相当于 [Page Up] 按键 (常用)           |
| [Ctrl] + [d]                                                 | 屏幕『向下』移动半页                                         |
| [Ctrl] + [u]                                                 | 屏幕『向上』移动半页                                         |
| +                                                            | 光标移动到非空格符的下一行                                   |
| -                                                            | 光标移动到非空格符的上一行                                   |
| n<space>                                                     | 那个 n 表示『数字』，例如 20 。按下数字后再按空格键，光标会向右移动这一行的 n 个字符。例如 20<space> 则光标会向后面移动 20 个字符距离。 |
| **0 或功能键[Home]**                                         | 这是数字『 0 』：移动到这一行的最前面字符处 (常用)           |
| **$ 或功能键[End]**                                          | 移动到这一行的最后面字符处(常用)                             |
| H                                                            | 光标移动到这个屏幕的最上方那一行的第一个字符                 |
| M                                                            | 光标移动到这个屏幕的中央那一行的第一个字符                   |
| L                                                            | 光标移动到这个屏幕的最下方那一行的第一个字符                 |
| **G**                                                        | 移动到这个档案的最后一行(常用)                               |
| nG                                                           | n 为数字。移动到这个档案的第 n 行。例如 20G 则会移动到这个档案的第 20 行(可配合 :set nu) |
| **gg**                                                       | 移动到这个档案的第一行，相当于 1G 啊！ (常用)                |
| **n<Enter>**                                                 | n 为数字。光标向下移动 n 行(常用)                            |
| 搜索替换                                                     |                                                              |
| /word                                                        | 向光标之下寻找一个名称为 word 的字符串。例如要在档案内搜寻 vbird 这个字符串，就输入 /vbird 即可！ (常用) |
| ?word                                                        | 向光标之上寻找一个字符串名称为 word 的字符串。               |
| n                                                            | 这个 n 是英文按键。代表重复前一个搜寻的动作。举例来说， 如果刚刚我们执行 /vbird 去向下搜寻 vbird 这个字符串，则按下 n 后，会向下继续搜寻下一个名称为 vbird 的字符串。如果是执行 ?vbird 的话，那么按下 n 则会向上继续搜寻名称为 vbird 的字符串！ |
| N                                                            | 这个 N 是英文按键。与 n 刚好相反，为『反向』进行前一个搜寻动作。 例如 /vbird 后，按下 N 则表示『向上』搜寻 vbird 。 |
| 使用 /word 配合 n 及 N 是非常有帮助的！可以让你重复的找到一些你搜寻的关键词！ |                                                              |
| :n1,n2s/word1/word2/g                                        | n1 与 n2 为数字。在第 n1 与 n2 行之间寻找 word1 这个字符串，并将该字符串取代为 word2 ！举例来说，在 100 到 200 行之间搜寻 vbird 并取代为 VBIRD 则： 『:100,200s/vbird/VBIRD/g』。(常用) |
| **:1,$s/word1/word2/g** 或 **:%s/word1/word2/g**             | 从第一行到最后一行寻找 word1 字符串，并将该字符串取代为 word2 ！(常用) |
| **:1,$s/word1/word2/gc** 或 **:%s/word1/word2/gc**           | 从第一行到最后一行寻找 word1 字符串，并将该字符串取代为 word2 ！且在取代前显示提示字符给用户确认 (confirm) 是否需要取代！(常用) |
| 删除、复制与贴上                                             |                                                              |
| x, X                                                         | 在一行字当中，x 为向后删除一个字符 (相当于 [del] 按键)， X 为向前删除一个字符(相当于 [backspace] 亦即是退格键) (常用) |
| nx                                                           | n 为数字，连续向后删除 n 个字符。举例来说，我要连续删除 10 个字符， 『10x』。 |
| **dd**                                                       | 删除游标所在的那一整行(常用)                                 |
| ndd                                                          | n 为数字。删除光标所在的向下 n 行，例如 20dd 则是删除 20 行 (常用) |
| **d1G**                                                      | 删除光标所在到第一行的所有数据                               |
| **dG**                                                       | 删除光标所在到最后一行的所有数据                             |
| d$                                                           | 删除游标所在处，到该行的最后一个字符                         |
| d0                                                           | 那个是数字的 0 ，删除游标所在处，到该行的最前面一个字符      |
| **yy**                                                       | 复制游标所在的那一行(常用)                                   |
| nyy                                                          | n 为数字。复制光标所在的向下 n 行，例如 20yy 则是复制 20 行(常用) |
| y1G                                                          | 复制游标所在行到第一行的所有数据                             |
| yG                                                           | 复制游标所在行到最后一行的所有数据                           |
| y0                                                           | 复制光标所在的那个字符到该行行首的所有数据                   |
| y$                                                           | 复制光标所在的那个字符到该行行尾的所有数据                   |
| **p, P**                                                     | p 为将已复制的数据在光标下一行贴上，P 则为贴在游标上一行！ 举例来说，我目前光标在第 20 行，且已经复制了 10 行数据。则按下 p 后， 那 10 行数据会贴在原本的 20 行之后，亦即由 21 行开始贴。但如果是按下 P 呢？ 那么原本的第 20 行会被推到变成 30 行。 (常用) |
| J                                                            | 将光标所在行与下一行的数据结合成同一行                       |
| c                                                            | 重复删除多个数据，例如向下删除 10 行，[ 10cj ]               |
| **u**                                                        | 复原前一个动作。(常用)                                       |
| **[Ctrl]+r**                                                 | 重做上一个动作。(常用)                                       |
| 这个 u 与 [Ctrl]+r 是很常用的指令！一个是复原，另一个则是重做一次～ 利用这两个功能按键，你的编辑，嘿嘿！很快乐的啦！ |                                                              |
| .                                                            | 不要怀疑！这就是小数点！意思是重复前一个动作的意思。 如果你想要重复删除、重复贴上等等动作，按下小数点『.』就好了！ (常用) |

### 第二部分：一般模式切换到编辑模式的可用的按钮说明

| 进入输入或取代的编辑模式                                     |                                                              |
| :----------------------------------------------------------- | ------------------------------------------------------------ |
| **i, I**                                                     | 进入输入模式(Insert mode)： i 为『从目前光标所在处输入』， I 为『在目前所在行的第一个非空格符处开始输入』。 (常用) |
| **a, A**                                                     | 进入输入模式(Insert mode)： a 为『从目前光标所在的下一个字符处开始输入』， A 为『从光标所在行的最后一个字符处开始输入』。(常用) |
| **o, O**                                                     | 进入输入模式(Insert mode)： 这是英文字母 o 的大小写。o 为在目前光标所在的下一行处输入新的一行； O 为在目前光标所在的上一行处输入新的一行！(常用) |
| r, R                                                         | 进入取代模式(Replace mode)： r 只会取代光标所在的那一个字符一次；R会一直取代光标所在的文字，直到按下 ESC 为止；(常用) |
| 上面这些按键中，在 vi 画面的左下角处会出现『--INSERT--』或『--REPLACE--』的字样。 由名称就知道该动作了吧！！特别注意的是，我们上面也提过了，你想要在档案里面输入字符时， 一定要在左下角处看到 INSERT 或 REPLACE 才能输入喔！ |                                                              |
| [Esc]                                                        | 退出编辑模式，回到一般模式中(常用)                           |

### 第三部分：一般模式切换到指令行模式的可用的按钮说明

| 指令行的储存、离开等指令                                     |                                                              |
| :----------------------------------------------------------- | ------------------------------------------------------------ |
| :w                                                           | 将编辑的数据写入硬盘档案中(常用)                             |
| :w!                                                          | 若文件属性为『只读』时，强制写入该档案。不过，到底能不能写入， 还是跟你对该档案的档案权限有关啊！ |
| :q                                                           | 离开 vi (常用)                                               |
| :q!                                                          | 若曾修改过档案，又不想储存，使用 ! 为强制离开不储存档案。    |
| 注意一下啊，那个惊叹号 (!) 在 vi 当中，常常具有『强制』的意思～ |                                                              |
| :wq                                                          | 储存后离开，若为 :wq! 则为强制储存后离开 (常用)              |
| ZZ                                                           | 这是大写的 Z 喔！如果修改过，保存当前文件，然后退出！效果等同于(保存并退出) |
| ZQ                                                           | 不保存，强制退出。效果等同于 **:q!**。                       |
| :w [filename]                                                | 将编辑的数据储存成另一个档案（类似另存新档）                 |
| :r [filename]                                                | 在编辑的数据中，读入另一个档案的数据。亦即将 『filename』 这个档案内容加到游标所在行后面 |
| :n1,n2 w [filename]                                          | 将 n1 到 n2 的内容储存成 filename 这个档案。                 |
| :! command                                                   | 暂时离开 vi 到指令行模式下执行 command 的显示结果！例如 『:! ls /home』即可在 vi 当中察看 /home 底下以 ls 输出的档案信息！ |
| vim 环境的变更                                               |                                                              |
| :set nu                                                      | 显示行号，设定之后，会在每一行的前缀显示该行的行号           |
| :set nonu                                                    | 与 set nu 相反，为取消行号！                                 |

特别注意，在 vi/vim 中，数字是很有意义的！数字通常代表重复做几次的意思！ 也有可能是代表去到第几个什么什么的意思。

举例来说，要删除 50 行，则是用 『50dd』 对吧！ 数字加在动作之前，如我要向下移动 20 行呢？那就是『20j』或者是『20↓』即可。



### vim 中批量添加注释

**方法一 ：块选择模式**

批量注释：

**Ctrl + v** 进入块选择模式，然后移动光标选中你要注释的行，再按大写的 **I** 进入行首插入模式输入注释符号如 **//** 或 **#**，输入完毕之后，按两下 **ESC**，**Vim** 会自动将你选中的所有行首都加上注释，保存退出完成注释。

取消注释：

**Ctrl + v** 进入块选择模式，选中你要删除的行首的注释符号，注意 **//** 要选中两个，选好之后按 **d** 即可删除注释，**ESC** 保存退出。

**方法二: 替换命令**

批量注释。

使用下面命令在指定的行首添加注释。

使用名命令格式： **:起始行号,结束行号s/^/注释符/g**（注意冒号）。

取消注释：

使用名命令格式： **:起始行号,结束行号s/^注释符//g**（注意冒号）。

例子：

1、在 10 - 20 行添加 **//** 注释

```
:10,20s#^#//#g
```

2、在 10 - 20 行删除 **//** 注释

```
:10,20s#^//##g
```

3、在 10 - 20 行添加 **#** 注释

```
:10,20s/^/#/g
```

4、在 **10 - 20** 行删除 # 注释

```
:10,20s/#//g
```

### vim快捷键补充（插入模式）

1. 自动补全：**ctrl** + **n**
2. **\p<** 插入一个include，并把光标置于<>中间
3. **\im** 插入主函数
4. **\ip** 插入printf，并自动添加**\n**，且把光标置于双引号中间

### vim快捷键补充（编辑模式）

1. dw 删除一个单词（配合b：将光标置于所在单词的首部）
2. yw 复制一个单词（配合p：粘贴）

### vim快捷键补充（插入与编辑模式通用）

1. \rr 运行程序
2. \rc 保存并编译程序（会生成二进制文件）





# Linux apt命令

apt（Advanced Packaging Tool）是一个在 Debian 和 Ubuntu 中的 Shell 前端软件包管理器。

apt 命令提供了查找、安装、升级、删除某一个、一组甚至全部软件包的命令，而且命令简洁而又好记。

apt 命令执行需要超级管理员权限(root)。

## apt 语法

```
  apt [options] [command] [package ...]
```

- **options：**可选，选项包括 -h（帮助），-y（当安装过程提示选择全部为"yes"），-q（不显示安装的过程）等等。
- **command：**要进行的操作。
- **package**：安装的包名。

------

## apt 常用命令

- 列出所有可更新的软件清单命令：**sudo apt update**

- **升级软件包：sudo apt upgrade**

  列出可更新的软件包及版本信息：**apt list --upgradeable**

  升级软件包，升级前先删除需要更新软件包：**sudo apt full-upgrade**

- **安装指定的软件命令：sudo apt install <package_name>**

  安装多个软件包：**sudo apt install <package_1> <package_2> <package_3>**

- **更新指定的软件命令：sudo apt update <package_name>**

- **显示软件包具体信息,例如：版本号，安装大小，依赖关系等等：sudo apt show <package_name>**

- **删除软件包命令：sudo apt remove <package_name>**

- 清理不再使用的依赖和库文件: **sudo apt autoremove**

- 移除软件包及配置文件: **sudo apt purge <package_name>**

- 查找软件包命令： **sudo apt search <keyword>**

- **列出所有已安装的包：apt list --installed**

- 列出所有已安装的包的版本信息：**apt list --all-versions**





​	
