FROM python:3.9.7-bullseye

WORKDIR app

RUN pip install poetry

COPY poetry.lock .
COPY pyproject.toml .
RUN poetry install
