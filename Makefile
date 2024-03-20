default: clear build-full

clear:
	@echo '>> clear <<'
	@rm -rf ./build

ensure-gox:
	@echo '>> ensure-gox <<'
	@go install github.com/mitchellh/gox@latest

build: ensure-gox
	@echo '>> build <<'
	@gox -osarch="linux/amd64 linux/arm64 darwin/amd64 darwin/arm64" -output "./build/{{.OS}}_{{.Arch}}/{{.Dir}}"

rebuild: ensure-gox
	@echo '>> rebuild <<'
	@gox -osarch="linux/amd64 linux/arm64 darwin/amd64 darwin/arm64" -output "./build/{{.OS}}_{{.Arch}}/{{.Dir}}" -rebuild

build-full: clear rebuild