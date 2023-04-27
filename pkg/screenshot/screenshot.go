package screenshot

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

type Response struct {
	Title string `json:"title"`
	Data  []byte `json:"data"`
}

func MakeScreenshot(url string) (*Response, error) {
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	page := browser.MustPage(url).MustWaitLoad()
	img, err := page.Screenshot(true, &proto.PageCaptureScreenshot{})
	if err != nil {
		return nil, err
	}

	title := page.MustElement("title").MustText()
	return &Response{
		Title: title,
		Data:  img,
	}, nil
}
