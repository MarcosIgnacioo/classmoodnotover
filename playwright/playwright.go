package pw

import (
	"fmt"
	"github.com/playwright-community/playwright-go"
	"log"
)

var expect = playwright.NewPlaywrightAssertions(10000)

func ClassroomScrap(browser *playwright.Browser, username string, password string) []playwright.Locator {
	classroom, err := (*browser).NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}

	classroom.Goto("https://accounts.google.com/ServiceLogin?continue=https%3A%2F%2Fclassroom.google.com&passive=true")
	classroom.Locator("#identifierId").Fill(fmt.Sprintf("%v@alu.uabcs.mx", username))
	classroom.Locator("button").Nth(2).Click()
	classroom.Locator("#username").Fill(username)
	classroom.Locator("#password").Fill(password)
	classroom.Locator("input").Nth(2).Click()

	expect.Locator(classroom.Locator("ol > li").Last()).ToBeVisible()

	classroomAssigments, _ := classroom.Locator("ol > li").All()
	return classroomAssigments
}

func MoodleScrap(browser *playwright.Browser, username string, password string, ms chan []playwright.Locator) {
	moodle, err := (*browser).NewPage()

	if err != nil {
		log.Fatalf("could not create moodle: %v", err)
	}

	if _, err = moodle.Goto("https://enlinea2024-1.uabcs.mx/login/"); err != nil {
		log.Fatalf("could not goto: %v", err)
	}
	moodle.Locator("#username").Fill(username)
	moodle.Locator("#password").Fill(password)
	moodle.Locator("#loginbtn").Click()

	expect.Locator(moodle.Locator(".multiline")).ToBeVisible()

	subjects, _ := moodle.Locator(".multiline").All()

	for _, s := range subjects {
		fmt.Println(s.TextContent())
	}
	ms <- subjects
}

func FuckAround(username string, password string) {
	// TODO: Crear un package con variables globales (Expect)
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	// playwright.BrowserTypeLaunchOptions{Headless: playwright.Bool(false)}
	//                                vv
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{Headless: playwright.Bool(false)})

	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}

	ms := make(chan []playwright.Locator)
	var result []playwright.Locator
	go MoodleScrap(&browser, username, password, ms)
	cs := ClassroomScrap(&browser, username, password)

	result = <-ms

	fmt.Println("desde go channel")
	for _, v := range result {
		fmt.Println(v.TextContent())
	}

	for _, v := range cs {
		fmt.Println(v.Locator("h2 > a > div").First().InnerText())
	}

	if err = browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}
	if err = pw.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}
}
