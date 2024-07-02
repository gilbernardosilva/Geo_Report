import './App.css';
import Login from './components/Login';
import 'bootstrap/dist/css/bootstrap.min.css';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import './configurations/i18n.js'
import Dashboard from './components/Dashboard/index.jsx';
import { AuthProvider, useAuth } from './hooks/AuthContext.jsx';
import AuthorityDashboard from './components/AuthorityDashboard/index.jsx';
import AdminDashboard from './components/Admin/AdminDashboard/index.jsx';
import Profile from './components/Profile/index.jsx';
import MyIssues from './components/Issues/index.jsx';
import { ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import { useLocation, Navigate } from 'react-router-dom';
import UserManagement from './components/Admin/Users/index.jsx';
import ReportTypePage from './components/Admin/Types/index.jsx';
import AdminIssues from './components/Admin/Issues/index.jsx';
import AdminArea from './components/Admin/Areas/index.jsx';
import AuthorityIssues from './components/AuthorityIssues/index.jsx'




function App() {
  return (
    <AuthProvider>
      <BrowserRouter>
        <AppRoutes />
      </BrowserRouter>
      <ToastContainer />
    </AuthProvider>
  );
}


function AppRoutes() {
  const { isLoggedIn, userRole } = useAuth();
  const location = useLocation();
  return (
    <Routes>
      <Route path="/" element={!isLoggedIn ? <Login /> : (
        <Navigate to={location.pathname} replace />
      )} />
      <Route path="/dashboard" element={isLoggedIn && userRole === 0 ? <Dashboard /> : <Login />} />
      <Route path="/authority" element={isLoggedIn && userRole === 1 || userRole === 2 ? <AuthorityDashboard /> : <Login />} />
      <Route path="/admin" element={isLoggedIn && userRole === 2 ? <AdminDashboard /> : <Login />} />
      <Route path="/profile" element={isLoggedIn ? <Profile /> : <Login />} />
      <Route path="/issues" element={isLoggedIn ? <MyIssues /> : <Login />} />
      <Route path="/authority/issues" element={isLoggedIn && userRole === 1 ? <AuthorityIssues /> : <Login />} />
      <Route path="/admin/users" element={isLoggedIn && userRole === 2 ? <UserManagement /> : <Login />} />
      <Route path="/admin/types" element={isLoggedIn && userRole === 2 ? <ReportTypePage /> : <Login />} />
      <Route path="/admin/issues" element={isLoggedIn && userRole === 2 ? <AdminIssues /> : <Login />} />
      <Route path="/admin/areas" element={isLoggedIn && userRole === 2 ? <AdminArea /> : <Login />} />

    </Routes>
  );
}

export default App;