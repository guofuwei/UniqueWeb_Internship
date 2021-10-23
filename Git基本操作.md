# Git基本操作

## git初始化

* 方式一:`git init`
* 方式二:`git clone <url>`

## git三部曲

1. `git add`
2. `git commit`
3. `git push`

##  git status用来查看文件状态

##  git diff参看文件和仓库里的有什么不同

1. git diff：是查看working tree与index的差别的。
2. git diff --cached：是查看index与repository的差别的。
3. git diff HEAD：是查看working tree和repository的差别的。其中：HEAD代表的是最近的一次commit的信息。

```shell
# ubuntu@VM-4-7-ubuntu:~/tmp/learngit$ git diff readme.txt
diff --git a/readme.txt b/readme.txt
index f362e74..f137c15 100644
--- a/readme.txt
+++ b/readme.txt
@@ -1,2 +1,2 @@
-Git is a version control system
+Git is a distributed version control system
 Git is a free software
```

##  git log可以查看提交历史

  *可以用`pretty=oneline`简化输出*

```shell
# ubuntu@VM-4-7-ubuntu:~/tmp/learngit$ git log
commit 6283feb8d6accd6bc8254c28231ecf2992743b3e (HEAD -> master)
Author: guofuwei <2625406970@qq.com>
Date:   Sat Oct 23 12:16:23 2021 +0800

    add GPL

commit 3f190079ee8c6d2cc8e3d4cb4b0bb90ea6afd7b8
Author: ubuntu <ubuntu@localhost.localdomain>
Date:   Sat Oct 23 12:11:09 2021 +0800

    add distributed

commit 8d767a0790bae38c72c15b5a0c953c619517eab5
Author: ubuntu <ubuntu@localhost.localdomain>
Date:   Sat Oct 23 12:03:32 2021 +0800

    add readme.txt
# ubuntu@VM-4-7-ubuntu:~/tmp/learngit$ git log --pretty=oneline
6283feb8d6accd6bc8254c28231ecf2992743b3e (HEAD -> master) add GPL
3f190079ee8c6d2cc8e3d4cb4b0bb90ea6afd7b8 add distributed
8d767a0790bae38c72c15b5a0c953c619517eab5 add readme.txt
```

##  git reset倒退版本

```shell
git reset --hard HEAD^ # 回退到上个版本 # HEAD^^上上个	#HEAD~100前一百个
git reset --hard 6283feb8d6accd6b# 回到指定版本

# ubuntu@VM-4-7-ubuntu:~/tmp/learngit$ git reset --hard HEAD^
HEAD is now at 3f19007 add distributed
# ubuntu@VM-4-7-ubuntu:~/tmp/learngit$ cat readme.txt 
Git is a distributed version control system
Git is a free software
```

##  git reflog可以参看命令历史

*一般用来回到未来*

```shell
# ubuntu@VM-4-7-ubuntu:~/tmp/learngit$ git reflog
6283feb (HEAD -> master) HEAD@{0}: reset: moving to 6283feb8d6accd6b
3f19007 HEAD@{1}: reset: moving to HEAD^
6283feb (HEAD -> master) HEAD@{2}: commit: add GPL
3f19007 HEAD@{3}: commit: add distributed
8d767a0 HEAD@{4}: commit (initial): add readme.txt
```

## *git restore <file>*可以放弃工作区的修改

````shell
# ubuntu@VM-4-7-ubuntu:~/tmp/learngit$ git status
On branch master
Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
	modified:   readme.txt
````

```shell
# ubuntu@VM-4-7-ubuntu:~/tmp/learngit$ git status
On branch master
Changes not staged for commit:
  (use "git add/rm <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
	deleted:    test.txt

no changes added to commit (use "git add" and/or "git commit -a")
# 可以继续用git restore <file>恢复删除的文件
```



## git restore --staged <file> 可以放弃暂存区的修改

* 注意：要想完全丢弃修改要重复git restore <file>再丢弃工作区的修改

```shell
# ubuntu@VM-4-7-ubuntu:~/tmp/learngit$ git add readme.txt 
# ubuntu@VM-4-7-ubuntu:~/tmp/learngit$ git status
On branch master
Changes to be committed:
  (use "git restore --staged <file>..." to unstage)
	modified:   readme.txt
```



## git远程操作

### 初始化

#### 1.生成 ssh公钥和私钥

`ssh-keygen -t rsa -C "you@example.com"`

#### 2.将ssh公钥添加到github中

#### 3.开始关联本地仓库和远程仓库

* 方式一：git clone克隆一个现成的git仓库

* 方式二：git remote add origin <url>设置本地仓库的上升流

  ​				git push -u origin main

* 删除远程仓库`git remote rm origin`





# Git分支管理

## 基本命令

* 创建分支	`git branch dev`
* 切换分支   `git checkout dev`
* 创建并切换分支  `git checkout -b dev`
* 查看分支  `git branch`
* 

