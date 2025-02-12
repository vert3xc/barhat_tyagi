FROM golang:latest

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./backend

EXPOSE 8080
CMD ["/app/server"]