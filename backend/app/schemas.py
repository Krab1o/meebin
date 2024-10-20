from pydantic import BaseModel, EmailStr, Field
from typing import Optional
from datetime import datetime

class UserBase(BaseModel):
    login: str
    mail: EmailStr
    name: str
    lastname: str
    surname: Optional[str] = None
    birthdate: datetime
    city: str

class UserCreate(UserBase):
    password: str

class UserOut(UserBase):
    id: int

    class Config:
        orm_mode = True

class TrashEventBase(BaseModel):
    photo_url: str
    address: str
    event_status: int
    time_called: datetime
    time_cleaned: datetime
    comment: str
    confirmation_photo_url: str
    price: int

class TrashEventCreate(TrashEventBase):
    utilizator_id: int

class TrashEventOut(TrashEventBase):
    id: int
    caller_id: int
    utilizator_id: int

    class Config:
        orm_mode = True
