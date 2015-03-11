package zlistutil

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

const (
	// V2ex
	V2EX_BASE_URL = "https://www.v2ex.com/api/topics/"
	V2EX_HOT      = "https://www.v2ex.com/api/topics/hot.json"
	V2EX_LATEST   = "https://www.v2ex.com/api/topics/latest.json"
	// Zhihu Daily
	DAILY_FETCH_NOW    = "http://news.at.zhihu.com/api/1.2/news/latest"
	DAILY_FETCH_BEFORE = "http://news.at.zhihu.com/api/1.2/news/before/"
	//Hacker News
	HACKER_NEWS_BASE_URL     = "https://news.ycombinator.com"
	HACKER_NEWS_BASE_API_URL = "https://hacker-news.firebaseio.com"
	HACKER_NEWS_TOP          = "https://hacker-news.firebaseio.com/v0/topstories.json"
	HACKER_NEWS_ITEM         = "https://hacker-news.firebaseio.com/v0/item/"
	// Jianshu
	JIANSHU_BASE_URL = "http://www.jianshu.com"
	JIANSHU_TOP_DAY  = "http://www.jianshu.com/trending/now"
	// Product Hunt
	PRODUCTHUNT_DAY = "https://api.producthunt.com/v1/posts"
)

type Item struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}

type DailyItems struct {
	News []struct {
		Title string `json:"title"`
		Url   string `json:"share_url"`
	}
}

type ProductHuntsItems struct {
	Posts []struct {
		Name    string `json:"name"`
		Tagline string `json:"tagline"`
		Url     string `json:"redirect_url"`
	}
}

func http_helper(method, url string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	req.Close = true
	// req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

func perror(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func FetchProductHunt(url string, num int) []Item {
	if num < 0 {
		return nil
	}
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Close = true
	req.Header.Set("Authorization", "Bearer 2dd283a9b3643bc72211c5c2b4aa085b7c9906d68194ea4805c00c46e7be01f4")
	res, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	perror(err)
	var productItems ProductHuntsItems
	err = json.Unmarshal(body, &productItems)
	perror(err)
	var items []Item
	num = min(num, len(productItems.Posts))
	for i := 0; i < num; i++ {
		var item Item
		item.Title = productItems.Posts[i].Name + " : " + productItems.Posts[i].Tagline
		item.Url = productItems.Posts[i].Url
		items = append(items, item)
	}
	return items
}

func FetchJianshu(url string, num int) []Item {
	if num < 0 {
		return nil
	}
	doc, err := goquery.NewDocument(url)
	perror(err)
	var items []Item
	doc.Find(".top-notes li").Each(func(i int, s *goquery.Selection) {
		var item Item
		item.Title = s.Find("h4").Text()
		item.Url, _ = s.Find("h4 a").Attr("href")
		item.Url = JIANSHU_BASE_URL + item.Url
		items = append(items, item)
	})
	num = min(num, len(items))
	return items[:num]
}
func FetchHackerNews(url string, num int) []Item {
	if num < 0 {
		return nil
	}
	body, err := http_helper("GET", url)
	perror(err)
	var ids []int
	err = json.Unmarshal(body, &ids)
	perror(err)
	// fmt.Println(ids)
	var items []Item
	num = min(num, len(ids))
	for i := 0; i < num; i++ {
		item_url := HACKER_NEWS_ITEM + strconv.Itoa(ids[i]) + ".json"
		// fmt.Println(item_url)
		body, err := http_helper("GET", item_url)
		perror(err)
		var item Item
		err = json.Unmarshal(body, &item)
		perror(err)
		if item.Url == "" { //Ask HN
			item.Url = HACKER_NEWS_BASE_URL + "/item?id=" + strconv.Itoa(ids[i])
		}
		items = append(items, item)
	}
	return items
}
func FetchV2ex(url string, num int) []Item {
	if num < 0 {
		return nil
	}
	res, err := http.Get(url)
	perror(err)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	perror(err)
	var items []Item
	err = json.Unmarshal(body, &items)
	perror(err)
	// fmt.Println(items)
	num = min(num, len(items))
	return items[:num]
}

func FetchZhihuDaily(url string, num int) []Item {
	if num < 0 {
		return nil
	}
	res, err := http.Get(url)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	perror(err)
	var daily DailyItems
	err = json.Unmarshal(body, &daily)
	perror(err)

	var items []Item
	num = min(num, len(daily.News))
	for i := 0; i < num; i++ {
		var item Item
		item.Title = daily.News[i].Title
		item.Url = daily.News[i].Url
		items = append(items, item)
	}
	return items

}