FROM cr.yandex/crpe9kgeql1v6b7gujfj/swagger-ui:latest AS swagger-ui

FROM golang:alpine AS build
WORKDIR /src
#COPY go.mod go.sum ./
#RUN go mod download
COPY . .
#WORKDIR /src/$service/cmd/app
#RUN go build -ldflags="-s -w" -o /src/app

FROM alpine
ARG service
WORKDIR /app
COPY --from=swagger-ui /swagger-ui swagger-ui
COPY --from=build /src/$service/cmd/app/app /app/app
COPY --from=build /src/$service/configs configs
COPY --from=build /src/$service/swagger/api swagger-ui/api
COPY --from=build /src/$service/swagger/swagger-initializer.js swagger-ui/swagger-initializer.js
CMD ["/app/app"]