import React, { createContext, useState, useContext  } from "react";

const AuthContext = createContext({
  token: null,
  isLoggedIn: false,
  login: (token) => {},
  logout: () => {},
  setLoggedIn: (value) => {},
});

export const AuthProvider = ({ children }) => {
  const [token, setToken] = useState(null);
  const [isLoggedIn, setIsLoggedIn] = useState(false);

  const loginHandler = (token) => {
    setToken(token);
    localStorage.setItem("token", token);
    setIsLoggedIn(true);
  };

  const logoutHandler = () => {
    setToken(null);
    localStorage.removeItem("token");
    isLoggedIn(false);
  };

  const contextValue = {
    token,
    isLoggedIn,
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