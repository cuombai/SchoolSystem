import React from "react";
import { Link } from "react-router-dom";
import "../styles/home.css";
import Navbar from "../components/Navbar";
import Footer from "../components/Footer";

const Home = () => {
  return (
    <>
    <Navbar/>
    <div className="home-container">
      <section className="hero">
        <div className="hero-content">
          <h1>Welcome to Kisumu School System</h1>
          <p>Empowering education through technology.</p>
          <Link to="/login" className="hero-cta">Get Started</Link>
        </div>
      </section>

      <section className="school-info">
        <h2>About Us</h2>
        <p>
          Kisumu School System is a modern platform for managing academic records, attendance, and communication between staff and students. Located in Kisumu County, Kenya, we serve learners with integrity and innovation.
        </p>
        <p><strong>Location:</strong> Kisumu CBD, Opposite Jubilee House</p>
        <p><strong>Contact:</strong> +254 700 000 001 | info@kisumuschool.ac.ke</p>
      </section>

      <section className="login-options">
        <h2>Access Your Dashboard</h2>
        <div className="login-buttons">
          <Link to="/login" className="login-btn">Sign in as Student</Link>
          <Link to="/login" className="login-btn">Sign in as Staff</Link>
        </div>
      </section>
    </div>
    <Footer/>
    </>
  );
};

export default Home;
