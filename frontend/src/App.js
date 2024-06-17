import './App.css';
import Login from './components/Login';
import 'bootstrap/dist/css/bootstrap.min.css'; 
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import './configurations/i18n.js'
import Dashboard from './components/Dashboard/index.jsx';


function AppRoutes() {
  return (
      <Routes>
          <Route path="/" element={<Login />} />
          <Route path="/dashboard" element={<Dashboard />} />
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