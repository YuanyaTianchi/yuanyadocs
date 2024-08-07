# Git

> [官网](https://git-scm.com/)；

# ---快开---

### Github Desktop

> [GitHub Desktop](https://desktop.github.com/)；

通过 Github Desktop 操作仓库，速度拉满。

### 安装

> [下载](https://git-scm.com/downloads)；

```shell
apt install -y git
```

### 配置 gitconfig

- 仓库配置：`<git-repository>/.git/config`；
- 全局配置：`~/.gitconfig`；

```shell
cat > ~/.gitconfig <<EOF
[alias]
  st = status
  co = commit
  br = branch
  ch = checkout
  lg = log --graph --abbrev-commit --pretty=oneline
  rebh = git rebase -i HEAD~
[core]
  # 关闭 quotepath，使中文不以八进制的字符编码显示，而是显示汉字
  quotepath=false
[http]
  # http 代理
  proxy = http://<proxy_server_ip>:<port>
[https]
  # https 代理
  proxy = https://<proxy_server_ip>:<port>
  # 关闭证书校验
  sslVerify = false
[user]
  name = Tianchi
  email = yuanyatianchi@gmail.com
EOF
# 查看
git config --list
```

### 配置 gitignore

> [.gitignore常用配置](https://gist.github.com/octocat/9257657)

```shell
.idea
```

### 配置 GitHub 可访问/加速

##### hosts

配置 GitHub 所需资源相关域名 hosts 使资源可访问，但不一定能使仓库操作加速。

配置 hosts 适用于任何无法加载资源的域名，通过 chrome/edge 插件 [Proxy SwitchyOmega](https://chromewebstore.google.com/detail/proxy-switchyomega/padekgcemlokbadohgkifijomclgjgif) 查看未加载资源域名（或通过 F12 查看，不推荐），然后访问 [IPADDRESS.COM](https://www.ipaddress.com/) 查询域名地址（或  [IP.cn](https://ip.cn/) ，不推荐）。

以下是常用条目（大部分稳定，如仍无法访问请查询后更新）：

```host
cat >> /etc/hosts << EOF


# search from ipaddress.com
# 用于加载 GitHub 页面样式
185.199.108.154 github.githubassets.com
# 其他视情况可选
#140.82.113.4    github.com
#151.101.1.194   github.global.ssl.fastly.net
#185.199.111.133 raw.githubusercontent.com
#172.253.115.97 googletagmanager.com
#172.253.122.94 recaptcha.net
EOF
```

windows 更改后仍无法访问请打开终端执行 `ipconfig /flushdns` 命令，仍然无法访问请**重启**系统。

# 基操

### rebase pull

避免产生 merge commit 记录

```shell
git pull --rebase <remote_name> <branch_name>
```

### 从 tag 切出分支

```shell

```

### 本地仓

### 远程仓

```shell
# 添加
git remote add <origin_name> <git@x.x:x/x.git or https://x.x/x/x.git>
# 删除

# 修改 url
git remote set-url <origin_name> <git@x.x:x/x.git or https://x.x/x/x.git>
# 查看
git remote
# 查看详细地址
git remote -v
```

- git remote remove `<repositoryName>`：删除远程的仓库的所有跟踪分支和配置设置
- git remote rename `<oldName>` `<newName>`：重命名远程仓库在本地的简称
- git remote show `<repositoryName>`：查看某个远程仓库的详细信息
- 
- git pull <--rebase> `<repositoryName>` `<branchName>` <--allow-unrelated-histories>：获取远程仓库项目文件
  - --allow-unrelated-histories：可选参数，可以合并两个独立启动仓库的历史
- 提交到本地仓库后再推送到远程仓库
- git push <--set-upstream> `<repositoryName>` `<branchName>`：推送到远程仓库
- git remote rm：删除源(origin)

## 概念

- 192.30.253.113 [github.com](http://github.com)：修改host，提高GitHub的push和pull速度
- 命令行符号

  - [ ]：可写可不写

  - <>：必须写且需要用你自己的内容替换
  - { }：必须在其中做出选择(选项之间以 | 隔开)
- 结构

  - 工作区
  - 暂存区
  - 本地库

https://www.yiibai.com/git

## 操作

### 推拉

#### pull

- `git pull [option]... <远程主机名> [branch-name]:[local-branch-name]`
  - -q：

```shell
git pull origin remoteBranch:localBranch #获取远程origin上的分支branch1，并合并到本地的分支branch2
```

#### push

#### clone

clone。一般建议使用ssh，https有可能出一些不知所以的问题

```shell
#ssh
git clone git@github.com:YuanyaTianchi/yuanyatianchi.git
#https
git clone https://github.com/YuanyaTianchi/yuanyatianchi.git
```

### 版本

#### log

- git log

```shell
git log --graph --abbrev-commit --pretty=oneline #图形化，hash值简化，单行
```

#### revert

https://juejin.im/post/6844903614767448072

重提回退版本

```shell
git revert -n xxxxxxx
git commit -m "恢复到xxxxxxx版本"
```

重提回退到指定版本，将此次回退作为新的提交重提，不删除该版本之后的提交。效果如下

```shell
$ git log
xxxxxx3 第3次提交
xxxxxx2 第2次提交
xxxxxx1 第1次提交
$ git revert -n xxxxxx2
$ git commit -m "恢复到第2次提交"
$ git log
xxxxxx4 恢复到第2次提交
xxxxxx3 第3次提交
xxxxxx2 第2次提交
xxxxxx1 第1次提交
```

推荐使用，尤其是在将回退版本的后面的提交已推送到远程仓库时，revert 后也能简单直接提交

#### reset

重设回退版本

```shell
# HEAD 直接指向指定版本号
git reset --hard xxxxxxx
# 前移 HEAD，几个^表示前移几个版本
git reset --hard HEAD^
```

重设回退到指定版本，该版本之后的提交都将被删除。效果如下

```shell
$ git log
xxxxxx3 第3次提交
xxxxxx2 第2次提交
xxxxxx1 第1次提交
$ git reset --hard HEAD^
$ git log
xxxxxx2 第2次提交
xxxxxx1 第1次提交
```

参数

- --soft： 回退后a分支修改的代码被保留并标记为add的状态（git status 是绿色的状态）
- --mixed： 重置索引，但不重置工作树，更改后的文件标记为未提交（add）的状态。默认操作。
- --hard： 重置索引和工作树，并且a分支修改的所有文件和中间的提交，没提交的代码都被丢弃了。
- --merge： 和--hard类似，只不过如果在执行reset命令之前你有改动一些文件并且未提交，merge会保留你的这些修改，hard则不会。【注：如果你的这些修改add过或commit过，merge和hard都将删除你的提交】
- --keep： 和--hard类似，执行reset之前改动文件如果是a分支修改了的，会提示你修改了相同的文件，不能合并。如果不是a分支修改的文件，会移除缓存区。git status还是可以看到保持了这些修改。

视情况使用，尤其不建议在将回退版本的后面的提交已推送到远程仓库时使用，否则本地版本落后于远程版本，只有强推才能覆盖推送，风险较大

### 分支

#### branch

#### merge

我的需求在 `feature/tianchi/xxx`分支上写完，要合并到 `develop`里面，需要先选择到开发分支再去merge我的分支，会是 `Merge branch 'feature/tianchi/xxx' into 'develop'`，千万**不要**在我的分支上去merge开发分支，会变成 `Merge branch 'develop' into 'feature/tianchi/xxx'`

```shell
#正确操作如下
git pull origin develop:develop
git checkout develop
git merge feature/tianchi/xxx
```

#### rebase

重定，变基(?)，指将分支移到一个新的基提交的过程

- rebase
  - -r：--rebase-merges简写。

```shell
git rebase -r
```

## 规范

https://www.jianshu.com/p/201bd81e7dc9?utm_source=oschina-app

### Commit message 的作用

提供更多的历史信息，方便快速浏览。

比如，下面的命令显示上次发布后的变动，每个commit占据一行。你只看行首，就知道某次 commit 的目的。

```bash
$ git log <last tag> HEAD --pretty=format:%s
```

可以过滤某些commit（比如文档改动），便于快速查找信息

```bash
$ git log <last release> HEAD --grep feature
```

可以直接从commit生成Change log。Change Log 是发布新版本时，用来说明与上一个版本差异的文档，详见后文。

其他优点

- 可读性好，清晰，不必深入看代码即可了解当前commit的作用。
- 为 Code Reviewing做准备
- 方便跟踪工程历史
- 让其他的开发者在运行 git blame 的时候想跪谢
- 提高项目的整体质量，提高个人工程素质

### Commit message 的格式

每次提交，Commit message 都包括三个部分：header，body 和 footer。

```xml
<type>(<scope>): <subject>
<BLANK LINE>
<body>
<BLANK LINE>
<footer>
```

其中，header 是必需的，body 和 footer 可以省略。
不管是哪一个部分，任何一行都不得超过72个字符（或100个字符）。这是为了避免自动换行影响美观。

#### Header

Header部分只有一行，包括三个字段：`type`（必需）、`scope`（可选）和 `subject`（必需）。

##### type

用于说明 commit 的类别，只允许使用下面7个标识。

- feat：新功能（feature）
- fix：修补bug
- docs：文档（documentation）
- style： 格式（不影响代码运行的变动）
- refactor：重构（即不是新增功能，也不是修改bug的代码变动）
- test：增加测试
- chore：构建过程或辅助工具的变动

如果type为 `feat`和 `fix`，则该 commit 将肯定出现在 Change log 之中。其他情况（`docs`、`chore`、`style`、`refactor`、`test`）由你决定，要不要放入 Change log，建议是不要。

##### scope

scope用于说明 commit 影响的范围，比如数据层、控制层、视图层等等，视项目不同而不同。

例如在 `Angular`，可以是 `$location`, `$browser`, `$compile`, `$rootScope`, `ngHref`, `ngClick`, `ngView`等。

如果你的修改影响了不止一个 `scope`，你可以使用 `*`代替。

##### subject

`subject`是 commit 目的的简短描述，不超过50个字符。

其他注意事项：

- 以动词开头，使用第一人称现在时，比如change，而不是changed或changes
- 第一个字母小写
- 结尾不加句号（.）

#### Body

Body 部分是对本次 commit 的详细描述，可以分成多行。下面是一个范例。

```php
More detailed explanatory text, if necessary.  Wrap it to 
about 72 characters or so. 

Further paragraphs come after blank lines.

- Bullet points are okay, too
- Use a hanging indent
```

有两个注意点:

- 使用第一人称现在时，比如使用change而不是changed或changes。
- 永远别忘了第2行是空行
- 应该说明代码变动的动机，以及与以前行为的对比。

#### Footer

Footer 部分只用于以下两种情况：

##### 不兼容变动

如果当前代码与上一个版本不兼容，则 Footer 部分以BREAKING CHANGE开头，后面是对变动的描述、以及变动理由和迁移方法。

```go
BREAKING CHANGE: isolate scope bindings definition has changed.

    To migrate the code follow the example below:

    Before:

    scope: {
      myAttr: 'attribute',
    }

    After:

    scope: {
      myAttr: '@',
    }

    The removed `inject` wasn't generaly useful for directives so there should be no code using it.
```

##### 关闭 Issue

如果当前 commit 针对某个issue，那么可以在 Footer 部分关闭这个 issue 。

```bash
Closes #234
```

#### Revert

还有一种特殊情况，如果当前 commit 用于撤销以前的 commit，则必须以revert:开头，后面跟着被撤销 Commit 的 Header。

```csharp
revert: feat(pencil): add 'graphiteWidth' option

This reverts commit 667ecc1654a317a13331b17617d973392f415f02.
```

Body部分的格式是固定的，必须写成 `This reverts commit <hash>`.，其中的hash是被撤销 commit 的 SHA 标识符。

如果当前 commit 与被撤销的 commit，在同一个发布（release）里面，那么它们都不会出现在 Change log 里面。如果两者在不同的发布，那么当前 commit，会出现在 Change log 的Reverts小标题下面。

### Commitizen

可以使用典型的git工作流程或通过使用CLI向导[Commitizen](https://link.jianshu.com?t=https%3A%2F%2Fgithub.com%2Fcommitizen%2Fcz-cli)来添加提交消息格式。

### 安装

```undefined
npm install -g commitizen
```

然后，在项目目录里，运行下面的命令，使其支持 Angular 的 Commit message 格式。

```kotlin
commitizen init cz-conventional-changelog --save --save-exact
```

以后，凡是用到 `git commit`命令，一律改为使用 `git cz`。这时，就会出现选项，用来生成符合格式的 Commit message。

![img](https:////upload-images.jianshu.io/upload_images/3827973-39053e8f0259dfda.png?imageMogr2/auto-orient/strip|imageView2/2/w/557/format/webp)

5.png

### validate-commit-msg

[validate-commit-msg](https://link.jianshu.com?t=https%3A%2F%2Fgithub.com%2Fkentcdodds%2Fvalidate-commit-msg) 用于检查项目的 Commit message 是否符合Angular规范。

该包提供了使用githooks来校验 `commit message`的一些二进制文件。在这里，我推荐使用[husky](https://link.jianshu.com?t=http%3A%2F%2Fnpm.im%2Fhusky)，只需要添加 `"commitmsg": "validate-commit-msg"`到你的 `package.json`中的 `nam scripts`即可.

当然，你还可以通过定义配置文件 `.vcmrc`来自定义校验格式。详细使用请见文档 [validate-commit-msg](https://link.jianshu.com?t=https%3A%2F%2Fgithub.com%2Fkentcdodds%2Fvalidate-commit-msg)

### 生成 Change log

如果你的所有 Commit 都符合 Angular 格式，那么发布新版本时， Change log 就可以用脚本自动生成。生成的文档包括以下三个部分：

- New features
- Bug fixes
- Breaking changes.

每个部分都会罗列相关的 commit ，并且有指向这些 commit 的链接。当然，生成的文档允许手动修改，所以发布前，你还可以添加其他内容。

[conventional-changelog](https://link.jianshu.com?t=https%3A%2F%2Fgithub.com%2Fajoslin%2Fconventional-changelog) 就是生成 Change log 的工具，运行下面的命令即可。

```ruby
$ npm install -g conventional-changelog
$ cd my-project
$ conventional-changelog -p angular -i CHANGELOG.md -w
```

## 本地库操作

### 本地库初始化

1. 在cmd中进到项目目录下
2. git init：本地库初始化。生成一个.git目录，该目录中存放的是本地库相关的子目录和文件
3. 设置签名：username和email用于区分不同开发人员的身份。这里设置的签名和登录远程库(代码托管中心)的账号密码没有任何关系。默认项目级别(仓库级别)仅在当前本地库范围内生效，项目级别信息保存在.git/config下。系统用户级别：指定--global，系统用户级别信息保存在系统用户家目录下的.config文件中。项目级别优先于系统用户级别。至少设置一个
   1. git config [--global] user.name `<username>`：设置用户名
   2. git config [--global] user.email `<email>`：设置Email地址
   3. cat .git/cogfig：查看本地库配置文件

### 基本操作

- git help <命令>：查看该命令文档

### 文件操作

- git status：查看状态。
  - on branch master表示在主分支上，no commits yet表示无提交内容
  - 红色文件表示未添加到暂存区中，绿色表示已添加到暂存区中
- git add `<filename>`：将文件从工作区添加到暂存区。unstage表示从暂存区中移除
- git rm `<filename>`：从暂存区删除
- git commit [-m `<description>`] [-a] `<filename>`：将文件从暂存区中提交到本地库，添加后会进入vim，写本次提交描述内容
  - -m：无需编辑vim，直接在后面写入描述内容，；
  - -a：无需git add操作，直接添加&提交，不过就不存在暂存区的撤销操作时间了
  - 结果内容：
    - root-commit：表示根提交(第一次提交)
    - 数字编号：暂时粗略认为是本次提交的版本号

### 版本移动

- git log [--pretty={oneline}] [--oneline]：查看日志版本，只显示当前及以前版本。这里有2个版本的记录，这里hash值表示该次提交的索引，HEAD是指向当前版本的指针
  - --pretty=oneline将每个版本只以一行显示
  - --oneline不仅一行显示还只显示部分hash值
- git reflog：查看日志，显示前后所有版本以及版本移动
  - HEAD@{n}表示移动到对应版本指针需要移动n步
- git reset --{hard/mixed/soft} HEAD `<headHash>`：索引移动(推荐)
  - --mixed在本地枯移动HEAD指针，重置暂存区，默认策略
  - --hard在本地库移动HEAD指针，重置暂存区，重置工作区（会删除文件），
  - --soft仅在本地库移动HEAD指针
- git reset  --hard HEAD^^^：^移动，只能后退，3个^即表示后退3个版本
- git reset  --hard HEAD~`<n>`：~移动，只能后退，n即表示后退n个版本
- git checkout `<versionNumber>` `<filename>`：选择某个版本的文件到工作区
- git diff [HEAD] [`<filename>`]：在修改工作区文件之后
  - 无HEAD：表示工作区与暂存区的该文件比较，显示文件变化
  - 有HEAD：表示工作区与本地库的该文件比较，显示文件变化
  - 无 `<filename>`：将比较当前工作区中的所有文件

### 分支管理

分支：在版本控制过程中，使用多条线同时推进多个任务。同时并行推进多个功能开发，提高开发效率。各个分支在开发过程中，如果一个分支开发失败，不会对其他分支有任何影响，删除重新开始即可

- git branch [-v] [-a]：
  - -v：显示版本号
  - -a：包括远程分支
- git branch `<branchName>`：创建分支
- git branch -D `<branchName>`：删除分支
- git checkout `<branchName>`：切换分支
- git merge `<branchName>`：合并某分支，将当前所在分支与某分支合并。
  - 分支合并冲突：当两个分支都修改了同一个文件中同一行的内容，合并时取舍哪个分支的该处内容git是无法判断的，只需要vim手动编辑后再添加提交(提交不能带文件名)即可

### clean

```sh
git clean -f
git clean -d -fx #强制删除Untracked files
```

## 合并工程

- 假如有master、test、develop分支，现在有一个需求从master切出来一个demand，需求子任务从demand切出来一个demand/child：`git pull origin demand:demand`，`git ch demand`，`git ch -b demand/child`
  - 任务开发完后，需要merge demand/child into develop。切到任务分支 `git ch demand/child`，拉取develop分支 `git pull origin develop:develop`，会有很多改变，会因版本而merge develop into demand/child的相关提示，中止merge `git merge --abort`，不能让merge develop into demand/child发生，这反了，切到dev `git ch develop`，merge demand/child into develop `git merge demand/child`，编译检查 `go build xxx`，提交之前再检查版本 `git lg -5`是否是merge demand/child into develop，推送 `git push origin develop:develop`
  - develop环境调试完，需要merge demand/child into demand。切到任务分支 `git ch demand/child`，直接拉取远程demand并合并到demand/child `git pull origin demand:demand/child`，因为demand本来就是demand/child的源，所以直接合并之后要再合并回去 `git push origin demand/child:demand`，不会有什么影响，也能将其他人的更改合并进来，如果develop也这么做的话，会将demand/child不需要的develop的内容合并进来，那就只有回退版本再重新操作了

## git flow

开发

```shell
# 创建 feature/<brance_name> 分支
git flow feature start brance_name
# 合并 feature/<brance_name> 到 develop，删除 feature/<brance_name>，并切换到 develop
git flow feature finish brance_name
```

发布

```shell
# 创建 release-1.1.1 分支
git flow release start 1.1.1
# 合并 release-1.1.1 到 master，使用 release-1.1.1 分支名称打 tag，回合 release-1.1.1 到 develop，删除 release-1.1.1，并切换到 develop
git flow release finish 1.1.1
```

紧急修复

```shell
# 创建 hotfix-1.1.1 分支
git flow hotfix start 1.1.1
# 合并 hotfix-1.1.1 到 master；
# 如果存在 release-1.1.1，则合并 hotfix-1.1.1 到 release-1.1.1，否则合并 hotfix-1.1.1 到 develop；
# 删除 hotfix-1.1.1，并切换到 develop
git flow release finish 1.1.1
```

# GitHub

### ssh

```shell
# ed25519算法
ssh-keygen -t ed25519 -C "yuanyatianchi@example.com"
# rsa算法
ssh-keygen -t rsa -b 4096 -C "yuanyatianchi@gmail.com"
```

### 搜索

匹配语法。

- 模糊匹配：`in:<field> <xxx>`，表示匹配某字段中包含 xxx 字符串
  - name：`in:name xxx`
  - readme：`in:readme xxx`
  - descriptio：`in:description xxx`
- 数字
  - stars:>1000 或 star>1000
  - forks:>1000 或 fork>1000
- 日期
  - pushed:>2019-04-15，表示2019年4月15日后有更新的
- 语言
  - language:go，表示 go 语言项目
