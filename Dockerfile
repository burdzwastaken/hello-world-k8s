FROM golang:1.9

ADD . /go/src/github.com/mulesoft-ops/sample-v2-spinnaker-app
WORKDIR /go/src/github.com/mulesoft-ops/sample-v2-spinnaker-app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /sample-v2-spinnaker-app .

FROM alpine
COPY --from=0 /sample-v2-spinnaker-app /sample-v2-spinnaker-app
ADD ./content /content
ENTRYPOINT /sample-v2-spinnaker-app
