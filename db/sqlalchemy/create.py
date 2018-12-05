import os
import sys

from sqlalchemy.ext.declarative import declarative_base

from sqlalchemy import create_engine
from declarative import Base            # This has to be here.   Otherwise, cannot find Person and address table
# Base = declarative_base()             # This cannot be here.
 
# Create an engine that stores data in the local directory's
# sqlalchemy_example.db file.
engine = create_engine('sqlite:///sqlalchemy_example.db')

# Create all tables in the engine. This is equivalent to "Create Table"
# statements in raw SQL.
Base.metadata.create_all(engine)

