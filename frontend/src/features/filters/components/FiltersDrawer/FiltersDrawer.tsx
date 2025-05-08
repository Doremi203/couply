import { Box } from '@mui/material';
import React, { useState, useEffect } from 'react';

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

  const [interestedIn, setInterestedIn] = useState<string>('Girls');
  const [distance, setDistance] = useState<number>(40);
  const [ageRange, setAgeRange] = useState<number[]>([20, 28]);
  const interestOptions = ['Sports', 'Travel', 'Music', 'Art', 'Food'];
  const [selectedInterests, setSelectedInterests] = useState<string[]>([]);
  const musicOptions = ['Rock', 'Pop', 'Hip Hop', 'Jazz', 'Classical'];
  const [selectedMusicPreferences, setSelectedMusicPreferences] = useState<string[]>([]);

  const [verificationStatus, setVerificationStatus] = useState<boolean>(false);

  const handleDistanceChange = (_event: Event, newValue: number | number[]) => {
    setDistance(newValue as number);
  };

  const handleAgeRangeChange = (_event: Event, newValue: number | number[]) => {
    setAgeRange(newValue as number[]);
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

  const toggleMusicPreference = (music: string) => {
    if (selectedMusicPreferences.includes(music)) {
      setSelectedMusicPreferences(selectedMusicPreferences.filter(item => item !== music));
    } else {
      setSelectedMusicPreferences([...selectedMusicPreferences, music]);
    }
  };

  const handleClearFilters = () => {
    setInterestedIn('Both');
    setDistance(40);
    setAgeRange([18, 28]);
    setSelectedInterests([]);
    setSelectedMusicPreferences([]);
    setVerificationStatus(false);
  };

  const genderOptions = [
    { label: 'Женщины', value: 'Girls' },
    { label: 'Мужчины', value: 'Boys' },
    { label: 'Оба', value: 'Both' },
  ];

  if (!open) return null;

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

          <ChipFilter
            title="Интересы"
            options={interestOptions}
            selectedOptions={selectedInterests}
            onToggle={toggleInterest}
          />

          <ChipFilter
            title="Музыкальные предпочтения"
            options={musicOptions}
            selectedOptions={selectedMusicPreferences}
            onToggle={toggleMusicPreference}
          />

          <ToggleFilter
            title="Статус верификации"
            description="Показывать только верифицированных пользователей"
            value={verificationStatus}
            onChange={handleVerificationToggle}
          />

          <FilterActions onContinue={onClose} />
        </Box>
      </div>
    </div>
  );
};

export default FiltersDrawer;
