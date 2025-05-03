import FavoriteBorderIcon from '@mui/icons-material/FavoriteBorder';
import HomeOutlinedIcon from '@mui/icons-material/HomeOutlined';
import PermIdentityOutlinedIcon from '@mui/icons-material/PermIdentityOutlined';
import { Link, useLocation } from 'react-router-dom';

import './navBar.css';

export const NavBar = () => {
  const location = useLocation();
  const currentPath = location.pathname;

  return (
    <div className="navBarContainer">
      <div className="navBarContent">
        <Link to="/home" className={`navItem ${currentPath === '/home' ? 'active' : ''}`}>
          <HomeOutlinedIcon
            style={{
              color:
                currentPath === '/home' ? 'var(--primary-color)' : 'var(--secondary-text-color)',
              width: '1.6rem',
              height: '1.6rem',
            }}
          />
        </Link>
        <Link to="/likes" className={`navItem ${currentPath === '/likes' ? 'active' : ''}`}>
          <FavoriteBorderIcon
            style={{
              color:
                currentPath === '/likes' ? 'var(--primary-color)' : 'var(--secondary-text-color)',
              width: '1.6rem',
              height: '1.6rem',
            }}
          />
        </Link>
        <Link to="/profile" className={`navItem ${currentPath === '/profile' ? 'active' : ''}`}>
          <PermIdentityOutlinedIcon
            style={{
              color:
                currentPath === '/profile' ? 'var(--primary-color)' : 'var(--secondary-text-color)',
              width: '1.6rem',
              height: '1.6rem',
            }}
          />
        </Link>
      </div>
    </div>
  );
};

export default NavBar;
