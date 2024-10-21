from fastapi import APIRouter, Depends, HTTPException
from sqlalchemy.orm import Session
from . import models, schemas, database, auth
from fastapi.security import OAuth2PasswordRequestForm

router = APIRouter()

@router.post("/register", response_model=schemas.UserOut)
def register(user: schemas.UserCreate, db: Session = Depends(database.get_db)):
    db_user = db.query(models.User).filter(models.User.login == user.login).first()
    if db_user:
        raise HTTPException(status_code=400, detail="Такой пользователь уже зарегистрирован")
    
    hashed_password = auth.get_password_hash(user.password)
    new_user = models.User(
        login=user.login, 
        mail=user.mail, 
        password=hashed_password, 
        name=user.name, 
        lastname=user.lastname, 
        surname=user.surname, 
        birthdate=user.birthdate, 
        city=user.city
    )
    
    db.add(new_user)
    db.commit()
    db.refresh(new_user)
    
    return new_user

@router.post("/login")
def login(form_data: OAuth2PasswordRequestForm = Depends(), db: Session = Depends(database.get_db)):
    user = db.query(models.User).filter(models.User.login == form_data.username).first()
    if not user or not auth.verify_password(form_data.password, user.password):
        raise HTTPException(status_code=400, detail="Неверный логин или пароль")
    
    access_token = auth.create_access_token(data={"sub": user.login})
    return {"access_token": access_token, "token_type": "bearer"}


# Включаем маршруты заявок
from .routers import events, users

main_router = APIRouter()
main_router.include_router(router, prefix="/auth", tags=["Authentication"])
main_router.include_router(events.router)
main_router.include_router(users.router)
