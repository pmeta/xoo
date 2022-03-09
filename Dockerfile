FROM --platform=linux/amd64 golang:1.17.8 as builder
ARG buildDir=/Go/src/github.com/skrbox/xoo
WORKDIR ${buildDir}
ENV GOPATH=/Go
ENV GOPROXY=https://goproxy.cn,direct
COPY . .
RUN make pkg

FROM --platform=linux/amd64 alpine:3.15.0 as runner
ARG buildDir=/Go/src/github.com/skrbox/xoo
ARG pkgDir
WORKDIR /app
COPY --from=builder --chown=bin ${buildDir}/_output/${pkgDir}/xoo /app
EXPOSE 80
CMD ["/app/xoo", "--meta.listen-addr=:80"]
