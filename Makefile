
FILES = $(shell find . -type f -name '*.go' -not -path './vendor/*')

gofmt:
	@gofmt -w $(FILES)
	@gofmt -r '&α{} -> new(α)' -w $(FILES)

deps:
	go get -u github.com/mgechev/revive

	go get -u github.com/altipla-consulting/content
	go get -u github.com/altipla-consulting/dateformatter
	go get -u github.com/altipla-consulting/datetime
	go get -u github.com/altipla-consulting/langs
	go get -u github.com/altipla-consulting/messageformat
	go get -u github.com/altipla-consulting/money
	go get -u github.com/ernestoalejo/aeimagesflags

test:
	revive -formatter friendly ./...
	go install .
