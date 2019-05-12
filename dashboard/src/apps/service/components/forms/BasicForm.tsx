import * as React from 'react';
import {connect} from 'react-redux';

import {Theme} from '@material-ui/core/styles/createMuiTheme';
import createStyles from '@material-ui/core/styles/createStyles';
import withStyles, {
  StyleRules,
  WithStyles,
} from '@material-ui/core/styles/withStyles';

import Button from '@material-ui/core/Button';
import CircularProgress from '@material-ui/core/CircularProgress';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogTitle from '@material-ui/core/DialogTitle';
import Grid from '@material-ui/core/Grid';
import TextField from '@material-ui/core/TextField';

import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import DoneIcon from '@material-ui/icons/Done';

import green from '@material-ui/core/colors/green';
import red from '@material-ui/core/colors/red';

import actions from '../../../../actions';
import logger from '../../../../lib/logger';

import Icon from '../../../../components/icons/Icon';

interface IBasicForm extends WithStyles<typeof styles> {
  targets;
  routes;
  data;
  selected;
  index;
  onClose;
  isSubmitting;
  title;
  rawData;
  submitButtonName;
  submitQueries;
}

class BasicForm extends React.Component<IBasicForm> {
  public state = {
    fieldMap: {},
  };

  public render() {
    const {
      classes,
      index,
      selected,
      isSubmitting,
      title,
      submitButtonName,
      onClose,
    } = this.props;
    logger.info('BasicForm', 'render', index, selected);

    const fields = this.renderFields();

    return (
      <div className={classes.root}>
        {title && <DialogTitle id="form-dialog-title">{title}</DialogTitle>}
        <DialogContent>
          <DialogContentText>{index.Description}</DialogContentText>
          {fields}
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
              <Grid container={true} item={true} xs={6} justify="flex-end">
                {isSubmitting ? (
                  <CircularProgress
                    size={24}
                    className={classes.buttonProgress}
                  />
                ) : (
                  <Button
                    variant="contained"
                    color="primary"
                    disabled={isSubmitting}
                    className={classes.button}
                    onClick={this.handleActionSubmit}>
                    <Icon kind={index.Icon} marginDirection={'right'} />
                    {submitButtonName}
                  </Button>
                )}
              </Grid>
            </Grid>
          </div>
        </DialogActions>
      </div>
    );
  }

  private renderFields = () => {
    const {classes, data, selected, index, rawData} = this.props;
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
      let isError = false;
      let helperText = '';
      if (fieldState) {
        if (fieldState.error !== '') {
          isError = true;
          helperText = fieldState.error;
        }
      }

      let autoFocus = false;
      if (i === 0) {
        autoFocus = true;
      }

      let disabled = false;

      let value = '';
      if (fieldState) {
        value = fieldState.value;
      } else {
        if (rawData) {
          value = rawData[field.Name];
          if (!field.Updatable) {
            disabled = true;
          }
        }
      }

      switch (field.Type) {
        case 'text':
          fields.push(
            <TextField
              id={field.Name}
              key={i}
              disabled={disabled}
              label={field.Name}
              autoFocus={autoFocus}
              margin="dense"
              type={field.Type}
              fullWidth={true}
              onChange={event => {
                this.handleTextFieldChange(event, field);
              }}
              value={value}
              helperText={helperText}
              error={isError}
            />,
          );
          break;
        case 'select':
          let options = field.Options;
          if (!options) {
            options = [];
            const d = data[field.DataKey];
            if (d) {
              for (let j = 0, l = d.length; j < l; j++) {
                options.push(d[j].Name);
              }
            } else {
              options.push('');
            }
          }
          if (!value || value === '') {
            value = options[0];
          }

          fields.push(
            <TextField
              select={true}
              key={i}
              label={field.Name}
              className={classes.textField}
              disabled={disabled}
              value={value}
              onChange={event => {
                this.handleSelectFieldChange(event, field);
              }}
              SelectProps={{
                MenuProps: {
                  className: classes.menu,
                },
                native: true,
              }}
              helperText="Please select"
              margin="normal"
              fullWidth={true}>
              {options.map(option => (
                <option key={option} value={option}>
                  {option}
                </option>
              ))}
            </TextField>,
          );
          break;
        default:
          fields.push(<span>FieldNotFound</span>);
          break;
      }
    }

    return fields;
  };

  private handleTextFieldChange = (event, field) => {
    const {fieldMap} = this.state;

    const text = event.target.value;
    let error = '';
    const len = text.length;

    if (field.Min) {
      if (len < field.Min) {
        error += `Please enter ${field.Min} or more charactors. `;
      } else if (len > field.Max) {
        error += `Please enter ${field.Max} or less charactors. `;
      }
    }

    if (field.RegExp) {
      const re = new RegExp(field.RegExp);
      if (!re.test(text)) {
        if (field.RegExpMsg) {
          error += field.RegExpMsg + ' ';
        } else {
          error += 'Invalid characters. ';
        }
      }
    }

    fieldMap[field.Name] = {
      error,
      type: field.Type,
      value: event.target.value,
    };

    this.setState({fieldMap});
  };

  private handleSelectFieldChange = (event, field) => {
    const {fieldMap} = this.state;
    fieldMap[field.Name] = {
      error: null,
      type: field.Type,
      value: event.target.value,
    };
    this.setState({fieldMap});
  };

  private handleActionSubmit = () => {
    const {
      index,
      data,
      rawData,
      selected,
      routes,
      targets,
      submitQueries,
    } = this.props;
    const {fieldMap} = this.state;
    const route = routes.slice(-1)[0];
    const fieldData = {};

    if (index.Fields) {
      // Validate
      // フォーム入力がなく、デフォルト値がある場合はセットする
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

        switch (field.Type) {
          case 'text':
            if (field.Require) {
              if (!value || value === '') {
                fieldMap[field.Name] = {
                  error: 'This is required',
                  type: field.Type,
                  value: '',
                };
              }
              fieldData[field.Name] = {value};
            }
            break;

          case 'select':
            if (!value) {
              let options = field.Options;
              if (!options) {
                options = [];
                const d = data[field.DataKey];
                if (d) {
                  for (let j = 0, l = d.length; j < l; j++) {
                    options.push(d[j].Name);
                  }
                } else {
                  options.push('');
                }
              }
              value = options[0];
            }
            fieldData[field.Name] = {value};
            break;
          default:
            break;
        }
      }
    }

    const items: any[] = [];
    if (index.SelectKey && selected) {
      for (let i = 0, len = selected.length; i < len; i++) {
        const item = {};
        item[index.SelectKey] = selected[i];
        items.push(item);
      }
      if (items.length === 0) {
        // TODO handle unknown error
        return;
      }
    } else {
      items.push({});
    }

    for (const key in fieldMap) {
      if (fieldMap[key].error && fieldMap[key].error !== '') {
        this.setState({fieldMap});
        return;
      }
    }

    submitQueries(index, items, fieldData, targets, route.match.params);
  };
}

function mapStateToProps(state, ownProps) {
  const {isSubmitting, isSubmitSuccess} = state.service;
  return {
    isSubmitSuccess,
    isSubmitting,
  };
}

function mapDispatchToProps(dispatch, ownProps) {
  const {queryKind} = ownProps;
  return {
    submitQueries: (index, items, fieldMap, targets, params) => {
      dispatch(
        actions.service.serviceSubmitQueries({
          action: index,
          fieldMap,
          items,
          params,
          queryKind,
          targets,
        }),
      );
    },
  };
}

const styles = (theme: Theme): StyleRules =>
  createStyles({
    button: {
      margin: theme.spacing.unit,
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
      margin: theme.spacing.unit,
      position: 'relative',
    },
  });

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(withStyles(styles, {withTheme: true})(BasicForm));
