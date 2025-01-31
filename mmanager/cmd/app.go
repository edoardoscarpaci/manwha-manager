package main

import (
	"fmt"
	"log"
	"mmanager/internal/drivers"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
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

	service, err := selenium.NewChromeDriverService("/home/edoardo/progetti/manwha-manager/mmanager/cmd/chromedriver", 4444)

	if err != nil {
		log.Fatal("Error:", err)
	}

	defer service.Stop()

	// configure the browser options
	caps := selenium.Capabilities{}
	caps.AddChrome(chrome.Capabilities{Args: []string{
		"--headless", // comment out this line for testing
		//"window-size=1920x1080",
		//"--no-sandbox",
		"--disable-dev-shm-usage",
		"disable-gpu",
	}})

	// create a new remote client with the specified options
	seleniumDriver, err := selenium.NewRemote(caps, "")
	if err != nil {
		log.Fatal("Error:", err)
	}

	fmt.Println(resources[0].GetAddress())
	page, err := driver.GetManwhaPage(resources[0], 0, seleniumDriver)

	if err != nil {
		panic(err)
	}

	for _, imageUrl := range page.ImageUrls {
		fmt.Println(imageUrl)
	}
}
