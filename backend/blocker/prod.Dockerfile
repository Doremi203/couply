FROM golang:alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
WORKDIR /src/blocker/cmd/blocker
RUN go build -ldflags="-s -w" -o /src/app

FROM alpine
WORKDIR /app
COPY --from=build /src/app ./
COPY --from=build /src/blocker/configs configs
CMD ["./app"]