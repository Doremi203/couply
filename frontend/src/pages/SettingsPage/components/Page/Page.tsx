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

  return (
    <div className={styles.page}>
      <PageHeader onBack={onBack} title="Настройки" />

      <div className={styles.cont}>
        <NotificationSettings />
        <SubscriptionSettings />
      </div>
    </div>
  );
};

export default SettingsPage;
