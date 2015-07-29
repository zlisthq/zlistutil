# zlistutil

## 简介

一个工具库，提供了获取 V2EX、知乎日报、HackerNews、简书、ProductHunt 等站点最热信息的方法。

## 使用方法

```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/zlisthq/zlistutil"
	"log"
)

func main() {
	items := zlistutil.GetItem(zlistutil.SITE_V2EX, zlistutil.V2EX_HOT, 10)
	json_items, err := json.Marshal(&items)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(json_items))
}
```

## 使用该工具的站点

- [zlist](http://zlist.whiteworld.me/)

## 添加站点

请提交待添加的网址到 [Issues](https://github.com/zlisthq/zlistutil/issues)

[待添加站点](https://github.com/zlisthq/zlistutil/labels/enhancement)
