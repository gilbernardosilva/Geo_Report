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
    // Function to update the interceptor
    const updateInterceptor = () => {
      // Remove the previous interceptor if it exists
      instance.interceptors.request.eject(instance.interceptors.request.handlers.length - 1);
      // Add the new interceptor
      instance.interceptors.request.use(
        (config) => {
          if (token) {
            config.headers["Authorization"] = `Bearer ${token}`;
          }
          return config;
        },
        (error) => Promise.reject(error)
      );
    };

    // Update the interceptor whenever the token changes
    updateInterceptor();

    return () => {
      // Remove the last interceptor on unmount
      instance.interceptors.request.eject(instance.interceptors.request.handlers.length - 1);
    };
  }, [token]);

  return instance;
}

