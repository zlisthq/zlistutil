package zlistutil

import (
	"testing"
)

func TestFetchProductHunt(t *testing.T) {
	var item_list []Item
	num := 1
	item_list = FetchProductHunt(PRODUCTHUNT_DAY, num)
	if len(item_list) != num {
		t.Error("Expected ", num, "got", len(item_list))
	} else if item_list == nil {
		t.Error("Expected not nil")
	}
}
func TestFetchJianshu(t *testing.T) {
	var item_list []Item
	num := 1
	item_list = FetchJianshu(JIANSHU_TOP_DAY, num)
	if len(item_list) != num {
		t.Error("Expected ", num, "got", len(item_list))
	} else if item_list == nil {
		t.Error("Expected not nil")
	}
}

func TestFetchNEXT(t *testing.T) {
	var item_list []Item
	num := 1
	item_list = FetchNEXT(NEXT, num)
	if len(item_list) != num {
		t.Error("Expected ", num, "got", len(item_list))
	} else if item_list == nil {
		t.Error("Expected not nil")
	}
}

func TestFetchHackerNews(t *testing.T) {
	var item_list []Item
	num := 1
	item_list = FetchHackerNews(HACKER_NEWS_TOP, num)
	if len(item_list) != num {
		t.Error("Expected ", num, "got", len(item_list))
	} else if item_list == nil {
		t.Error("Expected not nil")
	}
}
func TestFetchV2ex(t *testing.T) {
	var item_list []Item
	num := 1
	item_list = FetchV2ex(V2EX_HOT, num)
	if len(item_list) != num {
		t.Error("Expected ", num, "got", len(item_list))
	} else if item_list == nil {
		t.Error("Expected not nil")
	}
}
func TestFetchZhihuDaily(t *testing.T) {
	var item_list []Item
	num := 1
	item_list = FetchZhihuDaily(DAILY_FETCH_NOW, num)
	if len(item_list) != num {
		t.Error("Expected ", num, "got", len(item_list))
	} else if item_list == nil {
		t.Error("Expected not nil")
	}
}
func TestFetchWanqu(t *testing.T) {
	var item_list []Item
	num := 1
	item_list = FetchWanqu(WANQU, num)
	if len(item_list) != num {
		t.Error("Expected ", num, "got", len(item_list))
	} else if item_list == nil {
		t.Error("Expected not nil")
	}
}
func TestFetchPingWestNews(t *testing.T) {
	var item_list []Item
	num := 1
	item_list = FetchPingWestNews(PINGWEST_NEWS, num)
	if len(item_list) != num {
		t.Error("Expected ", num, "got", len(item_list))
	} else if item_list == nil {
		t.Error("Expected not nil")
	}
}
func TestFetchSolidot(t *testing.T) {
	var item_list []Item
	num := 1
	item_list = FetchSolidot(SOLIDOT, num)
	if len(item_list) != num {
		t.Error("Expected ", num, "got", len(item_list))
	} else if item_list == nil {
		t.Error("Expected not nil")
	}
}
func TestFetchGitHub(t *testing.T) {
	var item_list []Item
	num := 1
	item_list = FetchGitHub(GITHUB, num)
	if len(item_list) != num {
		t.Error("Expected ", num, "got", len(item_list))
	} else if item_list == nil {
		t.Error("Expected not nil")
	}
}
func TestFetchDoubanMoment(t *testing.T) {
	var item_list []Item
	num := 1
	item_list = FetchDoubanMoment(DOUBAN_MOMENT, num)
	if len(item_list) != num {
		t.Error("Expected ", num, "got", len(item_list))
	} else if item_list == nil {
		t.Error("Expected not nil")
	}
}
func TestFetchIfanr(t *testing.T) {
	var item_list []Item
	num := 1
	item_list = FetchIfanr(IFANR, num)
	if len(item_list) != num {
		t.Error("Expected ", num, "got", len(item_list))
	} else if item_list == nil {
		t.Error("Expected not nil")
	}
}
