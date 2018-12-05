
from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker
from declarative import Person, Base, Address

engine = create_engine('sqlite:///sqlalchemy_example.db')
Base.metadata.bind = engine

DBSession = sessionmaker()
DBSession.bind = engine
session = DBSession()

# Make a query to find all Persons in the database
print "----get  all()======================="
people = session.query(Person).all()
print type(people)
print "len:", len(people)

print ("first() person name ==========================")
# Return the first Person from all Persons in the database
person = session.query(Person).first()
print person.name

# Find all Address whose person field is pointing to the person object
print "find throught relation ======================== all"
people2 = session.query(Address).filter(Address.person == person).all()
print type(people2)
# Retrieve one Address whose person field is point to the person object
print "find throught relation ======================== one"
person2 = session.query(Address).filter(Address.person == person).one()
print type(person2)

print "find address by relation=============== one"
address = session.query(Address).filter(Address.person == person).one()
print address.post_code