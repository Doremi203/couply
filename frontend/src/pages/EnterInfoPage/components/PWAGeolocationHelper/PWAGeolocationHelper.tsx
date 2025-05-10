import { Dialog, Typography, Button } from '@mui/material';
import React from 'react';

import styles from './pwaGeolocationHelper.module.css';

interface PWAGeolocationHelperProps {
  open: boolean;
  onClose: () => void;
  isIOS: boolean;
}

/**
 * Component that provides detailed instructions for enabling geolocation in PWA mode
 */
const PWAGeolocationHelper: React.FC<PWAGeolocationHelperProps> = ({ open, onClose, isIOS }) => {
  return (
    <Dialog open={open} onClose={onClose} maxWidth="sm" fullWidth>
      <div className={styles.container}>
        <Typography variant="h6" className={styles.title}>
          Как включить геолокацию в PWA
        </Typography>

        {isIOS ? (
          <div className={styles.instructionsContainer}>
            <Typography variant="subtitle1" className={styles.subtitle}>
              Для iOS устройств:
            </Typography>
            <ol className={styles.instructionsList}>
              <li>Закройте это приложение (полностью выйдите из него)</li>
              <li>Откройте приложение &quot;Настройки&quot; на вашем устройстве</li>
              <li>Прокрутите вниз и найдите наше приложение в списке</li>
              <li>Нажмите на название приложения</li>
              <li>Найдите &quot;Местоположение&quot; и выберите &quot;При использовании&quot;</li>
              <li>Вернитесь в приложение и попробуйте снова</li>
            </ol>
            <div className={styles.imageContainer}>
              <img
                src="https://developer.apple.com/design/human-interface-guidelines/images/intro/platforms/platform-ios-dark_2x.png"
                alt="iOS Settings"
                className={styles.instructionImage}
              />
            </div>
          </div>
        ) : (
          <div className={styles.instructionsContainer}>
            <Typography variant="subtitle1" className={styles.subtitle}>
              Для Android устройств:
            </Typography>
            <ol className={styles.instructionsList}>
              <li>Откройте &quot;Настройки&quot; на вашем устройстве</li>
              <li>Выберите &quot;Приложения&quot; или &quot;Диспетчер приложений&quot;</li>
              <li>Найдите наше приложение в списке</li>
              <li>Нажмите &quot;Разрешения&quot;</li>
              <li>Найдите &quot;Местоположение&quot; и включите его</li>
              <li>Вернитесь в приложение и попробуйте снова</li>
            </ol>
            <div className={styles.imageContainer}>
              <img
                src="https://developer.android.com/static/images/guide/topics/permissions/runtime_permission_request_2x.png"
                alt="Android Permissions"
                className={styles.instructionImage}
              />
            </div>
          </div>
        )}

        <div className={styles.buttonContainer}>
          <Button variant="contained" color="primary" onClick={onClose} className={styles.button}>
            Понятно
          </Button>
        </div>
      </div>
    </Dialog>
  );
};

export default PWAGeolocationHelper;
