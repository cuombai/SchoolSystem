import React, { useContext } from "react";
import { AuthContext } from "../context/AuthProvider";
import { Link } from "react-router-dom";
import logo from "../assets/react.svg"; // Replace with your actual logo
import "../styles/Navbar.css";

const Navbar = () => {
  const { auth, logout } = useContext(AuthContext);

  return (
    <nav className="navbar">
      <div className="logo-container">
        <img src={logo} alt="School Logo" className="logo" />
        <h2 className="title">Kibali School</h2>
      </div>
      <ul className="nav-links">
        <li><Link to="/" className="link">Home</Link></li>
        <li><Link to="/contact" className="link">Contact Us</Link></li>
        <li><Link to="/about" className="link">About Us</Link></li>
        <li><Link to="/login" className="button">Login</Link></li>
        <li><Link to="/signup" className="button-outline">Sign Up</Link></li>
      </ul>
    </nav>
  );
};

export default Navbar;
