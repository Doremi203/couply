FROM golang:alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
WORKDIR /src/cmd/auth
RUN go build -ldflags="-s -w" -o /src/app

FROM alpine
WORKDIR /app
COPY --from=build /src/app ./
COPY --from=build /scripts/entrypoint.sh /entrypoint.sh
COPY configs configs
ENTRYPOINT ["/entrypoint.sh"]