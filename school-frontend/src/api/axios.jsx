import axios from "axios"

const api = axios.create({
    baseURL: "http://localhost:8080", //connecting to Go- backend
})

//Attach the JWt token generated in backend automatically 
api.interceptors.request.use((config) => {
    const token = localStorage.getItem("token");

    if (token) {
        config.headers.Authorization = `Bearer ${token}`;
    }
    return config
});

//global error handling

api.interceptors.request.use((config) => {
  // Skip attaching token for login or public routes
  if (!config.url.includes("/login")) {
    const token = localStorage.getItem("token");
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
  }
  return config;
});



export default api;