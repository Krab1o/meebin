from fastapi import FastAPI
from enum import Enum
from pydantic import BaseModel
from datetime import date

app = FastAPI()

class User(BaseModel):
    id: int
    login: str
    mail: str
    password: str
    name: str
    surname: str
    birthdate: date
    report_counter: int
    utilized_counter: int
    rating: float
    city: str



class TrashStatus(int, Enum):
    on_moderation = 0
    published = 1
    in_progress = 2
    done = 3


@app.get("/users")
async def get_users():
    return {"message": "Hello World"}

@app.get("/users/{user_id}")
async def get_user_by_id(user_id : int):
    return {"message": user_id}

@app.get("/events")
async def get_events():
    return {"message": "event1"}

@app.post("/users/")
async def create_user(user : User):
    return user