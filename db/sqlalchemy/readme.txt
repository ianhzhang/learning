https://www.pythoncentral.io/introductory-tutorial-python-sqlalchemy/

pip install sqlalchemy
pip install sqlalchemy_declarative


'''
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy import Column, ForeignKey, Integer, String
from sqlalchemy.orm import relationship, scoped_session, sessionmaker

engine = create_engine('db_location')
Base = declarative_base()
Base.metadata.create_all(engine)

or
engine = create_engine('db_location')
meta_data = MetaData(bind=engine)
db_session = scoped_session(
    sessionmaker(autocommit=False,
                 autoflush=False,
                 bind=engine)
)
Base = declarative_base(metadata=meta_data, bind=engine)
meta_data.create_all()

'''


class Person(Base):
    __tablename__ = 'person'

    id = Column(Integer, primary_key=True, autoincrement=True)
    name = Column(String(250), nullable=False)