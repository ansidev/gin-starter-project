FROM golang:alpine

LABEL maintainer="ansidev@ansidev.xyz"

RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

RUN adduser \
--disabled-password \
--gecos "" \
--home "/nonexistent" \
--shell "/sbin/nologin" \
--no-create-home \
--uid 1000 \
ansidev

WORKDIR /app

ENV GOPRIVATE=gitlab.com/ansidev/*

ONBUILD ARG APP_NAME
ONBUILD ARG DOCKER_NETRC

ONBUILD RUN echo "${DOCKER_NETRC}" > ~/.netrc

ONBUILD COPY ./app .

ONBUILD RUN go mod download && go mod verify
ONBUILD RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
-ldflags='-w -s -extldflags "-static"' -a \
-o /app/dist/${APP_NAME} .
