package pw

import (
	"fmt"
	"log"

	"github.com/MarcosIgnacioo/classmoodls/helpers/arraylist"
	"github.com/playwright-community/playwright-go"
)

var expect = playwright.NewPlaywrightAssertions(10000)

func ClassroomScrap(browser *playwright.Browser, username string, password string) []playwright.Locator {
	classroom, err := (*browser).NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
	fmt.Println("aqui")
	classroom.Goto("https://accounts.google.com/ServiceLogin?continue=https%3A%2F%2Fclassroom.google.com&passive=true")
	classroom.Locator("#identifierId").Fill(fmt.Sprintf("%v@alu.uabcs.mx", username))
	classroom.Locator("button").First().Click()
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
	fmt.Println("moodle:", subjects)

	for _, s := range subjects {
		fmt.Println(s.TextContent())
	}
	ms <- subjects
}

type ScrappedInfo struct {
	Moodle    []interface{}
	ClassRoom []interface{}
}

func NewScrappedInfo(md []interface{}, cr []interface{}) ScrappedInfo {
	return ScrappedInfo{
		Moodle:    md,
		ClassRoom: cr,
	}
}

func FuckAround(username string, password string) ScrappedInfo {
	// TODO: Crear un package con variables globales (Expect)
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	// playwright.BrowserTypeLaunchOptions{Headless: playwright.Bool(false)}
	//                                vv
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{Headless: playwright.Bool(true)})
	fmt.Println("wtf")

	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}
	fmt.Println("quwepedo")

	ms := make(chan []playwright.Locator)
	var result []playwright.Locator
	go MoodleScrap(&browser, username, password, ms)
	cs := ClassroomScrap(&browser, username, password)

	result = <-ms

	fmt.Println("desde go channel")

	moodleArray := arraylist.NewArrayList(10)
	classroomArray := arraylist.NewArrayList(10)

	for _, v := range result {
		hw, _ := v.TextContent()
		moodleArray.Push(hw)
	}

	for _, v := range cs {
		hw, _ := v.Locator("h2 > a > div").First().InnerText()
		classroomArray.Push(hw)
	}

	if err = browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}
	if err = pw.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}
	mArr := moodleArray.ArrayList[0:moodleArray.Length]
	cArr := classroomArray.ArrayList[0:classroomArray.Length]
	return NewScrappedInfo(mArr, cArr)
}
func Test() {
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	browser, err := pw.Chromium.Launch()
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}
	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
	if _, err = page.Goto("https://news.ycombinator.com"); err != nil {
		log.Fatalf("could not goto: %v", err)
	}
	entries, err := page.Locator(".athing").All()
	if err != nil {
		log.Fatalf("could not get entries: %v", err)
	}
	for i, entry := range entries {
		title, err := entry.Locator("td.title > span > a").TextContent()
		if err != nil {
			log.Fatalf("could not get text content: %v", err)
		}
		fmt.Printf("%d: %s\n", i+1, title)
	}
	if err = browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}
	if err = pw.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}
}
