# Node.js

## 安装配置

### windows

1. 下载：https://nodejs.org/zh-cn/download/

2. 解压：将node.exe所在目录添加到环境变量E:\it\front-end\Nodejs

   ```cmd
   # cmd测试
   node -v # 查看node版本
   npm -v # 查看npm版本，一般不是最新版本，可以在安装配置好后更新或重新全局安装
   ```

3. 配置`npm`的`prefix`和`cache`。它们的路径默认在C盘，修改以便将来查看

   ```cmd
   # 配置前需要先在`nodejs`的安装目录下创建`node_global`和`node_cache`文件夹
   npm config set prefix "E:\it\front_end\Nodejs\node_global" # 配置全局目录路径
   npm config set cache "E:\it\front_end\Nodejs\node_cache" # 配置缓存目录路径
   
   npm config list #查看基本配置
   npm config list -l #查看所有配置
   ```

4. 环境变量
   - `PATH`：`E:\it\front_end\nodejs`(必选，安装程序已经为你配好的)；`E:\it\front_end\nodejs\node_global`(可选，如果想要使用`node_global`下的命令程序就必须配置)

5. 淘宝镜像（可选）

   ```cmd
   # 修改npm的registry 配置
   npm config set registry "https://registry.npm.taobao.org"
   ```



## npm

-g表示全局安装，无-g在当前目录下安装到命令行所在目录的node_module目录下

##### 模块查看

```cmd
npm ls -g # -g表示全局目录，即查看prefix配置的node_global目录中的node_modules目录下。缺省则查看本地目录
```

##### 模块安装

```cmd
npm install 模块名 -g # 同上
npm install --global windows-build-tools #报错缺少某些工具时运行可以顺带安装构建工具
```

##### 模块删除

```cmd
npm uninstall 模块名 -g # 同上
```

##### 模块更新

```cmd
npm update 模块名 -g # 同上
```

##### 配置查看

```cmd
npm config ls #查看基本配置简写
npm config list #查看基本配置
npm config list -l #查看所有配置
```

##### 配置修改

```cmd
npm config set 配置名 配置值
npm config get 配置名
```

##### 版本查看

```cmd
npm -v
```

## Vue

```shell
npm install webpack -g #安装webpack
```

```shell
npm install -g @vue/cli #安装vue脚手架
```

```shell
vue init webpack vue-project-demo
#开始下载
#创建文件夹vue-projiect-demo作为项目目录
#项目名：随便取
#项目描述：随意
#作者信息：随意，直接回车
#环境：选择Runtime+Compiler，运行+编译环境
#安装vue-router：路由。Y
#安装ESLint：代码检查和语法规范。N
#安装单元测试：N
#e2e测试：N
#包管理工具：选择NPM
#等待下载结束
npm run dev #进入项目运行
```
