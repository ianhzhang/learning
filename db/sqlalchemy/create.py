
# https://docs.sqlalchemy.org/en/latest/orm/tutorial.html

import os
import sys

from sqlalchemy import create_engine
from declarative import Base            # This has to be here.   Otherwise, cannot find Person and address table
# declarative is our own definition.

engine = create_engine('sqlite:///sqlalchemy_example.db')
Base.metadata.create_all(engine)

