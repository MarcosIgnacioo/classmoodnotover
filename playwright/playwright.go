package pw

import (
	"fmt"
	"github.com/playwright-community/playwright-go"
	"log"
)

func FuckAround() {
	// TODO: Separar classroom y mooodle por funciones diferentes
	// TODO: Crear un package con variables globales (Expect)
	// TODO: Meter ambas cosas en go routines
	// TODO: Se podra hacer en HTMX?
	username := "marcosignc_21"
	password := "sopitasprecio"
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
	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}

	classroom, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}

	classroom.Goto("https://accounts.google.com/ServiceLogin?continue=https%3A%2F%2Fclassroom.google.com&passive=true")
	classroom.Locator("#identifierId").Fill(fmt.Sprintf("%v@alu.uabcs.mx", username))
	classroom.Locator("button").Nth(2).Click()
	classroom.Locator("#username").Fill(username)
	classroom.Locator("#password").Fill(password)
	classroom.Locator("input").Nth(2).Click()

	expect := playwright.NewPlaywrightAssertions(10000)
	expect.Locator(classroom.Locator("ol > li").Last()).ToBeVisible()

	classroomAssigments, _ := classroom.Locator("ol > li").All()
	for _, cA := range classroomAssigments {
		fmt.Println(cA.Locator("h2 > a > div").First().InnerText())
	}

	if _, err = page.Goto("https://enlinea2023-2.uabcs.mx/login/"); err != nil {
		log.Fatalf("could not goto: %v", err)
	}
	page.Locator("#username").Fill(username)
	page.Locator("#password").Fill(password)
	page.Locator("#loginbtn").Click()

	expect.Locator(page.Locator(".multiline")).ToBeVisible()

	subjects, _ := page.Locator(".multiline").All()

	for _, s := range subjects {
		fmt.Println(s.TextContent())
	}
	// page.Screenshot(playwright.PageScreenshotOptions{Path: playwright.String(fmt.Sprintf("%vpopo.png", rand.Intn(150)))})

	// entries, err := page.Locator(".athing").All()
	// if err != nil {
	// 	log.Fatalf("could not get entries: %v", err)
	// }
	// for i, entry := range entries {
	// 	title, err := entry.Locator("td.title > span > a").TextContent()
	// 	if err != nil {
	// 		log.Fatalf("could not get text content: %v", err)
	// 	}
	// 	fmt.Printf("%d: %s\n", i+1, title)
	// }
	// if err = browser.Close(); err != nil {
	// 	log.Fatalf("could not close browser: %v", err)
	// }
	// if err = pw.Stop(); err != nil {
	// 	log.Fatalf("could not stop Playwright: %v", err)
	// }
}
