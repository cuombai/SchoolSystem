import React, { useState } from "react";
import { useParams } from "react-router-dom";
import axios from "axios";

const ResetPassword = () => {
  const { token } = useParams();
  const [password, setPassword] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await axios.post(`/api/reset-password/${token}`, { password });
      alert("Password updated. You can now log in.");
    } catch (err) {
      alert("Reset failed.");
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <input
        type="password"
        placeholder="New password"
        onChange={(e) => setPassword(e.target.value)}
      />
      <button type="submit">Reset Password</button>
    </form>
  );
};

export default ResetPassword;
