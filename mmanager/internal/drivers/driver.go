package drivers

import "github.com/tebeka/selenium"

type ManwhaResource struct {
	id       int
	name     string
	address  string
	nChapter uint16
	imageUrl string
	pages    []*ManwhaPage
}

func (mr ManwhaResource) GetName() string {
	return mr.name
}

func (mr ManwhaResource) GetAddress() string {
	return mr.address
}
func (mr ManwhaResource) GetNChapters() uint16 {
	return mr.nChapter
}

func (mr ManwhaResource) GetImageUrl() string {
	return mr.imageUrl
}

func (mr ManwhaResource) GetPages() []*ManwhaPage {
	return mr.pages
}

func (mr *ManwhaResource) AddPage(page *ManwhaPage) {
	mr.pages = append(mr.pages, page)
}

type ManwhaPage struct {
	id         int
	pageNumber int
	ImageUrls  []string
}

type Manwha interface {
	getChapter(chapter uint32) ManwhaPage
}

type ManwhaDriver interface {
	GetDriverName() string
	GetBaseAddress() string
	ListComicsOnPage(page uint16) []*ManwhaResource
	GetManwhaPage(manwhaResource ManwhaResource, page uint16, seleniumDriver selenium.WebDriver) (*ManwhaPage, error)
}
