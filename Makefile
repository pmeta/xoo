
buildTime:=$(shell date "+%Y-%m-%d %H:%M:%S")
commitId:=$(shell git rev-parse --short HEAD)
branch:=$(shell git symbolic-ref --short -q HEAD)
version:=latest
imageName:=jeyrce/xoo:${version}
pkgDir:=xoo-${version}

.phony: all
all: image pkg
	@echo "make all"

.phony: binary
binary: init
	export GOOS=linux
	export GOPROXY=https://goproxy.cn,direct
	go build -ldflags " \
		-X 'main.commitId=${commitId}' \
		-X 'main.branch=${branch}' \
		-X 'main.version=${version}' \
		-X 'main.buildAt=${buildTime}' \
	" \
	-o _output/xoo xoo.go

.phony: image
image:
	docker build -t ${imageName} --build-arg pkgDir=${pkgDir} .
	docker push ${imageName}

.phony: pkg
pkg: binary
	mkdir -p "_output/${pkgDir}"
	cp _output/xoo \
		README.md \
	  	LICENSE \
	_output/${pkgDir} ; \
 	tar -cvf _output/${pkgDir}.tar.gz _output/${pkgDir}

.phony: init
init:
	rm -rf _output/**
