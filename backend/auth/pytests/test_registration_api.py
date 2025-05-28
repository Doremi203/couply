import uuid
import requests
import pytest

@pytest.fixture
def endpoint(auth_url):
    return f"{auth_url}/v1/register/basic"

def test_basic_register_success(endpoint, headers):
    """
    Успешная регистрация: валидный UUID в заголовке и корректный payload.
    Ожидаем 200 OK и пустой JSON-объект в ответе.
    """
    # Копируем headers и добавляем идемпотентный ключ
    hdr = headers.copy()
    hdr["Idempotency-Key"] = str(uuid.uuid4())

    payload = {
        "email": "user@example.com",
        "password": "SecureP@ssw0rd"
    }

    resp = requests.post(endpoint, json=payload, headers=hdr)
    assert resp.status_code == 200, f"Expected 200, got {resp.status_code}: {resp.text}"
    assert resp.headers.get("Content-Type", "").startswith("application/json")
    assert resp.json() == {}, "Expected empty JSON body for BasicRegisterResponseV1"

def test_basic_register_missing_idempotency_key(endpoint, headers):
    """
    Запрос без заголовка Idempotency-Key должен возвращать ошибку 400.
    """
    resp = requests.post(endpoint, json={
        "email": "user2@example.com",
        "password": "AnotherP@ss"
    }, headers=headers)  # не добавляем ключ
    assert resp.status_code == 400, f"Expected 400 Bad Request, got {resp.status_code}"

def test_basic_register_invalid_idempotency_key(endpoint, headers):
    """
    Запрос с некорректным UUID в заголовке должен вернуть 400.
    """
    hdr = headers.copy()
    hdr["Idempotency-Key"] = "not-a-uuid"

    resp = requests.post(endpoint, json={
        "email": "user3@example.com",
        "password": "ThirdP@ss"
    }, headers=hdr)
    assert resp.status_code == 400, f"Expected 400 on invalid UUID, got {resp.status_code}"