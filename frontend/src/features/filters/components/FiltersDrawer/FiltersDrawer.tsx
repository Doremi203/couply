import { Box } from '@mui/material';
import React, { useState, useEffect } from 'react';

import {
  useGetFilterQuery,
  useUpdateFilterMutation,
  useSearchUsersMutation,
} from '../../../../entities/search';
import { GenderPriority } from '../../../../entities/search/api/constants';
import { FilterResponse } from '../../../../entities/search/api/types';
import { Alcohol, Education, Smoking, Zodiac, Children, Goal } from '../../../../entities/user';
import { mapFiltersToApi } from '../../helpers/mapFiltersToApiFormat';
import { mapInterestsFromApiFormat } from '../../helpers/mapInterestsFromApiFormat';
import { mapInterestsToApiFormat } from '../../helpers/mapInterestsToApiFormat';
import {
  alcoholFromApi,
  alcoholOptions,
  alcoholToApi,
  childrenFromApi,
  childrenOptions,
  childrenToApi,
  educationFromApi,
  educationOptions,
  educationToApi,
  genderFromApi,
  genderOptions,
  genderToApi,
  goalFromApi,
  goalOptions,
  goalToApi,
  smokingFromApi,
  smokingOptions,
  smokingToApi,
  zodiacFromApi,
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
  initialFilterData?: FilterResponse;
};

export const FiltersDrawer: React.FC<Props> = ({ open, onClose, initialFilterData }) => {
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

  const { refetch } = useGetFilterQuery({}, { skip: !open });

  useEffect(() => {
    if (open) {
      refetch();
    }
  }, [open, refetch]);

  const [updateFilter] = useUpdateFilterMutation();
  const [searchUsers] = useSearchUsersMutation();

  const [selectedGender, setSelectedGender] = useState<string[]>([]);
  const [distance, setDistance] = useState<number>(40);
  const [ageRange, setAgeRange] = useState<number[]>([18, 28]);
  const [heightRange, setHeightRange] = useState<number[]>([170, 190]);
  const [selectedInterests, setSelectedInterests] = useState<string[]>([]);

  const [selectedZodiac, setSelectedZodiac] = useState<string[]>([]);
  const [selectedGoal, setSelectedGoal] = useState<string[]>([]);
  const [verificationStatus, setVerificationStatus] = useState<boolean>(false);
  const [premiumStatus, setPremiumStatus] = useState<boolean>(false);

  useEffect(() => {
    if (initialFilterData?.filter) {
      const filter = initialFilterData.filter;
      setSelectedGender([genderFromApi[filter.genderPriority]]);
      //@ts-ignore
      setDistance(filter.distanceKmRange.max);
      setAgeRange([filter.ageRange.min, filter.ageRange.max]);
      setHeightRange([filter.heightRange.min, filter.heightRange.max]);
      setVerificationStatus(filter.onlyVerified);
      setPremiumStatus(filter.onlyPremium);

      setSelectedInterests(mapInterestsFromApiFormat(filter.interest));

      if (filter.goal && goalFromApi[filter.goal]) {
        setSelectedGoal([goalFromApi[filter.goal]]);
      }

      if (filter.zodiac && zodiacFromApi[filter.zodiac]) {
        setSelectedZodiac([zodiacFromApi[filter.zodiac]]);
      }
      if (filter.education && educationFromApi[filter.education]) {
        setSelectedEducation([educationFromApi[filter.education]]);
      }
      if (filter.children && childrenFromApi[filter.children]) {
        setSelectedChildren([childrenFromApi[filter.children]]);
      }
      if (filter.alcohol && alcoholFromApi[filter.alcohol]) {
        setSelectedAlcohol([alcoholFromApi[filter.alcohol]]);
      }
      if (filter.smoking && smokingFromApi[filter.smoking]) {
        setSelectedSmoking([smokingFromApi[filter.smoking]]);
      }
    }
  }, [initialFilterData]);

  const handleDistanceChange = (_event: Event, newValue: number | number[]) => {
    setDistance(newValue as number);
  };

  const handleAgeRangeChange = (_event: Event, newValue: number | number[]) => {
    setAgeRange(newValue as number[]);
  };

  const handleHeightRangeChange = (_event: Event, newValue: number | number[]) => {
    setHeightRange(newValue as number[]);
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

  const handleGenderToggle = (value: string) => {
    setSelectedGender([value]);
  };

  const handleZodiacToggle = (value: string) => {
    setSelectedZodiac(prev => (prev.includes(value) ? [] : [value]));
  };

  const handleEducationToggle = (value: string) => {
    setSelectedEducation(prev => (prev.includes(value) ? [] : [value]));
  };

  const handleChildrenToggle = (value: string) => {
    setSelectedChildren(prev => (prev.includes(value) ? [] : [value]));
  };

  const handleAlcoholToggle = (value: string) => {
    setSelectedAlcohol(prev => (prev.includes(value) ? [] : [value]));
  };

  const handleSmokingToggle = (value: string) => {
    setSelectedSmoking(prev => (prev.includes(value) ? [] : [value]));
  };

  const handleGoalToggle = (value: string) => {
    setSelectedGoal(prev => (prev.includes(value) ? [] : [value]));
  };

  const handleClearFilters = () => {
    setSelectedGender(['Any']);
    setDistance(100);
    setAgeRange([18, 100]);
    setHeightRange([100, 250]);
    setSelectedInterests([]);
    setVerificationStatus(false);
    setSelectedEducation([]);
    setSelectedAlcohol([]);
    setSelectedChildren([]);
    setSelectedSmoking([]);
    setSelectedZodiac([]);
    setPremiumStatus(false);
    setVerificationStatus(false);
    setSelectedGoal([]);
  };

  if (!open) return null;

  const handleApplyFilters = async () => {
    const filterData = {
      genderPriority: mapFiltersToApi(selectedGender, genderToApi, GenderPriority.any)[0],
      minAge: ageRange[0],
      maxAge: ageRange[1],
      minHeight: heightRange[0],
      maxHeight: heightRange[1],
      minDistanceKm: 0,
      maxDistanceKm: distance,
      goal: mapFiltersToApi(selectedGoal, goalToApi, Goal.unspecified)[0],
      zodiac: mapFiltersToApi(selectedZodiac, zodiacToApi, Zodiac.unspecified)[0],
      education: mapFiltersToApi(selectedEducation, educationToApi, Education.unspecified)[0],
      children: mapFiltersToApi(selectedChildren, childrenToApi, Children.unspecified)[0],
      alcohol: mapFiltersToApi(selectedAlcohol, alcoholToApi, Alcohol.unspecified)[0],
      smoking: mapFiltersToApi(selectedSmoking, smokingToApi, Smoking.unspecified)[0],
      interest: mapInterestsToApiFormat(selectedInterests),
      onlyVerified: verificationStatus,
      onlyPremium: premiumStatus,
    };

    try {
      //@ts-ignore
      await updateFilter(filterData).unwrap();
      await searchUsers({
        offset: 0,
        limit: 20,
      }).unwrap();
      onClose();
    } catch (error) {
      console.error('Error updating filters:', error);
    }
  };

  return (
    <div className={styles.modalOverlay}>
      <div className={styles.modalContainer}>
        <Box className={styles.content}>
          <FilterHeader onBack={onClose} onClear={handleClearFilters} />

          <GenderFilter
            title="Пол"
            options={Object.values(genderOptions)}
            selectedOptions={selectedGender}
            onToggle={handleGenderToggle}
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
            onToggle={handleGoalToggle}
          />

          <InterestFilter
            title="Интересы"
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
