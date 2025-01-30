package main

import (
	"fmt"
	"mmanager/internal/drivers"
)

func main() {
	var driver drivers.AsuraDriver
	fmt.Println(driver.GetBaseAddress())

	resources, err := driver.ListComicsOnPage(1)
	if err != nil {
		panic(err)
	}

	for _, resource := range resources {
		fmt.Println(resource.GetName())
		fmt.Println(resource.GetNChapters())
		fmt.Println(resource.GetAddress())
		fmt.Println(resource.GetImageUrl())

	}
}
