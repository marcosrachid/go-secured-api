FROM golang:1.14.3-alpine as builder

WORKDIR /go/src/github.com/marcosrachid/go-secured-api

COPY go.mod go.sum ./

RUN go mod download

COPY ./cmd/api .
COPY ./internal ./internal
COPY ./pkg ./pkg

RUN go get -d -v ./internal/...
RUN go install -v ./internal/...
RUN go get -d -v ./pkg/...
RUN go install -v ./pkg/...

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api .


######## Start a new stage from scratch #######
FROM alpine:3.11.6  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /go/src/github.com/marcosrachid/go-secured-api/api .

EXPOSE 9090

# Command to run the executable
CMD ["./api"] 