import React from "react";
import { useContext } from "react";
import { AuthContext } from "../context/AuthProvider";
import "../styles/role.css";

// const RoleBadge = () => {
//   const { auth } = useContext(AuthContext);

//   const colors = {
//     admin: "bg-red-500",
//     teacher: "bg-blue-500",
//     student: "bg-green-500",
//   };

//   return (
//     <span className={`role-badge ${colors[auth.role]} text-white px-2 py-1 rounded`}>
//       {auth.role}
//     </span>
//   );
// };

const RoleBadge = ({role}) => {
  const colorMap = {
    admin: "red",
    teacher: "blue",
    student: "green",
  };

  return (
    <span style={{
      backgroundColor: colorMap[role] || "gray",
      color: "white",
      padding: "4px 8px",
      borderRadius: "4px",
      fontweight: "bold",
    }}>
      {role.toUpperCase()}
    </span>
  )
}


export default RoleBadge;
