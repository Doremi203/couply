import { useEffect } from 'react';
import { useSelector } from 'react-redux';
import { useNavigate } from 'react-router-dom';

import { RootState } from '../../../../app/store';
import { useGeolocation } from '../../../../shared/lib/hooks/useGeolocation';

import styles from './splashPage.module.css';

export const SplashPage = () => {
  const navigate = useNavigate();
  const { isBlocked } = useSelector((state: RootState) => state.blocking);

  const { updateUserLocation } = useGeolocation();

  // Get user's geolocation and update in backend if user is authorized
  useEffect(() => {
    const getUserLocation = async () => {
      const token = localStorage.getItem('token');

      // Only proceed if user is authorized and userId exists
      if (token) {
        // Try to update user location
        try {
          await updateUserLocation();
        } catch (error) {
          console.error('Failed to update user location:', error);
        }
      }
    };

    getUserLocation();
  }, [updateUserLocation]);

  // Navigate after splash screen
  useEffect(() => {
    const timer = setTimeout(() => {
      const token = localStorage.getItem('token');

      if (token) {
        if (isBlocked) {
          navigate('/blocked');
        } else {
          navigate('/home');
        }
      } else {
        navigate('/auth');
      }
    }, 4000);

    return () => clearTimeout(timer);
  }, [navigate, isBlocked]);

  return (
    <div className={styles.page}>
      <body>
        <div className={styles.loader} />
        <section>
          <div className={styles.content}>
            <h2>couply</h2>
            <h2>couply</h2>
          </div>
        </section>
      </body>
    </div>
  );
};

export default SplashPage;
