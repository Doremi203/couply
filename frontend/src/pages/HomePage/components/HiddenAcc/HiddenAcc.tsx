// @ts-nocheck
import { useEffect, useState } from 'react';
import { useGetUserMutation, useUpdateUserMutation } from '../../../../entities/user';
import { CustomButton } from '../../../../shared/components/CustomButton';

import styles from './hiddenAcc.module.css';

export const HiddenAcc = () => {
  const [updateUser] = useUpdateUserMutation();

  const [getUser] = useGetUserMutation();
  const [profileData, setUsersData] = useState([]);

  useEffect(() => {
    const fetchUsers = async () => {
      try {
        const res = await getUser({}).unwrap();

        console.log(res);

        setUsersData(res.user);

        // setUsersData(results.filter(user => user !== null));
      } catch (err) {
        console.error('Error fetching users:', err);
      }
    };

    fetchUsers();
  }, [getUser]);

  const handleUpdate = async () => {
    // @ts-ignore
    const userData = {
      name: profileData.name,
      age: profileData.age,
      bio: profileData.bio,
      latitude: profileData.latitude,
      longitude: profileData.longitude,
      gender: profileData.gender,
      isHidden: false,
      education: profileData.education,
      children: profileData.children,
      alcohol: profileData.alcohol,
      smoking: profileData.smoking,
      goal: profileData.goal,
      interest: profileData.interest,
      height: profileData.height,
      zodiac: profileData.zodiac,
      isVerified: profileData.isVerified,
      isPremium: profileData.isPremium,
      isBlocked: profileData.isBlocked,
      photos: profileData.photos,
      createdAt: profileData.createdAt,
      id: profileData.id,
    };

    await updateUser(userData).unwrap();
    window.location.reload();
  };

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
      <CustomButton className={styles.button} text="Показывать профиль" onClick={handleUpdate} />
    </div>
  );
};
