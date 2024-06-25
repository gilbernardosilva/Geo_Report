import axios from "axios";
import { useAuth } from "./../hooks/AuthContext";
import { useEffect } from "react";

const instance = axios.create({
    baseURL: process.env.REACT_APP_API_URL,
    headers: { "Content-Type": "application/json" },
});

export function useAxiosWithToken() {
  const { token } = useAuth();

  useEffect(() => {
    const updateInterceptor = () => {
      // Remove any existing interceptor before adding a new one
      instance.interceptors.request.handlers = []; // Clear the handlers array

      // Add the new interceptor
      instance.interceptors.request.use(
        (config) => {
          if (token) {
            config.headers['Authorization'] = `Bearer ${token}`;
          }
          return config;
        },
        (error) => Promise.reject(error)
      );
    };

    updateInterceptor();

  }, [token]);

  return instance;
}

