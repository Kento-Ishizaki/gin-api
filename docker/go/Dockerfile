FROM golang:1.16-alpine

WORKDIR /go/src

COPY ./src .

RUN apk update && \
  apk add git && \
  go get -u github.com/cosmtrek/air && \
  go build -o go/bin/air github.com/cosmtrek/air && \
  go get github.com/gin-gonic/gin && \
  go get github.com/jinzhu/gorm && \
  go get github.com/go-sql-driver/mysql

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]
