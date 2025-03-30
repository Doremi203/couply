import styles from "./authPage.module.css";
//import { Button } from 'keep-react'
import Button from "@mui/material/Button";
import { Stack, styled } from "@mui/material";
import { useNavigate } from "react-router-dom";

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
  backgroundColor: "white",
  color: "black",
  display: "flex",
  alignItems: "center",
  justifyContent: "center", // центрируем по горизонтали
  padding: "10px",
  minWidth: "150px",
  "&:hover": {
    backgroundColor: "#f0f0f0",
  },
  textTransform: "none", // если не хотите чтобы текст был в верхнем регистре
  width: "260px",
  fontFamily: "Jost",
});

export const AuthPage = () => {
  const navigate = useNavigate();

  const onClick = () => {
    navigate("/enterInfo");
  };
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
        <img src="logo.png" width="200px" height="150px" />

        <span className={styles.text}>
          {" "}
          Найди того, кто будет похож на тебя, как капля воды.{" "}
        </span>

        <Stack
          direction="column"
          spacing={2}
          sx={{
            justifyContent: "center",
            alignItems: "center",
          }}
        >
          <WhiteButton onClick={onClick}>
            <img
              src="image.png"
              width="20px"
              height="20px"
              style={{ marginRight: "20px" }}
            />
            LOGIN WITH GOOGLE
          </WhiteButton>

          <WhiteButton onClick={onClick}>
            <img
              src="vk.png"
              width="20px"
              height="20px"
              style={{ marginRight: "55px" }}
            />
            LOGIN WITH VK
          </WhiteButton>

          <WhiteButton onClick={onClick}>
            <img
              src="phone.png"
              width="20px"
              height="20px"
              style={{ marginRight: "25px" }}
            />
            LOGIN WITH PHONE
          </WhiteButton>

          {/* <button className={styles.loginButton}>
            <img src='../public/phone.png' width='20px' height='20px' style={{ marginRight: '25px' }} />
            LOGIN WITH PHONE
                </button> */}
        </Stack>
      </Stack>
    </div>
  );
};

export default AuthPage;
