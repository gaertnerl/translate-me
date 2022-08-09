import numpy as np
from similarity.embedder import embedder


def extract_similarity(similarity_matrix) -> float:
    return min(similarity_matrix[0][0], 1.0)


def calc_similarity(original: str, user_translation: str) -> float:

    original = embedder.get(original)
    user_translation = embedder.get(user_translation)
    original_user_sim = extract_similarity(
        np.inner(original, user_translation))

    return float(original_user_sim)
