# Building the binary of the App
FROM golang:latest AS build

# `boilerplate` should be replaced with your project name
WORKDIR /app

# Copy all the Code and stuff to compile everything
COPY . .

# Downloads all the dependencies in advance (could be left out, but it's more clear this way)
RUN go mod download

# Builds the application as a staticly linked one, to allow it to run on alpine
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .

# Moving the binary to the 'final Image' to make it smaller
FROM alpine:latest

RUN apk update && apk add --no-cache tzdata && rm -rf /var/cache/apk/*

# The time zone database needed by LoadLocation may not be present on all systems, especially 
# non-Unix systems. LoadLocation looks in the directory or uncompressed zip file named by the 
# ZONEINFO environment variable, if any, then looks in known installation locations on Unix 
# systems, and finally looks in $GOROOT/lib/time/zoneinfo.zip.
ADD https://github.com/golang/go/raw/master/lib/time/zoneinfo.zip /zoneinfo.zip

RUN chmod +r /zoneinfo.zip

ENV ZONEINFO /zoneinfo.zip

ENV TZ=Europe/London

WORKDIR /app

COPY --from=build /app .

# Set the PORT environment variable:
ENV PORT 8103

# Exposes port 8983 because our program listens on that port
EXPOSE 8103

ENTRYPOINT ["/app/start.sh"]