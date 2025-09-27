# Build Stage
FROM golang:1.25

WORKDIR /build

RUN go env -w GOCACHE=/go-cache

COPY go.mod go.sum* ./
RUN go mod download

COPY . .
RUN --mount=type=cache,target=go-cache CGO_ENABLED=0 GOOS=linux go build -v -o app .

CMD ["./app"]
