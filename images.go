package bing

import (
	"context"
	"fmt"
	"math"
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

// ImageResolution 图片分辨率
type ImageResolution string

const (
	ImageResolution240x320   ImageResolution = "240x320"
	ImageResolution320x240   ImageResolution = "320x240"
	ImageResolution240x400   ImageResolution = "240x400"
	ImageResolution400x240   ImageResolution = "400x240"
	ImageResolution480x640   ImageResolution = "480x640"
	ImageResolution640x480   ImageResolution = "640x480"
	ImageResolution480x800   ImageResolution = "480x800"
	ImageResolution800x480   ImageResolution = "800x480"
	ImageResolution600x800   ImageResolution = "600x800"
	ImageResolution800x600   ImageResolution = "800x600"
	ImageResolution720x1280  ImageResolution = "720x1280"
	ImageResolution1280x720  ImageResolution = "1280x720"
	ImageResolution768x1024  ImageResolution = "768x1024"
	ImageResolution1024x768  ImageResolution = "1024x768"
	ImageResolution768x1280  ImageResolution = "768x1280"
	ImageResolution1280x768  ImageResolution = "1280x768"
	ImageResolution768x1366  ImageResolution = "768x1366"
	ImageResolution1366x768  ImageResolution = "1366x768"
	ImageResolution1080x1920 ImageResolution = "1080x1920"
	ImageResolution1920x1080 ImageResolution = "1920x1080"
	ImageResolution1200x1920 ImageResolution = "1200x1920"
	ImageResolution1920x1200 ImageResolution = "1920x1200"
	ImageResolutionUHD       ImageResolution = "UHD"
)

// GetImageURL 返回bing壁纸图片地址
func GetImageURL(ctx context.Context, num int, shuffle bool, resolution ImageResolution) ([]string, error) {
	hcli := &http.Client{}
	n := 8
	pagecount := int(math.Floor(float64(num)/float64(n) + 0.5))
	if pagecount < 1 {
		pagecount = 1
	}
	imgs := []Image{}
	wg := sync.WaitGroup{}
	mux := sync.Mutex{}
	for idx := 0; idx < pagecount; idx++ {
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

	imgurls := []string{}
	for _, img := range imgs {
		fullURL := fmt.Sprintf("https://www.bing.com/%s_%v.jpg", img.Urlbase, resolution)
		imgurls = append(imgurls, fullURL)
	}

	if len(imgurls) > num {
		return imgurls[:num], nil
	}

	return imgurls, nil
}
