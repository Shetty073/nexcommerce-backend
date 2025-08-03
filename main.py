from fastapi import FastAPI
from config import settings

app = FastAPI(title=settings.project_name)

from api.v1 import example as example_v1
from api.v2 import example as example_v2

app.include_router(example_v1.router, prefix="/api/v1")
app.include_router(example_v2.router, prefix="/api/v2")

@app.get("/")
async def health():
    return {"message": "Ok"}
