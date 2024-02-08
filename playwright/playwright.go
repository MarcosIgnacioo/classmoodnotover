package pw

import (
	"errors"
	"fmt"
	"log"

	"github.com/MarcosIgnacioo/classmoodls/helpers/arraylist"
	"github.com/playwright-community/playwright-go"
)

var expect = playwright.NewPlaywrightAssertions(10000)

func ClassroomScrap(browser *playwright.Browser, username string, password string, cs chan []playwright.Locator) {
	classroom, err := (*browser).NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
	classroom.Goto("https://accounts.google.com/ServiceLogin?continue=https%3A%2F%2Fclassroom.google.com&passive=true")
	classroom.Locator("#identifierId").Fill(fmt.Sprintf("%v@alu.uabcs.mx", username))
	classroom.Locator("button").First().Click()
	classroom.Locator("#username").Fill(username)
	classroom.Locator("#password").Fill(password)
	classroom.Locator("input").Nth(2).Click()
	expect.Locator(classroom.Locator("ol > li").Last()).ToBeVisible()
	classroomAssigments, _ := classroom.Locator("ol > li").All()
	cs <- classroomAssigments
}

func MoodleScrap(browser *playwright.Browser, username string, password string) ([]interface{}, error) {
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
	url := moodle.URL()
	if url != "https://enlinea2024-1.uabcs.mx/my/" {
		err := errors.New("Credenciales incorrectas")
		return nil, err
	}

	expect.Locator(moodle.Locator(".multiline")).ToBeVisible()
	tabContent, _ := moodle.Locator(".event-name-container").All()

	subjects := arraylist.NewArrayList(10)
	for _, v := range tabContent {
		classSubject, _ := v.Locator("small").InnerText()
		anchorTag := v.Locator("a").First()
		assigmentTitle, assError := anchorTag.InnerText()
		if assError != nil {
			assigmentTitle = "No hay titulo"
		}
		link, linkErr := anchorTag.GetAttribute("href")
		if linkErr != nil {
			link = "No hay link"
		}
		subjects.Push(NewAssigment(classSubject, assigmentTitle, link))
	}
	return subjects.GetArray(), nil
}

type Assigment struct {
	ClassSubject string
	Title        string
	Link         string
}

func NewAssigment(c string, t string, l string) Assigment {
	return Assigment{ClassSubject: c, Title: t, Link: l}
}

type ScrappedInfo struct {
	Moodle    []interface{}
	ClassRoom []interface{}
}

type LoginError struct {
	ErrorMessage string
}

func NewLoginError(m string) *LoginError {
	return &LoginError{ErrorMessage: m}
}

func NewScrappedInfo(md []interface{}, cr []interface{}) *ScrappedInfo {
	return &ScrappedInfo{
		Moodle:    md,
		ClassRoom: cr,
	}
}

func FuckAround(username string, password string) (*ScrappedInfo, *LoginError) {
	// TODO: Crear un package con variables globales (Expect)
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	// playwright.BrowserTypeLaunchOptions{Headless: playwright.Bool(false)}
	//                                vv
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{Headless: playwright.Bool(true)})

	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}

	cs := make(chan []playwright.Locator)

	go ClassroomScrap(&browser, username, password, cs)
	ms, logErr := MoodleScrap(&browser, username, password)

	if logErr != nil {
		if err = browser.Close(); err != nil {
			log.Fatalf("could not close browser: %v", err)
		}
		if err = pw.Stop(); err != nil {
			log.Fatalf("could not stop Playwright: %v", err)
		}
		return nil, NewLoginError(logErr.Error())
	}

	moodleArray := arraylist.NewArrayList(10)
	classroomArray := arraylist.NewArrayList(10)
	prueba := <-cs
	for _, v := range prueba {
		hw, _ := v.Locator("h2 > a > div").First().InnerText()
		classroomArray.Push(hw)
	}
	for _, v := range ms {
		moodleArray.Push(v)
	}
	if err = browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}
	if err = pw.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}
	mArr := moodleArray.GetArray()
	cArr := classroomArray.GetArray()
	fmt.Println(mArr...)
	fmt.Println(cArr...)
	return NewScrappedInfo(mArr, cArr), nil
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
