import React from "react";
import CustomButton from "../../../../../../shared/components/CustomButton/CustomButton";

type FilterActionsProps = {
  onContinue: () => void;
  buttonText?: string;
};

const FilterActions: React.FC<FilterActionsProps> = ({
  onContinue,
  buttonText = "Continue"
}) => {
  return (
    <CustomButton
      text={buttonText}
      onClick={onContinue}
    />
  );
};

export default FilterActions;