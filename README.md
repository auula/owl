# Owl🦉

A dependency module feature scanning detection tool for static analysis.

---

[![DeepSource](https://deepsource.io/gh/auula/woodpecker.svg/?label=active+issues&show_trend=true&token=2dqhjlFmox_IfR5zuVpSv64Q)](https://deepsource.io/gh/auula/woodpecker/?ref=repository-badge)
[![codecov](https://codecov.io/gh/auula/woodpecker/branch/main/graph/badge.svg?token=0i8L7DuJlK)](https://codecov.io/gh/auula/woodpecker)
[![License](https://img.shields.io/badge/license-MIT-db5149.svg)](https://github.com/auula/bottle/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/auula/woodpecker)](https://goreportcard.com/report/github.com/auula/woodpecker)

---

### 介 绍

`Owl`是一款开源项目依赖分析工具，可以快速在指定的项目目录下查找符合某些特征的源代码文件或者依赖文件。为何开发了这款工具？例如很多时候我们项目太大，项目文件夹下有很多依赖文件，如一个`Java`项目引入了`log4j`这个`jar`依赖，在项目中某文件存在循环依赖问题。当某个依赖包出现了漏洞时，本工具能快速扫描项目目录下存在的可疑依赖文件，并且给出依赖文件所在的地址，帮助开发者能快速进行定位到可疑文件。


### 原 理

![](https://tva1.sinaimg.cn/large/e6c9d24egy1h2yvkgtmbwj20lo0ca0tl.jpg)