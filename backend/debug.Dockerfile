FROM cr.yandex/crpe9kgeql1v6b7gujfj/swagger-ui:latest AS swagger-ui

FROM golang:alpine AS delve
RUN go install github.com/go-delve/delve/cmd/dlv@latest

FROM golang:alpine AS build
ARG service
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
WORKDIR /src/$service/cmd/app
RUN go build -gcflags="all=-N -l" -o /src/app

FROM alpine
ARG service
WORKDIR /app
COPY --from=swagger-ui /swagger-ui swagger-ui
COPY --from=build /src/app ./
COPY --from=build /src/$service/configs configs
COPY --from=build /src/$service/swagger/api swagger-ui/api
COPY --from=build /src/$service/swagger/swagger-initializer.js swagger-ui/swagger-initializer.js

COPY --from=delve /go/bin/dlv /usr/local/bin/dlv
CMD ["dlv", "--listen=:40000", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "/app/app"]