import darkTheme from './themes/darkTheme';
import lightTheme from './themes/lightTheme';

import cyan from '@material-ui/core/colors/cyan';
import green from '@material-ui/core/colors/green';
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
    case 'Success':
      return green[100];
    case 'Critical':
      return red[100];
    case 'Warning':
      return orange[100];
    case 'Silenced':
      return cyan[100];
  }

  return grey[50];
}

function getFgColor(theme, color) {
  switch (color) {
    case 'Success':
      return green[300];
    case 'Critical':
      return red[500];
    case 'Warning':
      return orange[500];
    case 'Silenced':
      return cyan[300];
  }

  return grey[500];
}

export default {
  getBgColor,
  getFgColor,
  getTheme,
};
