INSTDIR?=/usr/local
BINDIR?=$(INSTDIR)/bin
GO?=go

GOSRC!=find . -name '*.go'
GOSRC+=go.mod go.sum

hugo-encrypt: $(GOSRC)
	$(GO) build \
		-o $@

# Exists in GNUMake but not in NetBSD make and others.
RM?=rm -f

clean:
	$(RM) hugo-encrypt

install: hugo-encrypt
	mkdir -p $(BINDIR)
	install -m755 $< $(BINDIR)/$<

