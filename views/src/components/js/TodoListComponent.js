import React from 'react';
import Todo from './Todo';

const TodoListComponent = ({todos, toggleTodo}) => {
  return todos.map((todo) => <Todo todo={todo} toggleTodo={toggleTodo} key={todo.id}/>)
}

export default TodoListComponent;
