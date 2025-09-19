import api from "./axios";

export const login = (credentials) => api.post("/login", credentials);

export const signup = (userData) => api.post("/signup", userData);
