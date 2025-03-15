import { useNavigate } from "react-router-dom";
import styles from "./splashPage.module.css";

export const SplashPage = () => {
  const navigate = useNavigate();

  setTimeout(() => {
    navigate("/auth");
  }, 3000);

  return (
    <div className={styles.page}>
      <body>
        <img
          src="logo.png"
          width="200px"
          height="150px"
          className={styles.logo}
        />

        <div className={styles.loader}></div>
      </body>
    </div>
  );
};

export default SplashPage;
