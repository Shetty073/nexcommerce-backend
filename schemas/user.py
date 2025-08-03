from fastapi_users import schemas
from typing import Optional
from uuid import UUID

class UserRead(schemas.BaseUser[UUID]):
    full_name: str
    phone_number: str
    role_id: UUID

class UserCreate(schemas.BaseUserCreate):
    full_name: str
    phone_number: str
    role_id: UUID

class UserUpdate(schemas.BaseUserUpdate):
    full_name: Optional[str] = None
    phone_number: Optional[str] = None
    role_id: Optional[UUID] = None
