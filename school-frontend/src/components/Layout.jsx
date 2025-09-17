import { useContext } from "react";
import { Outlet } from "react-router-dom";
import Navbar from "./Navbar";
import Sidebar from "./Sidebar";
import { AuthContext } from "../context/AuthProvider";
import "../styles/layout.css"; 


const Layout = () => {
  const { auth } = useContext(AuthContext);

  if (!auth.token) return null; // or redirect to login

  return (
    <div className="layout">
      <Navbar />
      <div className="main-content">
        <Sidebar />
        <div className="page-content">
          <Outlet />
        </div>
      </div>
    </div>
  );
};

export default Layout;
