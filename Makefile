BIN=bin/blockchain_demo
WalletFile=tmp/wallets.dat

test:
	@go test -v ./...


push:
	@git add -A && git commit -m "update" && git push origin master


build:
	@rm -rf ./${BIN}
	@rm -rf ./tmp && mkdir -p ./tmp
# @touch ./${WalletFile}
	@go build -ldflags='-s -w' -o ${BIN} ./main.go
	@echo 'done'


CMD ?=
run:
	@chmod +x ./${BIN}
	@./${BIN} ${CMD}


# make tag t=<your_version>
tag:
	@echo '${t}'
	@git tag -a ${t} -m "${t}" && git push origin ${t}

dtag:
	@echo 'delete ${t}'
	@git push --delete origin ${t} && git tag -d ${t}

.PHONY: test push build tag dtag run
