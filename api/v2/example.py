from fastapi import APIRouter

router = APIRouter(prefix="/example", tags=["Example v2"])

@router.post("/")
def example():
    return [{"id": 1, "name": "Alice (v2)"}, {"id": 2, "name": "Bob (v2)"}]
