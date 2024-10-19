# app/schemas.py

from pydantic import BaseModel, EmailStr, Field
from typing import Optional
from datetime import datetime
from enum import Enum

class UserBase(BaseModel):
    username: str
    email: EmailStr

class UserCreate(UserBase):
    password: str

class UserOut(UserBase):
    id: int

    class Config:
        orm_mode = True

class RequestStatus(str, Enum):
    available = "available"
    accepted = "accepted"
    completed = "completed"

class RequestBase(BaseModel):
    description: str
    address: str
    latitude: float
    longitude: float
    cleaning_time: datetime
    comment: Optional[str] = None

class RequestCreate(RequestBase):
    pass

class RequestUpdateStatus(BaseModel):
    status: RequestStatus

class RequestOut(RequestBase):
    id: int
    request_time: datetime
    status: RequestStatus
    creator_id: int
    accepted_by: Optional[int] = None

    class Config:
        orm_mode = True
