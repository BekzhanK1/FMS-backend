from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker

DATABASE_URL = "sqlite:///./test.db"

# Create a SQLite engine
engine = create_engine(DATABASE_URL, connect_args={"check_same_thread": False})

# Session factory for database connections
SessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine)


# Initialize the database
def init_db():
    from models import Base

    Base.metadata.create_all(bind=engine)
