import darkTheme from './themes/darkTheme';
import lightTheme from './themes/lightTheme';

import grey from '@material-ui/core/colors/grey';
import orange from '@material-ui/core/colors/orange';
import red from '@material-ui/core/colors/red';

function getTheme(theme) {
  switch (theme) {
    case 'Dark':
      return darkTheme;
  }
  return lightTheme;
}

// https://material-ui.com/customization/color/#color-palette
function getBgColor(theme, color) {
  switch (color) {
    case 'Critical':
      return red[100];
    case 'Warning':
      return orange[100];
  }

  return grey[50];
}

function getFgColor(theme, color) {
  switch (color) {
    case 'Critical':
      return red[500];
    case 'Warning':
      return orange[500];
  }

  return grey[500];
}

export default {
  getBgColor,
  getFgColor,
  getTheme,
};
