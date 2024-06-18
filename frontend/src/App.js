import './App.css';
import Login from './components/Login';
import 'bootstrap/dist/css/bootstrap.min.css';
import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';
import './configurations/i18n.js'
import Dashboard from './components/Dashboard/index.jsx';
import { useState } from 'react';



function App() {
  const [token, setToken] = useState(null);

  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={token ? <Dashboard setToken={setToken}/> : <Login setToken={setToken} />} />
        <Route
          path="/dashboard"
          element={token ? <Dashboard setToken={setToken}/> : <Navigate to="/" />}
        />
      </Routes>
    </BrowserRouter>
  );
}

export default App;