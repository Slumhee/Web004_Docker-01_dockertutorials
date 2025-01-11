# Web 004 Docker-01 INKKAPLUM SH

## 说明

本教程中全部文字版教程和代码为 B 站: [InkkaPlum 频道](https://space.bilibili.com/290859233) 和知乎: [Inkka Plum](https://www.zhihu.com/people/instead-opt)的相关教程所用, 仅供学习。

不得二次用于任何机构/个人再次录制 Docker / Nginx / Go / Gin / Gorm / Redis / MySQL / Vue 或其它任何语言, 框架, 架构, 工具等等教程中。

此外, 下文有任何问题或者需要改进的点, 请联系 UP 主。

## 前文

而桌面背景和 Code 背景没有变化, 毕竟 Up 主比较忙, 最近没有时间选择、更换壁纸, Code 背景是 Chinozo 的ベビーデーズ, 而桌面背景是 Harumakigohan 曲绘风格的 PJSK Jump More Jump 内容。

而值得一提的是, Code 的主题更换了一下, 换了一个非常蓝非常粉的主题, 亦为 Up 主一直以来最喜欢的三个颜色——蓝粉白之一。

如果觉得这个主题不错, 亦想采用, 这一款主题名字叫 "Doki Theme: Re:Zero: Beatrice", 欢迎尝试! 未来会考虑将这个主题和更多不同的背景进行搭配。

讲解方法和上一次的 Gin 内容相比, 有一定的小变化。

## Up 的话(前言)

这是一个比较综合的 Docker 教程, Up 顾及到了很多基本概念, 相信可以帮助各位学习者零基础入门、上手Docker。当然, 这个教程其中有一些内容涉及到了Go语言相关的重要知识和概念, 如果没学过Go语言也完全不必担心, 因为涉及的内容并不多, 如果感兴趣学习Go语言, 可以参考[本频道(哔哩哔哩(Bilibili): InkkaPlum 频道)](https://space.bilibili.com/290859233)的 Go 基础视频([两期 Fyne 内容](https://www.bilibili.com/video/BV1u142187Ps))

具体链接:

<https://www.bilibili.com/video/BV1u142187Ps>

<https://www.bilibili.com/video/BV1kz421i7iB>

强化 Go 语言知识点。

录制视频不仅是传递知识和解法, 也是对于 Up 自身的一个挑战和知识的再加强, 所以如果有任何不懂的地方, 尤其是视频没有讲清楚的地方, 欢迎私信问 Up 主, Up 主会尽力给出一个可用的解法。

当然, 请关注 Up 的 B 站(InkkaPlum频道)、知乎(InkkaPlum)、GitHub, 同时多多三连, 当然也可以充电支持 Up 主。大家的支持将给 Up 更多的动力, Up 也会努力给大家带来更好的视频。

同时, 所有课件和代码都在 GitHub 上开源地分享, 如果感到有帮助, 请给一个 Star, 并且别忘了关注 Up 的 Github。

## 再次注意

本教程中全部文字版教程和代码为 B 站: [InkkaPlum 频道](https://space.bilibili.com/290859233) 和知乎: [Inkka Plum](https://www.zhihu.com/people/instead-opt)的相关教程所用, 仅供学习。

不得二次用于任何机构/个人再次录制 Docker / Nginx / Go / Gin / Gorm / Redis / MySQL / Vue 或其它任何语言, 框架, 架构, 工具等等教程中。

此外, 下文有任何问题或者需要改进的点, 请联系 UP 主。

## 基本概念

### 1. 什么是 Docker?

我们不直接地介绍 Docker, 我们就来简单分享一下一下几个场景。

**请思考以下场景**:

在写代码的时候, 不仅仅是代码逻辑复杂, 往往于很多朋友而言, 环境配置就需要花费很多时间, 这有时并非技术不到位, 完全是因为不同的设备情况有可能截然不同, 因此环境配置常常让很多开发者感到一定的苦恼。

而且这个问题可能会在各种各样的情况下发生, 有可能是 Go 对应的环境配置, 也有可能是 MySQL 的安装和配置。

假设我们的案例项目规模越来越大, 我们可以让更多开发者一起来进行开发工作。在这个时候, 对于新成员, 很有可能需要花费不少时间来配置环境, 这亦让人感到苦恼。

我们可能听说过一句话, 叫做: "It works on my machine", 如果某些老旧的模块与当前环境不兼容, 那就麻烦了!

于是, 这个时候, 很多人想到, 能不能从根本上解决问题?

**定义**:

> > Docker is a tool that is used to automate the deployment of applications in lightweight containers so that applications can work efficiently in different environments in isolation.

**Docker 的重要核心含义**:

把一个个应用程序打包成一个个的集装箱, 然后如同这个架着 Gopher 的小鲸鱼一样, 运送到任何地方, 在不同的环境中隔离地运行。

### 2. 在我们的案例中, Docker 可以做什么?

依然参考我们的项目案例, 现在我们的前端毫无疑问是
`Vue+ElementPlus`, 而后端则是非常标准的`Gin+Gorm`的案例, 此外, 我们亦活用了`Mysql`和`Redis`。

**没有`Docker`的情况**:

假设没有 docker, 学习者 A 拿到了 up 提供的源码——前端需要配置`Nodejs`环境, 而后作为一个`vue`项目, 常需要装各种各样的 npm 依赖包。

顺其自然地, 接下来一定需要配置后端的 `go` 环境, 最后还需配置 `Redis` 和 `Mysql`。

**没有`Docker`, 换设备开发的情况**:

假如我们的项目换了个设备, 学习者 A 还想继续学习 Go 语言开发, 此时很有可能还要再来一遍。

因此**不要让配置环境阻碍你学习和开发的脚步**, 赶紧来试试看学习 Docker 吧!

### 3. Docker 和 Go 的联系, 基本概念概述

**Docker 的诞生和 Go 有很强的联系**:

Docker **使用 Go 语言进行开发实现**, 基于 Linux 内核的 `cgroup`、`namespace`, 以及 `OverlayFS` 类的 `Union FS` 等技术, 对进程进行封装隔离, 属于操作系统层面的虚拟化技术。

由于隔离的进程独立于宿主和其它的隔离的进程, 因此也称其为容器(Container)。

上面的概念比较复杂, 我们没必要完全理解, 但是我们注意到了一个词——**虚拟化**, 那么这个时候一些观众可能会联想到"虚拟机", 随之而来地, 会有一个疑问——`Docker`和虚拟机有什么区别呢?

我们可能活用过`VM Ware`体验过配置一个虚拟机, 比如说想要体验一下不一样的操作系统, 如`Windows XP`怀旧一下, 那么毫无疑问, 我们运行的这个`Windows Xp`操作系统毫无疑问是完整的——虚拟机可以提供一个完整的操作系统实例, 开发者可以在其中安装所有依赖和工具, 确保环境一致性。

**虚拟机的可能使用场景**:

最简单的案例, 我们创建一个安装配置完成 `MySQL`、`Redis`、`Nginx` 的虚拟机, 然后分发给团队成员。

**虚拟机的劣势**:

但是, 不难发现, 虚拟机会独占一部分内存和硬盘空间。它运行的时候, 其他程序就不能使用这些资源了。

哪怕虚拟机里面我们跑的是一个简单的 Go 案例, 占用资源非常小, 但是需要维持一个完整系统的运行, 虚拟机运行的是完整的操作系统, 因此会占用大量的 CPU、内存和存储资源。

### 4. 虚拟化(Virtualization)概述

**虚拟化的具体概念可如下理解**:

本质而言, **虚拟化是一种资源管理技术**, 虚拟化可以模拟硬件, 为在虚拟机中运行完整的操作系统提供硬件条件, 自然, 如前文所言, 一台物理服务器上可运行多台虚拟机, 这些虚拟机之间都是互相是隔离的, **一台虚拟机的操作系统崩溃不会影响其余的操作系统**, 虚拟机共享物理机 CPU、内存等。常用的`VM Ware`即为基于虚拟化技术的典型例。

**Hypervisor**:

而 Hypervisor 是虚拟化中非常重要的一个部分, Hypervisor 是一种运行在物理服务器和操作系统之间的中间层软件, 可以允许多个操作系统和应用共享一套基础物理硬件。

**可以将 Hypervisor 看做是虚拟环境中的"元"操作系统**, 可以协调访问服务器上的所有物理设备和虚拟机, 所以又称为虚拟机监视器(Virtual machine monitor(VMM)), 因此自然，当服务器启动并执行 Hypervisor 时, 会给每一台虚拟机分配适量的内存, CPU, 网络和磁盘资源, 并且加载所有虚拟机的客户操作系统。

### 5. 容器例——轻量级的虚拟化技术

**由于虚拟机存在这些缺点, 随着时间的变化, 另一种虚拟化技术被发展了出来——Linux 容器(Linux Containers)**。

其是一种轻量级的虚拟化技术, 因为它跟虚拟机比起来, 它少了一层 Hypervisor 层。

下图更加清晰地展示了这一点:

![alt text](image-1.png)

我们可以清晰地看出, 第一点它没有 Hypervisor 层, 另外一方面我们可以看到, 它是无`GUEST OS`的, 因为容器不需要运行一个完整的操作系统。反之, 完全可以只使用宿主机的系统, 也就是图中的`HOST OS`。

由于容器是进程级别的, 相比虚拟机有很多优势:

这一点我们就可以看出来, 其效率提高了, 速度更快了, 所需资源亦更少。自然, 我们也可以更加好的利用服务器的资源, 运行更多的容器。

一个只能运行如四五个虚拟机的设备, **也许可以解决数百个容器的运行**。

此外, 需要注意的是, **Docker 属于 Linux 容器的一种封装**, 提供简单易用的容器使用接口, 它是目前最流行的 Linux 容器解决方案。

### 6. 镜像(Image)、容器(Container)及仓库(Repository)的概念

**在具体操作 Docker 前, 不可避免地我们要先理清楚三个重点概念——镜像、容器和仓库**。

1. 首先, 清晰地, **镜像**是可以创建容器的**只读模板**, 其可以创建容器。

2. 而**镜像**每次运行之后会产生一个**容器**(`docker run`命令), 特点就是可读可写。

3. 现在请思考下面一个案例: 我的镜像保证了可创建一个能运行如`fyne` GUI 应用的环境, 我希望可把这个镜像给别人, 方便其余学习者利用或学习`fyne`, 那么此时, 即有了仓库的概念——也就是用来存放 Docker 镜像的地方。知名的公共仓库就是`Dockerhub`,我们可以得到重要的镜像 亦可以贡献重要的镜像。Dockerhub 的地址为: <https://hub.docker.com/>

4. 而 Dockerfile 则是生成镜像的配置文件, 用来书写自定义镜像的一些配置。

5. 当然, 我们亦可以以美食制作去理解Dockerfile, 镜像、容器等概念。

## 初次体验 Docker

### 1. 配置 Docker

**首先, 需要安装 Docker Desktop**, 地址如下:

<https://www.docker.com/products/docker-desktop/>

然后, 我们需要选择恰当的版本, 在 Up 主的例子中, 则是`Download for Windows - AMD64`, 这里需要结合你的版本进行选择。

![alt text](image-2.png)

点击后, 完成下载和安装的重要流程。

正常情况下, 如果一切正常, 在任务栏右下方, 点击`Show hidden icons`, 会看到 Docker 的小鲸鱼图标:

![alt text](image-3.png)

在启动 Docker 后, 我们亦可以通过`docker version`命令检查是否安装、启动成功。

正常情况的结果

```bash
Client:
 Version:
 API version:
 Go version:
 Git commit:
 Built:
 OS/Arch:
 Context:

Server: Docker Desktop 4.37.1
 Engine:
  Version:
  API version:
  Go version:
  Git commit:
  Built:
  OS/Arch:
  Experimental:
 containerd:
  Version:
  GitCommit:
 runc:
  Version:
  GitCommit:
 docker-init:
  Version:
  GitCommit:
```

大致地, 你将会看到`client – server`相关的重要输出内容。

不难发现, Docker 亦使用了这个架构模式。
Docker 使用 `client-server` 架构, `Docker client` 和 `Docker daemon` 进程通信。

**进一步理解**:

观察下图, 不难发现:

![alt text](image-4.png)

Docker 主要由三部分组成: `Docker Daemon`，`Docker Client`及`Registry`。`Docker Daemon`和`Docker Client`采用`client-server`的架构, `Docker Daemon`运行在宿主机上, `Docker Client`通过`REST API`与其交互，完成镜像、容器的管理。清晰地, 我们的`Docker Client`发送请求, 其发给`Docker Deamon`, `Docker Daemon`接受请求后, 处理并且返回。

而我们去构建镜像时, 镜像做好之后应该有一个统一存放位置, 我们称之为 Docker 仓库, `Registry`是存放各种`Docker`镜像的仓库(官方默认仓库<https://hub.docker.com>"), `Registry`分私有和公有两种。

**Repository和Registry的区别**:

**Repository**: 用于存储具体的`Docker`镜像, 起到的是仓库存储作用(一个集中的存储、分发镜像的服务)。

如`Tomcat`下面有很多个版本的镜像，它们共同组成了`Tomcat`的`Repository`, 我们通过`tag`来区分镜像版本。

>> 可以理解为相关镜像的集合(通常提供同一应用程序或服务的不同版本)。

**Registry**: 注册服务器, 管理镜像仓库, 起到的是服务器的作用, 比如官方的是`Docker Hub`, 我们一般通过`docker pull`默认是拉取官方镜像仓库的镜像。

>> 负责托管和分发镜像的服务。默认Registry是 Docker Hub。

一个 `Docker Registry`中可包含多个`仓库(Repository)`; 每个仓库可以包含多个`标签(Tag)`;每个`Tag`对应一个镜像。

如下图, `Redis`镜像内的`Tags`, 可以看到有很多不同的镜像选择。

![alt text](image-8.png)

### 2. 潜在的问题

我们在双击启动 Docker Desktop 时, 可能发现其出现了一些错误, **下面将提供两个简单的解决方案**。

#### a. 开启 Hyper-V 功能

在 Windows 例中, 请先开启 Hyper-V 功能, 具体操作如下:

以 Up 的 Win11 为例, 纵使 Win11, 很多设置在控制面板配置依然更加方便且具有普遍性, 因此即使是 Win7, 相信按照此步骤**亦可设置好**。

找到`programs`, 然后点击`turn windows feature on or off`——启用或者关闭 windows 功能即可。

![alt text](image-5.png)

![alt text](image-6.png)

然后, 勾选 HyperV, 重启即可。

![alt text](image-7.png)

#### b. 遇到`Unexpected WSL error`问题

**这通常意味着 Windows Subsystem for Linux (WSL)存在某种故障或配置问题**。

- **确保 WSL 是最新版本**。使用管理员权限运行 CMD, 执行命令来更新 WSL

```bash
wsl --update
# B站InkkaPlum频道
```

可能的结果:

```bash
Checking for updates.
The most recent version of Windows Subsystem for Linux is already installed.
```

更新完成后, **使用命令`wsl --shutdown`重启一下**, 以确保所有更新生效。

- 若依然无法解决或又遇到了其它问题, 则需考虑是不是根本没有安装`WSL`

自然, `wsl –-install`即可, 对应地, 可能需要设定如`Ubuntu`系统的用户名和密码。

然后完成对应配置即可。然后应该即可解决上述问题。

自然**若还有问题, 可评论区交流或者私信反馈**。

## 基本上手 Docker

### 1. 活用 Docker Hub

如前文所言, Docker Hub 上有大量的高质量的镜像可以用, 我们将阐述一下如何获取这些镜像。

命令是:

```bash
docker pull [选项] [Docker Registry Addr[:端口号]/]Repo Name[:标签]
```

如参考 Redis 的官方镜像:

<https://hub.docker.com/_/redis>

我们可以在页面中发现下面的命令。

参考下面的案例:

```bash
docker pull redis
```

此命令没有给出`Docker`镜像的 `Registry` 地址,因此默认将会从 Docker Hub, 获取镜像。

而镜像名称是 Redis, 因此将会获取官方镜像 Redis 仓库中最新的稳定版本的镜像。

等价于:

```bash
docker pull redis:latest
```

指定版本:

```bash
docker pull redis:7.4
```

默认标签, `dafault tag: latest`:

```bash
Using default tag: latest
latest: Pulling from library/redis
Digest:
Status: Image is up to date for redis:latest
docker.io/library/redis:latest

What's next:
    View a summary of image vulnerabilities and recommendations → docker scout quickview redis
```

### 2. 启动并且运行一个容器

**有了镜像后, 即可以此镜像为基础启动并运行一个容器**。

命令:

```bash
docker run --name my-redis -p 6379:6379 redis
```

**参数**:

1. `--name`: 给容器指定一个名字, 方便后续管理, 如果不指定的话, 往往会有一个自动生成的名字:![alt text](image-9.png)

2. `-p`: 将容器的`Redis`服务端口`6379` 映射到主机的`6379`端口。

**使用此命令运行容器后, `ターミナル`被容器占用**:

如下:

```bash
1:M 01 Jan 2025 18:45:32.990 * Server initialized
1:M 01 Jan 2025 18:45:32.990 * Ready to accept connections tcp
```

**解决方法**:

1. `-d`: 让容器在后台运行。

命令:

```bash
docker run -d --name my-redis-a -p 6379:6379 redis
```

结果:

```bash
60275b.......32c
```

**其它参数**:

1. 如`--rm`等。

### 3. 列出镜像、列出容器

命令:

```bash
docker image ls
```

Up 的例子:

```bash
REPOSITORY                      TAG       IMAGE ID       CREATED        SIZE

comprehensiveexample-backend    latest    7de....fa7   24 days ago     652MB
comprehensiveexample-frontend   latest    599....641   24 days ago     193MB
redis                           latest   87b....bd3   3 months ago   117MB
...
```

列表包含了`仓库名`、`标签`、`镜像 ID`、`创建时间` 以及`所占用的空间`。

Redis 在 Docker Hub 上看到的体积只有 40 余 MB, 但是在这里却超过了 100MB。

![alt text](image-10.png)

简单理解: `Docker Hub` 中显示的体积是压缩后的体积。

列出当下的所有容器:

```bash
docker ps -a
```

结果如下:

```bash
CONTAINER ID   IMAGE                           COMMAND                   CREATED          STATUS
PORTS                               NAMES
60275b5....c   redis                           "docker-entrypoint.s…"   30 minutes ago   Up 14 minutes               0.0.0.0:6379->6379/tcp              my-redis-a
...
```

### 删除本地镜像、删除本地仓库

**删除指定镜像**:

```bash
docker rmi <IMAGE_ID or REPOSITORY:TAG>
```

请结合实际情况, 尝试以下命令, 如:

```bash
docker rmi ea81....1379
```

或

```bash
docker rmi mariadb:latest
```

但是我们会发现, 如:

```bash
docker rmi redis:latest
```

结论:

```bash
Error response from daemon: conflict: unable to remove repository reference "redis:latest" (must force) - container 03e85....01 is using its referenced image 2724e....03
```

**因为此镜像现在正在被某个容器使用。需先删除对应的容器**。

删除所有未使用的镜像:

```bash
docker image prune
```

```bash
WARNING! This will remove all dangling images.
Are you sure you want to continue? [y/N] y
```

**删除容器**:

```bash
docker rm <CONTAINER_ID or CONTAINER_NAME>
```

例子:

```bash
docker rm flamboyant_noether
```

或

```bash
docker rm 03....b1
```

**批量删除全部停止的容器**:

```bash
docker container prune
```

```bash
WARNING! This will remove all stopped containers.
Are you sure you want to continue? [y/N] y
```

### 3. 重要的命令提示

需要注意的是, `docker run`是新建并启动一个容器。

如:

```bash
docker run my-redis-a
```

结论:

```bash
Unable to find image 'my-redis:latest' locally
docker: Error response from daemon: pull access denied for my-redis, repository does not exist or may require 'docker login': denied: requested access to the resource is denied.
See 'docker run --help'.
```

因此, 在启动已终止(`exited`)容器的时候, 正确的方法是:

`docker start`命令或`docker container start`。

案例如下:

```bash
docker start my-redis-a
# 或
docker container start my-redis-a
```

而反之, 则可以使用`docker stop`命令或`docker container stop`命令来终止一个容器。

```bash
docker stop my-redis-a
```

然后可以在 Docker desktop 中观察变化。

`docker restart` 命令用于重启容器:

```bash
docker restart my-redis-a
```

## 开始写 Dockerfile

我们将通过一个简单的 Go 案例来学习 Dockerfile 的书写。

> > 镜像是容器的基础

如我们之前的案例, 每回入力`docker run`的时候都会指定哪个镜像作为容器运行的基础。

`Dockerfile`的核心是帮助开发者定制合适的镜像。

其需以一个镜像为基础, 在其上进行定制。

### 学习`Dockerfile`内的指令

故我们将学习第一个`Dockerfile`指令, `FROM`指定**基础镜像**, 其是必备的指令且必须是第一条指令。

如:

```Dockerfile
FROM scratch
```

如果以`scratch`为基础镜像的话, 则意味着不以任何镜像为基础。

`RUN`指令——用来执行命令行命令。

如:

```Dockerfile
RUN go mod download
```

不难发现, 这个就是在`ターミナル`输入的内容。

`COPY`指令——用于复制文件。

如:

```Dockerfile
COPY package.json /usr/src/app/
```

最后, 我们了解下一个`Dockerfile`指令, `CMD`指令——容器启动指令

- `RUN`指令用于在**镜像构建过程中**执行命令, 而- `CMD`指令提供**容器启动时**的默认执行命令。

如:

```Dockerfile
CMD [ "./main" ]
# B站InkkaPlum频道
```

`RUN`和`CMD`的多种写法

- 以`shell`格式: `RUN <命令>`, 如直接在命令行中输入的命令一样。

- 以`executable(exec)`格式: RUN `["可执行文件", "参数1", "参数2"]`

最后, 我们将提及下一个指令`ENTRYPOINT`, 我们将辨析它和`CMD`的区别。

`ENTRYPOINT`是配置容器启动时运行的命令, 无法被命令中提供的额外参数替换。

#### `ENTRYPOINT`和`CMD`的一起活用

请思考, 假如我们在写`docker run`的时候添加了一些参数, 如下:

```bash
docker run IMAGE [COMMAND] [ARG...]
#B站Inkkaplum频道
```

Dockerfile 内的内容:

```Dockerfile
CMD ["echo", "Default message"]
```

我们敲的命令

```bash
docker run my-image
# InkkaPlum频道
```

观察输出:

```bash
Default message
```

但是倘若我们命令加上了参数, 我们希望输出为`Overridden message`。

```bash
docker run my-image echo "Overridden message"
```

观察输出:

```bash
Overridden message
```

我们可以把`CMD`的可执行文件部分写到`ENTRYPOINT`处。

如:

```Dockerfile
ENTRYPOINT ["echo"]
CMD ["Default Message"]
```

然后我们参考输出

```bash
docker run my-image
```

输出:

```bash
Default Message
```

同理,

```Dockerfile
ENTRYPOINT ["echo", "Default"]
CMD ["Message"]
```

最后结论亦是如此。

```bash
Default Message
```

这两个到底有什么区别呢?

例子:

Dockerfile 片段:

```Dockerfile
ENTRYPOINT ["echo", "EntryPoint:"]

CMD ["Default CMD"]
```

输出:

```bash
EntryPoint: Default CMD
# InkkaPlum频道
```

参数覆盖尝试:

```bash
docker run entrypoint-example Overridden CMD
```

输出:

```bash
EntryPoint: Overridden CMD
# InkkaPlum频道
```

#### 重要规则

1. 如果有多个`ENTRYPOINT`指令, 则最后一个会覆盖前面的
2. 如果有多个`CMD`指令, 则最后一个会覆盖前面的

## 实战

我们的要求很简单, 写出每个(或者大部分) Go 学习者第一次写的代码——`"Helloworld"`。

回忆:

1. `go mod init DockerTest`

2. 新建文件 `main.go`并且复制代码, 然后敲命令`go run .`

3. 结论: `Helloworld!`

```go
package main

import ("fmt")

func main(){
 fmt.Println("Helloworld!")
}
```

我们的工作具体流程

前情提要:

> > Starting in Go 1.13, module mode will be the default for all development. ——从 Go 1.13 开始，模块模式将成为所有开发的默认模式

因此我们将尽可能地利用它的优势去开发它。因此本质上, 我们只需要正确利用`go mod`进行包管理, 完全不需要有任何额外配置。

安装插件 Docker —— 其可以让我们在 VSC 内方便地管理容器、镜像, 并且在写 Dockerfile 的时候提供代码提示支持。

新建文件`Dockerfile`。

### 1. `FROM`指定基础镜像

我们这里参考官方镜像的解法, 其提供了下述案例代码:

```Dockerfile
FROM golang:1.23

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app ./...

CMD ["app"]
```

最直接的方法就是`golang:<版本号>`

`golang:<version>-alpine`, 这个镜像则是基于流行的`Alpine Linux`项目。好处是体积相对较小。

官方镜像:<https://hub.docker.com/_/golang/>

### 2. `COPY`指令的活用

1. 我们需要正确利用`go mod`进行包管理, 因此首当其冲地, 应该考虑把`go.mod`和`go.sum`文件 COPY 到我们镜像的对于目录下。

```Dockerfile
COPY go.mod .
```

用`WORKDIR`命令指定工作目录

```Dockerfile
WORKDIR /app
```

然后我们需要复制其它的文件, 在我们的案例里, 虽然只有一个`main.go`, 但是如果是一个比较复杂的项目, 那么我们不可能一个一个地复制, 解法:

```Dockerfile
COPY . .
```

### 3. 构建应用程序

我们采用`RUN`指令, 以`shell`模式来写:

```Dockerfile
RUN go build -o main .
```

### 4. 运行它

这里则是常规的`CMD`指令, 由于刚刚已经生成了一个可执行文件, 因此这里按照`exec`模式来写。

```Dockerfile
CMD ["./main"]
```

即可。

成品如下:

```Dockerfile
FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod .

COPY . .

RUN go build -o main .

CMD [ "./main" ]
```

### 5. 构建镜像并且测试

我们学习`docker build`命令

```bash
docker build [OPTIONS] PATH | URL
```

新参数解析: `-t`表示为镜像指定一个名称和标签。如`my-image:1.0`。

而后面的路径在我们的案例中这就只需要`.`即可, 因为当前目录即为 Dockerfile 的路径。

因此应该是

```bash
docker build -t golangtest:1.0 .
```

然后运行即可, myGO!

```bash
docker run --name my-go golangtest:1.0
```

输出:

```bash
Helloworld
```

## 进阶实战

简单尝试一下 Docker Compose

### 1. Docker Compose 介绍

1. 定义: Docker Compose 是一个用于定义和运行多容器 Docker 应用的工具。

2. 典型场景:

- 后端服务 (Gin)。
- 数据库服务 (MySQL)。
- 缓存服务 (Redis)。
- 前端服务 (Vue/React)。

Docker 容器的核心理念之一是**每个容器应该只包含一个主要的应用服务进程, 这一理念被称为"单进程容器"模型**, 我们需要利用`Docker Compose` 来解决多容器运行的问题。

实现模式: Docker Compose 将允许用户在一个`YAML`文件中定义和管理多个容器。(通常命名为 `docker-compose.yml`)。

Docker Compose 在 Windows 的案例中, 不需要单独去进行安装操作, 因为其作为 Docker Desktop 的一部分提供。

命令:

```bash
docker-compose --version
# InkkaPlum频道
```

结论:

```bash
Docker Compose version v2.30.3-desktop.1
# B站InkkaPlum频道
```

### 2. Docker Compose 文件基本格式

- 版本`version`: 指定`Compose`文件格式的版本, 版本决定了可用的配置项。*The attribute `version` is obsolete, it will be ignored, please remove it to avoid potential confusion*——这一行现在是可以被省略掉的，也建议省略。

- 服务`services`: 定义了应用中的每个容器(服务), 每个服务可以使用不同的镜像, 其中:

  - 镜像`image`: 从指定的镜像中启动容器

  - 构建`build`:

    - 上下文`context`: 指定构建镜像的 dockerfile 的上下文路径

    - `dockerfile`: 指定 Dockerfile 的文件名, 默认是 Dockerfile。

    - 端口`ports`: 映射容器和宿主机的端口

    - 环境变量`environment`: 设置服务运行所需的环境变量

- 网络`networks`: 容器间的网络连接。

- 卷`volumes`: 用于如数据持久化的数据卷定义。

- 依赖`depends_on`: 依赖配置项。

**简单的结构案例**:

```yml
services:
  backend:
    image:
  mysql:
    image:
  redis:
    image:
```

我们定义了三个服务, 分别是后端`backend`, `MySQL`和`Redis`。

现在我们有一个简单的 Go+Gorm+Redis+MySQL 的案例, 我们全部写在了一个`main.go`的文件内。

```go
package main

import (
 "github.com/gin-gonic/gin"
 "github.com/go-redis/redis/v8"
 "gorm.io/driver/mysql"
 "gorm.io/gorm"
 "context"
 "net/http"
)

var (
 db    *gorm.DB
 rdb   *redis.Client
 ctx   = context.Background()
)

type User struct {
 ID   uint   `json:"id" gorm:"primaryKey"`
 Name string `json:"name"`
}

func initMySQL() {
 dsn := "root:password@tcp(mysql:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
 var err error
 db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
 if err != nil {
  panic("failed to connect to MySQL: " + err.Error())
 }

 db.AutoMigrate(&User{})
}

func initRedis() {
 rdb = redis.NewClient(&redis.Options{
  Addr: "redis:6379",
  DB: 0,
  Password: "",
 })
}

func main() {

 // Initialize MySQL and Redis
 initMySQL()
 initRedis()

 r := gin.Default()

 // MySQL Route
 r.GET("/users", func(c *gin.Context) {
  var users []User
  result := db.Find(&users)
  if result.Error != nil {
   c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
   return
  }
  c.JSON(http.StatusOK, users)
 })

 // Redis Route
 r.GET("/redis", func(c *gin.Context) {
  // Set a value in Redis
  err := rdb.Set(ctx, "message", "Hello, Redis!", 0).Err()
  if err != nil {
   c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
   return
  }

  val, err := rdb.Get(ctx, "message").Result()
  if err != nil {
   c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
   return
  }
  c.JSON(http.StatusOK, gin.H{"message": val})
 })

 r.Run(":8080") // Run on port 8080
}
```

#### 写`docker-compose.yml`

`mysql` 服务提示, 这是我们定义的第二个服务, 即为 mysql

我们可以参考官方案例:

```yml
# Use root/example as user/password credentials
version: "3.1"

services:
  db:
    image: mysql
    restart: always
    environment:
      MYSQL_ROOT_PASSWORD: example
    # (this is just an example, not intended to be a production configuration)
```

>> 官方镜像对于环境变量的说明: When you start the mysql image, you can adjust the configuration of the MySQL instance by passing one or more environment variables on the docker run command line. Do note that none of the variables below will have any effect if you start the container with a data directory that already contains a database: any pre-existing database will always be left untouched on container startup.

`MYSQL_ROOT_PASSWORD`: MySQL 的 root 用户密码。
`MYSQL_DATABASE`: 初始化时自动创建的数据库名称。

不难发现其是 `key - value (kv)` 对格式的(键值对格式)

```yml
MYSQL_ROOT_PASSWORD: password
MYSQL_DATABASE: testdb
```

`restart`的作用: 官方案例是 `always`, **这里即为配置容器重启策略**。

### 3. Docker 数据卷概述

**定义**: `数据卷`是被设计用来持久化数据的, 它的生命周期独立于容器——故 Docker 不会在容器被删除后自动删除`数据卷`,并且也不存在垃圾回收这样的机制来处理没有任何容器引用的`数据卷`。

如果需要在删除容器的同时移除数据卷。可以在删除容器的时候使用`docker rm -v`这个命令即可避免产生任何多余的数据卷。

当然, `docker volume prune`此命令可以删除很多无用的数据卷, 减少不必要的空间占用。

结论:

```bash
WARNING! This will remove anonymous local volumes not used by at least one container.
Are you sure you want to continue? [y/N] y
Total reclaimed space: 0B
```

使用数据卷的好处:

1. 容器的文件系统是临时的, 容器停止或删除时, 容器内的数据会丢失。

2. 跨容器共享数据: 数据卷可让多个容器共享和读取数据。

在`docker compose`中, 我们可以用`volumes`指定数据卷

```yml
volumes:
  - mysql_data:/var/lib/mysql
```

### 4. 数据卷的种类

**命名卷**:

如下面的代码, 定义了一个名为`dbdata` 的命名卷, 其被`db`服务用来存储数据库数据。需要说明一下, 如果路径为数据卷名称, 必须在文件中配置数据卷。

在声明这个命名卷的存在。只写名称就足够了, 表示使用默认配置创建这个命名卷。

官方文档解释:

> > An entry under the top-level `volumes` section can be empty, in which case it uses the container engine's default configuration for creating a volume.

```yml
services:
  mysql:
    image: mysql
    volumes:
      - dbdata:/var/lib/mysql

volumes:
  dbdata:
```

**宿主机卷**:

将当前目录下的 html 文件夹挂载到了 web 服务的 `/usr/share/nginx/html`目录。

```yml
services:
  web:
    image: nginx
    volumes:
      - ./html:/usr/share/nginx/html
```

**匿名卷**:

backend 服务将创建一个匿名卷来存储 `/app/data` 目录的数据。

```yml
services:
  backend:
    image: backend
    volumes:
      - /app/data
```

采用第一种最流行的命名卷解决方法。

### 5. Redis 持久化复习

持久化

定义: 持久化是指将数据写入持久存储(durable storage), 如固态硬盘(SSD)。

Redis 提供了一系列选项。

1. `RDB(Redis Database)`: `RDB` 持久化通过在指定的时间间隔内创建数据集的快照来保存数据。

2. `AOF(Append Only File)`: `AOF` 持久化记录服务器接收到的每一个写操作，并将这些操作追加到日志文件中。

3. 无持久化: 完全禁用 `Redis` 的持久化机制, 这意味着 `Redis` 只会在内存中存储数据。

4. `AOF + RDB` 组合: 可以在同一个实例中同时启用 `RDB` 和 `AOF` 持久化。

### 6. Docker compose 的持久化设置方法

为了配置持久化, 我们需要修改配置文件, 在Docker的情况下, 我们应该怎么做呢?

1. 下载`Redis`官方提供的`redis.conf`配置文件——<https://github.com/redis/redis/blob/unstable/redis.conf>

RDB 部分:

找到

```conf
save 60 1000
```

AOF 部分

找到

将 no 改为 yes, 开启 AOF

```conf
appendonly yes
```

然后

```yml
volumes:
      - redis_data:/data
      - ./redis.conf:/usr/local/etc/redis/redis.conf
    command: 

volumes:
  redis_data:
```

`command`: command 字段指定了要在容器启动时运行的命令

案例:

```bash
command: mycommand arg1 arg2
```

案例解法

```yml
version: "3.3"
services:
  redis:
    image: redis:latest
    container_name: redis
    restart: always
    ports:
      - "6379:6379"
    volumes:
      - ./data:/data
      - ./redis.conf:/usr/local/etc/redis/redis.conf
      - ./logs:/logs
    #配置文件启动
    command: redis-server /usr/local/etc/redis/redis.conf
```

使用配置文件(`redis.conf`)启动 redis 服务，首先需要敲 `redis-server` 然后附上对应配置文件的具体路径即可。

`/usr/local/etc/redis/redis.conf` 是容器内部 Redis 默认配置文件路径。

```yml
services:
  backend:
    build:
      context: .
      dockerfile: Dockerfile
    ports:
      - "8080:8080" # Expose backend on port 8080
    depends_on:
      - mysql
      - redis
    environment:
      - DB_HOST=mysql
      - DB_PORT=3306
      - DB_USER=root
      - DB_PASSWORD=password
      - DB_NAME=testdb
      - REDIS_HOST=redis
      - REDIS_PORT=6379

  mysql:
    image: mysql:8.0
    container_name: mysql
    restart: always
    environment:
      MYSQL_ROOT_PASSWORD: password
      MYSQL_DATABASE: testdb
    ports:
      - "3306:3306"
    volumes:
      - mysql_data:/var/lib/mysql

  redis:
    image: redis:latest
    container_name: redis
    restart: always
    ports:
      - "6379:6379"
    volumes:
      - redis_data:/data 
      - ./redis.conf:/usr/local/etc/redis/redis.conf 
    command: ["redis-server", "/usr/local/etc/redis/redis.conf"]

volumes:
  mysql_data:
  redis_data:
```

### 7. Go 语言技巧——获取 Docker 环境变量

在 Go 中, 可以使用`os.Getenv`函数来获取`Docker`环境变量。

我们将改变原有的案例代码:

下面的代码示例演示了如何使用 os.Getenv 函数获取 Docker 环境变量。

```go
 dbHost := os.Getenv("DB_HOST")
 dbPort := os.Getenv("DB_PORT")
 dbUser := os.Getenv("DB_USER")
 dbPassword := os.Getenv("DB_PASSWORD")
 dbName := os.Getenv("DB_NAME")
```

### 8. 安全性提高

在 `docker-compose.yml` 中明文写入密码可能存在安全风险。可以使用 `.env` 文件来管理敏感信息, 当然这也是建议的做法。

`Compose` 模板文件亦支持动态读取主机的系统环境变量和**当前目录下的 `.env` 文件中的变量** (即若找到`.env`文件,它会自动加载其中的环境变量。)

`.env`文件例:

```conf
DB_HOST=mysql
DB_PORT=3306
DB_USER=root
DB_PASSWORD=password
DB_NAME=testdb
REDIS_HOST=redis
REDIS_PORT=6379
```

在 `compose yml` 文件内, 换成:

```yml
REDIS_PORT=${REDIS_PORT}
```

标准格式参考:

可通过`${VARIABLE_NAME}` 的形式引用 .env 文件中的变量。

### 10. Redis 配置文件进一步更改

检查 Redis 配置文件，确保没有绑定到只允许特定 IP 地址访问。如果 Redis 配置了 `bind 127.0.0.1`, 它只会接受来自本地的连接，故注释掉即可, 且需要将`protected mode`设置`no`。

启动用命令:

```bash
docker-compose up
```

测试:

`GET, http://localhost:8080/redis`

结论:

```json
{
  "message": "Hello, Redis!"
}
```

## Docker 的更核心的理解

以上, 我们便就学习了 Dockerfile 和 Docker compose 的重要写法、用法, 而接下来, 我将给各位更详细地介绍对于镜像更核心的理解概念, 以方便各位观看下一期 Docker 进阶课程。

### 1. Docker 层(Layer)

1. 镜像(Image)是由层(Layer)组成的——在拉取镜像时, 可以看到会按层去拉取

(`Rust`镜像例)
如:

```bash
38d98af84437: Downloading [===>                                               ]  12.88MB/191.4MB
```

若想查看某个镜像有哪些层, 可使用`docker inspect <image_name>`命令

如:

```bash
docker inspect mysql:8.0
```

输出片段:

```bash
 "Layers": [
                "sha256:7600fde...d99aedt5ef",
                "sha256:0af6d64...2fdaf9da6c",
                "sha256:...",
                "sha256:...",
                "sha256:...",
                "sha256:...",
                "sha256:...",
                "sha256:...",
                "sha256:...",
                "sha256:...",
                "sha256:..."
            ]
```

### 2. 镜像的层和层之间是什么关系

1. `copy-on-write (CoW)` 策略:(百科说明)

> > 写入时复制(Copy-on-write, (简称 COW))是一种计算机程序设计领域的优化策略。
> > 简而言之: 如果有多个调用者(callers)同时请求相同资源(如内存或磁盘上的数据存储), 他们会共同获取相同的指针指向相同的资源, 直到某个调用者试图修改资源的内容时, 系统才会真正复制一份专用副本(private copy)给该调用者, 而其他调用者所见到的最初的资源仍然保持不变。这过程对其他的调用者都是透明的。此作法主要的优点是如果调用者没有修改该资源, 就不会有副本(private copy)被创建, 因此多个调用者只是读取操作时可以共享同一份资源。

此策略非常重要, 在构建镜像和创建容器时都有用到。通过此概念, 我们进一步地理解镜像中的层和层。

我们讲采用`docker history golangtest:1.0` 来查看具体层:

可以看到, 如`RUN /bin/sh -c go build -o main .` 这一层我们基于 COPY 层的文件进行操作，COPY 层我们得到了如代码、必要的配置文件的内容。

```bash
f...cb   46 hours ago   CMD ["./main"]                                   0B        buildkit.dockerfile.v0
<missing>      46 hours ago   RUN /bin/sh -c go build -o main . # buildkit     31.8MB    buildkit.dockerfile.v0
<missing>      46 hours ago   COPY . . # buildkit                              218B      buildkit.dockerfile.v0
<missing>      46 hours ago   COPY go.mod . # buildkit                         29B       buildkit.dockerfile.v0
<missing>      46 hours ago   WORKDIR /app                                     0B        buildkit.dockerfile.v0
```

### 3. Dockerfile 层的大小问题

不难发现, `CMD`, `WORKDIR`这几层的大小都是 `0B`。

### 4. 以层为单位构建镜像的好处

**以层为单位可以更高效利用磁盘空间**。

假设:

我们一开始构建了镜像`x`, 依次有`i`, `n`, `k`, `a` 四层, 然后我们再去构建镜像`y`,依次有 `i`, `n`, `k`, `p` 四层, 不难发现, `i`,`n`,`k`三层**将会使用同一份文件**。

### 5. 层和容器之间的关系

- 继续理解`COW`策略:

> > 当某一层读取低层文件时, 如果出现了要修改的点, 易发现: 其会将文件拷贝过去再改, 大小可能会比就是比低层稍微大一点点, 因此低层的镜像毫无疑问是只读的, 而不难发现, 当镜像构建完成后, 其实每一层都是只读了。

这很好地解释了为什么"**镜像**是可以创建容器的**只读模板**"。

因此, 不难发现, 在创建容器时, 其实就是在镜像基础上增加了`container layer`, 而这一层是可读可写的。

而基于`COW`策略, 若容器想修改镜像中某个文件, 不难发现, 其会先把文件拷贝到`container layer`然后再修改。

## 结语

这是一个比较综合的 Docker 教程, Up 顾及到了很多基本概念, 做到了完全的零基础, 后续还会有一个单独的 Docker 进阶视频, 由于这一期视频已经拖得有点长了, 所以请期待下一期视频。

此外, 请关注 Up 的 B 站频道和知乎, 并且别忘了一键三连, 当然如果愿意, 欢迎给 Up 充电支持, 您的支持是 Up 前进的动力, 将会鼓励 Up 给各位带来更好的视频。

同时, 所有课件和代码都在 GitHub 上分享, 如果感到有帮助, 请给一个 Star 并关注 Up 的 Github。

扩充内容: 之后还会有一个进阶 Docker 教程! 敬请期待!

Up B 站 InkkaPlum 频道

<https://space.bilibili.com/290859233>

Up 知乎

<https://www.zhihu.com/people/instead-opt>

Up 掘金

<https://juejin.cn/user/3529872175284560>

Up GitHub

<https://github.com/Slumhee>

以上 祝学习成功!

Inkka Plum
