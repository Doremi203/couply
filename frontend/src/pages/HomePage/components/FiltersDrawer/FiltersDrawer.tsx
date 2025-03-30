import React, { useState } from "react";
import { Drawer, Box } from "@mui/material";
import styles from "./filtersDrawer.module.css";

// Import components
import FilterHeader from "./components/FilterHeader";
import GenderFilter from "./components/GenderFilter";
import SliderFilter from "./components/SliderFilter";
import ChipFilter from "./components/ChipFilter";
import ToggleFilter from "./components/ToggleFilter";
import FilterActions from "./components/FilterActions";

type Props = {
  open: boolean;
  onClose: () => void;
};

const FiltersDrawer: React.FC<Props> = ({ open, onClose }) => {
  // Gender preference
  const [interestedIn, setInterestedIn] = useState<string>("Girls");
  
  // Distance
  const [distance, setDistance] = useState<number>(40);
  
  // Age range
  const [ageRange, setAgeRange] = useState<number[]>([20, 28]);
  
  // Interests
  const interestOptions = ["Sports", "Travel", "Music", "Art", "Food"];
  const [selectedInterests, setSelectedInterests] = useState<string[]>([]);
  
  // Music preferences
  const musicOptions = ["Rock", "Pop", "Hip Hop", "Jazz", "Classical"];
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
    setInterestedIn("Both");
    setDistance(40);
    setAgeRange([20, 28]);
    setSelectedInterests([]);
    setSelectedMusicPreferences([]);
    setVerificationStatus(false);
  };

  const genderOptions = [
    { label: "Girls", value: "Girls" },
    { label: "Boys", value: "Boys" },
    { label: "Both", value: "Both" },
  ];

  return (
    <Drawer
      anchor="top"
      open={open}
      onClose={onClose}
      PaperProps={{
        sx: {
          borderBottomLeftRadius: '20px',
          borderBottomRightRadius: '20px',
          height: 'auto',
          maxHeight: '90vh',
          overflow: 'auto',
          width: '100%',
          maxWidth: '430px',
          margin: '0 auto',
          boxShadow: '0px 4px 10px rgba(0, 0, 0, 0.1)'
        }
      }}
    >
      <Box className={styles.container}>
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
    </Drawer>
  );
};

export default FiltersDrawer;
