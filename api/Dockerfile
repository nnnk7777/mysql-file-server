FROM golang:latest

WORKDIR /app

COPY . .
RUN go mod tidy
RUN go get -u github.com/cosmtrek/air

# CMD [ "go", "run", "main.go" ]
CMD ["air", "-c", ".air.toml"]