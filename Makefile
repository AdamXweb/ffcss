build:
	go mod tidy
	go build

test:
	make mocks-setup > /dev/null
	-go test -race -coverprofile=coverage.txt -covermode=atomic -v
	go get -u github.com/jandelgado/gcov2lcov
	gcov2lcov -infile=coverage.txt -outfile=coverage/lcov.info
	make mocks-teardown > /dev/null

install:
	@cp -v themes/*.yaml ~/.config/ffcss/themes/
	@cp -v ffcss ~/.local/bin/ffcss

format:
	gofmt -s -w **.go

mocks-setup:
# this is a code smell! ↓
	make install > /dev/null
	mkdir -p mocks/{zip-dropoff,cache-directory,homedir/.mozilla/firefox/667ekipp.default-release} testarea

mocks-teardown:
	rm -rf mocks/{zip-dropoff,cache-directory,homedir} testarea
