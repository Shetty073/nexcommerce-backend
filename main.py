# main.py
from fastapi import FastAPI
from config import settings
import importlib
import pkgutil

app = FastAPI(title=settings.project_name)

def auto_register_routers():
    import api  # Make sure api/__init__.py exists!

    for finder, name, ispkg in pkgutil.walk_packages(api.__path__, api.__name__ + "."):
        module = importlib.import_module(name)
        router = getattr(module, "router", None)
        if router:
            # api.v1.example → ["api", "v1", "example"]
            parts = name.split(".")
            # want prefix = /api/v1/example
            prefix = "/" + "/".join(parts[:-1])
            app.include_router(router, prefix=prefix)
            print("Mounted", name, "at", prefix)

auto_register_routers()

@app.get("/")
async def health():
    return {"message": "Ok"}
