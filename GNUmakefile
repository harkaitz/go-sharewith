.POSIX:
.SUFFIXES:
.PHONY: all clean install check
all:
PROJECT=go-sharewith
VERSION=1.0.0
PREFIX=/usr/local

##
build/sharewith$(EXE): gen/icon-orig-facebook.b64
build/sharewith$(EXE): gen/icon-orig-telegram.b64
build/sharewith$(EXE): gen/icon-orig-whatsapp.b64
build/sharewith$(EXE): gen/icon-orig-instagram.b64
build/sharewith$(EXE): gen/icon-orig-twitter.b64
gen/%.b64: img/%.png
	@mkdir -p $(dir $@)
	base64 $< > $@
## -- BLOCK:go --
build/sharewith$(EXE):
	mkdir -p build
	go build -o $@ $(GO_CONF) ./cmd/sharewith
all: build/sharewith$(EXE)
install: all
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	cp build/sharewith$(EXE) $(DESTDIR)$(PREFIX)/bin
clean:
	rm -f build/sharewith$(EXE)
## -- BLOCK:go --
