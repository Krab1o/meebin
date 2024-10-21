from fastapi import APIRouter, Depends, HTTPException, status
from sqlalchemy.orm import Session
from .. import models, database, schemas, auth
from typing import List

router = APIRouter(prefix="/users", tags=["Users"])

# Получить всех пользователей
@router.get("/", response_model=List[schemas.UserOut])
def get_users(db: Session = Depends(database.get_db)):
    users = db.query(models.User).all()
    return users

@router.get("/{user_id}")
def get_users(user_ud: int, db: Session = Depends(database.get_db)):
    users = db.query(models.User).all()
    return users

# Удалить пользователя
@router.delete("/{user_id}")
def delete_user(user_id: int, db: Session = Depends(database.get_db)):
    user = db.query(models.User).filter(models.User.id == user_id).first()
    if not user:
        raise HTTPException(status_code=404, detail="Пользователь не найден")
    db.delete(user)
    db.commit()
    return {"detail": "Пользователь удален"}

# Обновить информацию о пользователе
@router.put("/{user_id}", response_model=schemas.UserOut)
def update_user(user_id: int, user_update: schemas.UserCreate, db: Session = Depends(database.get_db)):
    user = db.query(models.User).filter(models.User.id == user_id).first()
    if not user:
        raise HTTPException(status_code=404, detail="Пользователь не найден")
    
    user.login = user_update.login
    user.mail = user_update.mail
    user.password = auth.get_password_hash(user_update.password)
    user.name = user_update.name
    user.lastname = user_update.lastname
    user.surname = user_update.surname, 
    user.birthdate = user_update.birthdate, 
    user.city = user_update.city        
    
    db.commit()
    db.refresh(user)
    return user


