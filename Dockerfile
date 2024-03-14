FROM golang:alpine AS build


COPY . /usr/src/yt-dl-rpc
WORKDIR /usr/src/yt-dl-rpc

RUN CGO_ENABLED=0 GOOS=linux go build -o yt-dl-rpc

FROM alpine:edge

VOLUME /downloads /config

WORKDIR /app

RUN apk update && \
    apk add psmisc ffmpeg yt-dlp --no-cache

COPY --from=build /usr/src/yt-dl-rpc/yt-dl-rpc /app

EXPOSE 3033
ENTRYPOINT [ "./yt-dl-rpc" , "--out", "/downloads", "--conf", "/config/config.yml"]
