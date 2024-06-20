import './App.css';
import Login from './components/Login';
import 'bootstrap/dist/css/bootstrap.min.css';
import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';
import './configurations/i18n.js'
import Dashboard from './components/Dashboard/index.jsx';
import { AuthProvider, useAuth } from './hooks/AuthContext.jsx';



function App() {
  return (
    <AuthProvider>
      <BrowserRouter>
        <AppRoutes />
      </BrowserRouter>
    </AuthProvider>
  );
}

function AppRoutes() {
  const { isLoggedIn } = useAuth();

  return (
    <Routes>
      <Route path="/" element={!isLoggedIn ? <Login /> : <Navigate to="/dashboard" />} />
      <Route path="/dashboard" element={isLoggedIn ? <Dashboard /> : <Navigate to="/" />} />
    </Routes>
  );
}

export default App;