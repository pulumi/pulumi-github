GOPKGS    = $(shell go list ./provider/... | grep -v /vendor/)

.PHONY: default
default: banner lint_quiet vet build test install

.PHONY: banner
banner:
	@echo "\033[1;37m===================\033[0m"
	@echo "\033[1;37mLumi GitHub Package\033[0m"
	@echo "\033[1;37m===================\033[0m"

.PHONY: gen
gen:
	@echo "\033[0;32mGEN:\033[0m"
	lumidl \
	    github idl/ \
		--recursive \
		--out-pack=pack/ \
		--out-rpc=rpc/

.PHONY: clean
clean:
	rm -rf ./pack/bin

.PHONY: build
build:
	@echo "\033[0;32mBUILD:\033[0m"
	cd pack/ && yarn link @pulumi/pulumi-fabric @pulumi/pulumi # link dependencies.
	cd pack/ && yarn run build
	go version
	cd provider/ && go build -i -o ../pack/bin/pulumi-provider-github # compile the resource provider

.PHONY: install
install:
	@echo "\033[0;32mINSTALL:\033[0m [${LUMILIB}]"
	cp pack/package.json pack/bin/package.json
	cd pack/bin/ && yarn link  # ensure NPM references resolve locally.

.PHONY: lint
lint:
	@echo "\033[0;32mLINT:\033[0m"
	golint -set_exit_status provider/...

.PHONY: lint_quiet
lint_quiet:
	@echo "\033[0;32mLINT (quiet):\033[0m"
	@echo "`golint provider/... | grep -v "or be unexported"`"
	@test -z "$$(golint provider/... | grep -v 'or be unexported')"
	@echo "\033[0;33mgolint was run quietly; to run with noisy errors, run 'make lint'\033[0m"

.PHONY: vet
vet:
	@echo "\033[0;32mVET:\033[0m"
	go tool vet -printf=false provider/

.PHONY: test
test:
	@echo "\033[0;32mTEST:\033[0m"
	go test -cover ${GOPKGS}

