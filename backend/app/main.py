# app/main.py

from fastapi import FastAPI
from . import models, database
from .router import main_router

app = FastAPI()

# Создаем все таблицы
models.Base.metadata.create_all(bind=database.engine)

app.include_router(main_router, prefix="/api")
