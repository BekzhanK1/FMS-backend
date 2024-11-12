from fastapi import FastAPI, WebSocket, WebSocketDisconnect, Request
from fastapi.templating import Jinja2Templates
from fastapi.responses import HTMLResponse
from app.chat_manager import ChatManager

app = FastAPI()
chat_manager = ChatManager()

templates = Jinja2Templates(directory="app/templates")

users = ["user1", "user2", "user3"]


@app.get("/", response_class=HTMLResponse)
async def list_users(request: Request):
    return templates.TemplateResponse("users-list.html", {"request": request, "users": users})


@app.get("/chat/{user_id}", response_class=HTMLResponse)
async def chat_page(request: Request, user_id: str):
    return templates.TemplateResponse("chat.html", {"request": request, "user_id": user_id})


@app.websocket("/ws/chat/{user_id}")
async def websocket_endpoint(websocket: WebSocket, user_id: str):
    await chat_manager.connect(user_id, websocket)
    try:
        while True:
            data = await websocket.receive_text()
            # Echo message back to the user and broadcast to the recipient
            await chat_manager.send_message_to_user(user_id, f"User says: {data}")
    except WebSocketDisconnect:
        chat_manager.disconnect(user_id, websocket)
