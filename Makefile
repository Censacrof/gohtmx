TW=./tools/tailwindcss
TW_IN=web/static/input.css
TW_OUT=web/static/tailwind.css
TW_CONTENT=web/**/*.{html,js}

build: gohtmx
.PHONY: build

run: build
	./gohtmx
.PHONY: build

$(TW):
	@echo "downloading tailwindcss cli..."
	mkdir -p $(@D)
	curl -Lo $@ https://github.com/tailwindlabs/tailwindcss/releases/download/v3.3.3/tailwindcss-linux-x64
	chmod +x $@

$(TW_CONTENT):

$(TW_OUT): $(TW_CONTENT)
	$(TW) -m --content $(TW_CONTENT) -i $(TW_IN) -o $@
.PHONY: $(TW_OUT)

gohtmx: $(TW_OUT)
	go build ./cmd/gohtmx
.PHONY: gohtmx


