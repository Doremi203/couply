import styles from './noUsersLeft.module.css';

export const NoUsersLeft = () => {
  return (
    <div className={styles.container}>
      <div className={styles.box}>
        <div className={`${styles.wave} ${styles.waveOne}`} />
        <div className={`${styles.wave} ${styles.waveTwo}`} />
      </div>
      <div className={styles.textOverlay}>
        <div className={styles.header}> Вы посмотрели всех пользователей</div>
        <div className={styles.text}>Загляните сюда позже</div>
      </div>
    </div>
  );
};
