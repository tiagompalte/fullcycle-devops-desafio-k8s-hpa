FROM golang:alpine as builder

WORKDIR $GOPATH/src/app

COPY src/fullcycle-devops-desafio-k8s-hpa/ .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/app .

FROM scratch

COPY --from=builder /go/bin/app /go/bin/app

ENTRYPOINT ["/go/bin/app"]

EXPOSE 8000
