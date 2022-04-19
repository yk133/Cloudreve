
## 增加新功能
#### 1 上传文件时files表中新增md5 
#### 2 支持通过md5搜索文件，这样文件就可以支持自动上传文件

## 适配的安卓移动端可以关注：
可自动新建目录并上传文件  
https://github.com/yk133/CloudreveApp  
  
    
  
  
  
<h1 align="center">
  <br>
  <a href="https://cloudreve.org/" alt="logo" ><img src="https://raw.githubusercontent.com/cloudreve/frontend/master/public/static/img/logo192.png" width="150"/></a>
  <br>
  Cloudreve
  <br>
</h1>

<h4 align="center">支持多家云存储驱动的公有云文件系统.</h4>

<p align="center">
  <a href="https://travis-ci.com/github/cloudreve/Cloudreve/">
    <img src="https://img.shields.io/travis/com/cloudreve/Cloudreve?style=flat-square"
         alt="travis">
  </a>
  <a href="https://codecov.io/gh/cloudreve/Cloudreve"><img src="https://img.shields.io/codecov/c/github/cloudreve/Cloudreve?style=flat-square"></a>
  <a href="https://goreportcard.com/report/github.com/cloudreve/Cloudreve">
      <img src="https://goreportcard.com/badge/github.com/cloudreve/Cloudreve?style=flat-square">
  </a>
  <a href="https://github.com/cloudreve/Cloudreve/releases">
    <img src="https://img.shields.io/github/v/release/cloudreve/Cloudreve?include_prereleases&style=flat-square">
  </a>
</p>

<p align="center">
  <a href="https://demo.cloudreve.org">演示站</a> •
  <a href="https://forum.cloudreve.org/">讨论社区</a> •
  <a href="https://docs.cloudreve.org/">文档</a> •
  <a href="https://github.com/cloudreve/Cloudreve/releases">下载</a> •
  <a href="https://t.me/cloudreve_official">Telegram 群组</a> •
  <a href="#scroll-许可证">许可证</a>
</p>

![Screenshot](https://raw.githubusercontent.com/cloudreve/docs/master/images/homepage.png)

## :sparkles: 特性

* :cloud: 支持本机、从机、七牛、阿里云 OSS、腾讯云 COS、又拍云、OneDrive (包括世纪互联版) 作为存储端
* :outbox_tray: 上传/下载 支持客户端直传，支持下载限速
* 💾 可对接 Aria2 离线下载，可使用多个从机机点分担下载任务
* 📚 在线 压缩/解压缩、多文件打包下载
* 💻 覆盖全部存储策略的 WebDAV 协议支持
* :zap: 拖拽上传、目录上传、流式上传处理
* :card_file_box: 文件拖拽管理
*  :family_woman_girl_boy:   多用户、用户组
* :link: 创建文件、目录的分享链接，可设定自动过期
* :eye_speech_bubble: 视频、图像、音频、文本、Office 文档在线预览
* :art: 自定义配色、黑暗模式、PWA 应用、全站单页应用
* :rocket: All-In-One 打包，开箱即用
* 🌈 ... ...

## :hammer_and_wrench: 部署

下载适用于您目标机器操作系统、CPU架构的主程序，直接运行即可。

```shell
# 解压程序包
tar -zxvf cloudreve_VERSION_OS_ARCH.tar.gz

# 赋予执行权限
chmod +x ./cloudreve

# 启动 Cloudreve
./cloudreve
```

以上为最简单的部署示例，您可以参考 [文档 - 起步](https://docs.cloudreve.org/) 进行更为完善的部署。

## :gear: 构建

自行构建前需要拥有 `Go >= 1.13`、`yarn`等必要依赖。

#### 克隆代码

```shell
git clone --recurse-submodules https://github.com/cloudreve/Cloudreve.git
```

#### 构建静态资源

```shell
# 进入前端子模块
cd assets
# 安装依赖
yarn install
# 开始构建
yarn run build
```

#### 嵌入静态资源

```shell
# 回到项目主目录
cd ../

# 安装 statik, 用于嵌入静态资源
go get github.com/rakyll/statik

# 开始嵌入
statik -src=assets/build/  -include=*.html,*.js,*.json,*.css,*.png,*.svg,*.ico -f
```

#### 编译项目

```shell
# 获得当前版本号、Commit
export COMMIT_SHA=$(git rev-parse --short HEAD)
export VERSION=$(git describe --tags)

# 开始编译
go build -a -o cloudreve -ldflags " -X 'github.com/cloudreve/Cloudreve/v3/pkg/conf.BackendVersion=$VERSION' -X 'github.com/cloudreve/Cloudreve/v3/pkg/conf.LastCommit=$COMMIT_SHA'"
```

你也可以使用项目根目录下的`build.sh`快速开始构建：

```shell
./build.sh  [-a] [-c] [-b] [-r]
	a - 构建静态资源
	c - 编译二进制文件
	b - 构建前端 + 编译二进制文件
	r - 交叉编译，构建用于release的版本
```

## :alembic: 技术栈

* [Go ](https://golang.org/) + [Gin](https://github.com/gin-gonic/gin)
* [React](https://github.com/facebook/react) + [Redux](https://github.com/reduxjs/redux) + [Material-UI](https://github.com/mui-org/material-ui)

## :scroll: 许可证

GPL V3

---
> GitHub [@HFO4](https://github.com/HFO4) &nbsp;&middot;&nbsp;
> Twitter [@abslant00](https://twitter.com/abslant00)
