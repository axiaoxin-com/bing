package bing

import (
	"context"
	"fmt"
	"net/http"
	"sort"
	"sync"

	"github.com/axiaoxin-com/goutils"
	"github.com/corpix/uarand"
)

// Image json images item
type Image struct {
	Startdate     string        `json:"startdate"`
	Fullstartdate string        `json:"fullstartdate"`
	Enddate       string        `json:"enddate"`
	URL           string        `json:"url"`
	Urlbase       string        `json:"urlbase"`
	Copyright     string        `json:"copyright"`
	Copyrightlink string        `json:"copyrightlink"`
	Title         string        `json:"title"`
	Quiz          string        `json:"quiz"`
	Wp            bool          `json:"wp"`
	Hsh           string        `json:"hsh"`
	Drk           int           `json:"drk"`
	Top           int           `json:"top"`
	Bot           int           `json:"bot"`
	Hs            []interface{} `json:"hs"`
}

// HPImageArchiveData HPImageArchive.aspx 接口返回结构
type HPImageArchiveData struct {
	Images   []Image `json:"images"`
	Tooltips struct {
		Loading  string `json:"loading"`
		Previous string `json:"previous"`
		Next     string `json:"next"`
		Walle    string `json:"walle"`
		Walls    string `json:"walls"`
	} `json:"tooltips"`
}

// GetImageURL 返回bing壁纸图片地址
func GetImageURL(ctx context.Context, num int, shuffle bool) ([]string, error) {
	hcli := &http.Client{}
	n := 8
	pagecount := num / n
	if pagecount < 1 {
		pagecount = 1
	}
	imgs := []Image{}
	wg := sync.WaitGroup{}
	mux := sync.Mutex{}
	for idx := 1; idx <= pagecount; idx++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			apiurl := fmt.Sprintf("https://www.bing.com/HPImageArchive.aspx?format=js&idx=%d&n=%d", idx, n)
			header := map[string]string{
				"user-agent": uarand.GetRandom(),
			}
			data := HPImageArchiveData{}
			if err := goutils.HTTPGET(ctx, hcli, apiurl, header, &data); err != nil {
				fmt.Println(err)
				return
			}
			mux.Lock()
			imgs = append(imgs, data.Images...)
			mux.Unlock()
		}(idx)
	}
	wg.Wait()

	if !shuffle {
		sort.Slice(imgs, func(i, j int) bool {
			if imgs[i].Startdate > imgs[j].Startdate {
				return true
			}
			return false
		})
	}

	baseURL := "https://www.bing.com"
	imgurls := []string{}
	for _, img := range imgs {
		fullURL := baseURL + img.URL
		imgurls = append(imgurls, fullURL)
	}

	if len(imgurls) > num {
		return imgurls[:num], nil
	}

	return imgurls, nil
}
