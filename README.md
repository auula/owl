# Owl🦉

A dependency module feature scanning detection tool for static analysis.

---

[![DeepSource](https://deepsource.io/gh/auula/owl.svg/?label=active+issues&show_trend=true&token=2dqhjlFmox_IfR5zuVpSv64Q)](https://deepsource.io/gh/auula/owl/?ref=repository-badge)
[![codecov](https://codecov.io/gh/auula/owl/branch/main/graph/badge.svg?token=0i8L7DuJlK)](https://codecov.io/gh/auula/owl)
[![License](https://img.shields.io/badge/license-MIT-db5149.svg)](https://github.com/auula/owl/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/auula/owl)](https://goreportcard.com/report/github.com/auula/owl)

---

### 介 绍

`Owl`是一款开源项目依赖分析工具，可以快速在指定的项目目录下查找符合某些特征的源代码文件或者依赖文件。为何开发了这款工具？例如很多时候我们项目太大，项目文件夹下有很多依赖文件，如一个`Java`项目引入了`log4j`这个`jar`依赖，在项目中某文件存在循环依赖问题。当某个依赖包出现了漏洞时，本工具能快速扫描项目目录下存在的可疑依赖文件，并且给出依赖文件所在的地址，帮助开发者能快速进行定位到可疑文件。


### 原 理

目前版本的功能比较简单，工作原理很简单，工具会对特定目录进行扫描通过内置的特征码算法匹配到特定文件，然后收集与其特征码匹配的文件地址，然后展示出来，也可以重定向到一个固定`json`文件中保存。

![](https://tva1.sinaimg.cn/large/e6c9d24egy1h2yvkgtmbwj20lo0ca0tl.jpg)

`Owl`类似于杀毒软件一样，和杀毒软件的工作原理差不多，`Owl`会根据依赖文件的特征码来扫描整个项目，和杀毒病毒库工作原理类似。当然如果严重按照杀毒软件那种标准做的话，可能涉及一些汇编相关的，目前`owl`功能实现还没有那么复杂，后面会版本会加入`codeql`代码分析引擎，通过`codeql`的数据库来做静态分析功能增强。

### 快速开始

如何使用`owl`？你可以克隆仓库然后通过如下命令：

```bash
git clone git@github.com:auula/owl.git
```
在仓库内部有一个`Makefile`文件可以快速帮助你构建相应平台的二进制文件，例如：

```bash
$: make help
make darwin	| Compile executable binary for MacOS platform
make linux	| Compile executable binary for Linux platform
make windows	| Compile executable binary for Windows platform
make clean	| Clean up executable binaries
```

`Owl`起因也是为`CodeAnalysis`所编写的特征检测工具，所以你也可以在：[`https://github.com/Tencent/CodeAnalysis`](https://github.com/Tencent/CodeAnalysis/tree/main/tools/feature_check) 这个项目下面的`tools`目录找到已经编译好的二进制可执行文件，下载对应平台的二进制文件即可。

### 如何使用

程序构建完成会得到一个二进制文件，程序名称为`owl`，如下为`owl`执行效果，一些子命令参数都已经列出：

```bash
$: ./owl

			 _____  _    _  __
			(  _  )( \/\/ )(  )
			 )(_)(  )    (  )(__
			(_____)(__/\__)(____) 🦉 v0.1.2

 A dependency module feature scanning detection tool for static analysis.


Usage:
  owl [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  hex         File hex encoding
  md5         Collection file md5
  run         Execute the scanner
  version     Version information

Flags:
  -h, --help   help for owl

Use "owl [command] --help" for more information about a command.
```

如果不知道子命令如何使用，可以在对应的子目录后面参入`--help`参数，即可得到帮助信息：

![](https://tva1.sinaimg.cn/large/e6c9d24egy1h2yz0laxdyj22ax0u07bb.jpg)

例如如果你要查找`log4j`，你首先要通过`owl`计算`log4j`特征码，命令如下：

```bash
$: ./owl md5 --path=/Users/ding/Downloads/log4j-1.2.17.jar
```

**注意这里的特征码计算必须使用`owl`程序的算法，因为`owl`里面的算法针对大文件我是采用分片方案计算的，提升程序运行速度，所有如果使用其他软件的算法那么就会出现问题！**

结果如下：

![](https://tva1.sinaimg.cn/large/e6c9d24egy1h2yz54cg72j22gm0e0af2.jpg)

你也可以使用十六进制字符串特征去查找：

```bash
$: ./owl hex --path=/Users/ding/Downloads/log4j-1.2.17.jar
```

程序会将对应的文件转成成十六进制字符串展示如下图：

![](https://tva1.sinaimg.cn/large/e6c9d24egy1h2yz7v68cbj217g0u0h0x.jpg)

如何就可以执行搜索了，搜索模式就指定`md5`和`hex`，命令如下：

```bash
$: ./owl run --dir=/Users/ding/Downloads/ --mode=md5 --code=04a41f0a068986f0f73485cf507c0f40
```

搜索得到具体依赖文件：

![](https://tva1.sinaimg.cn/large/e6c9d24egy1h2yze6emx3j21yq0dajwn.jpg)


**搜索结果如果过多，可以通过`--out`参数将结果重定向保存到文件中保存，文件格式为`json`！**

### SDK方式