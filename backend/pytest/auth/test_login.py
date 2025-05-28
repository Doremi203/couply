import uuid
import requests
import pytest
from faker import Faker

fake = Faker()

@pytest.fixture
def login_endpoint(auth_url):
    return f"{auth_url}/v1/login/basic"


def test_basic_login_success(login_endpoint, headers):
    """
    Успешный логин: сначала регистрируем нового пользователя, затем логинимся.
    Ожидаем 200 OK с полями token, expires_in и refresh_token.
    """
    # Регистрируем пользователя для последующего логина
    register_url = login_endpoint.replace('/login/basic', '/register/basic')
    reg_headers = headers.copy()
    reg_headers["Idempotency-Key"] = str(uuid.uuid4())
    email = fake.email()
    password = fake.password(length=12, special_chars=True, digits=True, upper_case=True)
    reg_payload = {"email": email, "password": password}
    reg_resp = requests.post(register_url, json=reg_payload, headers=reg_headers)
    assert reg_resp.status_code == 200, f"Registration failed: {reg_resp.status_code} {reg_resp.text}"

    # Выполняем логин
    resp = requests.post(login_endpoint, json=reg_payload, headers=headers)
    assert resp.status_code == 200, f"Expected 200 OK, got {resp.status_code}: {resp.text}"
    data = resp.json()
    # Проверяем схему ответа
    assert "token" in data and isinstance(data["token"], str) and data["token"], "Missing or invalid token"
    assert "expiresIn" in data and isinstance(data["expiresIn"], int), "Missing or invalid expires_in"
    assert "refreshToken" in data and isinstance(data["refreshToken"], dict), "Missing or invalid refresh_token"


def test_basic_login_invalid_credentials(login_endpoint, headers):
    """
    Попытка логина с несуществующими учётными данными: ожидаем 401 Unauthorized.
    """
    payload = {"email": fake.email(), "password": fake.password()}
    resp = requests.post(login_endpoint, json=payload, headers=headers)
    assert resp.status_code == 404, f"Expected 404 Not Found, got {resp.status_code}"


def test_basic_login_wrong_password(login_endpoint, headers):
    """
    Попытка логина для зарегистрированного пользователя с неверным паролем: ожидаем 401 Unauthorized.
    """
    # Регистрируем пользователя
    register_url = login_endpoint.replace('/login/basic', '/register/basic')
    reg_headers = headers.copy()
    reg_headers["Idempotency-Key"] = str(uuid.uuid4())
    email = fake.email()
    correct_password = fake.password(length=12, special_chars=True, digits=True, upper_case=True, lower_case=True)
    reg_payload = {"email": email, "password": correct_password}
    reg_resp = requests.post(register_url, json=reg_payload, headers=reg_headers)
    assert reg_resp.status_code == 200, f"Registration failed: {reg_resp.status_code} {reg_resp.text}"

    # Пытаемся залогиниться с неверным паролем
    wrong_password = correct_password + "x"
    resp = requests.post(
        login_endpoint,
        json={"email": email, "password": wrong_password},
        headers=headers
    )
    assert resp.status_code == 401, f"Expected 401 Unauthorized for wrong password, got {resp.status_code}"


def test_basic_login_missing_fields(login_endpoint, headers):
    """
    Логин-запрос без необходимых полей: ожидаем 400 Bad Request.
    """
    resp = requests.post(login_endpoint, json={"email": "", "password": "wrong_password"}, headers=headers)
    assert resp.status_code == 400, f"Expected 400 Bad Request, got {resp.status_code}"
