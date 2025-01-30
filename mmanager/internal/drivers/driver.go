package drivers

type ManwhaResource struct {
	name     string
	address  string
	nChapter uint16
	imageUrl string
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

type ManwhaPage struct {
	body []byte
}

type Manwha interface {
	getChapter(chapter uint32) ManwhaPage
}

type ManwhaDriver interface {
	GetBaseAddress() string
	ListComicsOnPage(page uint16) []*ManwhaResource
}
