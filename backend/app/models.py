# app/models.py

from sqlalchemy import Column, Integer, String, ForeignKey, Float, DateTime, Text, Enum
from sqlalchemy.orm import relationship
from database import Base
import enum
from datetime import datetime

class User(Base):
    __tablename__ = "users"

    id = Column(Integer, primary_key=True, index=True)
    username = Column(String, unique=True, index=True, nullable=False)
    email = Column(String, unique=True, index=True, nullable=False)
    hashed_password = Column(String, nullable=False)
    
    created_requests = relationship("Request", back_populates="creator")
    accepted_requests = relationship("Request", back_populates="accepted_by_user")

class RequestStatus(enum.Enum):
    available = "available"
    accepted = "accepted"
    completed = "completed"

class Request(Base):
    __tablename__ = "requests"

    id = Column(Integer, primary_key=True, index=True)
    description = Column(Text, nullable=False)
    address = Column(String, nullable=False)
    latitude = Column(Float, nullable=False)
    longitude = Column(Float, nullable=False)
    request_time = Column(DateTime, default=datetime.utcnow, nullable=False)
    cleaning_time = Column(DateTime, nullable=False)
    comment = Column(Text, nullable=True)
    status = Column(Enum(RequestStatus), default=RequestStatus.available, nullable=False)
    
    creator_id = Column(Integer, ForeignKey("users.id"))
    accepted_by = Column(Integer, ForeignKey("users.id"), nullable=True)
    
    creator = relationship("User", back_populates="created_requests", foreign_keys=[creator_id])
    accepted_by_user = relationship("User", back_populates="accepted_requests", foreign_keys=[accepted_by])
