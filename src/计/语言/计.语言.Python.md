

+++

title = "python"
description = "it.lang.python"
tags = ["it","lang","python"]

+++



# python

## pip

```shell
apt install python3-pip -y
```

使用--user用户安装时默认安装到 user base bin目录下，向~/.bashrc添加内容

```shell
export:$PATH:~/.local/bin
```



## virtualenv

### pipenv

> [pipenv](https://pipenv.pypa.io/en/latest/) 是一个用于帮助 Python 项目创建和管理 Python virtualenv 的工具，并在安装/卸载软件包时，在 Profile 中添加/删除软件包。
>
> [安装和使用](https://pipenv.pypa.io/en/latest/install/)。

```shell
# 使用 pip 安装 pipenv
$ pip3 install pipenv --user
```

```shell
$ mkdir pipenv-hello && cd pipenv-hello

# 创建虚拟环境。首次执行 pipenv install 将使用系统默认 python 版本创建虚拟环境，等同于 pipenv --python 3.10 && pipenv install
# 无 Pipfile 则创建，已有 Pipfile 则根据其配置安装包
$ pipenv install
Creating a virtualenv for this project...
Pipfile: /root/projects/yuanyatianchi.io/lang.python/pipenv-hello/Pipfile
Using /usr/bin/python3 (3.10.6) to create virtualenv...
⠦ Creating virtual environment...created virtual environment CPython3.10.6.final.0-64 in 369ms
  creator Venv(dest=/root/.local/share/virtualenvs/pipenv-hello-UuntEcb1, clear=False, no_vcs_ignore=False, global=False, describe=CPython3Posix)
  seeder FromAppData(download=False, pip=bundle, setuptools=bundle, wheel=bundle, via=copy, app_data_dir=/root/.local/share/virtualenv)
    added seed packages: pip==22.2.2, setuptools==65.3.0, wheel==0.37.1
  activators BashActivator,CShellActivator,FishActivator,NushellActivator,PowerShellActivator,PythonActivator

✔ Successfully created virtual environment! 
Virtualenv location: /root/.local/share/virtualenvs/pipenv-hello-UuntEcb1 # 虚拟环境位置
Creating a Pipfile for this project...
Pipfile.lock not found, creating...
Locking [packages] dependencies...
Locking [dev-packages] dependencies...
Updated Pipfile.lock (e4eef2)!
Installing dependencies from Pipfile.lock (e4eef2)...
To activate this project's virtualenv, run pipenv shell.
Alternatively, run a command inside the virtualenv with pipenv run.

# 卸载虚拟环境。不会删除 Pipfile
$ pipenv --rm
Removing virtualenv (/root/.local/share/virtualenvs/pipenv-hello-UuntEcb1)...
```

```shell
# 安装包
$ pipenv install requests

# 卸载包
$ pipenv uninstall requests
```

Pipfile

```toml
[[source]]
# 镜像源。默认为 https://pypi.org/simple，这里修改为阿里镜像源
url = "https://mirrors.aliyun.com/pypi/simple/"
```



## flask

> [flask](https://github.com/pallets/flask)；
>
> [教程](https://www.cainiaojc.com/flask)。



