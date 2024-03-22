ARG web_app_port

FROM golang:latest
WORKDIR /app
COPY ./src /app
ENV port=$web_app_port
RUN go get github.com/redis/go-redis/v9
EXPOSE $web_app_port
CMD ["go", "run", "."]
