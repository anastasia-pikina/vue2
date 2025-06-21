FROM golang:1.17-alpine
WORKDIR /app
COPY . .
RUN go mod init app && go mod tidy
RUN go build -o main .
CMD ["./main"]