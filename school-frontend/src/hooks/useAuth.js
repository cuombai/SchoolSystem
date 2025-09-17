import { useContext } from "react";
import { AuthContext } from "../context/AuthProvider";

const useAuth = () => {
  const { auth, login, logout } = useContext(AuthContext);
  return { auth, login, logout };
};

export default useAuth;
