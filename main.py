# main.py
from fastapi import FastAPI
from config import settings
from utils.routes_register import auto_register_routers

from auth.auth import (
    fastapi_users,
    auth_backend,
    get_user_manager,
)
from schemas.user import UserCreate, UserRead, UserUpdate


app = FastAPI(title=settings.project_name)


auto_register_routers(app=app)


@app.get("/")
async def health():
    return {"message": "Ok"}


# login/logout
app.include_router(
    fastapi_users.get_auth_router(auth_backend),
    prefix="/auth/jwt",
    tags=["auth"],
)


# registers
app.include_router(
    fastapi_users.get_register_router(UserRead, UserCreate),
    prefix="/auth",
    tags=["auth"],
)


# current user, update, etc.
app.include_router(
    fastapi_users.get_users_router(UserRead, UserUpdate),
    prefix="/users",
    tags=["users"],
)
