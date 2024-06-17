import './App.css';
import Login from './components/Login';
import 'bootstrap/dist/css/bootstrap.min.css'; 
import { BrowserRouter, Routes, Route } from 'react-router-dom';



function AppRoutes() {
  return (
      <Routes>
          <Route path="/" element={<Login />} />
      </Routes>
  );
}

function App() {
  return (
      <BrowserRouter>
          <AppRoutes />
      </BrowserRouter>
  );
}

export default App;