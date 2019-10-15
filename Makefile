BUILD_DATE := `date +%Y-%m-%d\ %H:%M`
GIT_VERSION := `git --no-pager describe --tags --dirty --always`
VERSIONFILE := pkg/setting/version.go
PACK_DATE := `date +%Y%m%d%H%M`

app:
	env GO111MODULE=on CGO_ENABLED=1 go build

cross:
	env GO111MODULE=on CC=arm-mac-gnueabihf-gcc CXX=arm-mac-gnueabihf-g++ CGO_ENABLED=1 GOOS=linux GOARCH=arm go build -ldflags "$(LDFLAGS)" -o ./swem

version:
	rm -f $(VERSIONFILE)
	@echo "package setting" > $(VERSIONFILE)
	@echo "const (" >> $(VERSIONFILE)
	@echo "  GitVersion = \"$(GIT_VERSION)\"" >> $(VERSIONFILE)
	@echo "  AppBuildTime = \"$(BUILD_DATE)\"" >> $(VERSIONFILE)
	@echo "  AppBuilder = \"$(USER)\"" >> $(VERSIONFILE)
	@echo ")" >> $(VERSIONFILE)

clean:
	rm -rf demo demo.db
