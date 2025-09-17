import React, { useState } from "react";
import axios from "axios";

const ForgotPassword = () => {
  const [email, setEmail] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await axios.post("/api/forgot-password", { email });
      alert("Password reset link sent to your email.");
    } catch (err) {
      alert("Error sending reset link.");
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <input
        type="email"
        placeholder="Enter your email"
        onChange={(e) => setEmail(e.target.value)}
      />
      <button type="submit">Send Reset Link</button>
    </form>
  );
};

export default ForgotPassword;
