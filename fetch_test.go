package zlistutil

import (
	"fmt"
	"testing"
)

func testGetItem(t *testing.T, site string, url string, num int) {
	fmt.Println("Test site:", site, ", url:", url)
	var item_list []Item
	item_list = GetItem(site, url, num)
	fmt.Println(item_list)
	if len(item_list) != num {
		t.Error("Site [", site, "] expected", num, "got", len(item_list))
	} else if item_list == nil {
		t.Error("Site :", site, "expected not nil")
	}
}

func TestOne(t *testing.T) {
	num := 2
	testGetItem(t, SITE_GUOKR, GUOKR_HANDPICK, num)
}

func TestAllGetItem(t *testing.T) {
	num := 2
	testGetItem(t, SITE_V2EX, V2EX_HOT, num)
	testGetItem(t, SITE_PRODUCTHUNT, PRODUCTHUNT_TODAY, num)
	testGetItem(t, SITE_JIANSHU, JIANSHU_TOP_DAY, num)
	testGetItem(t, SITE_NEXT, NEXT, num)
	testGetItem(t, SITE_HACKERNEWS, HACKER_NEWS_TOP, num)
	testGetItem(t, SITE_ZHIHUDAILY, DAILY_FETCH_NOW, num)
	testGetItem(t, SITE_WANQU, WANQU, num)
	testGetItem(t, SITE_PINGWEST, PINGWEST_NEWS, num)
	testGetItem(t, SITE_SOLIDOT, SOLIDOT, num)
	testGetItem(t, SITE_GITHUB, GITHUB, num)
	testGetItem(t, SITE_SOLIDOT, SOLIDOT, num)
	testGetItem(t, SITE_DOUBANMOMENT, DOUBAN_MOMENT, num)
	testGetItem(t, SITE_IFANR, IFANR, num)
	testGetItem(t, SITE_MINDSTORE, MINDSTORE, num)
	testGetItem(t, SITE_KICKSTARTER, KICKSTARTER, num)
	testGetItem(t, SITE_TOUTIAO, TOUTIAO, num)
	testGetItem(t, SITE_SEGMENTFAULT, SEGMENTFAULT_BLOG, num)
	testGetItem(t, SITE_THEPAPER, THEPAPER, num)
}
