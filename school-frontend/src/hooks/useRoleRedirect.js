import { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import useAuth from "./useAuth";

const useRoleRedirect = () => {
  const { auth } = useAuth();
  const navigate = useNavigate();

  useEffect(() => {
    if (auth.token && auth.role) {
      navigate(`/${auth.role}`);
    }
  }, [auth, navigate]);
};

export default useRoleRedirect;
