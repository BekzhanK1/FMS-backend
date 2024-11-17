import json
from fastapi import FastAPI, WebSocket, WebSocketDisconnect, Request
from fastapi.templating import Jinja2Templates
from fastapi.responses import HTMLResponse
from app.chat_manager import ChatManager

app = FastAPI()
chat_manager = ChatManager()

templates = Jinja2Templates(directory="app/templates")

# Farmer ID (can be dynamically retrieved from authentication in real-world applications)
FARMER_ID = "farmer1"


@app.get("/chats", response_class=HTMLResponse)
async def get_chats():
    """
    Display all active chat rooms for the farmer.
    """
    # Access the list of active rooms from ChatManager
    active_rooms = chat_manager.active_rooms.keys()

    # Filter only rooms related to the farmer
    farmer_rooms = [
        room for room in active_rooms if room.startswith(FARMER_ID)]

    # Prepare a list of clients from the room IDs
    client_ids = [room.split("-")[1] for room in farmer_rooms]

    # Return as a simple HTML response (you can enhance this as needed)
    html_content = "<h1>Active Chats</h1><ul>"
    for client_id in client_ids:
        html_content += f'<li><a href="/chat/{
            client_id}">Chat with {client_id}</a></li>'
    html_content += "</ul>"
    return HTMLResponse(content=html_content)


@app.get("/chats/{room_id}", response_class=HTMLResponse)
async def get_chat_room(request: Request, room_id: str):
    """
    Render the chat page for a specific room.
    The room_id is in the format 'farmer_id-client_id'.
    """
    # Split the room_id into farmer_id and client_id
    try:
        farmer_id, client_id = room_id.split("-")
    except ValueError:
        return HTMLResponse(content="<h1>Invalid room ID format</h1>", status_code=400)

    # Validate the farmer_id and client_id (optional based on your logic)
    if farmer_id != FARMER_ID:
        return HTMLResponse(content="<h1>Farmer not authorized</h1>", status_code=403)

    # Render the chat template
    return templates.TemplateResponse(
        "chat.html", {"request": request, "room_id": room_id,
                      "farmer_id": farmer_id, "client_id": client_id}
    )


@app.websocket("/ws/chat/{room_id}")
async def websocket_endpoint(websocket: WebSocket, room_id: str):
    """
    WebSocket endpoint for chat between farmer and client.
    The room_id is in the format 'farmer_id-client_id'.
    """
    try:
        farmer_id, client_id = room_id.split("-")
    except ValueError:
        await websocket.close(code=400)
        return

    # Connect to the chat room
    await chat_manager.connect(room_id, websocket)
    try:
        while True:
            # Receive a JSON message
            data = await websocket.receive_text()
            message_data = json.loads(data)

            # Validate the sender
            sender = message_data.get("sender")
            message = message_data.get("message")

            if sender not in ["farmer", "client"]:
                await websocket.send_text("Invalid sender type")
                continue

            # Broadcast the message with sender info
            await chat_manager.send_message_to_room(
                room_id, f"{sender.capitalize()} says: {message}"
            )
    except WebSocketDisconnect:
        chat_manager.disconnect(room_id, websocket)
