import * as React from 'react';
import {connect} from 'react-redux';

import {Theme} from '@material-ui/core/styles/createMuiTheme';
import withStyles, {
  WithStyles,
  StyleRules,
} from '@material-ui/core/styles/withStyles';
import createStyles from '@material-ui/core/styles/createStyles';

import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import DialogActions from '@material-ui/core/DialogActions';
import Grid from '@material-ui/core/Grid';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogTitle from '@material-ui/core/DialogTitle';

import CircularProgress from '@material-ui/core/CircularProgress';
import green from '@material-ui/core/colors/green';
import red from '@material-ui/core/colors/red';

import actions from '../../../../actions';

interface IBasicForm extends WithStyles<typeof styles> {
  targets;
  routes;
  data;
  index;
  onClose;
  isSubmitting;
  title;
  rawData;
  submitButtonName;
  submitQueries;
}

class BasicForm extends React.Component<IBasicForm> {
  state = {
    fieldMap: {},
  };

  handleTextFieldChange = (event, field) => {
    const {fieldMap} = this.state;
    let text = event.target.value;
    let error = '';
    let re = new RegExp(field.RegExp);
    let len = text.length;

    if (len < field.Min) {
      error += `Please enter ${field.Min} or more charactors. `;
    } else if (len > field.Max) {
      error += `Please enter ${field.Max} or less charactors. `;
    }
    if (!re.test(text)) {
      if (field.RegExpMsg) {
        error += field.RegExpMsg + ' ';
      } else {
        error += 'Invalid characters. ';
      }
    }

    fieldMap[field.Name] = {
      value: event.target.value,
      error: error,
      type: field.Type,
    };

    this.setState({fieldMap: fieldMap});
  };

  handleSelectFieldChange = (event, field) => {
    const {fieldMap} = this.state;
    fieldMap[field.Name] = {
      value: event.target.value,
      error: null,
      type: field.Type,
    };
    this.setState({fieldMap: fieldMap});
  };

  handleActionSubmit = () => {
    const {index, routes, targets, submitQueries} = this.props;
    const {fieldMap} = this.state;
    let route = routes.slice(-1)[0];

    // Validate
    // フォーム入力がなく、デフォルト値がある場合はセットする
    for (let i = 0, len = index.Fields.length; i < len; i++) {
      let field = index.Fields[i];
      switch (field.Type) {
        case 'text':
          if (field.Require) {
            if (!fieldMap[field.Name] || fieldMap[field.Name] === '') {
              fieldMap[field.Name] = {
                value: '',
                error: 'This is required',
                type: field.Type,
              };
            }
          }
          break;
        case 'select':
          if (!fieldMap[field.Name]) {
            fieldMap[field.Name] = {
              value: field.Options[0],
              error: null,
              type: field.Type,
            };
          }
          break;
        default:
          break;
      }
    }

    for (let key in fieldMap) {
      if (fieldMap[key].error && fieldMap[key].error !== '') {
        this.setState({fieldMap: fieldMap});
        return;
      }
    }

    submitQueries(index, fieldMap, targets, route.match.params);
  };

  render() {
    const {
      classes,
      data,
      index,
      onClose,
      isSubmitting,
      title,
      rawData,
      submitButtonName,
    } = this.props;
    const {fieldMap} = this.state;

    const fields: JSX.Element[] = [];
    for (let i = 0, len = index.Fields.length; i < len; i++) {
      let field = index.Fields[i];
      let fieldState = fieldMap[field.Name];
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
          fields.push(
            <TextField
              id={field.Name}
              key={field.Name}
              label={field.Name}
              autoFocus={autoFocus}
              margin="dense"
              type={field.Type}
              fullWidth
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
            let d = data[field.DataKey];
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
              select
              key={field.Name}
              label={field.Name}
              className={classes.textField}
              value={value}
              onChange={event => {
                this.handleSelectFieldChange(event, field);
              }}
              SelectProps={{
                native: true,
                MenuProps: {
                  className: classes.menu,
                },
              }}
              helperText="Please select"
              margin="normal"
              fullWidth>
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

    return (
      <div className={classes.root}>
        {title && <DialogTitle id="form-dialog-title">{title}</DialogTitle>}
        <DialogContent>
          <DialogContentText>{index.Description}</DialogContentText>
          {fields}
        </DialogContent>
        <DialogActions>
          <div className={classes.wrapper} style={{width: '100%'}}>
            <Grid container>
              <Grid container item xs={4} justify="flex-start">
                {onClose && (
                  <Button onClick={onClose} disabled={isSubmitting}>
                    Cancel
                  </Button>
                )}
              </Grid>
              <Grid container item xs={4} justify="center">
                {isSubmitting && (
                  <CircularProgress
                    size={24}
                    className={classes.buttonProgress}
                  />
                )}
              </Grid>
              <Grid container item xs={4} justify="flex-end">
                <Button
                  variant="contained"
                  color="primary"
                  disabled={isSubmitting}
                  onClick={this.handleActionSubmit}>
                  {submitButtonName}
                </Button>
              </Grid>
            </Grid>
          </div>
        </DialogActions>
      </div>
    );
  }
}

function mapStateToProps(state, ownProps) {
  const {isSubmitting, isSubmitSuccess} = state.service;
  return {
    isSubmitting: isSubmitting,
    isSubmitSuccess: isSubmitSuccess,
  };
}

function mapDispatchToProps(dispatch, ownProps) {
  const {queryKind} = ownProps;
  return {
    submitQueries: (index, fieldMap, targets, params) => {
      dispatch(
        actions.service.serviceSubmitQueries({
          queryKind: queryKind,
          action: index,
          fieldMap: fieldMap,
          targets: targets,
          params: params,
        }),
      );
    },
  };
}

const styles = (theme: Theme): StyleRules =>
  createStyles({
    root: {
      width: '100%',
    },
    wrapper: {
      margin: theme.spacing.unit,
      position: 'relative',
    },
    buttonSuccess: {
      backgroundColor: green[500],
      '&:hover': {
        backgroundColor: green[700],
      },
    },
    buttonFailed: {
      backgroundColor: red[500],
      '&:hover': {
        backgroundColor: red[700],
      },
    },
    fabProgress: {
      color: green[500],
      position: 'absolute',
      top: -6,
      left: -6,
      zIndex: 1,
    },
    buttonProgress: {
      color: green[500],
      position: 'absolute',
      top: '50%',
      left: '50%',
      marginTop: -12,
      marginLeft: -12,
    },
  });

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(withStyles(styles, {withTheme: true})(BasicForm));
