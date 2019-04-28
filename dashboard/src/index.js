import React from 'react';
import { render } from 'react-dom';

import Root from './containers/Root';

console.log(process.env.REACT_APP_API_SERVER_URL)

render(
  <Root />,
  document.getElementById('root')
);
