package imdb

import "testing"

func TestExportHyperlinks(t *testing.T) {
	for _, href := range []string{"/user/ur123/ratings/", "/fr/user/p.abc/ratings?ref_=x", "/user/p.abc/ratings"} {
		if !isAnyRatingsHyperlink(href) {
			t.Errorf("ratings not recognised: %s", href)
		}
	}
	if isAnyRatingsHyperlink("/list/ls1/") || isAnyRatingsHyperlink("/user/ur1/watchlist") {
		t.Error("false ratings match")
	}
	if !isWatchlistHyperlink("/fr/user/p.abc/watchlist/") || isWatchlistHyperlink("/list/ls1/") {
		t.Error("watchlist match")
	}
	if !isListHyperlink("/fr/list/ls599383799/?ref_=x") || !isListHyperlink("/list/ls1/") {
		t.Error("list match")
	}
	for href, want := range map[string]string{"/list/ls599383799/": "ls599383799", "/fr/list/ls42/?ref_=x": "ls42"} {
		if got, err := idExtract(href); err != nil || got != want {
			t.Errorf("idExtract(%s) = %s, %v", href, got, err)
		}
	}
}
