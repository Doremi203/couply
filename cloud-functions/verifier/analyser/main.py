import os

import cv2
import functions_framework
import grpc
import mediapipe as mp
import numpy as np
import requests
from insightface.app import FaceAnalysis
from mediapipe.tasks.python import vision as mp_vision
from mediapipe.tasks.python.core.base_options import BaseOptions
from mediapipe.tasks.python.vision import GestureRecognizer, GestureRecognizerOptions

import user_service_pb2
import user_service_pb2_grpc

model_dir = os.path.join(os.path.dirname(__file__), 'models', 'insightface_models')

face_model = FaceAnalysis(
    root=model_dir,
    name='buffalo_l',
    providers=['CPUExecutionProvider'],
)
face_model.prepare(ctx_id=0, det_size=(640, 640))

gesture_options = GestureRecognizerOptions(
    base_options=BaseOptions(model_asset_path="gesture_recognizer.task"),
    running_mode=mp_vision.RunningMode.IMAGE
)

gesture_recognizer = GestureRecognizer.create_from_options(gesture_options)


@functions_framework.http
def handler(request):
    data = request.get_json(silent=True)
    if not data:
        return {"error": "Bad Request: JSON body required"}, 400
    print('data: ', data)
    messages = data.get("messages")
    if len(messages) != 1:
        raise Exception("Bad Request: No message found")

    details = messages[0].get("details")

    bucket = details.get("bucket_id")
    key = details.get("object_id")
    print('bucket: ', bucket, 'key: ', key)
    expected_gesture = data.get("challenge", "Victory")
    #user_id = data.get("user_id", key.split("/")[0])
    token = fetch_iam_token()

    img_bytes = download_object(bucket, key, token)

    nparr = np.frombuffer(img_bytes, np.uint8)
    img = cv2.imdecode(nparr, cv2.IMREAD_COLOR)[:, :, ::-1]  # BGR→RGB

    faces = face_model.get(img)
    if len(faces) != 1:
        # send_verification_status(user_id, user_service_pb2.VerificationStatus.MANUAL)
        raise ValueError(f"Expected 1 face, found {len(faces)}")

    mp_image = mp.Image(image_format=mp.ImageFormat.SRGB, data=img)
    result = gesture_recognizer.recognize(mp_image)
    detected = result.gestures[0][0].category_name if result.gestures else None
    if detected != expected_gesture:
        # send_verification_status(user_id, user_service_pb2.VerificationStatus.FAIL)
        raise ValueError(f"Expected {expected_gesture}, got {detected}")

    # send_verification_status(user_id, user_service_pb2.VerificationStatus.PASS)

    return {
        'statusCode': 200,
    }


def download_object(bucket: str, key: str, token: str) -> bytes:
    url = f"https://storage.yandexcloud.net/{bucket}/{key}"
    resp = requests.get(
        url,
        headers={"Authorization": f"Bearer {token}"},
        timeout=10
    )
    resp.raise_for_status()
    return resp.content


def send_verification_status(user_id: str, status) -> bool:
    channel = grpc.insecure_channel("matcher.testing.couply.ru:5051")
    stub = user_service_pb2_grpc.UserServiceStub(channel)

    req = user_service_pb2.SetUserVerificationStatusByIDRequest(
        user_id=user_id,
        status=status
    )
    resp = stub.SetUserVerificationStatusByID(req, timeout=5)  # таймаут 5 с
    return resp.success

def fetch_iam_token():
    resp = requests.get(
        "http://169.254.169.254/computeMetadata/v1/instance/service-accounts/default/token",
        headers={"Metadata-Flavor": "Google"},
        timeout=5
    )
    resp.raise_for_status()
    return resp.json()["access_token"]