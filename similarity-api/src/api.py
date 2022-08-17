import os
from loadbalancing.register import register
from fastapi import FastAPI
from similarity.similarity_calc import calc_similarity

_VAR_API_PORT = "API_PORT"
_VAR_API_HOST = "API_HOST"
_VAR_LOADBALANCER_REGISTER_URL = "LOADBALANCER_REGISTER_URL"
_VAR_LOADBALANCER_API_KEY = "LOADBALANCER_API_KEY"


LOADBALANCER_REGISTER_URL = os.environ[_VAR_LOADBALANCER_REGISTER_URL]
LOADBALANCER_API_KEY = os.environ[_VAR_LOADBALANCER_API_KEY]
API_HOST = os.environ[_VAR_API_HOST]
API_PORT = os.environ[_VAR_API_PORT]


if __name__ == '__main__':

    register()

    app = FastAPI()

    def create_response(similarity_score: float) -> dict:
        return {"similarity": similarity_score}

    @app.get("/similarity/{sentence_1}/{sentence_2}")
    def similarity(sentence_1: str, sentence_2: str):
        score = calc_similarity(sentence_1, sentence_2)
        return create_response(score)
