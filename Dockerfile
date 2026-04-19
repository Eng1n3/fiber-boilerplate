# Building the binary of the App
FROM golang:1.26-alpine AS build

# `fiber_boilerplate` should be replaced with your project name
WORKDIR /go/src/fiber_boilerplate


# Copy all the Code and stuff to compile everything
COPY . .

# Downloads all the dependencies in advance (could be left out, but it's more clear this way)
RUN go mod download

# Builds the application as a staticly linked one, to allow it to run on alpine
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app cmd/main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o seeder ./seeder


# Moving the binary to the 'final Image' to make it smaller
FROM alpine:latest AS release

WORKDIR /app

# Create the `public` dir and copy all the assets into it
RUN mkdir ./static
COPY ./static ./static

# `fiber_boilerplate` should be replaced here as well
COPY --from=build /go/src/fiber_boilerplate/app .
COPY --from=build /go/src/fiber_boilerplate/seeder .

# `fiber_boilerplate` should be replaced here as well
# COPY --from=build /go/src/fiber_boilerplate/database/seeder/main ./seeder

# Add packages
RUN apk -U upgrade \
&& apk add --no-cache ca-certificates \
&& chmod +x /app/app \
&& chmod +x /app/seeder

# Add migrate tool to run the migrations in the entrypoint
RUN apk add --no-cache curl \
    && curl -L https://github.com/golang-migrate/migrate/releases/latest/download/migrate.linux-amd64.tar.gz \
    | tar xvz \
    && mv migrate /usr/local/bin/migrate
# Exposes port 8084 because our program listens on that port
EXPOSE 8084

CMD ["./app"]
