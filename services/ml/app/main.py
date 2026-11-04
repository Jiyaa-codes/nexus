from fastapi import FastAPI

app = FastAPI(title="Nexus ML Service", version="0.1.0")


@app.get("/healthz")
def health() -> dict[str, str]:
    return {"service": "ml", "status": "ok"}
