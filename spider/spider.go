package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func clearUrls(u string, urls []string) []string {
	var realurl []string
	var path2 string
	u1, err1 := url.Parse(u)
	if err1 != nil {
		log.Fatal(err1)
	}
	path1 := strings.Split(u1.Path, "/")
	path := strings.SplitAfterN(u1.Path, "/", len(path1))
	for i := 0; i < len(path1)-1; i++ {
		path2 = path2 + path[i]
	}
	//fmt.Println(u1.Path)
	//fmt.Println(len(path1))
	//fmt.Println(path2)

	for _, link := range urls {
		u2, err2 := url.Parse(link)
		if err2 != nil {
			log.Fatal(err2)
		}

		if link != "" && u2.Scheme != "" {
			realurl = append(realurl, link)
		} else if u2.Host != "" {
			link = u1.Scheme + ":" + link
			realurl = append(realurl, link)
		} else if u2.Path != "" && u2.Path[0] == '/' {
			//fmt.Println(u2.Path)
			link = u1.Scheme + "://" + u1.Host + link
			realurl = append(realurl, link)
		} else if u2.Path != "" && u2.Path[0] != '/' {
			link = u1.Scheme + "://" + u1.Host + path2 + link
			realurl = append(realurl, link)
		} else {
			continue
		}

		//fmt.Println(realurl)
		//fmt.Println(i, link)
	}

	return realurl
}

func RemoveDuplicatesAndEmpty(a []string) (ret []string) {
	sort.Strings(a)
	a_len := len(a)
	for i := 0; i < a_len; i++ {
		if (i > 0 && a[i-1] == a[i]) || len(a[i]) == 0 {
			continue
		}
		ret = append(ret, a[i])
	}
	return ret
}

func fetch(url string) ([]string, error) {
	//url := "http://daily.zhihu.com/"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return nil, err
	}
	var link []string
	//var ok bool
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		inlink, ok := s.Attr("src")
		if ok {
			link = append(link, inlink)
		}
	})
	//fmt.Println("real:", link)
	link = RemoveDuplicatesAndEmpty(link)
	return link, nil
}

func main() {
	//url := "http://daily.zhihu.com/"
	//url := "http://59.110.12.72:7070/golang-spider/img.html"
	url := os.Args[1]
	urls, err := fetch(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("pre url:")
	for _, u := range urls {
		fmt.Println(u)
	}

	fmt.Println("clear url:")
	s := clearUrls(url, urls)
	for _, u := range s {
		fmt.Println(u)
	}
}
