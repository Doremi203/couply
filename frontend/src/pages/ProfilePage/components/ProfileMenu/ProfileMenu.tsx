import DeleteForeverIcon from '@mui/icons-material/DeleteForever';
import LogoutIcon from '@mui/icons-material/Logout';
import React from 'react';
import { useNavigate } from 'react-router-dom';

import { useDeleteUserMutation } from '../../../../entities/user';

import styles from './profileMenu.module.css';

interface ProfileMenuProps {
  onEditProfileClick: () => void;
  onSettingsClick: () => void;
}

export const ProfileMenu: React.FC<ProfileMenuProps> = ({
  onEditProfileClick,
  onSettingsClick,
}) => {
  const naviagate = useNavigate();

  const handleLogOut = () => {
    naviagate('/auth');
    localStorage.removeItem('token');
  };

  const [deleteUser] = useDeleteUserMutation();

  const handleDeleteAccount = async () => {
    //@ts-ignore
    await deleteUser();
    naviagate('/auth');
    localStorage.removeItem('token');
  };
  return (
    <div className={styles.menuContainer}>
      <div className={styles.menuItem} onClick={onEditProfileClick}>
        <div className={`${styles.iconContainer} ${styles.profileIcon}`}>
          <svg width="24" height="24" viewBox="0 0 24 24" fill="#ff66a3">
            <path d="M12 12c2.21 0 4-1.79 4-4s-1.79-4-4-4-4 1.79-4 4 1.79 4 4 4zm0 2c-2.67 0-8 1.34-8 4v2h16v-2c0-2.66-5.33-4-8-4z" />
          </svg>
        </div>
        <div className={styles.menuText}>Редактировать профиль</div>
        <div className={styles.arrowIcon}>
          <svg width="24" height="24" viewBox="0 0 24 24" fill="currentColor">
            <path d="M8.59 16.59L13.17 12 8.59 7.41 10 6l6 6-6 6-1.41-1.41z" />
          </svg>
        </div>
      </div>

      <div className={styles.menuItem} onClick={onSettingsClick}>
        <div className={`${styles.iconContainer} ${styles.settingsIcon}`}>
          <svg width="24" height="24" viewBox="0 0 24 24" fill="#ff9933">
            <path d="M19.14,12.94c0.04-0.3,0.06-0.61,0.06-0.94c0-0.32-0.02-0.64-0.07-0.94l2.03-1.58c0.18-0.14,0.23-0.41,0.12-0.61 l-1.92-3.32c-0.12-0.22-0.37-0.29-0.59-0.22l-2.39,0.96c-0.5-0.38-1.03-0.7-1.62-0.94L14.4,2.81c-0.04-0.24-0.24-0.41-0.48-0.41 h-3.84c-0.24,0-0.43,0.17-0.47,0.41L9.25,5.35C8.66,5.59,8.12,5.92,7.63,6.29L5.24,5.33c-0.22-0.08-0.47,0-0.59,0.22L2.74,8.87 C2.62,9.08,2.66,9.34,2.86,9.48l2.03,1.58C4.84,11.36,4.8,11.69,4.8,12s0.02,0.64,0.07,0.94l-2.03,1.58 c-0.18,0.14-0.23,0.41-0.12,0.61l1.92,3.32c0.12,0.22,0.37,0.29,0.59,0.22l2.39-0.96c0.5,0.38,1.03,0.7,1.62,0.94l0.36,2.54 c0.05,0.24,0.24,0.41,0.48,0.41h3.84c0.24,0,0.44-0.17,0.47-0.41l0.36-2.54c0.59-0.24,1.13-0.56,1.62-0.94l2.39,0.96 c0.22,0.08,0.47,0,0.59-0.22l1.92-3.32c0.12-0.22,0.07-0.47-0.12-0.61L19.14,12.94z M12,15.6c-1.98,0-3.6-1.62-3.6-3.6 s1.62-3.6,3.6-3.6s3.6,1.62,3.6,3.6S13.98,15.6,12,15.6z" />
          </svg>
        </div>
        <div className={styles.menuText}>Настройки</div>
        <div className={styles.arrowIcon}>
          <svg width="24" height="24" viewBox="0 0 24 24" fill="currentColor">
            <path d="M8.59 16.59L13.17 12 8.59 7.41 10 6l6 6-6 6-1.41-1.41z" />
          </svg>
        </div>
      </div>

      <div className={styles.divider} />

      <div className={styles.menuItem}>
        <div className={`${styles.iconContainer} ${styles.inviteIcon}`}>
          <LogoutIcon />
        </div>
        <div className={styles.menuText} onClick={handleLogOut}>
          Выйти из аккаунта
        </div>
        <div className={styles.arrowIcon}>
          <svg width="24" height="24" viewBox="0 0 24 24" fill="currentColor">
            <path d="M8.59 16.59L13.17 12 8.59 7.41 10 6l6 6-6 6-1.41-1.41z" />
          </svg>
        </div>
      </div>

      <div className={styles.menuItem}>
        <div className={`${styles.iconContainer} ${styles.helpIcon}`}>
          <DeleteForeverIcon />
        </div>
        <div className={styles.menuText} onClick={handleDeleteAccount}>
          Удалить аккаунт
        </div>
        <div className={styles.arrowIcon}>
          <svg width="24" height="24" viewBox="0 0 24 24" fill="currentColor">
            <path d="M8.59 16.59L13.17 12 8.59 7.41 10 6l6 6-6 6-1.41-1.41z" />
          </svg>
        </div>
      </div>
    </div>
  );
};

export default ProfileMenu;
