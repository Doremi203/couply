import { Box } from '@mui/material';
import React, { useState, useEffect } from 'react';

// Import components
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
  // Prevent body scrolling when modal is open
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

  // Gender preference
  const [interestedIn, setInterestedIn] = useState<string>('Girls');

  // Distance
  const [distance, setDistance] = useState<number>(40);

  // Age range
  const [ageRange, setAgeRange] = useState<number[]>([20, 28]);

  // Interests
  const interestOptions = ['Sports', 'Travel', 'Music', 'Art', 'Food'];
  const [selectedInterests, setSelectedInterests] = useState<string[]>([]);

  // Music preferences
  const musicOptions = ['Rock', 'Pop', 'Hip Hop', 'Jazz', 'Classical'];
  const [selectedMusicPreferences, setSelectedMusicPreferences] = useState<string[]>([]);

  // Verification status
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
    setAgeRange([20, 28]);
    setSelectedInterests([]);
    setSelectedMusicPreferences([]);
    setVerificationStatus(false);
  };

  const genderOptions = [
    { label: 'Girls', value: 'Girls' },
    { label: 'Boys', value: 'Boys' },
    { label: 'Both', value: 'Both' },
  ];

  if (!open) return null;

  return (
    <div className={styles.modalOverlay}>
      <div className={styles.modalContainer}>
        <Box className={styles.content}>
          {/* Header */}
          <FilterHeader onBack={onClose} onClear={handleClearFilters} />

          {/* Gender Preference */}
          <GenderFilter
            value={interestedIn}
            options={genderOptions}
            onChange={handleGenderSelect}
          />

          {/* Distance Slider */}
          <SliderFilter
            title="Distance"
            value={distance}
            min={1}
            max={100}
            onChange={handleDistanceChange}
            unit="km"
          />

          {/* Age Range Slider */}
          <SliderFilter
            title="Age"
            value={ageRange}
            min={18}
            max={65}
            onChange={handleAgeRangeChange}
            valueLabelDisplay="auto"
          />

          {/* Interests */}
          <ChipFilter
            title="Interests"
            options={interestOptions}
            selectedOptions={selectedInterests}
            onToggle={toggleInterest}
          />

          {/* Music Preferences */}
          <ChipFilter
            title="Music Preferences"
            options={musicOptions}
            selectedOptions={selectedMusicPreferences}
            onToggle={toggleMusicPreference}
          />

          {/* Verification Status */}
          <ToggleFilter
            title="Verification Status"
            description="Only show verified profiles"
            value={verificationStatus}
            onChange={handleVerificationToggle}
          />

          {/* Continue Button */}
          <FilterActions onContinue={onClose} />
        </Box>
      </div>
    </div>
  );
};

export default FiltersDrawer;
