import requests


def register(loadbalancer_api_key, api_host, api_port):
    data = {
        "Key": loadbalancer_api_key,
        "Service": {
            "Port": api_port,
            "Host": api_host
        }
    }
    requests.post("http://bugs.python.org", data=data)
