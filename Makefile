#defining env variables
TARGETNAME = snake

ifeq ($(OS), Windows_NT)
	TARGETNAME = $(TARGETNAME).exe
endif

packages = ./pkg/game \

.PHONY: all
all: format code-quality test build

.PHONY: format
format: @$(foreach package,$(packages)), \
			goimports -w $(package)
			
.PHONY: code-quality
code-quality: @$(foreach package,$(packages)), \
				golint $(package)
				
.PHONY: test
test: @$(foreach package,$(packages)), \
				go test $(package)

.PHONY: build
test: @$(foreach package,$(packages)), \
				go build ./pkg -o $(TARGETNAME)