# main.py
from fastapi import FastAPI
from config import settings
from utils.routes_register import auto_register_routers

app = FastAPI(title=settings.project_name)

auto_register_routers(app=app)

@app.get("/")
async def health():
    return {"message": "Ok"}
