import FavoriteBorderIcon from '@mui/icons-material/FavoriteBorder';
import HomeOutlinedIcon from '@mui/icons-material/HomeOutlined';
import PermIdentityOutlinedIcon from '@mui/icons-material/PermIdentityOutlined';
import { Link, useLocation } from 'react-router-dom';
import { useSelector } from 'react-redux';

import { selectIsShowingAd } from '../../../entities/profile/model/profileSlice';
import './navBar.css';

export const NavBar = () => {
  const location = useLocation();
  const currentPath = location.pathname;
  const isShowingAd = useSelector(selectIsShowingAd);

  return (
    <div className={`navBarContainer ${isShowingAd ? 'disabled' : ''}`}>
      <div className="navBarContent">
        <Link
          to={isShowingAd ? '#' : '/home'}
          className={`navItem ${currentPath === '/home' ? 'active' : ''} ${isShowingAd ? 'disabled' : ''}`}
          onClick={e => isShowingAd && e.preventDefault()}
        >
          <HomeOutlinedIcon
            style={{
              color:
                currentPath === '/home' ? 'var(--primary-color)' : 'var(--secondary-text-color)',
              width: '1.8rem',
              height: '1.8rem',
              opacity: isShowingAd ? 0.5 : 1,
            }}
          />
        </Link>
        <Link
          to={isShowingAd ? '#' : '/likes'}
          className={`navItem ${currentPath === '/likes' ? 'active' : ''} ${isShowingAd ? 'disabled' : ''}`}
          onClick={e => isShowingAd && e.preventDefault()}
        >
          <FavoriteBorderIcon
            style={{
              color:
                currentPath === '/likes' ? 'var(--primary-color)' : 'var(--secondary-text-color)',
              width: '1.8rem',
              height: '1.8rem',
              opacity: isShowingAd ? 0.5 : 1,
            }}
          />
        </Link>
        <Link
          to={isShowingAd ? '#' : '/profile'}
          className={`navItem ${currentPath === '/profile' ? 'active' : ''} ${isShowingAd ? 'disabled' : ''}`}
          onClick={e => isShowingAd && e.preventDefault()}
        >
          <PermIdentityOutlinedIcon
            style={{
              color:
                currentPath === '/profile' ? 'var(--primary-color)' : 'var(--secondary-text-color)',
              width: '1.8rem',
              height: '1.8rem',
              opacity: isShowingAd ? 0.5 : 1,
            }}
          />
        </Link>
      </div>
    </div>
  );
};

export default NavBar;
