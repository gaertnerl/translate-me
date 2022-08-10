FROM python:3.10

WORKDIR /api

COPY ./requirements.txt /api

RUN pip install --no-cache-dir --upgrade -r /api/requirements.txt

COPY ./src /api/src

WORKDIR /api/src

EXPOSE 8000

CMD ["uvicorn", "api:app", "--host", "0.0.0.0", "--port", "8000"]