import { useState, useRef } from 'react';
import { v4 as uuidv4 } from 'uuid';
import TodoListComponent from '../components/js/TodoListComponent';

const TodoList = () => {
    const [todos, setTodos] = useState([]);
  
    const todoNameRef = useRef();
    
    const handleAddTodo = () => {
      const name = todoNameRef.current.value;
      setTodos((prevTodos) => {
        return [...prevTodos, { id: uuidv4(), name: name, completed: false }]
      })
    };
  
    const toggleTodo = (id) => {
      const newTodos = [...todos];
      const todo = newTodos.find((todo) => todo.id === id);
      todo.completed = !todo.completed;
      setTodos(newTodos);
    };
  
    const handleClear = () => {
      const newTodos = todos.filter((todo) => !todo.completed);
      setTodos(newTodos)
    }
  
    return (
      <>
        <TodoListComponent todos={todos} toggleTodo={toggleTodo}/>
        <label htmlFor="title">title</label>
        <input type="text" ref={todoNameRef} />
        <button onClick={handleAddTodo}>タスクを追加する</button>
        <button onClick={handleClear}>完了したタスクの削除</button>
        <div>残りのタスク:{todos.filter((todo) => !todo.completed).length}</div>
      </>
      );
  }

  export default TodoList;
