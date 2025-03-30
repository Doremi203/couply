
import { Link, useLocation } from "react-router-dom";
import HomeOutlinedIcon from "@mui/icons-material/HomeOutlined";
import FavoriteBorderIcon from "@mui/icons-material/FavoriteBorder";
import PermIdentityOutlinedIcon from "@mui/icons-material/PermIdentityOutlined";
import "./navBar.css";

export const NavBar = () => {
  const location = useLocation();
  const currentPath = location.pathname;

  return (
    <div className="navBarContainer">
      <div className="navBarContent">
        <Link
          to="/home"
          className={`navItem ${currentPath === "/home" ? "active" : ""}`}
        >
          <HomeOutlinedIcon style={{ color: currentPath === "/home" ? "#161F65" : "inherit" }} />
        </Link>
        <Link
          to="/likes"
          className={`navItem ${currentPath === "/likes" ? "active" : ""}`}
        >
          <FavoriteBorderIcon style={{ color: currentPath === "/likes" ? "#161F65" : "inherit" }} />
        </Link>
        <Link
          to="/profile"
          className={`navItem ${currentPath === "/profile" ? "active" : ""}`}
        >
          <PermIdentityOutlinedIcon style={{ color: currentPath === "/profile" ? "#161F65" : "inherit" }} />
        </Link>
      </div>
    </div>
  );
}

export default NavBar;