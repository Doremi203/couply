import { Box } from '@mui/material';
import React, { useState, useEffect } from 'react';

import { GenderPriority, useUpdateFilterMutation } from '../../../../entities/search';
import {
  Alcohol,
  Education,
  Goal,
  Hobby,
  Selfdevelopment,
  Smoking,
  Sport,
  Zodiac,
  Children,
} from '../../../../entities/user';
import { genderOptions, goalOptions, zodiacOptions } from '../constants';

import ChipFilter from './components/ChipFilter';
import FilterActions from './components/FilterActions';
import FilterHeader from './components/FilterHeader';
import GenderFilter from './components/GenderFilter';
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
  const [goal, setGoal] = useState<string[]>([]);
  const [verificationStatus, setVerificationStatus] = useState<boolean>(false);

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

  const toggleInterest = (interest: string) => {
    if (selectedInterests.includes(interest)) {
      setSelectedInterests(selectedInterests.filter(item => item !== interest));
    } else {
      setSelectedInterests([...selectedInterests, interest]);
    }
  };

  const handleGoalSelect = (value: string) => {
    // setInterestedIn(value);
    setGoal(value);
  };

  const handleClearFilters = () => {
    setInterestedIn('Both');
    setDistance(40);
    setAgeRange([18, 28]);
    setSelectedInterests([]);
    setVerificationStatus(false);
  };

  if (!open) return null;

  const handleApplyFilters = () => {
    updateFilter({
      genderPriority: GenderPriority.male,
      minAge: 18,
      maxAge: 100,
      minHeight: 100,
      maxHeight: 250,
      distance: 100,
      goal: Goal.unspecified,
      zodiac: Zodiac.unspecified,
      education: Education.unspecified,
      children: Children.unspecified,
      alcohol: Alcohol.unspecified,
      smoking: Smoking.unspecified,
      interest: {
        sport: [Sport.unspecified],
        selfDevelopment: [Selfdevelopment.unspecified],
        hobby: [Hobby.unspecified],
      },
      onlyVerified: false,
      onlyPremium: false,
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
            max={65}
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
            options={goalOptions}
            selectedOptions={goal}
            onToggle={handleGoalSelect}
          />

          <ChipFilter
            title="Интересы"
            options={interestOptions}
            selectedOptions={selectedInterests}
            onToggle={toggleInterest}
          />

          <ChipFilter
            title="Знаки зодиака"
            options={zodiacOptions}
            selectedOptions={selectedZodiac}
            onToggle={toggleInterest}
          />

          <ToggleFilter
            title="Статус верификации"
            description="Показывать только верифицированных пользователей"
            value={verificationStatus}
            onChange={handleVerificationToggle}
          />

          <FilterActions onContinue={handleApplyFilters} />
        </Box>
      </div>
    </div>
  );
};

export default FiltersDrawer;
