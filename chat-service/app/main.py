from fastapi import FastAPI, WebSocket, WebSocketDisconnect, Request
from fastapi.templating import Jinja2Templates
from fastapi.responses import HTMLResponse
from app.chat_manager import ChatManager

app = FastAPI()
chat_manager = ChatManager()

templates = Jinja2Templates(directory="app/templates")


@app.get("/chat", response_class=HTMLResponse)
def render_chat(request: Request):
    return templates.TemplateResponse("client.html", {"request": request})


@app.websocket("/ws/chat")
async def websocket_endpoint(websocket: WebSocket):
    await chat_manager.connect(websocket)
    try:
        while True:
            data = await websocket.receive_text()
            await chat_manager.broadcast(f"Client says: {data}")
    except WebSocketDisconnect:
        chat_manager.disconnect(websocket)
        await chat_manager.broadcast("A client disconnected.")
