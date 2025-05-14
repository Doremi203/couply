import React from 'react';

import { CustomInput } from '../../../../shared/components/CustomInput';
import { ToggleButtons } from '../../../../shared/components/ToggleButtons';
import { ProfileData } from '../../types';

import styles from './basicInfoForm.module.css';

interface BasicInfoFormProps {
  profileData: ProfileData;
  onInputChange: (field: string, value: string) => void;
}

export const BasicInfoForm: React.FC<BasicInfoFormProps> = ({ profileData, onInputChange }) => {
  return (
    <div className={styles.editSection}>
      <h3>Basic Information</h3>
      <div className={styles.editField}>
        <label>Name</label>
        <CustomInput
          type="text"
          placeholder="Your name"
          value={profileData.name}
          onChange={e => onInputChange('name', e.target.value)}
        />
      </div>
      <div className={styles.editField}>
        <label>Age</label>
        <CustomInput
          type="number"
          placeholder="Your age"
          value={profileData.age.toString()}
          onChange={e => onInputChange('age', e.target.value)}
        />
      </div>
      <div className={styles.editField}>
        <label>Date of Birth</label>
        <CustomInput
          type="date"
          placeholder="Date of birth"
          value={profileData.dateOfBirth}
          onChange={e => onInputChange('dateOfBirth', e.target.value)}
        />
      </div>
      <div className={styles.editField}>
        <label>Phone</label>
        <CustomInput
          type="tel"
          placeholder="Your phone number"
          value={profileData.phone}
          onChange={e => onInputChange('phone', e.target.value)}
        />
      </div>
      <div className={styles.editField}>
        <label>Email</label>
        <CustomInput
          type="email"
          placeholder="Your email"
          value={profileData.email}
          onChange={e => onInputChange('email', e.target.value)}
        />
      </div>
      <div className={styles.editField}>
        <label>Gender</label>
        <ToggleButtons
          options={[
            { label: 'Female', value: 'female' },
            { label: 'Male', value: 'male' },
          ]}
          value={profileData.gender}
          onSelect={value => onInputChange('gender', value)}
        />
      </div>
    </div>
  );
};

export default BasicInfoForm;
