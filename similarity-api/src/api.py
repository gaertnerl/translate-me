from fastapi import FastAPI
from similarity.similarity_calc import calc_similarity


app = FastAPI()


def create_response(similarity_score: float) -> dict:
    return {"similarity": similarity_score}


@app.get("/similarity/{sentence_1}/{sentence_2}")
def similarity(sentence_1: str, sentence_2: str):
    score = calc_similarity(sentence_1, sentence_2)
    return create_response(score)
