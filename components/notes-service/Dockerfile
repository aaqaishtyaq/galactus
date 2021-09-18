FROM golang:1.16-alpine as builder
RUN apk add --no-cache ca-certificates git
RUN apk add build-base
WORKDIR /go/src/github.com/aaqaishtyaq/galactus/src/notes-service
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN GOOS=linux go build -o app .

FROM alpine as release
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /go/src/github.com/aaqaishtyaq/galactus/src/notes-service/app .
ENV NOTES_DIR=/notes
RUN : \
    && mkdir -p logs \
    && adduser -S -D -H -h ./logs user \
    && :
USER user
CMD [ "./app" ]
