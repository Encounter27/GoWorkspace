PROJECTS=onebigproject/...


all:
	cd src/ && go install $(PROJECTS)

test:
	cd src/ && go test -coverprofile=coverage.out $(PROJECTS)
	go tool cover -html=src/coverage.out

# generate XML report in Cobertura format
test.xml:
	cd src/ && gocov test $(PROJECTS) | gocov-xml > coverage.xml

clean:

clean-all: