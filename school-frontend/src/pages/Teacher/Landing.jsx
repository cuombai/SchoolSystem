// const TeacherLanding = () => {
//     return <h1>Login successful</h1>
// }

// export default TeacherLanding;
import React from "react";
import RoleBadge from "../../components/RoleBadge";
import LogoutButton from "../../components/LogoutButton";
import useAuth from "../../hooks/useAuth";

const TeacherLanding = () => {
  const { auth } = useAuth();

  return (
    <div style={{ padding: "2rem", textAlign: "center" }}>
      <h1>Login successful</h1>
      <p>Welcome, <RoleBadge role={auth.role} /></p>
      <LogoutButton />
    </div>
  );
};

export default TeacherLanding;
