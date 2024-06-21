import React, { createContext, useState, useContext } from "react";
import { jwtDecode } from "jwt-decode";

const AuthContext = createContext({
  token: null,
  isLoggedIn: false,
  userInfo: null,
  login: (token) => {},
  logout: () => {},
  setLoggedIn: (value) => {},
});

export const AuthProvider = ({ children }) => {
  const [token, setToken] = useState(null);
  const [userInfo, setUserInfo] = useState(null);
  const [isLoggedIn, setIsLoggedIn] = useState(() => {
    const storedToken = localStorage.getItem("token");
    return !!storedToken;  
});
    
  const loginHandler = (token) => {
    setToken(token);
    localStorage.setItem("token", token);
    const decodedToken = jwtDecode(token);
    setUserInfo(decodedToken); 
    localStorage.setItem("userInfo", JSON.stringify(decodedToken));
    setIsLoggedIn(true);
  };

  const logoutHandler = () => {
    setToken(null);
    localStorage.removeItem("token");
    setUserInfo(null);
    localStorage.removeItem("userInfo");
    setIsLoggedIn(false);
  };

  const contextValue = {
    token,
    isLoggedIn,
    userInfo,
    login: loginHandler,
    logout: logoutHandler,
    setLoggedIn: setIsLoggedIn,
  };

  return (
    <AuthContext.Provider value={contextValue}>{children}</AuthContext.Provider>
  );
};

export function useAuth() {
  return useContext(AuthContext);
}