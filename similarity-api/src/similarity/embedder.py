from similarity.custom_types import Embedding
import tensorflow_hub as hub
import tensorflow_text
from threading import Lock
import sys


class CachedEmbedder:

    def __init__(self, cache_size: int) -> None:
        self.__c = {}
        self.__embed = hub.load(
            "https://tfhub.dev/google/universal-sentence-encoder-multilingual/3")
        self.__lock = Lock()
        self.__cache_size = cache_size

    def size_exceeded(self) -> bool:
        return len(self.__c) > self.__cache_size

    def store(self, string: str, embedding: Embedding):
        exceeded = self.size_exceeded()
        with self.__lock:
            if exceeded:
                self.__c = {}
            self.__c[string] = embedding

    def get(self, string: str) -> Embedding:
        with self.__lock:
            stored = self.__c.get(string)
        if stored is not None:
            return stored
        embedding = self.__embed([string])
        self.store(string, embedding)
        return embedding


embedder = CachedEmbedder(5000)
