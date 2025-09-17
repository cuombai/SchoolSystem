import React from "react";
import { useState, useContext } from "react";
import { login } from "../api/auth";
import { AuthContext } from "../context/AuthProvider";
import { useNavigate } from "react-router-dom";
import useRoleRedirect from "../hooks/useRoleRedirect";
import "../styles/form.css";
import { Link } from "react-router-dom";


const Login = () => {
  const { login: setAuth, auth } = useContext(AuthContext);
  const [form, setForm] = useState({ email: "", password: "" });
  const navigate = useNavigate();

  useRoleRedirect(); // ✅ called at top level

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const res = await login(form);
      const { token } = res.data;
      const payload = JSON.parse(atob(token.split(".")[1]));
      setAuth({ token, role: payload.role, userID: payload.userID });
      // useRoleRedirect handles navigation by role
    } catch (err) {
      alert("Login failed");
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <input
        type="email"
        placeholder="Email"
        onChange={(e) => setForm({ ...form, email: e.target.value })}
      />
      <input
        type="password"
        placeholder="Password"
        onChange={(e) => setForm({ ...form, password: e.target.value })}
      />
      <button type="submit">Login</button>

    <Link to="/forgot-password">Forgot password?</Link>

    </form>
  );
};

export default Login;
