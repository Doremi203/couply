import { useNavigate } from 'react-router-dom';

import styles from './splashPage.module.css';

export const SplashPage = () => {
  const navigate = useNavigate();

  setTimeout(() => {
    navigate('/auth');
  }, 4000);

  return (
    <div className={styles.page}>
      <body>
        {/* <img src="logo.png" width="200px" height="150px" className={styles.logo} /> */}

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
