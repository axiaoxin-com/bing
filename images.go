package bing

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"

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
func GetImageURL(num int, shuffle bool) ([]string, error) {
	hcli := &http.Client{
		Timeout: time.Second * 60 * 5,
	}
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
			if err := goutils.HTTPGET(context.Background(), hcli, apiurl, header, &data); err != nil {
				fmt.Println(err)
				return
			}
			mux.Lock()
			imgs = append(imgs, data.Images...)
			mux.Unlock()
		}(idx)
	}
	wg.Wait()

	baseURL := "https://www.bing.com"
	imgurls := []string{}
	for _, img := range imgs {
		fullURL := baseURL + img.URL
		imgurls = append(imgurls, fullURL)
	}

	if shuffle {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(imgurls), func(i, j int) {
			imgurls[i], imgurls[j] = imgurls[j], imgurls[i]
		})
	}
	return imgurls, nil
}