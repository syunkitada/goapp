import * as React from 'react';

import DetailsIcon from '@material-ui/icons/Details';
import EditIcon from '@material-ui/icons/Edit';
import AddBoxIcon from '@material-ui/icons/AddBox';
import DeleteIcon from '@material-ui/icons/Delete';


function getIcon(icon) {
  switch(icon) {
  case "Detail":
    return <DetailsIcon />;
  case "Update":
    return <EditIcon />;
  case "Create":
    return <AddBoxIcon />;
  case "Delete":
    return <DeleteIcon />;
  default:
    return <span>IconNotFound</span>
  }
}

export default {
  getIcon,
}
