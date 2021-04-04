FROM golang:1.13-alpine AS builder

WORKDIR /go/src/github.com/homenitor/back/
COPY . .

ENV GO111MODULE=on

RUN GOARM=6 CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/back main.go


FROM arm32v6/alpine

COPY --from=builder /go/bin/back /go/bin/back

EXPOSE 3000
ENTRYPOINT [ "/go/bin/back"]
