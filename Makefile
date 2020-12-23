APP_NAME := demo
APP_VERSION := 0.0.1
BUILD_DATE := `date +%Y-%m-%d\ %H:%M`
GIT_VERSION := `git --no-pager describe --tags --dirty --always`
VERSIONFILE := pkg/setting/version.go
PACK_DATE := `date +%Y%m%d%H%M`

app: version
	env GO111MODULE=on CGO_ENABLED=0 go build

cross: version
	env GO111MODULE=on CC=arm-mac-linux-gnueabihf-gcc CXX=arm-mac-linux-gnueabihf-g++ CGO_ENABLED=1 GOOS=linux GOARCH=arm go build -ldflags "$(LDFLAGS)" -o ./demo

test:
	mockery -dir=internal/po -name=PO -output=internal/po/mocks
	go test -v ./internal/serv
#	follow command will run all subdirectory test even those do not contain ant test, but output ? xxxxx [no test files]
#	go test -v ./...

version:
	rm -rf $(VERSIONFILE)
	@echo "package setting\\n" > $(VERSIONFILE)
	@echo "const (" >> $(VERSIONFILE)
	@echo "	AppName = \"$(APP_NAME)\"" >> $(VERSIONFILE)
	@echo "	AppVersion = \"$(APP_VERSION)\"" >> $(VERSIONFILE)
	@echo "	AppVer = \"v$(APP_VERSION) $(GIT_VERSION). Built by $(USER) at $(BUILD_DATE)\"" >> $(VERSIONFILE)
	@echo "	GitVersion = \"$(GIT_VERSION)\"" >> $(VERSIONFILE)
	@echo "	AppBuildTime = \"$(BUILD_DATE)\"" >> $(VERSIONFILE)
	@echo "	AppBuilder = \"$(USER)\"" >> $(VERSIONFILE)
	@echo ")" >> $(VERSIONFILE)

clean:
	rm -rf demo demo.db
	rm -rf internal/po/mocks
