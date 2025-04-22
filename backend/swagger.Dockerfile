FROM alpine AS swagger-ui
RUN apk add --no-cache curl tar jq

RUN SWAGGER_VERSION=$(curl -s https://api.github.com/repos/swagger-api/swagger-ui/releases/latest | jq -r .tag_name) && \
    echo "Latest Swagger UI version: $SWAGGER_VERSION" && \
    curl -L "https://github.com/swagger-api/swagger-ui/archive/refs/tags/${SWAGGER_VERSION}.tar.gz" -o swagger-ui.tar.gz && \
    mkdir swagger-ui && \
    tar -xzf swagger-ui.tar.gz -C swagger-ui --strip-components=2 swagger-ui-${SWAGGER_VERSION#v}/dist