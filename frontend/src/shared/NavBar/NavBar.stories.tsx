// NavBar.stories.js

import React from 'react';
import { BrowserRouter as Router } from 'react-router-dom';
import NavBar from './NavBar';

export default {
  title: 'Components/NavBar',
  component: NavBar,
};

const Template = (args) => (
  <Router>
    <NavBar {...args} />
  </Router>
);

export const Default = Template.bind({});
Default.args = {};

