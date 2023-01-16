

+++

title = "rust.base"
description = "rust.base"
tags = ["it","rust"]

+++

# rust.base

官网：https://www.rust-lang.org/zh-CN/



https://mp.weixin.qq.com/s/VheT027N71946GyoT3BxTw，![img](file:///C:\Users\yuanya\AppData\Roaming\Tencent\QQTempSys\8LDO48C$8@[GWU0353$FOVS.png)https://mp.weixin.qq.com/s/kHVGK2b0oAidORE5C8Yopg ，![img](file:///C:\Users\yuanya\AppData\Roaming\Tencent\QQTempSys\8LDO48C$8@[GWU0353$FOVS.png)https://mp.weixin.qq.com/s/muLkr6ICfsbsRc-BVVtwGg ，![img](file:///C:\Users\yuanya\AppData\Roaming\Tencent\QQTempSys\8LDO48C$8@[GWU0353$FOVS.png)https://mp.weixin.qq.com/s/dtYnwOy3FkGjnOV7vshzvg ，![img](file:///C:\Users\yuanya\AppData\Roaming\Tencent\QQTempSys\8LDO48C$8@[GWU0353$FOVS.png)https://mp.weixin.qq.com/s/y5WJbP8qcvdm4WASXEEibQ ，![img](file:///C:\Users\yuanya\AppData\Roaming\Tencent\QQTempSys\8LDO48C$8@[GWU0353$FOVS.png)https://mp.weixin.qq.com/s/a0pwSZwPziD-JHnTi5KGiQ ，![img](file:///C:\Users\yuanya\AppData\Roaming\Tencent\QQTempSys\8LDO48C$8@[GWU0353$FOVS.png)https://mp.weixin.qq.com/s/tO8CFBDZ2UsOonqARxOQUA ，![img](file:///C:\Users\yuanya\AppData\Roaming\Tencent\QQTempSys\8LDO48C$8@[GWU0353$FOVS.png)https://mp.weixin.qq.com/s/jUZkFMaC5FThEMHICBwllg ，![img](file:///C:\Users\yuanya\AppData\Roaming\Tencent\QQTempSys\8LDO48C$8@[GWU0353$FOVS.png)https://mp.weixin.qq.com/s/nrGagd49cIucvq-ULR9Oow ，![img](file:///C:\Users\yuanya\AppData\Roaming\Tencent\QQTempSys\8LDO48C$8@[GWU0353$FOVS.png)https://mp.weixin.qq.com/s/501wCGxATJ0J7kGZqsKiJw ，![img](file:///C:\Users\yuanya\AppData\Roaming\Tencent\QQTempSys\8LDO48C$8@[GWU0353$FOVS.png)https://mp.weixin.qq.com/s/ga5EES5wxfvQscJJaw97fw ，![img](file:///C:\Users\yuanya\AppData\Roaming\Tencent\QQTempSys\8LDO48C$8@[GWU0353$FOVS.png)https://mp.weixin.qq.com/s/I73xlEMRabE_0jcYS0NDFg ，![img](file:///C:\Users\yuanya\AppData\Roaming\Tencent\QQTempSys\8LDO48C$8@[GWU0353$FOVS.png)https://mp.weixin.qq.com/s/r5sh9AsezQX42w4uoTEdSA ，![img](file:///C:\Users\yuanya\AppData\Roaming\Tencent\QQTempSys\8LDO48C$8@[GWU0353$FOVS.png)https://mp.weixin.qq.com/s/lUsPmV4APjkYXj8ruOIkhg

## 环境

### rust

#### linux

```shell
curl https://sh.rustup.rs -sSf | sh
```



#### win

https://blog.csdn.net/cnds123/article/details/105770367/

安装 rust：https://www.rust-lang.org/zh-CN/tools/install 。用如下方式检查是否安装成功

```shell
# rustup 是 Rust 的工具链安装器
$ rustup -v
rustup 1.24.3 (ce5817a94 2021-05-31)
info: This is the version for the rustup toolchain manager, not the rustc compiler.
info: The currently active `rustc` version is `rustc 1.54.0 (a178d0322 2021-07-26)`
```

安装 C++ build tools：根据官方提示，在 win 上编译 rust 还需要安装[Microsoft C++ 生成工具](https://visualstudio.microsoft.com/zh-hans/visual-cpp-build-tools/)（否则编译时会报错error: linker \`link.exe\` not found），根据提示一直下一步，直到到选择安装内容，选择`使用 C++ 的桌面开发`即可

### ide

使用 goland 2021.2.3版本，搭配 [jetbrains 的 rust 插件](https://plugins.jetbrains.com/plugin/8182-rust/versions/stable/138558) Compatibility with goland。

### 中科大镜像源

```shell
cat > ~/.cargo/config <<EOF
[source.crates-io]
# registry = "https://github.com/rust-lang/crates.io-index"
replace-with = 'ustc'
[source.ustc]
registry = "https://mirrors.ustc.edu.cn/crates.io-index"
EOF
```



### rustup

rustup 是一个 Rust 工具链安装器，专门用于安装 Rust，管理 Rust 很方便：安装、升级、卸载等，还可以切换版本，包括 nightly，beta 和 stable。rustup 帮助文档如下

```shell
$ rustup
rustup 1.24.3 (ce5817a94 2021-05-31)
The Rust toolchain installer

USAGE:
    rustup [FLAGS] [+toolchain] <SUBCOMMAND>

FLAGS:
    -v, --verbose    Enable verbose output
    -q, --quiet      Disable progress output
    -h, --help       Prints help information
    -V, --version    Prints version information

ARGS:
    <+toolchain>    release channel (e.g. +stable) or custom toolchain to set override

SUBCOMMANDS:
    show           Show the active and installed toolchains or profiles
    update         Update Rust toolchains and rustup
    check          Check for updates to Rust toolchains and rustup
    default        Set the default toolchain
    toolchain      Modify or query the installed toolchains
    target         Modify a toolchain's supported targets
    component      Modify a toolchain's installed components
    override       Modify directory toolchain overrides
    run            Run a command with an environment configured for a given toolchain
    which          Display which binary will be run for a given command
    doc            Open the documentation for the current toolchain
    self           Modify the rustup installation
    set            Alter rustup settings
    completions    Generate tab-completion scripts for your shell
    help           Prints this message or the help of the given subcommand(s)

DISCUSSION:
    Rustup installs The Rust Programming Language from the official
    release channels, enabling you to easily switch between stable,
    beta, and nightly compilers and keep them updated. It makes
    cross-compiling simpler with binary builds of the standard library
    for common platforms.

    If you are new to Rust consider running `rustup doc --book` to
    learn Rust.
```



## hello

https://mp.weixin.qq.com/s/dUFJMEzOJGwM5YY8cP9MBQ

使用 ide 创建 rust 项目，默认选项即可

```rust
fn main() {
    println!("Hello, world!");
}
```

```shell
# rustc 是 Rust 的编译器，但实际上很少使用
$ rustc main.rs
$ ./main.exe
Hello, world!
```



### cargo

实际中应该一直使用 Rust 的生成工具和依赖管理器 cargo 这个工具。cargo 是 Rust 的包管理器，它能够下载 Rust 包的依赖，编译包，制作可分发的包，并将它们上传到 crates.io（Rust 社区的包注册中心）上，帮助文档如下

```shell
$ cargo
Rust's package manager

USAGE:
    cargo [+toolchain] [OPTIONS] [SUBCOMMAND]

OPTIONS:
    -V, --version                  Print version info and exit
        --list                     List installed commands
        --explain <CODE>           Run `rustc --explain CODE`
    -v, --verbose                  Use verbose output (-vv very verbose/build.rs output)
    -q, --quiet                    No output printed to stdout
        --color <WHEN>             Coloring: auto, always, never
        --frozen                   Require Cargo.lock and cache are up to date
        --locked                   Require Cargo.lock is up to date
        --offline                  Run without accessing the network
        --config <KEY=VALUE>...    Override a configuration value (unstable)
    -Z <FLAG>...                   Unstable (nightly-only) flags to Cargo, see 'cargo -Z help' for details
    -h, --help                     Prints help information

Some common cargo commands are (see all commands with --list):
    build, b    Compile the current package
    check, c    Analyze the current package and report errors, but don't build object files
    clean       Remove the target directory
    doc         Build this package's and its dependencies' documentation
    new         Create a new cargo package
    init        Create a new cargo package in an existing directory
    run, r      Run a binary or example of the local package
    test, t     Run the tests
    bench       Run the benchmarks
    update      Update dependencies listed in Cargo.lock
    search      Search registry for crates
    publish     Package and upload this package to the registry
    install     Install a Rust binary. Default location is $HOME/.cargo/bin
    uninstall   Uninstall a Rust binary

See 'cargo help <command>' for more information on a specific command.
```

常用的命令如下：

- 使用 `cargo new` 创建新的 package（包），包括可执行的和普通包。
- 使用 `cargo build` 构建你的包。
- 使用 `cargo run` 生成和运行包。
- 使用 `cargo test` 测试你的包。
- 使用 `cargo check` 进行包分析，并报告错误。
- 使用 `cargo doc` 为你的包（以及依赖包）生成文档。
- 使用 `cargo publish` 将包发布到 crates.io。
- 使用 `cargo install` 安装 Rust 可执行程序。

使用 cargo 创建项目（当然使用ide创建项目都会自动生成）

```shell
$ cargo new hello-cargo
     Created binary (application) `hello-cargo` package
$ tree hello-cargo
hello-cargo
├── Cargo.toml
└── src
    └── main.rs
```

Cargo.toml 是 Rust 的清单文件，用于保存项目和依赖的元数据信息，类似 Go Module 中的 go.mod 文件。内容如下

```toml
[package]
name = "yuanyatianchirust"
version = "0.1.0"
edition = "2018"

# See more keys and their definitions at https://doc.rust-lang.org/cargo/reference/manifest.html

[dependencies]
```

edition 字段目前可以是 2015 和 2018，默认是 2018，具体什么区别，可以认为 2018 是一个 Rust 大版本（虽然向下兼容）；authors 字段怎么获取的，可以参考 https://doc.rust-lang.org/cargo/commands/cargo-new.html，里面有详细的解释。（也可以通过 cargo help new 帮助中查到）；cargo 还生成了 git 相关的隐藏文件和文件夹：`.git` 和 `.gitignore`。也就是说，默认情况下，该项目就通过 git 进行版本控制，可以通过 `--vcs` 选项控制。

最后是 Rust 源代码。Cargo 要求，源代码必须在 src 目录下

使用 ide 的 run 按钮执行程序，用的是 cargo run 这个命令：先编译，显示编译完成相关信息，然后运行。`--package` 指定要运行的目标包名， `--bin` 指定要运行的目标二进制文件名。（实际上，针对当前 hello-cargo 项目，执行运行 cargo run 效果是一样的）。项目根目录会生成一个 target 目录，里面的文件很多，具体每个文件的作用现在不知，一般也不用去知晓，别劝退~

```shell
C:/Users/yuanya/.cargo/bin/cargo.exe run --color=always --package yuanyatianchirust --bin yuanyatianchirust
    Finished dev [unoptimized + debuginfo] target(s) in 0.01s
     Running `target\debug\yuanyatianchirust.exe`
Hello, world!
```

可以在终端输入如下命令，生产环境运行的程序应该始终使用 `--release` 选项。这时，在 target 目录下会生成一个 release 目录，而不是 debug 目录。

```shell
$ cargo run --release
   Compiling yuanyatianchirust v0.1.0 (D:\it\rust\rustproject\yuanyatianchirust)
    Finished release [optimized] target(s) in 0.23s
     Running `target\release\yuanyatianchirust.exe`
Hello, world!
```



