FROM golang:1.19

WORKDIR /app
COPY go.mod ../../producer
COPY go.sum ../../producer
RUN go mod download

COPY ../../producer .

# RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.# sh | sh -s -- -b $(go env GOPATH)/bin

CMD ["go", "run", "/app/cmd/main/main.go"]