package main

//Github api interface.

//=========================

// type Response struct {
// 	Total_count        int
// 	Incomplete_results bool
// 	Items              []Item
// }

// type Item struct {
// 	Id                uint
// 	Name              string
// 	Full_name         string
// 	Owner             Owner
// 	Private           bool
// 	Html_url          string
// 	Description       string
// 	Fork              bool
// 	Url               string
// 	Forks_url         string
// 	Keys_url          string
// 	Collaborators_url string
// 	Teams_url         string
// 	Hooks_url         string
// 	Issue_events_url  string
// 	Events_url        string
// 	Assignees_url     string
// 	Branches_url      string
// 	Tags_url          string
// 	Blobs_url         string
// 	Git_tags_url      string
// 	Git_refs_url      string
// Trees_url         string
// Statuses_url      string
// Languages_url     string
// Stargazers_url    string
// Contributors_url  string
// Subscribers_url   string
// Subscription_url  string
// Commits_url       string
// Git_commits_url   string
// Comments_url      string
// Issue_comment_url string
// Contents_url      string
// Compare_url       string
// Merges_url        string
// Archive_url       string
// Downloads_url     string
// Issues_url        string
// Pulls_url         string
// Milestones_url    string
// Notifications_url string
// Labels_url        string
// Releases_url      string
// Created_at        string
// Updated_at        string
// Pushed_at         string
// 	Git_url           string
// 	Ssh_url           string
// 	Clone_url         string
// 	Svn_url           string
// 	Homepage          string
// 	Size              uint
// 	Stargazers_count  uint
// 	Watchers_count    uint
// 	Language          string
// 	Has_issues        bool
// 	Has_downloads     bool
// 	Has_wiki          bool
// 	Has_pages         bool
// 	Forks_count       uint
// 	Mirror_url        string
// 	Open_issues_count uint
// 	Forks             uint
// 	Open_issues       uint
// 	Watchers          uint
// 	Default_branch    string
// 	Score             float64
// }
// type Owner struct {
// 	Login               string
// 	Id                  uint
// 	Avatar_url          string
// 	Gravatar_id         string
// 	Url                 string
// 	Html_url            string
// 	Followers_url       string
// 	Following_url       string
// 	Gists_url           string
// 	Starred_url         string
// 	Subscriptions_url   string
// 	Organizations_url   string
// 	Repos_url           string
// 	Events_url          string
// 	Received_events_url string
// 	Type                string
// 	Site_admin          bool
// }

import "net/http"
import "fmt"
import "encoding/json"

type Response struct {
	Results []struct {
		Path     string
		Synopsis string
	}
}

func search(s string) []string {
	var result []string
	res, err := http.Get("http://api.godoc.org/search?q=" + s)
	if err != nil {
		fmt.Println(fmt.Sprintf("\033[33m%s%s", "Can't find package named:", s))
		return nil
	}
	defer res.Body.Close()
	var packages Response
	err = json.NewDecoder(res.Body).Decode(&packages)
	if err != nil {
		fmt.Println(fmt.Sprintf("\033[33m%s", "Can't parse remote json"))
		return nil
	}
	for _, pkg := range packages.Results {
		result = append(result, pkg.Path)
		fmt.Println(fmt.Sprintf("\033[32mPath:\t\t%s", pkg.Path))
		fmt.Println(fmt.Sprintf("\033[32mSynopsis:\t%s", pkg.Synopsis))
		fmt.Println(fmt.Sprintf("\033[32mGodoc:\t\t%s", "godoc.org/pkg/"+pkg.Path))
		fmt.Println(fmt.Sprintf("\033[36mTo install it, go to \033[35m%s\033[36m, or use \033[35m`go get %s`", pkg.Path, pkg.Path))
		println()
	}
	return result
}
