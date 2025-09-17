import { useContext } from "react";
import { AuthContext } from "../context/AuthContext";

const Navbar = () => {
  const { auth, logout } = useContext(AuthContext);

  return (
    <nav className="navbar">
      <span className="logo">SchoolSystem</span>
      <span className="role">{auth.role?.toUpperCase()}</span>
      <button onClick={logout}>Logout</button>
    </nav>
  );
};

export default Navbar;
