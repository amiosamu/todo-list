# Step 1: Modules caching
FROM golang:alpine as modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

# Step 2: Builder
FROM golang:alpine as builder
COPY --from=modules /go/pkg /go/pkg
COPY . /todo
WORKDIR /todo
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -tags migrate -o /bin/todo ./cmd/todo

# Step 3: Final
FROM scratch
COPY --from=builder /todo/config /config
COPY --from=builder /bin/todo /todo

CMD ["/todo"]