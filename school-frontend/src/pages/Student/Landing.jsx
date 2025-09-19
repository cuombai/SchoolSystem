// const StudentLanding = () => {
//     return <h1>Login Successful</h1>
// }

// export default StudentLanding
import React from "react";
import RoleBadge from "../../components/RoleBadge";
import LogoutButton from "../../components/LogoutButton";
import useAuth from "../../hooks/useAuth";


const StudentLanding = () => {
    const {auth} = useAuth();

  return (
    <div style={{
        padding: "2rem",
        textAlign: "center"
    }}>
        <h1>Login Successful</h1>
        <p>Welcome, <RoleBadge role={auth.role}/></p>
        <LogoutButton/>
    </div>
  );
};

export default StudentLanding;
