import importlib
import pkgutil
from fastapi.applications import FastAPI

def auto_register_routers(app: FastAPI):
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