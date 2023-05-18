import './App.css';
import TodoList from './pages/TodoList';
import { BrowserRouter, Routes, Route } from 'react-router-dom';

function App() {
  return (
      <BrowserRouter>
        <Routes>
          <Route path='/home' element={<TodoList />}>
          </Route>
        </Routes>
      </BrowserRouter>
    );
}

export default App;
