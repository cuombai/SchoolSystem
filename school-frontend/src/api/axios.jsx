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

api.interceptors.response.use(
  (res) => res,
  (err) => {
    console.error("API Error:", err.response?.data || err.message);
    return Promise.reject(err);
  }
);


export default api;