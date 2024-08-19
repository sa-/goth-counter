.PHONY: dev
dev:
	@tailwindcss -o assets/css/generated-tailwind.css
	@templ generate --watch --proxy="http://localhost:42069" --cmd="go run ."

.PHONY: build
build:
	@tailwindcss -o assets/css/generated-tailwind.css
	@templ generate
	@go build -o tmp/app

.PHONY: fmt
	@templ fmt .
	

.PHONY: clean
clean:
	@rm -rf tmp

.PHONY: deps
deps:
	@go install github.com/a-h/templ/cmd/templ@latest
	@go install github.com/cosmtrek/air@latest
