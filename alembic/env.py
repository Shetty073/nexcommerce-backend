from logging.config import fileConfig
from sqlalchemy import pool
from sqlalchemy.ext.asyncio import create_async_engine
from alembic import context
from models.base import Base
from models.user import User
from models.role import Role
from config import settings
import asyncio
import fastapi_users_db_sqlalchemy

# this is the Alembic Config object
config = context.config

# Override the sqlalchemy.url with your settings
config.set_main_option("sqlalchemy.url", settings.database_url)

# Configure imports for migration files
def include_name(name, type_, parent_names):
    """Include all detected changes in migrations"""
    return True

def render_item(type_, obj, autogen_context):
    """Add custom imports to migration files"""
    if type_ == "type" and hasattr(obj, "__module__"):
        if "fastapi_users_db_sqlalchemy" in str(obj):
            autogen_context.imports.add("import fastapi_users_db_sqlalchemy")
    return False

# Interpret the config file for Python logging.
if config.config_file_name is not None:
    fileConfig(config.config_file_name)

# add your model's MetaData object here
target_metadata = Base.metadata

def run_migrations_offline() -> None:
    """Run migrations in 'offline' mode."""
    url = settings.database_url
    context.configure(
        url=url,
        target_metadata=target_metadata,
        literal_binds=True,
        dialect_opts={"paramstyle": "named"},
        include_name=include_name,
        render_item=render_item,
    )

    with context.begin_transaction():
        context.run_migrations()

def do_run_migrations(connection):
    context.configure(
        connection=connection, 
        target_metadata=target_metadata,
        include_name=include_name,
        render_item=render_item,
    )

    with context.begin_transaction():
        context.run_migrations()

async def run_async_migrations():
    """Create an Engine and associate a connection with the context."""
    connectable = create_async_engine(
        settings.database_url,
        poolclass=pool.NullPool,
    )

    async with connectable.connect() as connection:
        await connection.run_sync(do_run_migrations)

    await connectable.dispose()

def run_migrations_online() -> None:
    """Run migrations in 'online' mode."""
    asyncio.run(run_async_migrations())

if context.is_offline_mode():
    run_migrations_offline()
else:
    run_migrations_online()
