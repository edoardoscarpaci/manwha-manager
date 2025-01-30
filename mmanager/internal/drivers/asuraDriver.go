package drivers

import (
	"errors"
	"fmt"
	"mmanager/internal/manwhaparser"
	"mmanager/internal/requests"
	"net/http"
	"slices"
	"strconv"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

var neededAttributes = [...]html.Attribute{{Namespace: "", Key: "class", Val: "grid grid-cols-2 sm:grid-cols-2 md:grid-cols-5 gap-3 p-4"}}
var chapterInformationAttr = [...]html.Attribute{{Namespace: "", Key: "class", Val: "block w-[100%] h-auto  items-center"}}

var nameAttr html.Attribute = html.Attribute{Namespace: "", Key: "class", Val: "block text-[13.3px] font-bold"}
var chapterAttr html.Attribute = html.Attribute{Namespace: "", Key: "class", Val: "text-[13px] text-[#999]"}
var ratingAttr html.Attribute = html.Attribute{Namespace: "", Key: "class", Val: "flex text-[12px] text-[#999]"}

const nodeTag atom.Atom = atom.Div
const baseAddress string = "https://asuracomic.net/"

type AsuraDriver struct {
}

func (ad AsuraDriver) GetBaseAddress() string {
	return baseAddress
}

func (ad AsuraDriver) ListComicsOnPage(page uint16) ([]*ManwhaResource, error) {
	var baseAddress string = ad.GetBaseAddress()
	var compiledAddress string = fmt.Sprintf("%sseries?page=%d&genres=&status=-1&types=-1&order=desc", baseAddress, page)

	resp := make(chan *http.Response)
	go requests.GetRequest(compiledAddress, resp)

	response := <-resp
	node, err := html.Parse(response.Body)

	if err != nil {
		panic(err)
	}

	foundNode, err := manwhaparser.FindTag(node, nodeTag, neededAttributes[:])

	if err != nil {
		panic(err)
	}

	if foundNode == nil {
		return nil, errors.New("no comic found in the page")

	}

	manwhaResources, err := getManwhaResourcesFromNode(foundNode)
	if err != nil {
		panic(err)

	}

	return manwhaResources, nil
}

func getManwhaResourcesFromNode(startNode *html.Node) ([]*ManwhaResource, error) {
	var manwhaResources []*ManwhaResource = make([]*ManwhaResource, 0)
	for node := range startNode.ChildNodes() {
		manwhaResource, err := parseNodeForManwhaResource(node)
		if err != nil {
			fmt.Println("Error: " + err.Error())
			continue
		}
		manwhaResources = append(manwhaResources, manwhaResource)
	}
	return manwhaResources, nil
}

func parseNodeForManwhaResource(node *html.Node) (*ManwhaResource, error) {
	if node.DataAtom != atom.A {
		return nil, errors.New("not the right node type")
	}

	var manwhaResource *ManwhaResource = new(ManwhaResource)
	hrefIndex := slices.IndexFunc(node.Attr, func(attribute html.Attribute) bool { return attribute.Key == "href" })

	if hrefIndex == -1 {
		return nil, errors.New("node doesn't have href attribute ")
	}

	nodeHref := node.Attr[hrefIndex]
	manwhaResource.address = baseAddress + nodeHref.Val
	manwhaInformationNode, err := manwhaparser.FindTag(node, atom.Div, chapterInformationAttr[:])
	if err != nil {
		panic(err)
	}
	for childNode := range manwhaInformationNode.ChildNodes() {
		for _, attr := range childNode.Attr {
			if attr == nameAttr {
				manwhaResource.name = childNode.FirstChild.Data
				break
			} else if attr == chapterAttr {
				chapterNumber, err := strconv.Atoi(childNode.LastChild.Data)
				if err != nil {
					panic(err)
				}

				manwhaResource.nChapter = uint16(chapterNumber)
				break
			}
		}
	}

	imgNode, err := manwhaparser.FindTag(node, atom.Img, nil)
	if err != nil {
		return manwhaResource, errors.New("no node img found")
	}

	srcIndex := slices.IndexFunc(imgNode.Attr, func(attribute html.Attribute) bool { return attribute.Key == "src" })
	if srcIndex == -1 {
		fmt.Println(imgNode.Attr)
		return nil, errors.New("node img doesn't have src attribute ")
	}

	manwhaResource.imageUrl = imgNode.Attr[srcIndex].Val

	return manwhaResource, nil
}
