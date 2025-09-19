import React from "react";
import { Routes, Route } from "react-router-dom";
import Login from "../pages/Login";
import Layout from "../components/Layout";
import ProtectedRoute from "../components/ProtectedRoute";

// Admin pages
import Dashboard from "../pages/Admin/Dashboard";
import Users from "../pages/Admin/Users";
import AuditLogs from "../pages/Admin/AuditLogs";

// Teacher pages
import Attendance from "../pages/Teacher/Attendance";
import Grades from "../pages/Teacher/Grades";
import ClassList from "../pages/Teacher/ClassList";

// Student pages
import Performance from "../pages/Student/Performance";
import Transcript from "../pages/Student/Transcript";
import Fees from "../pages/Student/Fees";
import ForgotPassword from "../pages/ForgotPassword";
import ResetPassword from "../pages/ResetPassword";
import Home from "../pages/Home";
import AdminLanding from "../pages/Admin/Landing";
import StudentLanding from "../pages/Student/Landing";
import TeacherLanding from "../pages/Teacher/Landing";
import ContactUs from "../components/ContactUs";
import AboutUs from "../components/AboutUs";
import SignUp from "../components/SignUp";





const AppRouter = () => (
  <Routes>
    <Route path="/" element={<Home/>}/>
    {/* <Route path="/home" element={<HomePage />} /> */}

    {/* Public route */}
    <Route path="/login" element={<Login />} />
    <Route path="/contact" element={<ContactUs />} />
    <Route path="/about" element={<AboutUs />} />
    <Route path="/signup" element={<SignUp />} />

    
    
    {/* protected Routes */}
    <Route path= "/admin" element={<AdminLanding/>}/>
    <Route path= "/teacher" element={<TeacherLanding/>}/>
    <Route path= "/student" element={<StudentLanding/>}/>
    {/* Protected layout wrapper */}
    <Route element={<ProtectedRoute><Layout /></ProtectedRoute>}>
      {/* Admin routes */}
      <Route path="/admin" element={<Dashboard />} />
      <Route path="/admin/users" element={<Users />} />
      <Route path="/admin/audit" element={<AuditLogs />} />

      {/* Teacher routes */}
      <Route path="/teacher/attendance" element={<Attendance />} />
      <Route path="/teacher/grades" element={<Grades />} />
      <Route path="/teacher/class/:id" element={<ClassList />} />

      {/* Student routes */}
      <Route path="/student/performance" element={<Performance />} />
      <Route path="/student/transcript" element={<Transcript />} />
      <Route path="/student/fees" element={<Fees />} />
    </Route>

    {/* Optional: fallback route */}
    <Route path="*" element={<Login />} />
    {/* password reset */}
    <Route path="/forgot-password" element={<ForgotPassword />} />
    <Route path="/reset-password/:token" element={<ResetPassword />} />
    
  </Routes>
);

export default AppRouter;
