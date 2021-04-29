package main

import "fmt"

func main() {
	osCheck()
	flags()
	if wantsLicense {
		fmt.Println(touchThisAndYourMomDiesInHerSleepTonight)
	}
	fmt.Println("The TurtleCoin Shop is open :)")
	checkDirs()
	go restAPI()
	select {}
}
