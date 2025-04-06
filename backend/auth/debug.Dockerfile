FROM alpine AS swagger-ui
RUN apk add --no-cache curl tar jq

RUN SWAGGER_VERSION=$(curl -s https://api.github.com/repos/swagger-api/swagger-ui/releases/latest | jq -r .tag_name) && \
    echo "Latest Swagger UI version: $SWAGGER_VERSION" && \
    curl -L "https://github.com/swagger-api/swagger-ui/archive/refs/tags/${SWAGGER_VERSION}.tar.gz" -o swagger-ui.tar.gz && \
    mkdir swagger-ui && \
    tar -xzf swagger-ui.tar.gz -C swagger-ui --strip-components=2 swagger-ui-${SWAGGER_VERSION#v}/dist

FROM golang:alpine AS delve
RUN go install github.com/go-delve/delve/cmd/dlv@latest

FROM golang:alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
WORKDIR /src/cmd/auth
RUN go build -gcflags="all=-N -l" -o /src/app

FROM alpine
WORKDIR /app
COPY --from=build /src/app ./
COPY configs configs
COPY --from=swagger-ui /swagger-ui swagger-ui
COPY swagger/swagger-initializer.js swagger-ui/swagger-initializer.js
COPY swagger/api swagger-ui/api

COPY --from=delve /go/bin/dlv /usr/local/bin/dlv
CMD ["dlv", "--listen=:40000", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "/app/app"]