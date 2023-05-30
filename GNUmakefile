all:

all: gen/icon-orig-facebook.b64
all: gen/icon-orig-telegram.b64
all: gen/icon-orig-whatsapp.b64
all: gen/icon-orig-instagram.b64
all: gen/icon-orig-twitter.b64

gen/%.b64: img/%.png
	@mkdir -p $(dir $@)
	base64 $< > $@
## -- AUTO-GO --
GO_PROGRAMS += bin/sharewith$(EXE) 
.PHONY all-go: $(GO_PROGRAMS)
all:     all-go
install: install-go
clean:   clean-go
deps:
bin/sharewith$(EXE): deps 
	go build -o $@ $(SHAREWITH_FLAGS) $(GO_CONF) ./cmd/sharewith
install-go:
	install -d $(DESTDIR)$(PREFIX)/bin
	cp bin/sharewith$(EXE) $(DESTDIR)$(PREFIX)/bin
clean-go:
	rm -f $(GO_PROGRAMS)
## -- AUTO-GO --
## -- AUTO-SERVICE --

## -- AUTO-SERVICE --
