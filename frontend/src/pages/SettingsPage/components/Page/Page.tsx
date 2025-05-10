import { useNavigate } from 'react-router-dom';

import PageHeader from '../../../../shared/components/PageHeader';
import NotificationSettings from '../NotificationSettings';

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
        {/* <div className={styles.section}> */}
        {/* <h3>Уведомления</h3> */}
        <NotificationSettings />
        {/* </div> */}
      </div>
    </div>
  );
};

export default SettingsPage;
