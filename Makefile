.PHONY: all \
        clean \
        help \
        lint \

PROJECT = readlink-proxy
VERSION = 0.1.0

all: $(PROJECT)

$(PROJECT): clean
	go build .

help:
	@echo "clean  remove artifacts"
	@echo "dist   cross-compile binaries for distribution"
	@echo "help   show this page"
	@echo "lint   check style with golint"

clean:
	go clean -x
	$(RM) -r dist/

dist:
	gox --osarch="linux/amd64 darwin/amd64 windows/amd64" --output "dist/$(PROJECT)-$(VERSION)-{{ .OS }}_{{ .Arch }}"

lint:
	golint ./...
