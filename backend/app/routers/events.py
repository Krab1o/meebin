# app/routers/events.py

from fastapi import APIRouter, Depends, HTTPException, status
from sqlalchemy.orm import Session
from typing import List
from datetime import timedelta
from .. import models, schemas, database, auth
from fastapi.security import OAuth2PasswordBearer
from jose import JWTError, jwt

router = APIRouter(prefix="/events", tags=["Events"])

oauth2_scheme = OAuth2PasswordBearer(tokenUrl="api/login")

def get_db():
    db = database.SessionLocal()
    try:
        yield db
    finally:
        db.close()

# Функция для получения текущего пользователя
def get_current_user(token: str = Depends(oauth2_scheme), db: Session = Depends(get_db)):
    credentials_exception = HTTPException(
        status_code=status.HTTP_401_UNAUTHORIZED,
        detail="Невалидные учетные данные",
        headers={"WWW-Authenticate": "Bearer"},
    )
    try:
        payload = jwt.decode(token, auth.SECRET_KEY, algorithms=[auth.ALGORITHM])
        username: str = payload.get("sub")
        if username is None:
            raise credentials_exception
    except JWTError:
        raise credentials_exception
    user = db.query(models.User).filter(models.User.username == username).first()
    if user is None:
        raise credentials_exception
    return user

@router.post("/", response_model=schemas.TrashEventOut)
def create_event(
    event: schemas.TrashEventCreate, 
    db: Session = Depends(database.get_db), 
    current_user: models.User = Depends(auth.get_current_user)
    ):
    new_event = models.TrashEvent(
        photo_url=event.photo_url,
        address=event.address,
        caller_id=current_user.id,
        utilizator_id=event.utilizator_id,
        event_status=event.event_status,
        time_called=event.time_called,
        time_cleaned=event.time_cleaned,
        comment=event.comment,
        confirmation_photo_url=event.confirmation_photo_url,
        price=event.price
    )
    db.add(new_event)
    db.commit()
    db.refresh(new_event)
    return new_event

@router.get("/", response_model=List[schemas.TrashEventOut])
def get_available_events(db: Session = Depends(get_db)):
    events = db.query(models.TrashEvent).filter(models.TrashEvent.status == models.TrashStatus.available).all()
    return events

@router.get("/my", response_model=List[schemas.TrashEventOut])
def get_my_events(db: Session = Depends(get_db), current_user: models.User = Depends(get_current_user)):
    events = db.query(models.TrashEvent).filter(models.TrashEvent.creator_id == current_user.id).all()
    return events

@router.get("/accepted", response_model=List[schemas.TrashEventOut])
def get_accepted_events(db: Session = Depends(get_db), current_user: models.User = Depends(get_current_user)):
    events = db.query(models.TrashEvent).filter(models.TrashEvent.accepted_by == current_user.id).all()
    return events

@router.get("/completed", response_model=List[schemas.TrashEventOut])
def get_completed_events(db: Session = Depends(get_db), current_user: models.User = Depends(get_current_user)):
    events = db.query(models.TrashEvent).filter(models.TrashEvent.status == models.TrashStatus.completed, models.TrashEvent.accepted_by == current_user.id).all()
    return events

@router.post("/{event_id}/accept", response_model=schemas.TrashEventOut)
def accept_event(event_id: int, db: Session = Depends(get_db), current_user: models.User = Depends(get_current_user)):
    event = db.query(models.TrashEvent).filter(models.TrashEvent.id == event_id, models.TrashEvent.status == models.TrashStatus.available).first()
    if not event:
        raise HTTPException(status_code=404, detail="Заявка не найдена или уже принята")
    event.status = models.TrashStatus.accepted
    event.accepted_by = current_user.id
    db.commit()
    db.refresh(event)
    return event

@router.post("/{event_id}/complete", response_model=schemas.TrashEventOut)
def complete_event(event_id: int, db: Session = Depends(get_db), current_user: models.User = Depends(get_current_user)):
    event = db.query(models.TrashEvent).filter(models.TrashEvent.id == event_id, models.TrashEvent.accepted_by == current_user.id).first()
    if not event:
        raise HTTPException(status_code=404, detail="Заявка не найдена или вы её не приняли")
    event.status = models.TrashStatus.completed
    db.commit()
    db.refresh(event)
    return event

# Получить историю заявок пользователя
@router.get("/users/{user_id}/history", response_model=List[schemas.TrashEventOut])
def get_user_request_history(user_id: int, db: Session = Depends(get_db)):
    requests = db.query(models.TrashEvent).filter(
        (models.TrashEvent.caller_id == user_id) | (models.TrashEvent.utilizator_id == user_id),
        models.TrashEvent.status == models.TrashStatus.completed
    ).all()
    return requests
