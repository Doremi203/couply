FROM cr.yandex/crpe9kgeql1v6b7gujfj/swagger-ui:latest AS swagger-ui

FROM golang:alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
WORKDIR /src/blocker/cmd/blocker
RUN go build -ldflags="-s -w" -o /src/app

FROM alpine
WORKDIR /app
COPY --from=swagger-ui /swagger-ui swagger-ui
COPY --from=build /src/app ./
COPY --from=build /src/blocker/configs configs
COPY --from=build /src/blocker/swagger/api swagger-ui/api
COPY --from=build /src/blocker/swagger/swagger-initializer.js swagger-ui/swagger-initializer.js
CMD ["./app"]