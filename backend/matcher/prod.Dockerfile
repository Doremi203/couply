FROM golang:alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
WORKDIR /src/matcher/cmd/matcher
RUN go build -ldflags="-s -w" -o /src/app

FROM alpine
WORKDIR /app
COPY --from=build /src/app ./
COPY --from=build /src/matcher/configs configs
CMD ["./app"]