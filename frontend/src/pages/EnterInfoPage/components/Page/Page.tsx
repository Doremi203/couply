import React, { useState } from "react";
import styles from "./enterInfo.module.css";
import { CustomInput } from "../../../../shared/components/CustomInput";
import { CustomButton } from "../../../../shared/components/CustomButton";
import KeyboardBackspaceIcon from "@mui/icons-material/KeyboardBackspace";
import { ToggleButtons } from "../../../../shared/components/ToggleButtons";
import { useNavigate } from "react-router-dom";

export const EnterInfoPage = () => {
  const [currentStep, setCurrentStep] = useState(0);
  const navigate = useNavigate();
  
  // State for form values
  const [name, setName] = useState("");
  const [birthDate, setBirthDate] = useState("");
  const [userGender, setUserGender] = useState("");
  const [preferredGender, setPreferredGender] = useState("");

  const nextStep = () => {
    if (currentStep === sections.length - 1) {
      // If we're on the last step, navigate to home page
      navigate("/home");
    } else {
      // Otherwise, go to the next step
      setCurrentStep((prevStep) => prevStep + 1);
    }
  };

  const prevStep = () => {
    if (currentStep > 0) {
      setCurrentStep((prevStep) => prevStep - 1);
    }
  };

  const handleUserGenderSelect = (value: string) => {
    setUserGender(value);
  };

  const handlePreferredGenderSelect = (value: string) => {
    setPreferredGender(value);
  };

  // Check if the current step's form is valid
  const isCurrentStepValid = () => {
    switch (currentStep) {
      case 0:
        return name.trim() !== "";
      case 1:
        return birthDate !== "";
      case 2:
        return userGender !== "" && preferredGender !== "";
      default:
        return false;
    }
  };

  const sections = [
    <div key="nameSection">
      <h2>Как вас зовут?</h2>
      <CustomInput
        placeholder="Введите имя"
        type="text"
        value={name}
        onChange={(e) => setName(e.target.value)}
      />
      <CustomButton
        onClick={nextStep}
        text={"Дальше"}
        disabled={!isCurrentStepValid()}
        className={styles.nextButton}
      />
    </div>,
    <div key="birthDateSection">
      <h2>Дата рождения</h2>
      <CustomInput
        placeholder="Выберите дату рождения"
        type="date"
        value={birthDate}
        onChange={(e) => setBirthDate(e.target.value)}
      />
      <CustomButton
        onClick={nextStep}
        text={"Дальше"}
        disabled={!isCurrentStepValid()}
      />
    </div>,
    <div key="datingSettingsSection">
      <h2>Настройки дейтинга</h2>
      <div>
        <label>Ваш пол:</label>
        <ToggleButtons
          options={[
            { label: "Женский", value: "female" },
            { label: "Мужской", value: "male" },
          ]}
          onSelect={handleUserGenderSelect}
          value={userGender}
        />
      </div>
      <div>
        <label>Кого вам показывать:</label>
        <ToggleButtons
          options={[
            { label: "Женщин", value: "female" },
            { label: "Мужчин", value: "male" },
            { label: "Всех", value: "other" },
          ]}
          onSelect={handlePreferredGenderSelect}
          value={preferredGender}
        />
      </div>
      <CustomButton
        onClick={nextStep}
        text={"Дальше"}
        disabled={!isCurrentStepValid()}
      />
    </div>,
  ];

  return (
    <div className={styles.container}>
      <div className={styles.backIcon} onClick={prevStep}>
        <KeyboardBackspaceIcon />
      </div>
      {sections[currentStep]}
    </div>
  );
};

export default EnterInfoPage;
