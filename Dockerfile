FROM golang:1.14-alpine
COPY . /go/src/go-grafana
WORKDIR /go/src/go-grafana
RUN go mod vendor
RUN cd cmd/ && go build -o go-grafana-app .
CMD ["./cmd/go-grafana-app"]