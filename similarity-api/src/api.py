from similarity.similarity_calc import calc_similarity
from time import sleep
from fastapi import FastAPI
from loadbalancing.register import register
from tkinter import N
from argparse import ArgumentError
import os
import log


class Envs:

    def assertAndGetEnv(name):
        val = os.getenv(name)
        if val is None or val == '':
            raise ArgumentError(f"expected arg {name} is not present")
        return val

    _VAR_API_PORT = "API_PORT"
    _VAR_API_HOST = "API_HOST"
    _VAR_LOADBALANCER_REGISTER_URL = "LOADBALANCER_REGISTER_URL"
    _VAR_LOADBALANCER_API_KEY = "LOADBALANCER_API_KEY"

    LOADBALANCER_REGISTER_URL = assertAndGetEnv(_VAR_LOADBALANCER_REGISTER_URL)
    LOADBALANCER_API_KEY = assertAndGetEnv(_VAR_LOADBALANCER_API_KEY)
    API_HOST = assertAndGetEnv(_VAR_API_HOST)
    API_PORT = assertAndGetEnv(_VAR_API_PORT)


register(Envs.LOADBALANCER_REGISTER_URL,
         Envs.LOADBALANCER_API_KEY,
         Envs.API_PORT)

app = FastAPI()


def create_response(similarity_score: float) -> dict:
    return {"similarity": similarity_score}


@app.get("/similarity/{sentence_1}/{sentence_2}")
def similarity(sentence_1: str, sentence_2: str):
    score = calc_similarity(sentence_1, sentence_2)
    return create_response(score)
