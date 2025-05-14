import DeleteForeverIcon from '@mui/icons-material/DeleteForever';
import LogoutIcon from '@mui/icons-material/Logout';
import React from 'react';
import { useNavigate } from 'react-router-dom';

import { useDeleteUserMutation } from '../../../../entities/user';

import styles from './profileMenu.module.css';

interface ProfileMenuProps {
  onEditProfileClick: () => void;
  onMyStatsClick: () => void;
  onSettingsClick: () => void;
  onInviteFriendClick: () => void;
  onHelpClick: () => void;
}

export const ProfileMenu: React.FC<ProfileMenuProps> = ({
  onEditProfileClick,
  onSettingsClick,
  onInviteFriendClick,
  onHelpClick,
}) => {
  const naviagate = useNavigate();

  const handleLogOut = () => {
    naviagate('/auth');
    localStorage.removeItem('token');
  };

  const [deleteUser] = useDeleteUserMutation();

  const handleDeleteAccount = async () => {
    await deleteUser();
    naviagate('/auth');
    localStorage.removeItem('token');
  };
  return (
    <div className={styles.menuContainer}>
      {/* Edit profile */}
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

      {/* My stats */}
      {/* <div className={styles.menuItem} onClick={onMyStatsClick}>
        <div className={`${styles.iconContainer} ${styles.statsIcon}`}>
          <svg width="24" height="24" viewBox="0 0 24 24" fill="#6666ff">
            <path d="M9 17H7V10H9V17ZM13 17H11V7H13V17ZM17 17H15V13H17V17ZM19 19H5V5H19V19ZM19 3H5C3.9 3 3 3.9 3 5V19C3 20.1 3.9 21 5 21H19C20.1 21 21 20.1 21 19V5C21 3.9 20.1 3 19 3Z" />
          </svg>
        </div>
        <div className={styles.menuText}>My stats</div>
        <div className={styles.arrowIcon}>
          <svg width="24" height="24" viewBox="0 0 24 24" fill="currentColor">
            <path d="M8.59 16.59L13.17 12 8.59 7.41 10 6l6 6-6 6-1.41-1.41z" />
          </svg>
        </div>
      </div> */}

      {/* Settings */}
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

      {/* Divider */}
      <div className={styles.divider} />

      {/* Invite a friend */}
      <div className={styles.menuItem} onClick={onInviteFriendClick}>
        <div className={`${styles.iconContainer} ${styles.inviteIcon}`}>
          {/* <svg width="24" height="24" viewBox="0 0 24 24" fill="#666666">
            <path d="M15 12c2.21 0 4-1.79 4-4s-1.79-4-4-4-4 1.79-4 4 1.79 4 4 4zm-9-2V7H4v3H1v2h3v3h2v-3h3v-2H6zm9 4c-2.67 0-8 1.34-8 4v2h16v-2c0-2.66-5.33-4-8-4z" />
          </svg> */}
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

      {/* Help */}
      <div className={styles.menuItem}>
        <div className={`${styles.iconContainer} ${styles.helpIcon}`}>
          {/* <svg width="24" height="24" viewBox="0 0 24 24" fill="#666666">
            <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm1 17h-2v-2h2v2zm2.07-7.75l-.9.92C13.45 12.9 13 13.5 13 15h-2v-.5c0-1.1.45-2.1 1.17-2.83l1.24-1.26c.37-.36.59-.86.59-1.41 0-1.1-.9-2-2-2s-2 .9-2 2H8c0-2.21 1.79-4 4-4s4 1.79 4 4c0 .88-.36 1.68-.93 2.25z" />
          </svg> */}
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
