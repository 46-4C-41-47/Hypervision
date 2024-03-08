FROM golang:latest
WORKDIR /app
COPY ./src /app
RUN go get github.com/redis/go-redis/v9
EXPOSE 3000
CMD ["go", "run", "."]