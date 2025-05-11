FROM golang:alpine AS build
ARG service
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
WORKDIR /src/$service/cmd/app
RUN go build -ldflags="-s -w" -o /src/app

FROM alpine
ARG service
WORKDIR /app
COPY --from=build /src/app ./
COPY --from=build /src/$service/configs configs
CMD ["./app"]