import { Slider, styled } from '@mui/material';

// Custom styled slider with blue track
export const CustomSlider = styled(Slider)({
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