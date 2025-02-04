import React from 'react'
import styles from './authPage.module.css'
//import { Button } from 'keep-react'
import { ShoppingCart } from 'phosphor-react'
import Button, { ButtonProps } from '@mui/material/Button';
import { Stack, styled } from '@mui/material';
import { white } from '@mui/material/colors';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCoffee } from '@fortawesome/free-solid-svg-icons';
import { Logo } from './logo'



/*const WhiteButton = styled(Button)({
    backgroundColor: 'white',
    color: 'black',
    '&:hover': {
      backgroundColor: '#f0f0f0', // изменяем цвет при наведении
    },
    borderRadius: '30px',
    borderColor: 'white',
    width: '260px',
    height: '40px',
    //textAlign: 'center'
  });*/

const WhiteButton = styled(Button)({
  backgroundColor: 'white',
  color: 'black',
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'center', // центрируем по горизонтали
  padding: '10px',
  minWidth: '150px',
  '&:hover': {
    backgroundColor: '#f0f0f0',
  },
  textTransform: 'none' // если не хотите чтобы текст был в верхнем регистре
});



export default function AuthPage() {
  return (
    <div className={styles.page}>

      <Stack
        direction="column"
        spacing={2}
        sx={{
          justifyContent: "center",
          alignItems: "center",
        }}
      >

        <img src='../public/logo.png' width='200px' height='150px' />

        <span className={styles.text}> Найди того, кто будет похож на тебя, как капля воды. </span>

        <Stack
          direction="column"
          spacing={2}
          sx={{
            justifyContent: "center",
            alignItems: "center",
          }}
        >
          <WhiteButton>
            <img src='../public/image.png' width='20px' height='20px' style={{ marginRight: '20px' }} />
            LOGIN WITH GOOGLE</WhiteButton>

          <WhiteButton>
            <img src='../public/vk.png' width='20px' height='20px' style={{ marginRight: '55px' }} />
            LOGIN WITH VK</WhiteButton>

          <WhiteButton>
            <img src='../public/phone.png' width='20px' height='20px' style={{ marginRight: '25px' }} />
            LOGIN WITH PHONE</WhiteButton>


        </Stack>

      </Stack>


    </div>
  )

}



