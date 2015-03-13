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
