import os
import pytest
from dotenv import load_dotenv

for v in ("HTTP_PROXY","http_proxy","HTTPS_PROXY","https_proxy"):
    os.environ.pop(v, None)
os.environ["NO_PROXY"] = "localhost,127.0.0.1"

load_dotenv(dotenv_path=os.path.join(os.path.dirname(__file__), ".env"))

def pytest_addoption(parser):
    parser.addoption(
        "--auth-url",
        action="store",
        default=os.getenv("AUTH_URL", "http://localhost:9090"),
        help="Base URL для auth-сервиса (можно задать через ENV: AUTH_URL)"
    )
    parser.addoption(
        "--notificator-url",
        action="store",
        default=os.getenv("NOTIFICATOR_URL", "http://localhost:8002"),
        help="Base URL для notificator-сервиса (можно задать через ENV: NOTIFICATOR_URL)"
    )

@pytest.fixture(scope="session")
def auth_url(request):
    return request.config.getoption("--auth-url")

@pytest.fixture(scope="session")
def notificator_url(request):
    return request.config.getoption("--notificator-url")

@pytest.fixture
def headers():
    return {
        "Content-Type": "application/json",
    }

