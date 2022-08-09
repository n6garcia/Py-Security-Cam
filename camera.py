import cv2
import requests
from datetime import datetime
from time import sleep

def sendImage(frame):
    imencoded = cv2.imencode(".jpg", frame)[1]

    now = datetime.now()
    seq = now.strftime("%Y%m%d%H%M%S")

    file = {'file': (seq, imencoded.tobytes(), 'image/jpeg')}

    try:
        response = requests.post("https://noeldev.site/cam", files=file, timeout=5)
        return response
    except requests.exceptions.RequestException as e:
        print(e)

    return

def takeImage():
    cap = cv2.VideoCapture(0)
    ret, frame = cap.read()
    print(sendImage(frame))
    cap.release()


while 1:
    takeImage()
    sleep(5)