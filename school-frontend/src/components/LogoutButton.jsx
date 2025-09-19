// import { useNavigate, useNavigation } from "react-router-dom";
// import {useAuth} from "../hooks/useAuth";

// const LogoutButton = () => {
//     const {logout} = useAuth();
//     const navigate = useNavigate();

//     const handleLogout = () => {
//         logout();//Clear auth context
//         navigate("/"); //redirect to school homepage
//     };

//     return (
//         <button onClick={handleLogout} style={{
//             marginTop: "20px",
//             padding: "8px 16px",
//             backgroundColor: "#333",
//             color: "#fff",
//             border: "none",
//             borderRadius: "4px",
//             cursor: "pointer"
//         }}>
//             Logout
//         </button>
//     )
// }

// export default LogoutButton;

// LogoutButton.jsx
import React from "react";
import { useContext } from "react";
import useAuth from "../hooks/useAuth";
import { useNavigate } from "react-router-dom";



const LogoutButton = () => {
  const { login: setAuth } = useAuth();
  const navigate = useNavigate();

  const {logout} = useAuth();
  const handleLogout = () => {
    logout(); // Clear auth state
    navigate("/"); // Redirect to homepage
  };

  return (
    <button onClick={handleLogout} style={{
         marginTop: "20px",
            padding: "8px 16px",
            backgroundColor: "#333",
            color: "#fff",
            border: "none",
            borderRadius: "4px",
            cursor: "pointer"
    }}>
      Logout
    </button>
  );
};

export default LogoutButton;
