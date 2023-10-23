TW=./tools/tailwindcss
TW_IN=./web/static/style/input.css
TW_OUT=./web/static/style/tailwind.css
TW_CONTENT=./{web,cmd}/**/*.{html,js}
TW_CONFIG=./tailwind.config.js

AIR=./tools/air

build: gohtmx
.PHONY: build

run: build
	./gohtmx
.PHONY: build

gohtmx: $(TW_OUT)
	go build ./cmd/gohtmx
.PHONY: gohtmx

dev: $(AIR)
	$(AIR)

# tailwindcss standalone cli tool
$(TW):
	@echo "downloading tailwindcss cli..."
	mkdir -p $(@D)
	curl -Lo $@ https://github.com/tailwindlabs/tailwindcss/releases/download/v3.3.3/tailwindcss-linux-x64
	chmod +x $@

$(TW_CONTENT):

$(TW_OUT): $(TW) $(TW_CONTENT)
	$(TW) -c $(TW_CONFIG) -m -i $(TW_IN) -o $@
.PHONY: $(TW_OUT)

# air
$(AIR):
	@echo "downloading air..."
	mkdir -p $(@D)
	curl -SfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(@D)
