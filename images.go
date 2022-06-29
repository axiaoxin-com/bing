package bing

import (
	"context"
	"fmt"
	"net/http"

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

// FullURL 返回完整url
func (i *Image) FullURL(resolution ...ImageResolution) string {
	return GetImageFullURL(i.Urlbase, resolution...)
}

// GetImageFullURL 返回完整url
func GetImageFullURL(urlbase string, resolution ...ImageResolution) string {
	rslt := ImageResolution1920x1080
	if len(resolution) > 0 {
		rslt = resolution[0]
	}
	return fmt.Sprintf("https://www.bing.com%s_%v.jpg", urlbase, rslt)
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

// GetImages 返回HPImageArchive的Images
func GetImages(ctx context.Context) ([]Image, error) {
	hcli := &http.Client{}
	apiurl := "https://www.bing.com/HPImageArchive.aspx?format=js&idx=0&n=8"
	header := map[string]string{
		"user-agent": uarand.GetRandom(),
	}
	data := HPImageArchiveData{}
	if err := goutils.HTTPGET(ctx, hcli, apiurl, header, &data); err != nil {
		return nil, err
	}

	return data.Images, nil
}
