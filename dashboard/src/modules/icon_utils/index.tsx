import * as React from 'react';

// https://material.io/resources/icons/?style=baseline

import AddBoxIcon from '@material-ui/icons/AddBox';
import DeleteIcon from '@material-ui/icons/Delete';
import DetailsIcon from '@material-ui/icons/Details';
import EditIcon from '@material-ui/icons/Edit';

import ChatIcon from '@material-ui/icons/Chat';
import CloudIcon from '@material-ui/icons/Cloud';
import CloudCircleIcon from '@material-ui/icons/CloudCircle';
import CloudQueueIcon from '@material-ui/icons/CloudQueue';
import HelpOutlineIcon from '@material-ui/icons/HelpOutline';
import HomeIcon from '@material-ui/icons/Home';
import ReceiptIcon from '@material-ui/icons/Receipt';

function getIcon(icon) {
  switch (icon) {
    case 'Detail':
      return <DetailsIcon />;
    case 'Update':
      return <EditIcon />;
    case 'Create':
      return <AddBoxIcon />;
    case 'Delete':
      return <DeleteIcon />;
    default:
      return <span>IconNotFound</span>;
  }
}

function getServiceIcon(icon) {
  switch (icon) {
    case 'Home':
      return <HomeIcon key={icon} />;
    case 'HomeProject':
      return <HomeIcon key={icon} />;
    case 'Chat':
      return <ChatIcon key={icon} />;
    case 'Wiki':
      return <ReceiptIcon key={icon} />;
    case 'Ticket':
      return <ReceiptIcon key={icon} />;
    case 'ResourcePhysicalAdmin':
      return <CloudIcon key={icon} />;
    case 'ResourceVirtualAdmin':
      return <CloudQueueIcon key={icon} />;
    case 'ResourcePhysical':
      return <CloudIcon key={icon} />;
    case 'ResourceVirtual':
      return <CloudQueueIcon key={icon} />;
    case 'ResourceMonitor':
      return <CloudCircleIcon key={icon} />;
    default:
      return <HelpOutlineIcon key={icon} />;
  }
}

export default {
  getIcon,
  getServiceIcon,
};
