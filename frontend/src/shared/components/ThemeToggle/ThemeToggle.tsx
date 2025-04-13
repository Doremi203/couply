import DarkModeIcon from '@mui/icons-material/DarkMode';
import LightModeIcon from '@mui/icons-material/LightMode';
import React from 'react';

import { useTheme } from '../../lib/context/ThemeContext';
import { IconButton } from '../IconButton';

import styles from './themeToggle.module.css';

interface ThemeToggleProps {
  className?: string;
}

export const ThemeToggle: React.FC<ThemeToggleProps> = ({ className }) => {
  const { theme, toggleTheme } = useTheme();

  return (
    <div className={`${styles.themeToggle} ${className || ''}`}>
      <IconButton onClick={toggleTheme} touchFriendly={true}>
        {theme === 'dark' ? <LightModeIcon /> : <DarkModeIcon />}
      </IconButton>
    </div>
  );
};

export default ThemeToggle;
