package zlistutil

import (
	"encoding/json"
	// "fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	// "log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	//SITE NAME
	SITE_V2EX         = "v2ex"
	SITE_ZHIHUDAILY   = "zhihu_daily"
	SITE_HACKERNEWS   = "hackernews"
	SITE_JIANSHU      = "jianshu"
	SITE_PRODUCTHUNT  = "producthunt"
	SITE_NEXT         = "next"
	SITE_WANQU        = "wanqu"
	SITE_PINGWEST     = "pingwest"
	SITE_SOLIDOT      = "solidot"
	SITE_GITHUB       = "github"
	SITE_DOUBANMOMENT = "douban_moment"
	SITE_IFANR        = "ifanr"
	SITE_MINDSTORE    = "mindstore"
	SITE_KICKSTARTER  = "kickstarter"
	SITE_GADGETHUNT   = "gadgethunt"
	SITE_TOUTIAO      = "toutiao"
	SITE_SEGMENTFAULT = "segmentfault"
	SITE_THEPAPER     = "thepaper"
	SITE_GUOKR        = "guokr"

	/* URL */
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
	PRODUCTHUNT_TODAY     = "https://api.producthunt.com/v1/posts"
	PRODUCTHUNT_YESTERDAY = "https://api.producthunt.com/v1/posts?days_ago=1"
	//36kr NEXT
	NEXT_BASE_URL = "http://next.36kr.com"
	NEXT          = "http://next.36kr.com/posts"
	//湾区日报
	WANQU = "http://wanqu.co"
	//PingWest快讯
	PINGWEST_NEWS = "http://news.pingwest.com/"
	//Solidot
	SOLIDOT = "http://www.solidot.org"
	//GitHub
	GITHUB_BASE_URL = "https://github.com"
	GITHUB          = "https://github.com/trending"
	//豆瓣 一刻
	DOUBAN_MOMENT = "http://moment.douban.com/api/stream/date/"
	//ifanr 观察
	IFANR = "http://www.ifanr.com"
	//mindstore
	MINDSTORE = "http://mindstore.io"
	//Kickstarter technology newest
	KICKSTARTER_BASE_URL = "https://www.kickstarter.com"
	KICKSTARTER          = "https://www.kickstarter.com/discover/categories/technology?sort=newest"
	//gadgethunt daily
	GADGETHUNT = "http://www.gadgethunt.club/daily"
	//toutiao.io
	TOUTIAO = "http://toutiao.io"
	//segmentfault
	SEGMENTFAULT_BASE_URL = "http://segmentfault.com"
	SEGMENTFAULT_BLOG     = "http://segmentfault.com/blogs"
	//thepaper
	THEPAPER = "http://www.thepaper.cn/"
	//果壳精选
	GUOKR_HANDPICK = "http://apis.guokr.com/handpick/article.json?ad=1&category=all&retrieve_type=by_since"
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
type MomentItems struct {
	Posts []struct {
		Title string `json:"title"`
		Url   string `json:"url"`
	}
}
type ProductHuntsItems struct {
	Posts []struct {
		Name    string `json:"name"`
		Tagline string `json:"tagline"`
		Url     string `json:"redirect_url"`
	}
}

type GuokrItems struct {
	Result []struct {
		Title string `json:"title"`
		Url   string `json:"page_source"`
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

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func getDate() string {
	const layout = "2006-01-02"
	t := time.Now()
	return t.Format(layout)
}

func GetItem(site string, url string, num int) []Item {
	m := map[string]func(string, int) []Item{
		SITE_V2EX:         fetchV2ex,
		SITE_ZHIHUDAILY:   fetchZhihuDaily,
		SITE_HACKERNEWS:   fetchHackerNews,
		SITE_JIANSHU:      fetchJianshu,
		SITE_PRODUCTHUNT:  fetchProductHunt,
		SITE_NEXT:         fetchNEXT,
		SITE_WANQU:        fetchWanqu,
		SITE_PINGWEST:     fetchPingWestNews,
		SITE_SOLIDOT:      fetchSolidot,
		SITE_GITHUB:       fetchGitHub,
		SITE_DOUBANMOMENT: fetchDoubanMoment,
		SITE_IFANR:        fetchIfanr,
		SITE_MINDSTORE:    fetchMindStore,
		SITE_KICKSTARTER:  fetchKickstarter,
		SITE_TOUTIAO:      fetchToutiao,
		SITE_SEGMENTFAULT: fetchSegmentfault,
		SITE_THEPAPER:     fetchThePaper,
		SITE_GUOKR:        fetchGuokr,
	}
	return m[site](url, num)
}

func fetchProductHunt(url string, num int) []Item {
	var item_list []Item
	item_list = fetchProductHuntHelper(url, num)
	if len(item_list) == 0 && url == PRODUCTHUNT_TODAY {
		item_list = fetchProductHuntHelper(PRODUCTHUNT_YESTERDAY, num)
	}
	return item_list
}

func fetchProductHuntHelper(url string, num int) []Item {
	items := []Item{}
	if num < 0 {
		return items
	}
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Close = true
	req.Header.Set("Authorization", "Bearer 2dd283a9b3643bc72211c5c2b4aa085b7c9906d68194ea4805c00c46e7be01f4")
	res, err := client.Do(req)
	if err != nil {
		return items
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return items
	}
	var productItems ProductHuntsItems
	err = json.Unmarshal(body, &productItems)
	if err != nil {
		return items
	}
	num = min(num, len(productItems.Posts))
	for i := 0; i < num; i++ {
		var item Item
		item.Title = productItems.Posts[i].Name + " : " + productItems.Posts[i].Tagline
		item.Url = productItems.Posts[i].Url
		items = append(items, item)
	}
	return items
}

func fetchJianshu(url string, num int) []Item {
	items := []Item{}
	if num < 0 {
		return items
	}
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return items
	}
	doc.Find(".article-list li").Each(func(i int, s *goquery.Selection) {
		var item Item
		item.Title = s.Find("h4").Text()
		item.Url, _ = s.Find("h4 a").Attr("href")
		item.Url = JIANSHU_BASE_URL + item.Url
		items = append(items, item)
	})
	num = min(num, len(items))
	return items[:num]
}

func fetchNEXT(url string, num int) []Item {
	items := []Item{}
	if num < 0 {
		return items
	}
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return items
	}
	doc.Find(".post").First().Find(".product-item").Each(func(i int, s *goquery.Selection) {
		var item Item
		item.Title = s.Find(".product-url a").Text() + " : " + s.Find(".post-tagline").Text()
		item.Url, _ = s.Find(".product-url a").Attr("href")
		item.Url = NEXT_BASE_URL + item.Url
		items = append(items, item)
	})
	num = min(num, len(items))
	return items[:num]
}

func fetchHackerNews(url string, num int) []Item {
	items := []Item{}
	if num < 0 {
		return items
	}
	body, err := http_helper("GET", url)
	if err != nil {
		return items
	}
	var ids []int
	err = json.Unmarshal(body, &ids)
	if err != nil {
		return items
	}
	num = min(num, len(ids))
	for i := 0; i < num; i++ {
		item_url := HACKER_NEWS_ITEM + strconv.Itoa(ids[i]) + ".json"
		body, err := http_helper("GET", item_url)
		if err != nil {
			return items
		}
		var item Item
		err = json.Unmarshal(body, &item)
		if err != nil {
			return items
		}
		if item.Url == "" { //Ask HN
			item.Url = HACKER_NEWS_BASE_URL + "/item?id=" + strconv.Itoa(ids[i])
		}
		items = append(items, item)
	}
	return items
}
func fetchV2ex(url string, num int) []Item {
	items := []Item{}
	if num < 0 {
		return items
	}
	res, err := http.Get(url)
	if err != nil {
		return items
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return items
	}
	err = json.Unmarshal(body, &items)
	if err != nil {
		return items
	}
	num = min(num, len(items))
	return items[:num]
}

func fetchZhihuDaily(url string, num int) []Item {
	items := []Item{}
	if num < 0 {
		return items
	}
	res, err := http.Get(url)
	if err != nil {
		return items
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return items
	}
	var daily DailyItems
	err = json.Unmarshal(body, &daily)
	if err != nil {
		return items
	}
	num = min(num, len(daily.News))
	for i := 0; i < num; i++ {
		var item Item
		item.Title = daily.News[i].Title
		item.Url = daily.News[i].Url
		items = append(items, item)
	}
	return items
}

func fetchWanqu(url string, num int) []Item {
	items := []Item{}
	if num < 0 {
		return items
	}
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return items
	}
	doc.Find(".list-group-item").Each(func(i int, s *goquery.Selection) {
		var item Item
		item.Title = s.Find("a").Text()
		item.Url, _ = s.Find("a").Attr("href")
		items = append(items, item)
	})
	num = min(num, len(items))
	return items[:num]

}
func fetchPingWestNews(url string, num int) []Item {
	items := []Item{}
	if num < 0 {
		return items
	}
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return items
	}
	doc.Find(".item_title").Each(func(i int, s *goquery.Selection) {
		var item Item
		item.Title = s.Find(".topic").Text()
		item.Url, _ = s.Find(".topic").Attr("href")
		items = append(items, item)
		// fmt.Print(i)
	})
	num = min(num, len(items))
	return items[:num]

}

func fetchSolidot(url string, num int) []Item {
	items := []Item{}
	if num < 0 {
		return items
	}
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return items
	}
	doc.Find(".bg_htit").Each(func(i int, s *goquery.Selection) {
		var item Item
		item.Title = s.Find("a").Last().Text()
		item.Url, _ = s.Find("a").Last().Attr("href")
		item.Url = SOLIDOT + item.Url
		items = append(items, item)
	})
	num = min(num, len(items))
	return items[:num]
}
func fetchGitHub(url string, num int) []Item {
	items := []Item{}
	if num < 0 {
		return items
	}
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return items
	}
	doc.Find(".repo-list-item").Each(func(i int, s *goquery.Selection) {
		var item Item
		// item.Title = s.Find(".repo-list-name a").Text()
		item.Url, _ = s.Find(".repo-list-name a").Attr("href")
		item.Title = item.Url[1:] + " : " + strings.TrimSpace(s.Find(".repo-list-description").Text())
		item.Url = GITHUB_BASE_URL + item.Url
		items = append(items, item)
	})
	num = min(num, len(items))
	return items[:num]
}
func fetchDoubanMoment(url string, num int) []Item {
	items := []Item{}
	if num < 0 {
		return items
	}
	url += getDate()
	res, err := http.Get(url)
	if err != nil {
		return items
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return items
	}
	var moment MomentItems
	err = json.Unmarshal(body, &moment)
	if err != nil {
		return items
	}
	num = min(num, len(moment.Posts))
	for i := 0; i < num; i++ {
		var item Item
		item.Title = moment.Posts[i].Title
		item.Url = moment.Posts[i].Url
		items = append(items, item)
	}
	return items

}
func fetchIfanr(url string, num int) []Item {
	items := []Item{}
	if num < 0 {
		return items
	}
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return items
	}
	doc.Find(".post-item-content h2").Each(func(i int, s *goquery.Selection) {
		var item Item
		item.Title = strings.TrimSpace(s.Find("a").Text())
		item.Url, _ = s.Find("a").Attr("href")
		items = append(items, item)
	})
	num = min(num, len(items))
	return items[:num]
}

func fetchMindStore(url string, num int) []Item {
	items := []Item{}
	if num < 0 {
		return items
	}
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return items
	}
	doc.Find(".mind-list-ul li").Each(func(i int, s *goquery.Selection) {
		var item Item
		item.Title = strings.TrimSpace(s.Find(".mind-title a").Text()) + " : " + strings.TrimSpace(s.Find(".mind-des").Text())
		item.Url, _ = s.Find(".mind-title a").Attr("href")
		items = append(items, item)
	})
	num = min(num, len(items))
	return items[:num]
}

func fetchKickstarter(url string, num int) []Item {
	items := []Item{}
	if num < 0 {
		return items
	}
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return items
	}
	doc.Find(".project").Each(func(i int, s *goquery.Selection) {
		var item Item
		item.Title = s.Find(".project-card-content a").Text() + " : " + strings.TrimSpace(s.Find(".project-blurb").Text())
		item.Url, _ = s.Find(".project-thumbnail a").Attr("href")
		item.Url = KICKSTARTER_BASE_URL + item.Url
		items = append(items, item)
	})
	num = min(num, len(items))
	return items[:num]
}

func fetchToutiao(url string, num int) []Item {
	items := []Item{}
	if num < 0 {
		return items
	}
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return items
	}
	doc.Find(".post").Each(func(i int, s *goquery.Selection) {
		var item Item
		item.Title = s.Find(".content a").Text()
		item.Url, _ = s.Find(".content a").Attr("href")
		items = append(items, item)
	})
	num = min(num, len(items))
	return items[:num]
}

func fetchSegmentfault(url string, num int) []Item {
	items := []Item{}
	if num < 0 {
		return items
	}
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return items
	}
	doc.Find(".stream-list__item").Each(func(i int, s *goquery.Selection) {
		var item Item
		item.Title = s.Find(".title a").Text()
		item.Url, _ = s.Find(".title a").Attr("href")
		item.Url = SEGMENTFAULT_BASE_URL + item.Url
		items = append(items, item)
	})
	num = min(num, len(items))
	return items[:num]
}

func fetchThePaper(url string, num int) []Item {
	items := []Item{}
	if num < 0 {
		return items
	}

	resp, err := http.Get(url)
	if err != nil {
		return items
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return items
	}
	src := string(body)
	//将HTML标签全转换成小写
	//re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	//src = re.ReplaceAllStringFunc(src, strings.ToLower)

	//去除STYLE
	re, _ := regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")

	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")

	re, _ = regexp.Compile("<a(.*?)href=\"newsDetail_forward_[0-9]+[^>]*>(.*?)</a>")
	aresult := re.FindAllString(src, 10)

	for _, str := range aresult {
		var item Item
		re, _ = regexp.Compile(">(.*)<")
		title := re.FindString(str)
		item.Title = title[1 : len(title)-1]
		re, _ = regexp.Compile("newsDetail_forward_[0-9]+")
		item.Url = THEPAPER + re.FindString(str)
		items = append(items, item)
	}

	num = min(num, len(items))
	return items[:num]
}

func fetchGuokr(url string, num int) []Item {
	items := []Item{}
	if num < 0 {
		return items
	}
	res, err := http.Get(url + "&limit=" + strconv.Itoa(num))
	if err != nil {
		return items
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return items
	}
	var guokr GuokrItems
	err = json.Unmarshal(body, &guokr)
	if err != nil {
		return items
	}
	num = min(num, len(guokr.Result))
	for i := 0; i < num; i++ {
		var item Item
		item.Title = guokr.Result[i].Title
		item.Url = guokr.Result[i].Url
		items = append(items, item)
	}
	return items
}
