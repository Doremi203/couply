import { Slider, styled } from '@mui/material';

export const CustomSlider = styled(Slider)({
  color: 'var(--primary-color)',
  height: 5,
  '& .MuiSlider-track': {
    border: 'none',
    backgroundColor: 'var(--primary-color)',
  },
  '& .MuiSlider-thumb': {
    height: 15,
    width: 15,
    backgroundColor: 'var(--primary-color)',
    boxShadow: '0 3px 6px var(--shadow-color)',
    '&:focus, &:hover, &.Mui-active, &.Mui-focusVisible': {
      boxShadow: '0 3px 8px var(--shadow-color)',
    },
  },
  '& .MuiSlider-rail': {
    backgroundColor: 'var(--button-background)',
  },
});
