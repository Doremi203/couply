import React, { useState } from "react";
import styles from "./enterInfo.module.css";
import { CustomInput } from "../../../../shared/CustomInput";
import { CustomButton } from "../../../../shared/CustomButton";
import { ButtonWithIcon } from "../../../../shared/ButtonWithIcon";
import KeyboardBackspaceIcon from "@mui/icons-material/KeyboardBackspace";
import { ToggleButtons } from "../../../../shared/ToggleButtons";

export const EnterInfoPage = () => {
  const [currentStep, setCurrentStep] = useState(0);

  const nextStep = () => {
    setCurrentStep((prevStep) => (prevStep + 1) % sections.length);
  };

  const handleSelect = (value: string) => {
    console.log("Selected:", value);
  };

  const sections = [
    <div key="nameSection">
      <h2>Как вас зовут?</h2>
      <CustomInput placeholder="Введите имя" type="text" />
      <CustomButton onClick={nextStep} text={"Дальше"} />
    </div>,
    <div key="birthDateSection">
      <h2>Дата рождения</h2>
      {/* <input type="date" />
      <button onClick={nextStep}>Дальше</button> */}
      <CustomInput placeholder="Введите имя" type="date" />
      <CustomButton onClick={nextStep} text={"Дальше"} />
    </div>,
    <div key="datingSettingsSection">
      <h2>Настройки дейтинга</h2>
      <div>
        <label>Ваш пол:</label>
        <button>Женский</button>
        <button>Мужской</button>
      </div>
      <div>
        <label>Кого вам показывать:</label>
        <button>Женщин</button>
        <button>Мужчин</button>
        <button>Всех</button>
        <ToggleButtons
          options={[
            { label: "Женский", value: "female" },
            { label: "Мужской", value: "male" },
            { label: "Другой", value: "other" },
          ]}
          onSelect={handleSelect}
        />
      </div>
      <CustomButton onClick={nextStep} text={"Дальше"} />
    </div>,
  ];

  return (
    <div>
      <KeyboardBackspaceIcon />
      {sections[currentStep]}
    </div>
  );
};

export default EnterInfoPage;
