from sqlalchemy import String
from sqlalchemy.orm import Mapped, mapped_column, relationship
from uuid import uuid4, UUID
from models.base import Base
from typing import List, TYPE_CHECKING

if TYPE_CHECKING:
    from models.user import User

class Role(Base):
    __tablename__ = "roles"
    
    id: Mapped[UUID] = mapped_column(primary_key=True, default=uuid4)
    name: Mapped[str] = mapped_column(String(50), unique=True)
    
    # Optional: back reference to users
    users: Mapped[List["User"]] = relationship("User", back_populates="role")
