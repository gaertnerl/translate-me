import requests


def register(loadbalancer_register_url, loadbalancer_api_key, api_port):
    data = {
        "Protocol": "http",
        "Key": loadbalancer_api_key,
        "Port": api_port
    }
    resp = requests.post(loadbalancer_register_url, json=data)
    if resp.status_code != 200:
        raise RuntimeError(
            "could not register at loadbalancer, received response with status " + str(resp.status_code))
