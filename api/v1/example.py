from fastapi import APIRouter

router = APIRouter(prefix="/example", tags=["Example v1"])

@router.post("/")
def example():
    return [{"id": 1, "name": "Alice (v1)"}]
