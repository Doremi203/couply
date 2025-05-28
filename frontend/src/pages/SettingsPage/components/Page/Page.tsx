import { useNavigate } from 'react-router-dom';

import PageHeader from '../../../../shared/components/PageHeader';
import NotificationSettings from '../NotificationSettings';
import SubscriptionSettings from '../SubscriptionSettings';

import styles from './settings.module.css';

export const SettingsPage = () => {
  const navigate = useNavigate();

  const onBack = () => {
    navigate('/profile');
  };

  const openTermsOfService = () => {
    navigate('/terms');
  };

  const openPrivacyPolicy = () => {
    navigate('/privacy');
  };

  return (
    <div className={styles.page}>
      <PageHeader onBack={onBack} title="Настройки" />

      <div className={styles.cont}>
        <NotificationSettings />
        <SubscriptionSettings />
        <div className={styles.terms}>
          <div className={styles.link1} onClick={openPrivacyPolicy}>
            {' '}
            Политика конфиденциальности
          </div>
          <div className={styles.link2} onClick={openTermsOfService}>
            {' '}
            Пользовательское соглашение
          </div>
        </div>
      </div>
    </div>
  );
};

export default SettingsPage;
