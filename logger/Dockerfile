FROM golang:1.19

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
COPY producer.log data_logs:/logs

# RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.# sh | sh -s -- -b $(go env GOPATH)/bin

CMD ["go", "run", "/app/cmd/main/main.go"]