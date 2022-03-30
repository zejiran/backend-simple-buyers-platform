FROM golang:1.17.8-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

#RUN go install

RUN go install github.com/cespare/reflex@latest

EXPOSE 3100

CMD reflex -g '*.go' go run main.go --start-service
