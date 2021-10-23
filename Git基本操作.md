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

*`git log --graph --pretty=oneline --abbrev-commit`可以更好的查看log日志*

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

### 常见命令

* `git remote -v`查看远程分支的详细情况
* `git pull`从远程仓库拉取
* `git branch --set-upstream branch-name origin/branch-name`设置上传流

* `git checkout -b dev origin/dev`创建并切换分支，并关联本地和远程的上传流



# Git分支管理

## 基本命令

* 创建分支	`git branch dev`
* 删除分支    `git branch -d`
* 切换分支   `git checkout dev`或者`git switch dev`
* 创建并切换分支  `git checkout -b dev`或者`git switch -c dev`
* 查看分支  `git branch`
* 合并分支  `git merge `

```shell
例如:
git checkout master
git merge dev
```

## 冲突的解决

*当同时在两个分支上修改了文件并commit后，就无法执行快速合并，会产生冲突，我们必须解决冲突才能继续合并*

* 使用`git log --graph`查看合并的日志

```shell
# ubuntu@VM-4-7-ubuntu:~/tmp/learngit$ git status
On branch master
Your branch is ahead of 'origin/master' by 2 commits.
  (use "git push" to publish your local commits)

You have unmerged paths.
  (fix conflicts and run "git commit")
  (use "git merge --abort" to abort the merge)

Unmerged paths:
  (use "git add <file>..." to mark resolution)
	both modified:   readme.txt
# 接下来我们可以用vim打开文件，手动修改文件，解决冲突，然后add,commit就行了
# ubuntu@VM-4-7-ubuntu:~/tmp/learngit$ git log --graph --pretty=oneline --abbrev-commit
*   7090366 (HEAD -> master) confilct fixed
|\  
| * 2af215c (feature1) add quick
* | db260e3 master add
|/  
* 5e01d5b create a new branch is quick;
* f8bcaf1 (origin/master) add test.txt
* 2a7ae56 git tracks changes
* 81bdc93 create License,add a line
* 6283feb add GPL
* 3f19007 add distributed
* 8d767a0 add readme.txt
```

## 分支管理策略

### 禁用Fast forward模式

```shell
# merge时使用--no-ff 参数，注意加上 -m 参数说明提交信息

# ubuntu@VM-4-7-ubuntu:~/tmp/learngit$ git merge --no-ff dev -m 'no foucs is good'
Merge made by the 'recursive' strategy.
 readme.txt | 1 +
 1 file changed, 1 insertion(+)
```

### bug分支：暂存工作状态和复制提交

常见命令：

* `git stash`暂存工作区
* `git stash list`参看暂存工作区的情况
* `git stash pop`取消暂存并删除暂存工作区
* `git stash apply`取消暂存
* `git stash drop`删除暂存工作区
* `git cherry-pick <commit_id>`复制提交到当前分支

```shell
# ubuntu@VM-4-7-ubuntu:~/tmp/learngit$ git stash
Saved working directory and index state WIP on dev: a18c371 fix confilct is so easy

# ubuntu@VM-4-7-ubuntu:~/tmp/learngit$ git stash pop
On branch dev
Changes to be committed:
  (use "git restore --staged <file>..." to unstage)
	new file:   hello.py

Dropped refs/stash@{0} (2466f2f3a8d466849c71d8d8c3d4d63a03c24e8f)
```

*在master分支上修复的bug，想要合并到当前dev分支，可以用`git cherry-pick <commit>`命令，把bug提交的修改“复制”到当前分支，避免重复劳动。*

```shell
# ubuntu@VM-4-7-ubuntu:~/tmp/learngit$ cat readme.txt 
Git is a distributed version control system
Git is a free software distributed under the GPL
Git has a mutable index called stage
Git tracks changes
create a new branch is quick and simple
fix a confilct is so easy
# 此时dev分支上的bug没有消除

# ubuntu@VM-4-7-ubuntu:~/tmp/learngit$ git reflog
c89775f (HEAD -> dev) HEAD@{0}: commit: add hello.py
a18c371 HEAD@{1}: checkout: moving from master to dev
5f222f9 (master) HEAD@{2}: merge bug-01: Fast-forward
8ec373c HEAD@{3}: checkout: moving from bug-01 to master
5f222f9 (master) HEAD@{4}: commit: fix bug01
# 5f222f9 就是修复bug的提交
# ubuntu@VM-4-7-ubuntu:~/tmp/learngit$ git cherry-pick 5f222f9
[dev 632cedd] fix bug01
 Date: Sat Oct 23 16:05:43 2021 +0800
 1 file changed, 1 insertion(+), 1 deletion(-)
 # 完成bug修复
```

### feature分支：git branch -D name 强制删除

如果要丢弃一个没有被合并过的分支，可以通过`git branch -D <name>`强行删除



### 多人协作时要注意的事项

一般的流程

1. `git clone url`克隆一个仓库下来，注意：这时只有一个主分支master
2. `git checkout -b dev origin/dev`创建dev分支，并将本地和远程的dev分支关联起来

```shell
# ubuntu@VM-4-7-ubuntu:~/tmp/learngit2/learngit$ git checkout -b dev origin/dev
Branch 'dev' set up to track remote branch 'dev' from 'origin'.
Switched to a new branch 'dev'
```

3. 这时就可以愉快的向dev分支提交了

同时修改同一个文件的问题

```shell
ubuntu@VM-4-7-ubuntu:~/tmp/learngit$ git push origin dev
To github.com:guofuwei/learngit.git
 ! [rejected]        dev -> dev (fetch first)
error: failed to push some refs to 'git@github.com:guofuwei/learngit.git'
hint: Updates were rejected because the remote contains work that you do
hint: not have locally. This is usually caused by another repository pushing
hint: to the same ref. You may want to first integrate the remote changes
hint: (e.g., 'git pull ...') before pushing again.
hint: See the 'Note about fast-forwards' in 'git push --help' for details.

ubuntu@VM-4-7-ubuntu:~/tmp/learngit$ git pull
remote: Enumerating objects: 8, done.
remote: Counting objects: 100% (8/8), done.
remote: Compressing objects: 100% (2/2), done.
remote: Total 6 (delta 1), reused 3 (delta 1), pack-reused 0
Unpacking objects: 100% (6/6), 878 bytes | 878.00 KiB/s, done.
From github.com:guofuwei/learngit
   632cedd..2420bb1  dev        -> origin/dev
 * [new branch]      main       -> origin/main
There is no tracking information for the current branch.
Please specify which branch you want to merge with.
See git-pull(1) for details.

    git pull <remote> <branch>

If you wish to set tracking information for this branch you can do so with:

    git branch --set-upstream-to=origin/<branch> dev
    
# 这时git pull会提醒你设置上传流
# ubuntu@VM-4-7-ubuntu:~/tmp/learngit$ git branch --set-upstream-to=origin/dev dev
Branch 'dev' set up to track remote branch 'dev' from 'origin'.
# ubuntu@VM-4-7-ubuntu:~/tmp/learngit$ git pull
Auto-merging hello.py
CONFLICT (content): Merge conflict in hello.py
Automatic merge failed; fix conflicts and then commit the result.
# 设置上传流之后提醒冲突，手动解决冲突然后提交
# 解决冲突之后再add,commit，push完成
```

### 使用git rebase进行变基，完成git log的整理化

- rebase操作可以把本地未push的分叉提交历史整理成直线；
- rebase的目的是使得我们在查看历史提交的变化时更容易，因为分叉的提交需要三方对比。

```shell
# ubuntu@VM-4-7-ubuntu:~/tmp/learngit$ git log --pretty=oneline --graph --abbrev-commit
* a0b478b (HEAD -> master) add author email
*   9697f63 (origin/master) Merge branch 'dev'
|\  
| *   ccefa6d (origin/dev, dev) fix the conflict
| |\  
* a0b478b (HEAD -> master) add author email
*   9697f63 (origin/master) Merge branch 'dev'
|\  
| *   ccefa6d (origin/dev, dev) fix the conflict
| |\  
* a0b478b (HEAD -> master) add author email
*   9697f63 (origin/master) Merge branch 'dev'
|\  
| *   ccefa6d (origin/dev, dev) fix the conflict
| |\  
| | * 2420bb1 2 add a line in hello.py
| * | 1b5236c add a line in hello.py
| |/  
| * 632cedd fix bug01
| * c89775f add hello.py
* | 5f222f9 fix bug01
* | 8ec373c no foucs is good
............

```

* 我有点不太懂。。。



# Git标签管理

## 常见操作

* `git tag <name>`/`git tag <name> commit_id`这样就可以打下一个新标签
  * 可以添加-a指定标签名，-m指定说明文字
  * 示例:`git tag -a v0.9 -m 'version v0.9 released' 632cedd`
* `git tag`查看所有的标签
* `git show <tagname>`查看标签信息
* `git tag -d <tagname>`删除标签
* `git push origin <tagname>`推送标签到远程
  * `git push origin --tags`一次性推送所有的标签

* `git push origin :refs/tags/<tagname>`可以删除一个远程标签(之前记得先删除本地标签)

# Others

## 忽略特殊文件

* 新建**.gitignore**文件

* 我们可以使用github上现成的.gitignore文件,[地址](https://github.com/github/gitignore)

.gitgnore的简单规则

* 直接写入文件名，直接忽略该文件
* 可以用`*.env`等形式进行匹配
* 把指定文件排除在`.gitignore`规则外的写法就是`!`+文件名
* 可以用`git check-ignore`命令检查.gitgnore文件



### 配置别名

目的：简化命令

* `git config --global --edit`查看总体配置文件/可以直接`vim .git/config`看配置文件

```shell
# 配置status 为 st
git config --global alias.st status
```

