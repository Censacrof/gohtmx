TW=./tools/tailwindcss
TW_IN=web/static/input.css
TW_OUT=web/static/tailwind.css

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

$(TW_OUT):
	$(TW) -i $(TW_IN) -o $@

gohtmx: $(TW_OUT)
	go build ./cmd/gohtmx
.PHONY: gohtmx


