FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /api-go-std ./cmd/api/main.go

EXPOSE 8081

# Run
CMD ["/api-go-std"]