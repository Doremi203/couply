import React, { useState } from "react";
import { Drawer, Box, Slider, styled } from "@mui/material";
import styles from "./filtersDrawer.module.css";
import ToggleButtons from "../../../../shared/ToggleButtons/ToggleButtons";
import CustomInput from "../../../../shared/CustomInput/CustomInput";
import CustomButton from "../../../../shared/CustomButton/CustomButton";
import KeyboardArrowRightIcon from '@mui/icons-material/KeyboardArrowRight';

type Props = {
  open: boolean;
  onClose: () => void;
};

// Custom styled slider with blue track
const CustomSlider = styled(Slider)({
  color: '#202C83',
  height: 8,
  '& .MuiSlider-track': {
    border: 'none',
    backgroundColor: '#202C83',
  },
  '& .MuiSlider-thumb': {
    height: 24,
    width: 24,
    backgroundColor: '#202C83',
    border: '2px solid #fff',
    boxShadow: '0 3px 6px rgba(0,0,0,0.16)',
    '&:focus, &:hover, &.Mui-active, &.Mui-focusVisible': {
      boxShadow: '0 3px 8px rgba(0,0,0,0.3)',
    },
  },
  '& .MuiSlider-rail': {
    backgroundColor: '#E0E0E0',
  },
});

const FiltersDrawer: React.FC<Props> = ({ open, onClose }) => {
  // Gender preference
  const [interestedIn, setInterestedIn] = useState<string>("Girls");
  
  // Location
  const [location, setLocation] = useState<string>("Chicago, USA");
  
  // Distance
  const [distance, setDistance] = useState<number>(40);
  
  // Age range
  const [ageRange, setAgeRange] = useState<number[]>([20, 28]);
  
  // Interests (new)
  const [interests, setInterests] = useState<string[]>([]);
  
  // Music preferences (new)
  const [musicPreferences, setMusicPreferences] = useState<string[]>([]);
  
  // Verification status (new)
  const [verificationStatus, setVerificationStatus] = useState<boolean>(false);

  const handleDistanceChange = (event: Event, newValue: number | number[]) => {
    setDistance(newValue as number);
  };

  const handleAgeRangeChange = (event: Event, newValue: number | number[]) => {
    setAgeRange(newValue as number[]);
  };

  const handleGenderSelect = (value: string) => {
    setInterestedIn(value);
  };

  const handleVerificationToggle = () => {
    setVerificationStatus(!verificationStatus);
  };

  const handleClearFilters = () => {
    setInterestedIn("Both");
    setLocation("Chicago, USA");
    setDistance(40);
    setAgeRange([20, 28]);
    setInterests([]);
    setMusicPreferences([]);
    setVerificationStatus(false);
  };

  const genderOptions = [
    { label: "Girls", value: "Girls" },
    { label: "Boys", value: "Boys" },
    { label: "Both", value: "Both" },
  ];

  return (
    <Drawer
      anchor="bottom"
      open={open}
      onClose={onClose}
      PaperProps={{
        sx: {
          borderTopLeftRadius: '20px',
          borderTopRightRadius: '20px',
          height: 'auto',
          maxHeight: '90vh',
          overflow: 'auto'
        }
      }}
    >
      <Box className={styles.container}>
        {/* Header */}
        <div className={styles.header}>
          <button className={styles.backButton} onClick={onClose}>
            ←
          </button>
          <h2 className={styles.title}>Filters</h2>
          <button className={styles.clearButton} onClick={handleClearFilters}>
            Clear
          </button>
        </div>

        {/* Gender Preference */}
        <div className={styles.section}>
          <h3 className={styles.sectionTitle}>Interested in</h3>
          <ToggleButtons
            options={genderOptions}
            onSelect={handleGenderSelect}
            value={interestedIn}
          />
        </div>

        {/* Location */}
        <div className={styles.section}>
          <h3 className={styles.sectionTitle}>Location</h3>
          <div className={styles.locationContainer}>
            <CustomInput
              type="text"
              placeholder="Enter location"
              value={location}
              onChange={(e) => setLocation(e.target.value)}
              className={styles.locationInput}
            />
            <KeyboardArrowRightIcon className={styles.locationArrow} />
          </div>
        </div>

        {/* Distance */}
        <div className={styles.section}>
          <div className={styles.sliderHeader}>
            <h3 className={styles.sectionTitle}>Distance</h3>
            <span className={styles.sliderValue}>{distance}km</span>
          </div>
          <CustomSlider
            value={distance}
            onChange={handleDistanceChange}
            aria-labelledby="distance-slider"
            min={1}
            max={100}
          />
        </div>

        {/* Age */}
        <div className={styles.section}>
          <div className={styles.sliderHeader}>
            <h3 className={styles.sectionTitle}>Age</h3>
            <span className={styles.sliderValue}>{ageRange[0]}-{ageRange[1]}</span>
          </div>
          <CustomSlider
            value={ageRange}
            onChange={handleAgeRangeChange}
            valueLabelDisplay="auto"
            aria-labelledby="range-slider"
            min={18}
            max={65}
          />
        </div>

        {/* Interests */}
        <div className={styles.section}>
          <h3 className={styles.sectionTitle}>Interests</h3>
          <div className={styles.chipContainer}>
            {/* Placeholder for interests selection */}
            <div className={styles.chip}>Sports</div>
            <div className={styles.chip}>Travel</div>
            <div className={styles.chip}>Music</div>
            <div className={styles.chip}>Art</div>
            <div className={styles.chip}>Food</div>
          </div>
        </div>

        {/* Music Preferences */}
        <div className={styles.section}>
          <h3 className={styles.sectionTitle}>Music Preferences</h3>
          <div className={styles.chipContainer}>
            {/* Placeholder for music preferences selection */}
            <div className={styles.chip}>Rock</div>
            <div className={styles.chip}>Pop</div>
            <div className={styles.chip}>Hip Hop</div>
            <div className={styles.chip}>Jazz</div>
            <div className={styles.chip}>Classical</div>
          </div>
        </div>

        {/* Verification Status */}
        <div className={styles.section}>
          <div className={styles.toggleContainer}>
            <h3 className={styles.sectionTitle}>Verification Status</h3>
            <label className={styles.switch}>
              <input
                type="checkbox"
                checked={verificationStatus}
                onChange={handleVerificationToggle}
              />
              <span className={styles.slider}></span>
            </label>
          </div>
          <p className={styles.verificationText}>
            Only show verified profiles
          </p>
        </div>

        {/* Continue Button */}
        <div className={styles.buttonContainer}>
          <CustomButton
            text="Continue"
            onClick={onClose}
          />
        </div>
      </Box>
    </Drawer>
  );
};

export default FiltersDrawer;
