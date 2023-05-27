import { useEffect, useRef, useState } from 'react';
import axios from 'axios';
import TodoList from '../components/js/TodoList';

const ShowTodoList = () => {
    const [todos, setTodos] = useState([]);
  
    const todoNameRef = useRef();

    useEffect(() => {
      axios
        .get('http://localhost:8080/tasks')
        .then((res) => {
          console.log(res.data);
        })
    })
  
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
        <TodoList todos={todos} toggleTodo={toggleTodo}/>
        <label htmlFor="title">title</label>
        <input type="text" ref={todoNameRef} />
        <button onClick={handleClear}>完了したタスクの削除</button>
        <div>残りのタスク:{todos.filter((todo) => !todo.completed).length}</div>
      </>
      );
  }

  export default ShowTodoList;
