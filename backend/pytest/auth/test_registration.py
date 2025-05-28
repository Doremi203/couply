import uuid
import requests
import pytest
from faker import Faker

fake = Faker()

@pytest.fixture
def registration_endpoint(auth_url):
    return f"{auth_url}/v1/register/basic"

def test_basic_register_success(registration_endpoint, headers):
    """
    Успешная регистрация: валидный UUID в заголовке и корректный payload.
    Ожидаем 200 OK и пустой JSON-объект в ответе.
    """
    # Копируем headers и добавляем идемпотентный ключ
    hdr = headers.copy()
    hdr["Idempotency-Key"] = str(uuid.uuid4())

    payload = {
        "email": fake.email(),
        "password": fake.password(length=12, special_chars=True, digits=True, upper_case=True)
    }

    resp = requests.post(registration_endpoint, json=payload, headers=hdr)
    assert resp.status_code == 200, f"Expected 200, got {resp.status_code}: {resp.text}"
    assert resp.headers.get("Content-Type", "").startswith("application/json")
    assert resp.json() == {}, "Expected empty JSON body for BasicRegisterResponseV1"

def test_basic_register_missing_idempotency_key(registration_endpoint, headers):
    """
    Запрос без заголовка Idempotency-Key должен возвращать ошибку 400.
    """
    resp = requests.post(registration_endpoint, json={
        "email": "user2@example.com",
        "password": "AnotherP@ss"
    }, headers=headers)  # не добавляем ключ
    assert resp.status_code == 400, f"Expected 400 Bad Request, got {resp.status_code}"

def test_basic_register_invalid_idempotency_key(registration_endpoint, headers):
    """
    Запрос с некорректным UUID в заголовке должен вернуть 400.
    """
    hdr = headers.copy()
    hdr["Idempotency-Key"] = "not-a-uuid"

    resp = requests.post(registration_endpoint, json={
        "email": "user3@example.com",
        "password": "ThirdP@ss"
    }, headers=hdr)
    assert resp.status_code == 400, f"Expected 400 on invalid UUID, got {resp.status_code}"


@pytest.mark.parametrize("password,error_msg", [
    ("Ab1!a", "password must be at least 8 characters long"),
    ("A" * 17 + "!", "password must be at most 16 characters long"),
    ("Abcdefgh1A", "password must contain at least one special character (_!@#?)"),
    ("abcdef!1", "password must contain at least one uppercase letter"),
])
def test_register_invalid_passwords(registration_endpoint, headers, password, error_msg):
    """
    Проверяем, что при невалидных паролях функция NewPassword возвращает ошибку,
    а HTTP-эндпоинт регистрации возвращает 400 Bad Request с соответствующим сообщением.
    """
    hdr = headers.copy()
    hdr["Idempotency-Key"] = str(uuid.uuid4())
    payload = {"email": fake.email(), "password": password}
    resp = requests.post(registration_endpoint, json=payload, headers=hdr)

    assert resp.status_code == 400, (
        f"Expected 400 for password '{password}', got {resp.status_code}"
    )
    # Проверяем, что текст ошибки содержит сообщение от NewPassword
    try:
        data = resp.json()
        error_text = data.get("error") or data.get("message", "")
    except ValueError:
        error_text = resp.text
    assert error_msg in error_text, (
        f"Expected error message '{error_msg}', got '{error_text}'"
    )
