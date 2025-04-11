import { Stack } from '@mui/material';
import { useNavigate } from 'react-router-dom';

import { ButtonWithIcon } from '../../../../shared/components/ButtonWithIcon';

import styles from './authPage.module.css';


export const AuthPage = () => {
  const navigate = useNavigate();

  const onClick = () => {
    navigate('/enterInfo');
  };
  return (
    <div className={styles.page}>
      <Stack
        direction="column"
        spacing={2}
        sx={{
          justifyContent: 'center',
          alignItems: 'center',
        }}
      >
        <img src="logo.png" width="200px" height="150px" />

        <span className={styles.text}>
          {' '}
          Найди того, кто будет похож на тебя, как капля воды.{' '}
        </span>

        <Stack
          direction="column"
          spacing={2}
          sx={{
            justifyContent: 'center',
            alignItems: 'center',
          }}
        >

          <ButtonWithIcon onClick={onClick} icon={ <img
              src="image.png"
              width="20px"
              height="20px"
            />} text="login with google" />
          <ButtonWithIcon onClick={onClick} icon={<img
              src="vk.png"
              width="20px"
              height="20px"
            />} text="login with vk" />
          <ButtonWithIcon onClick={onClick} icon={<img
              src="phone.png"
              width="20px"
              height="20px"
            />} text="login with phone" />
        </Stack>
      </Stack>
    </div>
  );
};

export default AuthPage;
