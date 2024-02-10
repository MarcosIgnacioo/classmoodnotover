
#!/bin/bash



# Instalar Playwright para Go
go run github.com/playwright-community/playwright-go/cmd/playwright@latest install --with-deps

# O alternativamente, puedes utilizar go install
# go install github.com/playwright-community/playwright-go/cmd/playwright@latest

# Instalar Playwright
playwright install --with-deps
