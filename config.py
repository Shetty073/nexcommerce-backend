from pydantic_settings import BaseSettings, SettingsConfigDict
from pydantic import Field
import yaml

class Settings(BaseSettings):
    # ==== Your .env variables ====
    debug: bool = Field(False, env="DEBUG")
    database_url: str = Field(..., env="DATABASE_URL")

    # ==== Additional values coming from yaml ====
    project_name: str = "NexCommerce"

    model_config = SettingsConfigDict(env_file=".env", extra="ignore")

    def load_from_yaml(self, yaml_file: str = "config.yaml"):
        try:
            with open(yaml_file) as f:
                data = yaml.safe_load(f)
                for key, value in data.items():
                    if hasattr(self, key):
                        setattr(self, key, value)
        except FileNotFoundError:
            pass

settings = Settings()
settings.load_from_yaml()   # ✔ loads yaml values (overwrite env/defaults)
