import { CustomButton } from '../../../../shared/components/CustomButton';

import styles from './hiddenAcc.module.css';

// @ts-ignore

export const HiddenAcc = () => {
  return (
    <div className={styles.container}>
      <div className={styles.box}>
        <div className={`${styles.wave} ${styles.waveOne}`} />
        <div className={`${styles.wave} ${styles.waveTwo}`} />
      </div>
      <div className={styles.textOverlay}>
        <div className={styles.header}> Ваш профиль скрыт</div>
        <div className={styles.text}>
          Активируйте профиль, чтобы продолжить искать свою каплю в океане людей
        </div>
      </div>
      {/** @ts-ignore */}
      <CustomButton className={styles.button} text="Активировать профиль" />
    </div>
  );
};
