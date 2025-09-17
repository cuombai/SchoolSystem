import { useContext } from "react";
import { AuthContext } from "../context/AuthProvider";
import { Navigate } from "react-router-dom";

const ProtectedRoute = ({ children, role }) => {
  const { auth } = useContext(AuthContext);

  if (!auth.token) return <Navigate to="/login" />;
  if (role && auth.role !== role) return <Navigate to="/unauthorized" />;

  return children;
};

export default ProtectedRoute;
