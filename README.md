# Owlð¦

A dependency module feature scanning detection tool for static analysis.

---

[![DeepSource](https://deepsource.io/gh/auula/owl.svg/?label=active+issues&show_trend=true&token=2dqhjlFmox_IfR5zuVpSv64Q)](https://deepsource.io/gh/auula/owl/?ref=repository-badge)
[![codecov](https://codecov.io/gh/auula/owl/branch/main/graph/badge.svg?token=0i8L7DuJlK)](https://codecov.io/gh/auula/owl)
[![License](https://img.shields.io/badge/license-MIT-db5149.svg)](https://github.com/auula/owl/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/auula/owl)](https://goreportcard.com/report/github.com/auula/owl)

---

### ä» ç»

`Owl`æ¯ä¸æ¬¾å¼æºé¡¹ç®ä¾èµåæå·¥å·ï¼å¯ä»¥å¿«éå¨æå®çé¡¹ç®ç®å½ä¸æ¥æ¾ç¬¦åæäºç¹å¾çæºä»£ç æä»¶æèä¾èµæä»¶ãä¸ºä½å¼åäºè¿æ¬¾å·¥å·ï¼ä¾å¦å¾å¤æ¶åæä»¬é¡¹ç®å¤ªå¤§ï¼é¡¹ç®æä»¶å¤¹ä¸æå¾å¤ä¾èµæä»¶ï¼å¦ä¸ä¸ª`Java`é¡¹ç®å¼å¥äº`log4j`è¿ä¸ª`jar`ä¾èµï¼å¨é¡¹ç®ä¸­ææä»¶å­å¨å¾ªç¯ä¾èµé®é¢ãå½æä¸ªä¾èµååºç°äºæ¼æ´æ¶ï¼æ¬å·¥å·è½å¿«éæ«æé¡¹ç®ç®å½ä¸å­å¨çå¯çä¾èµæä»¶ï¼å¹¶ä¸ç»åºä¾èµæä»¶æå¨çå°åï¼å¸®å©å¼åèè½å¿«éè¿è¡å®ä½å°å¯çæä»¶ã


### å ç

ç®åçæ¬çåè½æ¯è¾ç®åï¼å·¥ä½åçå¾ç®åï¼å·¥å·ä¼å¯¹ç¹å®ç®å½è¿è¡æ«æéè¿åç½®çç¹å¾ç ç®æ³å¹éå°ç¹å®æä»¶ï¼ç¶åæ¶éä¸å¶ç¹å¾ç å¹éçæä»¶å°åï¼ç¶åå±ç¤ºåºæ¥ï¼ä¹å¯ä»¥éå®åå°ä¸ä¸ªåºå®`json`æä»¶ä¸­ä¿å­ã

![](https://tva1.sinaimg.cn/large/e6c9d24egy1h2yvkgtmbwj20lo0ca0tl.jpg)

`Owl`ç±»ä¼¼äºææ¯è½¯ä»¶ä¸æ ·ï¼åææ¯è½¯ä»¶çå·¥ä½åçå·®ä¸å¤ï¼`Owl`ä¼æ ¹æ®ä¾èµæä»¶çç¹å¾ç æ¥æ«ææ´ä¸ªé¡¹ç®ï¼åææ¯çæ¯åºå·¥ä½åçç±»ä¼¼ãå½ç¶å¦æä¸¥æ ¼æç§ææ¯è½¯ä»¶é£ç§æ ååçè¯ï¼å¯è½æ¶åä¸äºæ±ç¼ç¸å³çï¼ç®å`owl`åè½å®ç°è¿æ²¡æé£ä¹å¤æï¼åé¢ä¼çæ¬ä¼å å¥`codeql`ä»£ç åæå¼æï¼éè¿`codeql`çæ°æ®åºæ¥åéæåæåè½å¢å¼ºã

### å¿«éå¼å§

å¦ä½ä½¿ç¨`owl`ï¼ä½ å¯ä»¥åéä»åºç¶åéè¿å¦ä¸å½ä»¤ï¼

```bash
git clone git@github.com:auula/owl.git
```
å¨ä»åºåé¨æä¸ä¸ª`Makefile`æä»¶å¯ä»¥å¿«éå¸®å©ä½ æå»ºç¸åºå¹³å°çäºè¿å¶æä»¶ï¼ä¾å¦ï¼

```bash
$: make help
make darwin	| Compile executable binary for MacOS platform
make linux	| Compile executable binary for Linux platform
make windows	| Compile executable binary for Windows platform
make clean	| Clean up executable binary
```

`Owl`èµ·å ä¹æ¯ä¸º`CodeAnalysis`æç¼åçç¹å¾æ£æµå·¥å·ï¼æä»¥ä½ ä¹å¯ä»¥å¨ï¼[`https://github.com/Tencent/CodeAnalysis`](https://github.com/Tencent/CodeAnalysis/tree/main/tools/owl) è¿ä¸ªé¡¹ç®ä¸é¢ç`tools`ç®å½æ¾å°å·²ç»ç¼è¯å¥½çäºè¿å¶å¯æ§è¡æä»¶ï¼ä¸è½½å¯¹åºå¹³å°çäºè¿å¶æä»¶å³å¯ã

### å¦ä½ä½¿ç¨

ç¨åºæå»ºå®æä¼å¾å°ä¸ä¸ªäºè¿å¶æä»¶ï¼ç¨åºåç§°ä¸º`owl`ï¼å¦ä¸ä¸º`owl`æ§è¡ææï¼ä¸äºå­å½ä»¤åæ°é½å·²ç»ååºï¼

```bash
$: ./owl

			 _____  _    _  __
			(  _  )( \/\/ )(  )
			 )(_)(  )    (  )(__
			(_____)(__/\__)(____) ð¦ v0.1.2

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

å¦æä¸ç¥éå­å½ä»¤å¦ä½ä½¿ç¨ï¼å¯ä»¥å¨å¯¹åºçå­å½ä»¤åé¢åå¥`--help`åæ°ï¼å³å¯å¾å°å¸®å©ä¿¡æ¯ï¼

![](https://tva1.sinaimg.cn/large/e6c9d24egy1h2yz0laxdyj22ax0u07bb.jpg)

ä¾å¦å¦æä½ è¦æ¥æ¾`log4j`ï¼ä½ é¦åè¦éè¿`owl`è®¡ç®`log4j`ç¹å¾ç ï¼å½ä»¤å¦ä¸ï¼

```bash
$: ./owl md5 --path=/Users/ding/Downloads/log4j-1.2.17.jar
```

**æ³¨æè¿éçç¹å¾ç è®¡ç®å¿é¡»ä½¿ç¨`owl`ç¨åºçç®æ³ï¼å ä¸º`owl`éé¢çç®æ³éå¯¹å¤§æä»¶ææ¯éç¨åæ°æ®åæ¹æ¡è®¡ç®çï¼æåç¨åºè¿è¡éåº¦ï¼æä»¥å¦æä½¿ç¨å¶ä»è½¯ä»¶çç®æ³é£ä¹å°±ä¼åºç°é®é¢ï¼**

ç»æå¦ä¸ï¼

![](https://tva1.sinaimg.cn/large/e6c9d24egy1h2yz54cg72j22gm0e0af2.jpg)

ä½ ä¹å¯ä»¥ä½¿ç¨åå­è¿å¶å­ç¬¦ä¸²ç¹å¾å»æ¥æ¾ï¼

```bash
$: ./owl hex --path=/Users/ding/Downloads/log4j-1.2.17.jar
```

ç¨åºä¼å°å¯¹åºçæä»¶è½¬æåå­è¿å¶å­ç¬¦ä¸²å±ç¤ºï¼å¦ä¸å¾ï¼

![](https://tva1.sinaimg.cn/large/e6c9d24egy1h2yz7v68cbj217g0u0h0x.jpg)

ç°å¨å°±å¯ä»¥ä½¿ç¨æ«æå¨è¿è¡æ«æäºï¼å¹éæ¨¡å¼å¯ä»¥æå®ä¸º`md5`æè`hex`ï¼æªæ¥å¯è½ä¼æ·»å è·å¤çæ¨¡å¼ï¼å½ä»¤å¦ä¸ï¼

```bash
$: ./owl run --dir=/Users/ding/Downloads/ --mode=md5 --code=04a41f0a068986f0f73485cf507c0f40
```

æç´¢å¾å°å·ä½ä¾èµæä»¶ï¼

![](https://tva1.sinaimg.cn/large/e6c9d24egy1h2yze6emx3j21yq0dajwn.jpg)


**æç´¢ç»æå¦æè¿å¤ï¼å¯ä»¥éè¿`--out`åæ°å°ç»æéå®åä¿å­å°æä»¶ä¸­ä¿å­ï¼æä»¶æ ¼å¼ä¸º`json`ï¼**

### SDKæ¹å¼

ä¸é¢ä»ç»å®æ¯`command line`æ¹å¼è¿è¡çï¼`owl`ç¨åºæ¬èº«å°±æ¯ä¸ä¸ª`command line`ï¼æ ¸å¿é»è¾å¨ [`github.com/auula/owl/scan`](https://github.com/auula/owl/tree/main/scan) è¿ä¸ªåä¸­ç¼åçï¼å¦ææ³äºæ¬¡å¼åï¼é£ä¹å°±å¯ä»¥ç´æ¥ä½¿ç¨`go get github.com/auula/owl` å®è£è¿ä¸ªæ¨¡åå°ä½ é¡¹ç®éé¢ï¼ç¶åç´æ¥éè¿ç¡¬ç¼ç çæ¹å¼è¿è¡èªå®ä¹ç¼ç¨ï¼


ä¸ä¸ªç®åå®ä¾ï¼éè¿èªå®ä¹ä»£ç æ¹å¼è¿è¡ä¾èµæä»¶æ«æåæ¶éï¼


```go
package main

import (
    "fmt"

    "github.com/auula/owl/scan"
)

func main() {
    // åå»ºæ«æå¨
    scanner := new(scan.Scanner)
    // è®¾ç½®æ«æå¨è·¯å¾     
    scanner.SetPath("github.com/auula/owl") 
    // è¿åå¯¹åºè·¯å¾æææä»¶ç¹å¾ç 
    res, _ := scanner.List() 
    fmt.Println(res)

    // è®¾ç½®æå®çå¹éå¨ï¼å¶ä»å¹éå¨æ¥çAPIææ¡£
    scanner.SetMatcher(new(scan.Md5Matcher))
    // æç´¢åå«ç¹å¾ç æä»¶ï¼è¿åæä»¶è®°å½éå
    res, _ = scanner.Search("xxxx")

    // æå¼ä¸ä¸ªæä»¶æè¿°ç¬¦
    file, _ := os.OpenFile("res.json", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
    // å°ç»æä¿å­å°æå®æä»¶ä¸­
    scanner.Output(file, res)
}
```

**ä»¥ä¸å°±æ¯éè¿`SDK`æ¹å¼èªå®ä¹ç¼ç å®æä¾èµç¹å¾æ£æµã**


### å¶ä»

æé®é¢æ¬¢è¿æ`issue`ï¼å·¥å·ä¸éçè¯è®°å¾æä¸ä¸ª`â­`ï¼å¦å¤æ´å¼ºä»£ç åæå·¥å·ä½¿ç¨ï¼[`https://github.com/Tencent/CodeAnalysis`](https://github.com/Tencent/CodeAnalysis)ã
