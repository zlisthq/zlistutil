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

## 待添加站点

- ~~Product Hunt~~
- ~~简书~~
- ~~NEXT~~
- ~~Hacker News~~
- ~~V2EX~~
- ~~知乎日报~~
- ~~湾区日报~~
- ~~PingWest News~~
- ~~Solidot~~
- ~~GitHub~~
- ~~豆瓣一刻~~
- ~~ifanr 观察~~
- ~~mindstore~~
- ~~Show HN and Ask HN~~
- ~~kickstarter:Technology sorted by Newest~~
- ~~toutiao.io~~
- 一个
- Gadget Hunt
- Beta List:newest
- Startup List
- news.mydrivers.com/blog/
- techmeme.com
- 点名时间：新品
- 澎湃新闻
- 少数派
- 什么值得买
- CC98


