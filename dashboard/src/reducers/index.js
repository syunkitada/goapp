import { combineReducers } from 'redux'

import auth from '../apps/auth/reducers';
import service from '../apps/service/reducers';
import home from '../apps/home/reducers';
import resourcePhysical from '../apps/project/resource/physical/reducers';
import resourceVirtual from '../apps/project/resource/virtual/reducers';
import monitor from '../apps/project/monitor/reducers';

export default combineReducers({
  auth,
  service,
  home,
  resourcePhysical,
  resourceVirtual,
  monitor,
});
