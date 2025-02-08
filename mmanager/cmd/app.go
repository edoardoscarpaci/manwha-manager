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
	databaseDriver := drivers.SqlLiteDatabase{DatabasePath: "/home/edoardo/progetti/manwha-manager/mmanager/db.sqlite"}
	err := databaseDriver.InitDatabase(true)
	if err != nil {
		panic(err)
	}

	fmt.Println(driver.GetBaseAddress())

	resources, err := driver.ListComicsOnPage(0)
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
	customUserAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36 Edg/91.0.864.59"
	// configure the browser options
	caps := selenium.Capabilities{}
	caps.AddChrome(chrome.Capabilities{Args: []string{
		"--headless", // comment out this line for testing
		//"window-size=1920x1080",
		//"--no-sandbox",
		"--disable-dev-shm-usage",
		"--user-agent=" + customUserAgent,
		//"--proxy-server=44.219.175.186",
		"disable-gpu",
	}})

	// create a new remote client with the specified options
	seleniumDriver, err := selenium.NewRemote(caps, "")
	if err != nil {
		log.Fatal("Error:", err)
	}
	for _, resource := range resources {
		firstResource := resource
		page, err := driver.GetManwhaPage(firstResource, 1, seleniumDriver)
		if err != nil {
			panic(err)
		}

		firstResource.AddPage(page)

		for _, imageUrl := range page.ImageUrls {
			fmt.Println(imageUrl)
		}

		err = databaseDriver.AddManwhaResource(firstResource, driver.GetDriverName(), true)
		if err != nil {
			panic(err)
		}

	}

}
