#defining env variables
TARGETNAME = snake

ifeq ($(OS), Windows_NT)
	TARGETNAME = $(TARGETNAME).exe
endif

packages = ./game \

.PHONY: all
all: format code-quality test build

.PHONY: format
format: 
	@$(foreach package,$(packages), \
			goimports -w $(package))
			
.PHONY: code-quality
code-quality: 
	@$(foreach package,$(packages), \
				golint $(package))
				
.PHONY: test
test: 
	@$(foreach package,$(packages), \
				go test $(package))

.PHONY: build
build: 
	go build -o $(TARGETNAME)
