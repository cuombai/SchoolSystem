import React, { useState } from "react";
import RoleBadge from "../../components/RoleBadge";
import LogoutButton from "../../components/LogoutButton";
import useAuth from "../../hooks/useAuth";
import logo from "../../assets/react.svg"; // Replace with actual logo
import "../../styles/studentLanding.css";

const StudentLanding = () => {
  const { auth } = useAuth();
  const [activeTab, setActiveTab] = useState("dashboard");

  return (
    <div className="student-layout">
      {/* Vertical Sidebar */}
      <aside className="sidebar">
        <img src={logo} alt="School Logo" className="sidebar-logo" />
        <div className="user-info">
          <p className="user-name">{auth?.name || "Student Name"}</p>
          <RoleBadge role={auth.role} />
        </div>
        <LogoutButton />
      </aside>

      {/* Main Content */}
      <main className="main-content">
        {/* Top Horizontal Navbar */}
        <nav className="top-nav">
          <button onClick={() => setActiveTab("dashboard")} className={activeTab === "dashboard" ? "active" : ""}>Dashboard</button>
          <button onClick={() => setActiveTab("fees")} className={activeTab === "fees" ? "active" : ""}>Fees</button>
          <button onClick={() => setActiveTab("funzone")} className={activeTab === "funzone" ? "active" : ""}>Funzone</button>
        </nav>

        {/* Dynamic Content */}
        <section className="content-area">
          {activeTab === "dashboard" && (
            <div className="dashboard">
              <h2>Academic Overview</h2>
              <p>📊 Graphs and stats about your school performance will appear here.</p>
              {/* You can embed charts or progress bars here */}
            </div>
          )}

          {activeTab === "fees" && (
            <div className="fees">
              <h2>Fee Statement</h2>
              <p><strong>Amount Paid:</strong> KES 45,000</p>
              <p><strong>Balance:</strong> KES 15,000</p>
              <p><strong>Status:</strong> Partially Paid</p>
              {/* Later: fetch from backend /internal/finance */}
            </div>
          )}

          {activeTab === "funzone" && (
            <div className="funzone">
              <h2>🎉 Funzone</h2>
              <p>Games, quizzes, and student activities coming soon!</p>
            </div>
          )}
        </section>
      </main>
    </div>
  );
};

export default StudentLanding;
