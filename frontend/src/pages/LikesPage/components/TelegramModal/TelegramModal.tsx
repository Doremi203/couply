import React from 'react';
import { createPortal } from 'react-dom';

import { useSetTelegramMutation } from '../../../../entities/telegram/api/telegramApi';

import styles from './telegramModal.module.css';

interface TelegramModalProps {
  isOpen: boolean;
  onClose: () => void;
}

export const TelegramModal: React.FC<TelegramModalProps> = ({ isOpen, onClose }) => {
  const [setTelegram] = useSetTelegramMutation();

  const onTelegramAuth = (user: any) => {
    console.log('user', user);
    setTelegram(user)
      .unwrap()
      .then(data => {
        console.log('Success:', data);
        // Close the modal immediately on success
        onClose();
      })
      .catch(error => {
        console.error('Error:', error);
        onClose();
      });
  };

  React.useEffect(() => {
    if (!isOpen) return;

    window.onTelegramAuth = onTelegramAuth;

    const script = document.createElement('script');
    script.src = 'https://telegram.org/js/telegram-widget.js?22';
    script.async = true;
    script.setAttribute('data-telegram-login', 'couply_auth_bot');
    script.setAttribute('data-size', 'large');
    script.setAttribute('data-onauth', 'onTelegramAuth(user)');

    document.getElementById('telegram-login-container')?.appendChild(script);

    return () => {
      const container = document.getElementById('telegram-login-container');
      if (container) {
        container.innerHTML = '';
      }

      window.onTelegramAuth = () => {};
    };
  }, [isOpen, onClose, onTelegramAuth]);

  if (!isOpen) return null;

  return createPortal(
    <div className={styles.modalOverlay}>
      <div className={styles.modalContent}>
        <h2 className={styles.title}>Подключите Telegram</h2>
        <p className={styles.message}>Чтобы смотреть лайки и мэтчи, нужно подключить телеграмм</p>

        <div className={styles.telegramContainer}>
          <div id="telegram-login-container" className={styles.telegramLoginButton} />
        </div>
      </div>
    </div>,
    document.body,
  );
};

declare global {
  interface Window {
    onTelegramAuth: (user: any) => void;
  }
}

export default TelegramModal;
