FROM golang:1.18-alpine AS base
WORKDIR /app

ENV GO111MODULE="on"
ENV GOOS="linux"
ENV CGO_ENABLED=0

# System dependencies
RUN apk update \
    && apk add --no-cache \
    ca-certificates \
    git \
    && update-ca-certificates

### Development with nodemon and debugger
FROM base AS dev
WORKDIR /app
COPY go.mod .
COPY go.sum .

# Dependencies
RUN go mod download \
    && go mod verify

# Install nodemon
RUN apk add --update nodejs npm \
    && apk add --update npm

# Update path
ENV PATH /app/node_modules/.bin:$PATH
RUN npm install nodemon -g --loglevel notice

EXPOSE 8080

CMD [ "nodemon", "--config", "nodemon.json" ]
