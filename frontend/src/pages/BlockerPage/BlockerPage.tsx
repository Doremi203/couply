import React from 'react';
import { useNavigate } from 'react-router-dom';
import { useSelector } from 'react-redux';
import { RootState } from '../../app/store';
import styles from './blockerPage.module.css';
import { reasonsFromApi } from '../../entities/blocker/constants';

export const BlockerPage: React.FC = () => {
  const navigate = useNavigate();
  const { reasons, message } = useSelector((state: RootState) => state.blocking);

  console.log(reasons);

  const handleLogOut = () => {
    navigate('/auth');
    localStorage.removeItem('token');
  };

  return (
    <div className={styles.container}>
      <img src="pv3.jpg" width="200px" height="150px" alt="Logo" className={styles.logo} />
      <div className={styles.content}>
        <h1 className={styles.title}>Аккаунт заблокирован</h1>
        <p className={styles.message}>К сожалению, ваш аккаунт был заблокирован.</p>
        {reasons.length > 0 && (
          <>
            <p className={styles.message}>Причины блокировки:</p>
            <ul className={styles.reasonsList}>
              {reasons.map((reason, index) => (
                <li key={index} className={styles.reasonItem}>
                  {reasonsFromApi[reason]}
                </li>
              ))}
            </ul>
          </>
        )}
        <button className={styles.logoutButton} onClick={handleLogOut}>
          Выйти из аккаунта
        </button>
      </div>
    </div>
  );
};

export default BlockerPage;
