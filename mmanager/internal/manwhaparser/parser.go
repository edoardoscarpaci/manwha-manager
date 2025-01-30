package manwhaparser

import (
	"errors"
	"fmt"
	"slices"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func hasEverything[T comparable](s1 []T, target []T) bool {

	if len(s1) == 0 || len(s1) > len(target) {
		return false
	}

	for _, val := range s1 {

		if !slices.Contains(target, val) {
			return false
		}
	}

	return true
}

func FindTag(startNode *html.Node, tag atom.Atom, attributes []html.Attribute) (*html.Node, error) {
	if startNode == nil {
		return nil, errors.New("no start node given")
	}

	for node := range startNode.Descendants() {

		if node.Type == html.ElementNode && node.DataAtom == tag && (len(attributes) == 0 || hasEverything(attributes, node.Attr)) {
			return node, nil
		}

	}

	return nil, fmt.Errorf("no node found with tag %s and requested attributes", string(tag))
}
