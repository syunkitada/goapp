import * as React from 'react';
import {connect} from 'react-redux';

import {Theme} from '@material-ui/core/styles/createMuiTheme';
import createStyles from '@material-ui/core/styles/createStyles';
import withStyles, {
  StyleRules,
  WithStyles,
} from '@material-ui/core/styles/withStyles';

import Button from '@material-ui/core/Button';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogTitle from '@material-ui/core/DialogTitle';
import Grid from '@material-ui/core/Grid';

import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableRow from '@material-ui/core/TableRow';

import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import DoneIcon from '@material-ui/icons/Done';

import green from '@material-ui/core/colors/green';
import red from '@material-ui/core/colors/red';

import logger from '../../../../lib/logger';

interface IBasicView extends WithStyles<typeof styles> {
  targets;
  routes;
  data;
  selected;
  index;
  onClose;
  isSubmitting;
  title;
  rawData;
  submitQueries;
}

class BasicView extends React.Component<IBasicView> {
  public state = {
    fieldMap: {},
  };

  public render() {
    const {classes, index, selected, isSubmitting, title, onClose} = this.props;
    logger.info('BasicView', 'render', index, selected);

    const fields = this.renderFields();

    return (
      <div className={classes.root}>
        {title && <DialogTitle id="form-dialog-title">{title}</DialogTitle>}
        <DialogContent>
          <DialogContentText>{index.Description}</DialogContentText>
          <Table className={classes.table}>
            <TableBody>{fields}</TableBody>
          </Table>
        </DialogContent>
        <DialogActions>
          <div className={classes.wrapper} style={{width: '100%'}}>
            <Grid container={true}>
              <Grid container={true} item={true} xs={6} justify="flex-start">
                {onClose && (
                  <Button onClick={onClose} disabled={isSubmitting}>
                    Cancel
                  </Button>
                )}
              </Grid>
              <Grid container={true} item={true} xs={6} justify="flex-end" />
            </Grid>
          </div>
        </DialogActions>
      </div>
    );
  }

  private renderFields = () => {
    const {selected, index, rawData} = this.props;
    const {fieldMap} = this.state;
    const fields: JSX.Element[] = [];

    if (selected) {
      const listItems: JSX.Element[] = [];
      for (let i = 0, len = selected.length; i < len; i++) {
        const s = selected[i];
        listItems.push(
          <ListItem key={s}>
            <ListItemIcon>
              <DoneIcon />
            </ListItemIcon>
            <ListItemText primary={s} />
          </ListItem>,
        );
      }

      fields.push(<List key={'selected'}>{listItems}</List>);
    }

    if (!index.Fields) {
      return fields;
    }

    for (let i = 0, len = index.Fields.length; i < len; i++) {
      const field = index.Fields[i];
      const fieldState = fieldMap[field.Name];

      let value = '';
      if (fieldState) {
        value = fieldState.value;
      } else {
        if (rawData) {
          value = rawData[field.Name];
        }
      }

      switch (field.Kind) {
        case 'text':
          fields.push(
            <TableRow key={field.Name}>
              <TableCell>{field.Name}</TableCell>
              <TableCell style={{width: '100%'}}>{value}</TableCell>
            </TableRow>,
          );
          break;
        case 'select':
          fields.push(
            <TableRow key={field.Name}>
              <TableCell>{field.Name}</TableCell>
              <TableCell style={{width: '100%'}}>{value}</TableCell>
            </TableRow>,
          );
          break;
        default:
          fields.push(<span>FieldNotFound</span>);
          break;
      }
    }

    return fields;
  };
}

function mapStateToProps(state, ownProps) {
  return {};
}

function mapDispatchToProps(dispatch, ownProps) {
  return {};
}

const styles = (theme: Theme): StyleRules =>
  createStyles({
    button: {
      margin: theme.spacing(1),
    },
    buttonFailed: {
      '&:hover': {
        backgroundColor: red[700],
      },
      backgroundColor: red[500],
    },
    buttonProgress: {
      color: green[500],
      left: '50%',
      marginLeft: -12,
      marginTop: -12,
      position: 'absolute',
      top: '50%',
    },
    buttonSuccess: {
      '&:hover': {
        backgroundColor: green[700],
      },
      backgroundColor: green[500],
    },
    fabProgress: {
      color: green[500],
      left: -6,
      position: 'absolute',
      top: -6,
      zIndex: 1,
    },
    root: {
      width: '100%',
    },
    wrapper: {
      margin: theme.spacing(1),
      position: 'relative',
    },
  });

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(withStyles(styles, {withTheme: true})(BasicView));