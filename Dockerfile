# syntax=docker/dockerfile:1

FROM golang:1.22.0-alpine
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY . .

# Build
RUN go build -o ./bin/goter-daily-notification-cronjob
 

CMD [ "./bin/goter-daily-notification-cronjob" ]