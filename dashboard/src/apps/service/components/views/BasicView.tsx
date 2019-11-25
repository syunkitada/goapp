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
import Typography from '@material-ui/core/Typography';

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

import LineGraphCard from '../cards/LineGraphCard';

import ExpansionPanel from '@material-ui/core/ExpansionPanel';
import ExpansionPanelDetails from '@material-ui/core/ExpansionPanelDetails';
import ExpansionPanelSummary from '@material-ui/core/ExpansionPanelSummary';
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';

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
  handleChange;
}

class BasicView extends React.Component<IBasicView> {
  public state = {
    fieldMap: {},
  };

  public render() {
    const {classes, index, selected, isSubmitting, title, onClose} = this.props;
    // const {expanded} = this.state;
    logger.info('BasicView', 'render', index, selected);

    const fields = this.renderFields();
    const panels = this.renderPanels();

    return (
      <div className={classes.root}>
        {title && <DialogTitle id="form-dialog-title">{title}</DialogTitle>}
        <DialogContent>
          <DialogContentText>{index.Description}</DialogContentText>
          <Table className={classes.table}>
            <TableBody>{fields}</TableBody>
          </Table>
          {panels}
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

  private handleChange = (event, isExpanded) => {
    console.log('handleChange');
  };

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

  private renderPanels = () => {
    const {classes, index, rawData} = this.props;
    console.log('debug renderpanels', index, rawData);
    console.log('debug index', index.PanelsGroups);

    const panelsGroups: JSX.Element[] = [];
    for (let i = 0, len = index.PanelsGroups.length; i < len; i++) {
      console.log('debug panel');
      const panelsGroup = index.PanelsGroups[i];
      const panels: JSX.Element[] = [];
      if (panelsGroup.DataType === 'MetricsGroups') {
        console.log('debug data', rawData[panelsGroup.DataKey]);
        const metricsGroups = rawData[panelsGroup.DataKey];
        for (let j = 0, jlen = metricsGroups.length; j < jlen; j++) {
          const metricsGroup = metricsGroups[j];
          const cards: JSX.Element[] = [];
          for (let x = 0, xlen = metricsGroup.Metrics.length; x < xlen; x++) {
            const metric = metricsGroup.Metrics[x];
            console.log('DEBUG metrics', metric);
            cards.push(
              <Grid key={metric.Name} item={true} xs={6}>
                <LineGraphCard data={metric} />
              </Grid>,
            );
          }
          panels.push(
            <ExpansionPanel
              key={metricsGroup.Name}
              expanded={true}
              onChange={this.handleChange}
              className={classes.expansionPanel}>
              <ExpansionPanelSummary
                expandIcon={<ExpandMoreIcon />}
                aria-controls="panel1bh-content"
                id="panel1bh-header"
                className={classes.expansionPanelSummary}>
                <Typography variant="subtitle1">{metricsGroup.Name}</Typography>
              </ExpansionPanelSummary>
              <ExpansionPanelDetails className={classes.expansionPanelDetail}>
                <Grid container={true} spacing={2}>
                  {cards}
                </Grid>
              </ExpansionPanelDetails>
            </ExpansionPanel>,
          );
        }
      }

      panelsGroups.push(
        <div key={panelsGroup.Name}>
          <hr />
          <Typography variant="subtitle1">{panelsGroup.Name}</Typography>
          {panels}
        </div>,
      );
    }

    return <div>{panelsGroups}</div>;
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
    expansionPanel: {
      border: '1px solid rgba(0, 0, 0, .125)',
      boxShadow: 'none',
    },
    expansionPanelDetail: {
      boxShadow: 'none',
    },
    expansionPanelSummary: {
      borderBottom: '1px solid rgba(0, 0, 0, .125)',
      boxShadow: 'none',
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
