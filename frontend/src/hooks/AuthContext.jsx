import React, { createContext, useState, useContext, useEffect } from "react";
import { jwtDecode } from "jwt-decode";
import PropTypes from 'prop-types';

const AuthContext = createContext({
  token: null,
  isLoggedIn: false,
  userInfo: null,
  userRole: -1,
  login: (token) => { },
  logout: () => { },
  setLoggedIn: (value) => { },
});

export const AuthProvider = ({ children }) => {
  const [token, setToken] = useState(null);
  const [userInfo, setUserInfo] = useState(null);
  const [userRole, setUserRole] = useState(null);

  const [isLoggedIn, setIsLoggedIn] = useState(() => {
    const storedToken = localStorage.getItem("token");
    return !!storedToken;
  });


  useEffect(() => {
    const storedToken = localStorage.getItem("token");

    if (storedToken) {
      try {
        const decodedToken = jwtDecode(storedToken);
        debugger;
        if (decodedToken.eat < Date.now() / 1000) {
          logoutHandler();
        } else {
          setToken(storedToken);
          setUserInfo(decodedToken);
          setUserRole(decodedToken.role);
          setIsLoggedIn(true);
        }
      } catch (error) {
        console.error("Invalid token:", error);
        localStorage.removeItem("token");
        localStorage.removeItem("userInfo");
        setUserRole(-1);
        setIsLoggedIn(false);
      }
    }
  }, []);

  const loginHandler = (token) => {
    setToken(token);
    localStorage.setItem("token", token);
    const decodedToken = jwtDecode(token);
    setUserInfo(decodedToken);
    setUserRole(decodedToken.role);
    localStorage.setItem("userInfo", JSON.stringify(decodedToken));
    setIsLoggedIn(true);
  };

  const logoutHandler = () => {
    debugger;
    setToken(null);
    localStorage.removeItem("token");
    setUserInfo(null);
    localStorage.removeItem("userInfo");
    setUserRole(-1);
    setIsLoggedIn(false);
  };

  const contextValue = {
    token,
    isLoggedIn,
    userInfo,
    userRole,
    login: loginHandler,
    logout: logoutHandler,
    setLoggedIn: setIsLoggedIn,
  };

  return (
    <AuthContext.Provider value={contextValue}>{children}</AuthContext.Provider>
  );
};

AuthProvider.propTypes = {
  children: PropTypes.node.isRequired,
};

export function useAuth() {
  return useContext(AuthContext);
}