import { useNavigate } from 'react-router-dom';

import styles from './splashPage.module.css';

export const SplashPage = () => {
  const navigate = useNavigate();

  setTimeout(() => {
    // Check if user is authenticated (token exists in localStorage)
    const token = localStorage.getItem('token');

    if (token) {
      // If authenticated, redirect to home page
      navigate('/home');
    } else {
      // If not authenticated, redirect to auth page
      navigate('/auth');
    }
  }, 4000);

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
