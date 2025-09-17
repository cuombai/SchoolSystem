import { useContext } from "react";
import { AuthContext } from "../context/AuthContext";
import { Link } from "react-router-dom";

const Sidebar = () => {
  const { auth } = useContext(AuthContext);

  const links = {
    admin: [
      { label: "Dashboard", path: "/admin" },
      { label: "Users", path: "/admin/users" },
      { label: "Audit Logs", path: "/admin/audit" },
    ],
    teacher: [
      { label: "Attendance", path: "/teacher/attendance" },
      { label: "Grades", path: "/teacher/grades" },
      { label: "Class List", path: "/teacher/class/1" },
    ],
    student: [
      { label: "Performance", path: "/student/performance" },
      { label: "Transcript", path: "/student/transcript" },
      { label: "Fees", path: "/student/fees" },
    ],
  };

  return (
    <aside className="sidebar">
      {links[auth.role]?.map((link) => (
        <Link key={link.path} to={link.path}>{link.label}</Link>
      ))}
    </aside>
  );
};

export default Sidebar;
