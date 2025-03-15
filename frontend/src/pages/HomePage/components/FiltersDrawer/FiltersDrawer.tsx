import React, { useState } from "react";
import {
  Drawer,
  Box,
  Typography,
  Button,
  Slider,
  TextField,
  ToggleButtonGroup,
  ToggleButton,
  styled,
} from "@mui/material";
import styles from "./filtersDrawer.module.css";

type Props = {
  open: boolean;
  onClose: () => void;
};

function valuetext(value: number) {
  return `${value} km`;
}

const CustomToggleButton = styled(ToggleButton)(({ theme }) => ({
  "&.Mui-selected": {
    color: "#fff",
    backgroundColor: "#202C83", // Кастомный цвет для выбранного состояния
    "&:hover": {
      backgroundColor: "#202C83", // Цвет при наведении в выбранном состоянии
    },
    //padding: '10px'
    width: "100px",
    height: "60px",
    // borderRadius: 15
  },
  "&:not(.Mui-selected)": {
    color: "#000", // Цвет текста для невыбранного состояния
    //backgroundColor: 'white', // Кастомный цвет для невыбранного состояния
    // '&:hover': {
    //   backgroundColor: '#d5d5d5', // Цвет при наведении в невыбранном состоянии
    // },
    // borderColor: '#E8E6EA',
    // border: '0.5px solid #E8E6EA',
    border: "1px solid grey",
    width: "100px",
    height: "60px",
    //borderRadius: 15
  },
}));

const CustomToggleButtonGroup = styled(ToggleButtonGroup)({
  borderRadius: 15, // Установите нужное вам значение
  overflow: "hidden",
  //width: '600px',
  marginLeft: "20px",
});

const FilterDrawer: React.FC<Props> = ({ open, onClose }) => {
  const [interestedIn, setInterestedIn] = useState<string>("Girls");
  const [location, setLocation] = useState<string>("Chicago, USA");
  const [distance, setDistance] = useState<number>(40);
  const [ageRange, setAgeRange] = useState<number[]>([20, 28]);

  const handleDistanceChange = (event: Event, newValue: number | number[]) => {
    setDistance(newValue as number);
  };

  const handleAgeRangeChange = (event: Event, newValue: number | number[]) => {
    setAgeRange(newValue as number[]);
  };

  const [interest, setInterest] = useState<string>("Girls");

  const handleChange = (
    event: React.MouseEvent<HTMLElement>,
    newInterest: string | null
  ) => {
    if (newInterest !== null) {
      setInterest(newInterest);
    }
  };

  return (
    <Drawer anchor="top" open={open} onClose={onClose}>
      <Box sx={{ width: 300, padding: 2 }}>
        {/* <Typography variant="h6" textAlign="center">Filters</Typography> */}
        <h3 className={styles.header}>Фильтры</h3>
        <CustomToggleButtonGroup
          value={interest}
          exclusive
          onChange={handleChange}
          aria-label="interest"
        >
          <CustomToggleButton value="Girls" aria-label="girls">
            Женщины
          </CustomToggleButton>
          <CustomToggleButton value="Boys" aria-label="boys">
            Мужчины
          </CustomToggleButton>
          <CustomToggleButton value="Both" aria-label="both">
            Неважно
          </CustomToggleButton>
        </CustomToggleButtonGroup>
        <Box mt={2}>
          <Typography>Location</Typography>
          <TextField
            value={location}
            onChange={(e) => setLocation(e.target.value)}
            variant="outlined"
            fullWidth
          />
        </Box>
        <Box mt={2}>
          <Typography>Distance</Typography>
          <Slider
            value={distance}
            onChange={handleDistanceChange}
            aria-labelledby="distance-slider"
            getAriaValueText={valuetext}
            max={100}
          />
        </Box>
        <Box mt={2}>
          <Typography>Age</Typography>
          <Slider
            value={ageRange}
            onChange={handleAgeRangeChange}
            valueLabelDisplay="auto"
            aria-labelledby="range-slider"
            max={100}
          />
        </Box>
        <Box mt={3} textAlign="center">
          <Button variant="contained" color="primary" onClick={onClose}>
            Continue
          </Button>
        </Box>
      </Box>
    </Drawer>
  );
};

export default FilterDrawer;
