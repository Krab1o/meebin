from sqlalchemy import Column, Integer, String, BigInteger, ForeignKey, Float, TIMESTAMP
from sqlalchemy.orm import relationship
from .database import Base

class User(Base):
    __tablename__ = "Users"

    id = Column(Integer, primary_key=True, index=True, unique=True)
    login = Column(String(50), nullable=False)
    mail = Column(String(80), nullable=False)
    password = Column(String(255), nullable=False)
    name = Column(String(60), nullable=False)
    lastname = Column(String(60), nullable=False)
    surname = Column(String(60), nullable=True)
    birthdate = Column(TIMESTAMP, nullable=False, default='-1')
    report_counter = Column(BigInteger, default=0, nullable=False)
    utilized_counter = Column(BigInteger, default=0, nullable=False)
    rating = Column(Float, nullable=True)
    city = Column(String(60), nullable=False)

    roles = relationship("Role", secondary="Users_Roles", back_populates="users")
    called_events = relationship("TrashEvent", foreign_keys="[TrashEvent.caller_id]", back_populates="caller")
    utilized_events = relationship("TrashEvent", foreign_keys="[TrashEvent.utilizator_id]", back_populates="utilizator")


class TrashEvent(Base):
    __tablename__ = "TrashEvent"

    id = Column(Integer, primary_key=True, index=True, unique=True)
    photo_url = Column(String(100), nullable=False)
    address = Column(String(255), nullable=False)
    caller_id = Column(BigInteger, ForeignKey("Users.id"), nullable=False)
    utilizator_id = Column(BigInteger, ForeignKey("Users.id"), nullable=False)
    event_status = Column(BigInteger, ForeignKey("TrashStatus.id"), default=0, nullable=False)
    time_called = Column(TIMESTAMP, nullable=False)
    time_cleaned = Column(TIMESTAMP, nullable=False)
    comment = Column(String(255), nullable=False)
    confirmation_photo_url = Column(String(255), nullable=False)
    price = Column(BigInteger, nullable=False)

    caller = relationship("User", foreign_keys=[caller_id], back_populates="called_events")
    utilizator = relationship("User", foreign_keys=[utilizator_id], back_populates="utilized_events")
    status = relationship("TrashStatus")


class Role(Base):
    __tablename__ = "Roles"

    id = Column(Integer, primary_key=True, index=True, unique=True)
    title = Column(String(60), nullable=False)

    users = relationship("User", secondary="Users_Roles", back_populates="roles")


class UsersRoles(Base):
    __tablename__ = "Users_Roles"

    id_roles = Column(BigInteger, ForeignKey("Roles.id"), primary_key=True)
    id_users = Column(BigInteger, ForeignKey("Users.id"), primary_key=True)


class TrashStatus(Base):
    __tablename__ = "TrashStatus"

    id = Column(Integer, primary_key=True, index=True, unique=True)
    title = Column(String(60), nullable=False)
