from typing import Dict, List
from fastapi import WebSocket


class ChatManager:
    def __init__(self):
        # Dictionary to track chat rooms and their connections
        self.active_rooms: Dict[str, List[WebSocket]] = {}

    async def connect(self, room_id: str, websocket: WebSocket):
        """Add a connection to a chat room."""
        await websocket.accept()
        if room_id not in self.active_rooms:
            self.active_rooms[room_id] = []
        self.active_rooms[room_id].append(websocket)

    def disconnect(self, room_id: str, websocket: WebSocket):
        """Remove a connection from a chat room."""
        if room_id in self.active_rooms:
            self.active_rooms[room_id].remove(websocket)
            if not self.active_rooms[room_id]:  # Clean up empty rooms
                del self.active_rooms[room_id]

    async def send_message_to_room(self, room_id: str, message: str):
        """Broadcast a message to all connections in a chat room."""
        if room_id in self.active_rooms:
            for connection in self.active_rooms[room_id]:
                await connection.send_text(message)

