import grpc
import logging
from concurrent import futures
import market_pb2_grpc
from service import CategoryService, ProductService
from database import init_db


# Configure logging
logging.basicConfig(
    level=logging.INFO,  # Set the logging level
    format="%(asctime)s - %(levelname)s - %(message)s",  # Customize the log format
    handlers=[
        logging.StreamHandler(),  # Output to the console
    ],
)


def start_grpc_server():
    # Initialize the database
    logging.info("Initializing the database...")
    init_db()

    # Create the gRPC server
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))

    # Register the services
    market_pb2_grpc.add_CategoryServiceServicer_to_server(CategoryService(), server)
    market_pb2_grpc.add_ProductServiceServicer_to_server(ProductService(), server)

    # Bind the server to a port
    port = 50051
    server.add_insecure_port(f"[::]:{port}")
    logging.info(f"gRPC server is starting on port {port}...")

    # Start the server
    server.start()
    logging.info(f"gRPC server is now running on port {port}.")

    # Keep the server running
    server.wait_for_termination()


if __name__ == "__main__":
    logging.info("Starting market-service...")
    start_grpc_server()
