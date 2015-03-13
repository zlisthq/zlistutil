# zlistutil

## 简介

一个工具，提供了获取 V2EX、知乎日报、HackerNews、简书、ProductHunt 等站点最热信息的方法。

## 使用方法

```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/WhiteWorld/zlistutil"
	"log"
)

func main() {
	items := zlistutil.FetchJianshu(zlistutil.JIANSHU_TOP_DAY, 10)
	json_items, err := json.Marshal(&items)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(json_items))
}
```

## 使用该工具的站点

- [zlist](http://zlist.whiteworld.me/)

## 待添加站点

- ~~NEXT~~
- MindStore
- PingWest
- 好奇心日报
- 弯曲日报

