package zlistutil

import (
	"testing"
)

func TestGetItem(t *testing.T) {
	var item_list []Item
	num := 1
	item_list = GetItem(SITE_V2EX, V2EX_HOT, num)
	if len(item_list) != num {
		t.Error("Expected ", num, "got", len(item_list))
	} else if item_list == nil {
		t.Error("Expected not nil")
	}
	item_list = GetItem(SITE_PRODUCTHUNT, PRODUCTHUNT_TODAY, num)
	if len(item_list) != num {
		t.Error("Expected ", num, "got", len(item_list))
	} else if item_list == nil {
		t.Error("Expected not nil")
	}
	item_list = GetItem(SITE_JIANSHU, JIANSHU_TOP_DAY, num)
	if len(item_list) != num {
		t.Error("Expected ", num, "got", len(item_list))
	} else if item_list == nil {
		t.Error("Expected not nil")
	}
	item_list = GetItem(SITE_NEXT, NEXT, num)
	if len(item_list) != num {
		t.Error("Expected ", num, "got", len(item_list))
	} else if item_list == nil {
		t.Error("Expected not nil")
	}
	item_list = GetItem(SITE_HACKERNEWS, HACKER_NEWS_TOP, num)
	if len(item_list) != num {
		t.Error("Expected ", num, "got", len(item_list))
	} else if item_list == nil {
		t.Error("Expected not nil")
	}
	item_list = GetItem(SITE_ZHIHUDAILY, DAILY_FETCH_NOW, num)
	if len(item_list) != num {
		t.Error("Expected ", num, "got", len(item_list))
	} else if item_list == nil {
		t.Error("Expected not nil")
	}
	item_list = GetItem(SITE_WANQU, WANQU, num)
	if len(item_list) != num {
		t.Error("Expected ", num, "got", len(item_list))
	} else if item_list == nil {
		t.Error("Expected not nil")
	}
	item_list = GetItem(SITE_PINGWEST, PINGWEST_NEWS, num)
	if len(item_list) != num {
		t.Error("Expected ", num, "got", len(item_list))
	} else if item_list == nil {
		t.Error("Expected not nil")
	}
	item_list = GetItem(SITE_SOLIDOT, SOLIDOT, num)
	if len(item_list) != num {
		t.Error("Expected ", num, "got", len(item_list))
	} else if item_list == nil {
		t.Error("Expected not nil")
	}
	item_list = GetItem(SITE_GITHUB, GITHUB, num)
	if len(item_list) != num {
		t.Error("Expected ", num, "got", len(item_list))
	} else if item_list == nil {
		t.Error("Expected not nil")
	}
	item_list = GetItem(SITE_SOLIDOT, SOLIDOT, num)
	if len(item_list) != num {
		t.Error("Expected ", num, "got", len(item_list))
	} else if item_list == nil {
		t.Error("Expected not nil")
	}
	item_list = GetItem(SITE_DOUBANMOMENT, DOUBAN_MOMENT, num)
	if len(item_list) != num {
		t.Error("Expected ", num, "got", len(item_list))
	} else if item_list == nil {
		t.Error("Expected not nil")
	}
	item_list = GetItem(SITE_IFANR, IFANR, num)
	if len(item_list) != num {
		t.Error("Expected ", num, "got", len(item_list))
	} else if item_list == nil {
		t.Error("Expected not nil")
	}
	item_list = GetItem(SITE_MINDSTORE, MINDSTORE, num)
	if len(item_list) != num {
		t.Error("Expected ", num, "got", len(item_list))
	} else if item_list == nil {
		t.Error("Expected not nil")
	}
	item_list = GetItem(SITE_KICKSTARTER, KICKSTARTER, num)
	if len(item_list) != num {
		t.Error("Expected ", num, "got", len(item_list))
	} else if item_list == nil {
		t.Error("Expected not nil")
	}
}
