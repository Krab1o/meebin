# app/routers/requests.py

from fastapi import APIRouter, Depends, HTTPException, status
from sqlalchemy.orm import Session
from typing import List
from datetime import timedelta
from .. import models, schemas, database, auth
from fastapi.security import OAuth2PasswordBearer
from jose import JWTError, jwt

router = APIRouter(prefix="/requests", tags=["Requests"])

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

@router.post("/", response_model=schemas.RequestOut)
def create_request(request: schemas.RequestCreate, db: Session = Depends(get_db), current_user: models.User = Depends(get_current_user)):
    new_request = models.Request(
        description=request.description,
        address=request.address,
        latitude=request.latitude,
        longitude=request.longitude,
        cleaning_time=request.cleaning_time,
        comment=request.comment,
        creator_id=current_user.id
    )
    db.add(new_request)
    db.commit()
    db.refresh(new_request)
    return new_request

@router.get("/", response_model=List[schemas.RequestOut])
def get_available_requests(db: Session = Depends(get_db)):
    requests = db.query(models.Request).filter(models.Request.status == models.RequestStatus.available).all()
    return requests

@router.get("/my", response_model=List[schemas.RequestOut])
def get_my_requests(db: Session = Depends(get_db), current_user: models.User = Depends(get_current_user)):
    requests = db.query(models.Request).filter(models.Request.creator_id == current_user.id).all()
    return requests

@router.get("/accepted", response_model=List[schemas.RequestOut])
def get_accepted_requests(db: Session = Depends(get_db), current_user: models.User = Depends(get_current_user)):
    requests = db.query(models.Request).filter(models.Request.accepted_by == current_user.id).all()
    return requests

@router.get("/completed", response_model=List[schemas.RequestOut])
def get_completed_requests(db: Session = Depends(get_db), current_user: models.User = Depends(get_current_user)):
    requests = db.query(models.Request).filter(models.Request.status == models.RequestStatus.completed, models.Request.accepted_by == current_user.id).all()
    return requests

@router.post("/{request_id}/accept", response_model=schemas.RequestOut)
def accept_request(request_id: int, db: Session = Depends(get_db), current_user: models.User = Depends(get_current_user)):
    request = db.query(models.Request).filter(models.Request.id == request_id, models.Request.status == models.RequestStatus.available).first()
    if not request:
        raise HTTPException(status_code=404, detail="Заявка не найдена или уже принята")
    request.status = models.RequestStatus.accepted
    request.accepted_by = current_user.id
    db.commit()
    db.refresh(request)
    return request

@router.post("/{request_id}/complete", response_model=schemas.RequestOut)
def complete_request(request_id: int, db: Session = Depends(get_db), current_user: models.User = Depends(get_current_user)):
    request = db.query(models.Request).filter(models.Request.id == request_id, models.Request.accepted_by == current_user.id).first()
    if not request:
        raise HTTPException(status_code=404, detail="Заявка не найдена или вы её не приняли")
    request.status = models.RequestStatus.completed
    db.commit()
    db.refresh(request)
    return request
