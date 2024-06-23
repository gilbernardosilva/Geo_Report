import './App.css';
import Login from './components/Login';
import 'bootstrap/dist/css/bootstrap.min.css';
import { BrowserRouter, Routes, Route  } from 'react-router-dom';
import './configurations/i18n.js'
import Dashboard from './components/Dashboard/index.jsx';
import { AuthProvider, useAuth } from './hooks/AuthContext.jsx';
import AuthorityDashboard from './components/AuthorityDashboard/index.jsx';
import Profile from './components/Profile/index.jsx';
import MyIssues from './components/Issues/index.jsx';




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
      <Route path="/" element={!isLoggedIn ? < Login/> : <Dashboard />} />
      <Route path="/dashboard" element={isLoggedIn ? <Dashboard /> : <Login />} />
      <Route path="/authority" element={isLoggedIn ? <AuthorityDashboard /> : <Login />} />
      <Route path="/profile" element={isLoggedIn ? <Profile /> : <Login />} />
      <Route path="/issues" element={isLoggedIn ? <MyIssues /> : <Login />} />
    </Routes>
  );
}

export default App;