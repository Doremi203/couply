import { Box } from '@mui/material';
import React, { useState, useEffect } from 'react';

import { useUpdateFilterMutation } from '../../../../entities/search';
import { Alcohol, Education, Smoking, Zodiac, Children } from '../../../../entities/user';
import { createToggleHandler } from '../../helpers/createToggleHandler';
import { mapFiltersToApi } from '../../helpers/mapFiltersToApiFormat';
import { mapInterestsToBackendFormat } from '../../helpers/mapInterestsToApiFormat';
import {
  alcoholOptions,
  alcoholToApi,
  childrenOptions,
  childrenToApi,
  educationOptions,
  educationToApi,
  genderOptions,
  genderToApi,
  goalOptions,
  goalToApi,
  smokingOptions,
  smokingToApi,
  zodiacOptions,
  zodiacToApi,
} from '../constants';

import ChipFilter from './components/ChipFilter';
import FilterActions from './components/FilterActions';
import FilterHeader from './components/FilterHeader';
import GenderFilter from './components/GenderFilter';
import InterestFilter from './components/InterestFilter/InterestFilter';
import SliderFilter from './components/SliderFilter';
import ToggleFilter from './components/ToggleFilter';
import styles from './filtersDrawer.module.css';

type Props = {
  open: boolean;
  onClose: () => void;
};

export const FiltersDrawer: React.FC<Props> = ({ open, onClose }) => {
  useEffect(() => {
    if (open) {
      document.body.style.overflow = 'hidden';
    } else {
      document.body.style.overflow = '';
    }
    return () => {
      document.body.style.overflow = '';
    };
  }, [open]);

  const [updateFilter] = useUpdateFilterMutation();

  const [interestedIn, setInterestedIn] = useState<string>('Girls');
  const [distance, setDistance] = useState<number>(40);
  const [ageRange, setAgeRange] = useState<number[]>([18, 28]);
  const [heightRange, setHeightRange] = useState<number[]>([170, 190]);
  const interestOptions = ['Спорт', 'Путешествия', 'Музыка', 'Искусство', 'Гастрономия'];
  const [selectedInterests, setSelectedInterests] = useState<string[]>([]);

  const [selectedZodiac, setSelectedZodiac] = useState<string[]>([]);
  const [selectedGoal, setSelectedGoal] = useState<string[]>([]);
  const [verificationStatus, setVerificationStatus] = useState<boolean>(false);
  const [premiumStatus, setPremiumStatus] = useState<boolean>(false);

  const handleDistanceChange = (_event: Event, newValue: number | number[]) => {
    setDistance(newValue as number);
  };

  const handleAgeRangeChange = (_event: Event, newValue: number | number[]) => {
    setAgeRange(newValue as number[]);
  };

  const handleHeightRangeChange = (_event: Event, newValue: number | number[]) => {
    setHeightRange(newValue as number[]);
  };

  const handleGenderSelect = (value: string) => {
    setInterestedIn(value);
  };

  const handleVerificationToggle = () => {
    setVerificationStatus(!verificationStatus);
  };

  const handlePremiumToggle = () => {
    setPremiumStatus(!premiumStatus);
  };

  const [selectedEducation, setSelectedEducation] = useState<string[]>([]);
  const [selectedChildren, setSelectedChildren] = useState<string[]>([]);
  const [selectedAlcohol, setSelectedAlcohol] = useState<string[]>([]);
  const [selectedSmoking, setSelectedSmoking] = useState<string[]>([]);

  const handleZodiacToggle = createToggleHandler(setSelectedZodiac, selectedZodiac);
  const handleEducationToggle = createToggleHandler(setSelectedEducation, selectedEducation);
  const handleChildrenToggle = createToggleHandler(setSelectedChildren, selectedChildren);
  const handleAlcoholToggle = createToggleHandler(setSelectedAlcohol, selectedAlcohol);
  const handleSmokingToggle = createToggleHandler(setSelectedSmoking, selectedSmoking);

  const handleGoalSelect = (value: string) => {
    setSelectedGoal(value);
  };

  const handleClearFilters = () => {
    setInterestedIn('Both');
    setDistance(40);
    setAgeRange([18, 28]);
    setSelectedInterests([]);
    setVerificationStatus(false);
  };

  if (!open) return null;

  // console.log(selectedGoal);
  // console.log(Goal[selectedGoal]);
  console.log(mapFiltersToApi(selectedZodiac, zodiacToApi, Zodiac.unspecified));
  // console.log(goalToApi.selectedGoal);
  console.log(goalToApi[selectedGoal]);

  console.log(selectedZodiac);
  console.log(mapInterestsToBackendFormat(selectedInterests));

  const handleApplyFilters = () => {
    updateFilter({
      genderPriority: genderToApi[interestedIn],
      minAge: ageRange[0],
      maxAge: ageRange[1],
      minHeight: heightRange[0],
      maxHeight: heightRange[1],
      distance: distance,
      goal: goalToApi[selectedGoal],
      // zodiac: mapFiltersToApi(selectedZodiac, zodiacToApi, Zodiac.unspecified),
      zodiac: Zodiac.unspecified, //TODO
      education: mapFiltersToApi(selectedEducation, educationToApi, Education.unspecified)[0],
      children: mapFiltersToApi(selectedChildren, childrenToApi, Children.unspecified)[0],
      alcohol: mapFiltersToApi(selectedAlcohol, alcoholToApi, Alcohol.unspecified)[0],
      smoking: mapFiltersToApi(selectedSmoking, smokingToApi, Smoking.unspecified)[0],

      // Группа интересов с множественным выбором
      interest: mapInterestsToBackendFormat(selectedInterests),
      // interest: null,
      onlyVerified: verificationStatus,
      onlyPremium: premiumStatus,
    });
    onClose();
  };

  return (
    <div className={styles.modalOverlay}>
      <div className={styles.modalContainer}>
        <Box className={styles.content}>
          <FilterHeader onBack={onClose} onClear={handleClearFilters} />

          <GenderFilter
            value={interestedIn}
            options={genderOptions}
            onChange={handleGenderSelect}
          />

          <SliderFilter
            title="Дистанция"
            value={distance}
            min={1}
            max={100}
            onChange={handleDistanceChange}
            unit="km"
          />

          <SliderFilter
            title="Возраст"
            value={ageRange}
            min={18}
            max={99}
            onChange={handleAgeRangeChange}
            valueLabelDisplay="auto"
          />

          <SliderFilter
            title="Рост"
            value={heightRange}
            min={150}
            max={240}
            onChange={handleHeightRangeChange}
            valueLabelDisplay="auto"
          />

          <ChipFilter
            title="Цель"
            options={Object.values(goalOptions)}
            selectedOptions={selectedGoal}
            onToggle={handleGoalSelect}
          />

          <InterestFilter
            title="Интересы"
            options={interestOptions}
            selectedOptions={selectedInterests}
            onSelect={selected => setSelectedInterests(selected)}
          />

          <ChipFilter
            title="Знаки зодиака"
            options={Object.values(zodiacOptions)}
            selectedOptions={selectedZodiac}
            onToggle={handleZodiacToggle}
          />

          <ChipFilter
            title="Образование"
            options={Object.values(educationOptions)}
            selectedOptions={selectedEducation}
            onToggle={handleEducationToggle}
          />

          <ChipFilter
            title="Дети"
            options={Object.values(childrenOptions)}
            selectedOptions={selectedChildren}
            onToggle={handleChildrenToggle}
          />

          <ChipFilter
            title="Алкоголь"
            options={Object.values(alcoholOptions)}
            selectedOptions={selectedAlcohol}
            onToggle={handleAlcoholToggle}
          />

          <ChipFilter
            title="Курение"
            options={Object.values(smokingOptions)}
            selectedOptions={selectedSmoking}
            onToggle={handleSmokingToggle}
          />

          <ToggleFilter
            title="Только верифицированные"
            description="Показывать только верифицированных пользователей"
            value={verificationStatus}
            onChange={handleVerificationToggle}
          />

          <ToggleFilter
            title="Только премиум"
            description="Показывать только премиум пользователей"
            value={premiumStatus}
            onChange={handlePremiumToggle}
          />
          <FilterActions onContinue={handleApplyFilters} />
        </Box>
      </div>
    </div>
  );
};

export default FiltersDrawer;
