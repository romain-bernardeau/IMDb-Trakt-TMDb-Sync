package imdb

import "testing"

func TestNormalizeHref(t *testing.T) {
	for href, want := range map[string]string{
		"/fr/user/p.abc/ratings":   "/user/p.abc/ratings",
		"/de/user/ur1/watchlist/":  "/user/ur1/watchlist/",
		"/en-us/user/ur1/ratings/": "/user/ur1/ratings/",
		"/pt-br/list/ls42/?ref_=x": "/list/ls42/?ref_=x",
		"/user/ur1/ratings/":       "/user/ur1/ratings/",
		"/list/ls42/":              "/list/ls42/",
	} {
		if got := normalizeHref(href); got != want {
			t.Errorf("normalizeHref(%s) = %s, want %s", href, got, want)
		}
	}
}

func TestExportListID(t *testing.T) {
	for _, userID := range []string{"ur123", "p.abc"} {
		c := &client{userID: userID, watchlistID: "ls7"}
		for href, want := range map[string]string{
			"/list/ls599383799/":                "ls599383799",
			"/fr/list/ls42/?ref_=x":             "ls42",
			"/user/" + userID + "/watchlist/":   "ls7",
			"/fr/user/" + userID + "/watchlist": "ls7",
			"/user/other/watchlist":             "",
			"/user/" + userID + "/ratings":      "",
		} {
			if got, err := c.exportListID(href); err != nil || got != want {
				t.Errorf("exportListID(%s) = %q, %v, want %q", href, got, err, want)
			}
		}
	}
}

func TestRatingsHyperlink(t *testing.T) {
	for _, userID := range []string{"ur123", "p.abc"} {
		for _, href := range []string{"/user/" + userID + "/ratings/", "/fr/user/" + userID + "/ratings?ref_=x", "/en-us/user/" + userID + "/ratings"} {
			if !isRatingsHyperlink(normalizeHref(href), userID) {
				t.Errorf("ratings not recognised for %s: %s", userID, href)
			}
		}
		for _, href := range []string{"/user/other/ratings", "/fr/user/other/ratings", "/user/" + userID + "/watchlist", "/list/ls1/"} {
			if isRatingsHyperlink(normalizeHref(href), userID) {
				t.Errorf("false ratings match for %s: %s", userID, href)
			}
		}
	}
}
