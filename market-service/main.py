import grpc
from concurrent import futures
import market_pb2_grpc
from service import CategoryService, ProductService
from database import init_db


def start_grpc_server():
    # Initialize the database
    init_db()

    # Create the gRPC server
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))

    # Register the CategoryService
    market_pb2_grpc.add_CategoryServiceServicer_to_server(CategoryService(), server)
    market_pb2_grpc.add_ProductServiceServicer_to_server(ProductService(), server)

    # Bind the server to a port
    server.add_insecure_port("[::]:50051")
    print("gRPC server running on port 50051...")

    # Start the server
    server.start()
    server.wait_for_termination()


if __name__ == "__main__":
    start_grpc_server()
