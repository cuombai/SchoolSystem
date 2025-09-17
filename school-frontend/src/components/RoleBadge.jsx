import { useContext } from "react";
import { AuthContext } from "../context/AuthContext";
import "../styles/role.css";

const RoleBadge = () => {
  const { auth } = useContext(AuthContext);

  const colors = {
    admin: "bg-red-500",
    teacher: "bg-blue-500",
    student: "bg-green-500",
  };

  return (
    <span className={`role-badge ${colors[auth.role]} text-white px-2 py-1 rounded`}>
      {auth.role}
    </span>
  );
};

export default RoleBadge;
