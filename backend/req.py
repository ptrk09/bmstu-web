import requests

url = "http://localhost:5005"
method = "/v1/auth/sign-in"

data = requests.post(
    url + method,
    headers = {
        'accept': '*/*',
        'Content-Type': 'application/json',
    },
    data={
        'login': 'edlavr',
        'password': '12345678'
    }
)

print(data.json)

