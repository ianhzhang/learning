https://www.youtube.com/watch?v=p7VujaALaGQ

Type/Objects: Sort of like serializers for Graphql.
Schema: The container of all the graphql types at your disposal.
Resolver: The function that is executed to give back the data for a single attribute of a type/object.
Query: What you use to get o set data using graphql.
Mutations: Subset of queries that is used to change/modify data.

The ORM Models
* All models are SQLAlchemy Models
* They inherit from Base.
* They also inherit from PrimaryKeyMixin.


class PrimaryKeyIdMixin:
    id = Column(Integer, primary_key=True, autoincrement = True)

class User(Base, PrimaryKeyIdMixin):
    __tablename__ = "users"

    username = Column(String(255))
    email = Column(String(length=255))
    todos = relationship('Todo, back_populates='user')

class Todo(Base, PrimaryKeyIdMixin):
    __tablename__ = 'todos'
    title = Column(String(length=255))

    user_id = Column(Integer, ForeignKey('users.id))
    user = relationship('User', back_populates='todos')

    items = relationship('TodoItem', back='todo')
    meta_info = Column(JSONB)

class TodoItem(Base, PrimaryKeyIdMixin):
    __tablename__ = 'todos_items'

    description = Column(String)
    is_done = Column(Boolean)
    priority  = Column(Integer)
    todo_id = Column(Integer, ForeignKey('todos.id'))
    todo = relationship('Todo', back_populates='items')


class TodoObject(SQLAlchemyObjectType):
    meta_info = graphene.Field(JSONString)

    def resolve_meta_info(self, args, context, info):
        return self.meta_info

    class Meta:
        model = Todo

    
