#!/bin/bash
set -e

echo "[INFO] Получение IAM токена из Metadata Service..."
export YC_TOKEN=$(curl -sf -H Metadata-Flavor:Google 169.254.169.254/computeMetadata/v1/instance/service-accounts/default/token | jq -r .access_token)
echo "[INFO] Получение секрета из Yandex Lockbox..."
export db-id=$(curl -sf -H Metadata-Flavor:Google 169.254.169.254/latest/user-data | yq .datasource.secrets.db | tr -d \")
export DATABASE_USER=$(curl -sf -H "Authorization: Bearer $YC_TOKEN" "https://payload.lockbox.api.cloud.yandex.net/lockbox/v1/secrets/${db-id}/payload" | jq -r .entries[0].textValue)
export DATABASE_PASSWORD=$(curl -sf -H "Authorization: Bearer $YC_TOKEN" "https://payload.lockbox.api.cloud.yandex.net/lockbox/v1/secrets/${db-id}/payload" | jq -r .entries[1].textValue)
echo "[INFO] Секреты успешно загружены"

echo "[INFO] Запуск Auth"
./app