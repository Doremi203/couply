import { Dialog, DialogTitle, DialogContent, DialogActions } from '@mui/material';
import React, { useState, useEffect } from 'react';

import { CustomButton } from '../../../../../../shared/components/CustomButton';
import {
  foodDrinkOptions,
  hobbyOptions,
  moviesTVOptions,
  musicOptions,
  personalityTraitsOptions,
  petsOptions,
  selfdevelopmentOptions,
  sportOptions,
} from '../../../constants';
import styles from '../../filtersDrawer.module.css';
import ChipFilter from '../ChipFilter';

type InterestFilterProps = {
  title: string;
  options: string[];
  selectedOptions: string[];
  onSelect: (selected: string[]) => void;
};

const InterestFilter: React.FC<InterestFilterProps> = ({
  title,
  options,
  selectedOptions,
  onSelect,
}) => {
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [tempSelected, setTempSelected] = useState<string[]>([]);

  useEffect(() => {
    if (isModalOpen) {
      setTempSelected(selectedOptions);
    }
  }, [isModalOpen, selectedOptions]);

  const handleOpenModal = () => setIsModalOpen(true);
  const handleCloseModal = () => setIsModalOpen(false);

  const handleSave = () => {
    onSelect(tempSelected);
    handleCloseModal();
  };

  const toggleInterest = (interest: string) => {
    setTempSelected(prev =>
      prev.includes(interest) ? prev.filter(item => item !== interest) : [...prev, interest],
    );
  };

  const getAllInterestOptions = () => {
    return [
      ...Object.values(sportOptions),
      ...Object.values(selfdevelopmentOptions),
      ...Object.values(hobbyOptions),
      ...Object.values(musicOptions),
      ...Object.values(moviesTVOptions),
      ...Object.values(foodDrinkOptions),
      ...Object.values(personalityTraitsOptions),
      ...Object.values(petsOptions),
    ];
  };

  return (
    <div className={styles.cont}>
      <h3 className={styles.sectionTitle}>{title}</h3>
      <div className={styles.section}>
        <div className={styles.chipContainer}>
          {selectedOptions
            .filter(option => getAllInterestOptions().includes(option))
            .map((option, index) => (
              <div
                key={index}
                className={`${styles.chip} ${selectedOptions.includes(option) ? styles.chipSelected : ''}`}
              >
                {option}
              </div>
            ))}
        </div>
      </div>

      <CustomButton
        text={selectedOptions.length ? 'Изменить' : 'Добавить'}
        onClick={handleOpenModal}
        className={styles.button}
        variant="outlined"
      />

      <Dialog fullScreen open={isModalOpen} onClose={handleCloseModal}>
        <DialogTitle sx={{ p: 1.5, borderBottom: '1px solid #eee', fontFamily: 'Jost' }}>
          {title}
        </DialogTitle>

        <DialogContent sx={{ p: 0, padding: '20px', marginTop: '10px' }} className={styles.content}>
          <ChipFilter
            title="Спорт"
            options={Object.values(sportOptions)}
            selectedOptions={tempSelected}
            onToggle={toggleInterest}
          />

          <ChipFilter
            title="Саморазвитие"
            options={Object.values(selfdevelopmentOptions)}
            selectedOptions={tempSelected}
            onToggle={toggleInterest}
          />

          <ChipFilter
            title="Хобби"
            options={Object.values(hobbyOptions)}
            selectedOptions={tempSelected}
            onToggle={toggleInterest}
          />

          <ChipFilter
            title="Музыка"
            options={Object.values(musicOptions)}
            selectedOptions={tempSelected}
            onToggle={toggleInterest}
          />

          <ChipFilter
            title="Фильмы"
            options={Object.values(moviesTVOptions)}
            selectedOptions={tempSelected}
            onToggle={toggleInterest}
          />

          <ChipFilter
            title="Еда"
            options={Object.values(foodDrinkOptions)}
            selectedOptions={tempSelected}
            onToggle={toggleInterest}
          />

          <ChipFilter
            title="Персональные качества"
            options={Object.values(personalityTraitsOptions)}
            selectedOptions={tempSelected}
            onToggle={toggleInterest}
          />

          <ChipFilter
            title="Домашние животные"
            options={Object.values(petsOptions)}
            selectedOptions={tempSelected}
            onToggle={toggleInterest}
          />
        </DialogContent>

        <DialogActions sx={{ p: 2, borderTop: '1px solid #eee' }}>
          <CustomButton text="Отмена" onClick={handleCloseModal} variant="outlined" />
          <CustomButton text="Сохранить" onClick={handleSave} sx={{ ml: 2 }} />
        </DialogActions>
      </Dialog>
    </div>
  );
};

export default InterestFilter;
