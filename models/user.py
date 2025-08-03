from fastapi_users.db import SQLAlchemyBaseUserTableUUID
from sqlalchemy import String, ForeignKey
from sqlalchemy.orm import Mapped, mapped_column, relationship
from uuid import UUID
from models.base import Base
from models.role import Role

class User(SQLAlchemyBaseUserTableUUID, Base):
    __tablename__ = "users"
    
    full_name: Mapped[str] = mapped_column(String(100))
    phone_number: Mapped[str] = mapped_column(String(20))
    
    role_id: Mapped[UUID] = mapped_column(ForeignKey("roles.id"))
    role: Mapped[Role] = relationship("Role", back_populates="users")
