import React, {Component} from 'react';
import { connect } from 'react-redux';

import { withStyles } from '@material-ui/core/styles';

import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import Grid from '@material-ui/core/Grid';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogTitle from '@material-ui/core/DialogTitle';

import classNames from 'classnames';
import CircularProgress from '@material-ui/core/CircularProgress';
import green from '@material-ui/core/colors/green';
import red from '@material-ui/core/colors/red';

import actions from '../../../../actions'


class FormDialog extends Component {
  state = {
    fieldMap: {},
  };

  componentWillMount() {
    console.log("FormDialog.componentWillMount")
  }

  handleTestFieldChange = (event, field) => {
    console.log("DEBUG handleTestFieldChange", field.Name, event.target.value)
    const { fieldMap } = this.state;

    // TODO Validate
    fieldMap[field.Name] = {value: event.target.value, error: null, type: field.Type}

    this.setState({fieldMap: fieldMap})
  };

  handleSelectFieldChange = (event, field) => {
    console.log("DEBUG handleSelectFieldChange", field.Name, event.target.value)
    const { fieldMap } = this.state;

    // TODO Validate
    fieldMap[field.Name] = {value: event.target.value, error: null, type: field.Type}

    this.setState({fieldMap: fieldMap})
  };

  handleActionSubmit = () => {
    const { action, routes, targets, submitQueries } = this.props
    const { fieldMap } = this.state
    console.log("DEBUG handleActionSubmit", action.Name, action.DataKind, fieldMap, targets)
    let route = routes.slice(-1)[0]
    console.log(route)
    submitQueries(action, fieldMap, targets, route.match.params)
  };

  render() {
    const { classes, data, open, action, onClose, isSubmitting, isSubmitSuccess, isSubmitFailed } = this.props
    const { fieldMap } = this.state;

    const buttonClassname = classNames({
      [classes.buttonSuccess]: isSubmitSuccess,
    });

		console.log("DEBUG FormDialog", open)

    let fields = []
    for (let i = 0, len = action.Fields.length; i < len; i++) {
      let field = action.Fields[i]
      let autoFocus = false
      if (i === 0) {
        autoFocus = true
      }
      switch (field.Type) {
      case "text":
        fields.push(
          <TextField key={field.Name} id={field.Name} label={field.Name}
            autoFocus={autoFocus} margin="dense" type={field.Type} fullWidth
            onChange={event => {this.handleTestFieldChange(event, field)}}
          />
        )
        break
      case "select":
        let f = fieldMap[field.Name]
        let options = field.Options
        if (!options) {
          console.log(data)
          options = []
          let d = data[field.DataKey]
          if (d) {
            for (let j = 0, l = d.length; j < l; j++) {
              options.push(d[j].Name)
            }
          } else {
            options.push("")
          }
        }
        if (!f) {
          f = options[0]
        }

        console.log(f)
        fields.push(
          <TextField
            select
            key={field.Name}
            label={field.Name}
            className={classes.textField}
            value={f.value}
            onChange={event => {this.handleSelectFieldChange(event, field)}}
            SelectProps={{
              native: true,
              MenuProps: {
                className: classes.menu,
              },
            }}
            helperText="Please select"
            margin="normal"
            fullWidth
          >
            {options.map(option => (
              <option key={option} value={option}>
                {option}
              </option>
            ))}
          </TextField>
        )
        break
      default:
        fields.push(
          <span>FieldNotFound</span>
        )
        break
      }
    }

		return (
        <div>
          <Dialog
            open={open}
            onClose={onClose}
            aria-labelledby="form-dialog-title"
          >
            <DialogTitle id="form-dialog-title">{ action.Name } { action.DataKind }</DialogTitle>
            <DialogContent>
              <DialogContentText>{ action.Description }</DialogContentText>
              {fields}
            </DialogContent>
            <DialogActions>
              <div className={classes.wrapper} style={{width: '100%'}}>
                <Grid container>
                  <Grid container item xs={4} justify="flex-start">
                    <Button onClick={onClose} disabled={isSubmitting}>
                      Cancel
                    </Button>
                  </Grid>
                  <Grid container item xs={4} justify="center">
                    {isSubmitting && <CircularProgress size={24} className={classes.buttonProgress} />}
                    {isSubmitSuccess && <Button
                      variant="contained"
                      color="primary"
                      className={classes.buttonSuccess} >Success</Button> }
                    {isSubmitFailed && <Button
                      variant="contained"
                      color="primary"
                      className={classes.buttonFailed} >Failed</Button> }
                  </Grid>
                  <Grid container item xs={4} justify="flex-end">
                    <Button
                      variant="contained"
                      color="primary"
                      disabled={isSubmitting}
                      onClick={this.handleActionSubmit}
                    >
                      Submit
                    </Button>
                  </Grid>
                </Grid>
              </div>

            </DialogActions>
          </Dialog>
        </div>
      );
  }
}

function mapStateToProps(state, ownProps) {
  const { isSubmitting, isSubmitSuccess, isSubmitFailed } = state.service;
  return {
    isSubmitting: isSubmitting,
    isSubmitSuccess: isSubmitSuccess,
    isSubmitFailed: isSubmitFailed,
  }
}

function mapDispatchToProps(dispatch, ownProps) {
  return {
    submitQueries: (action, fieldMap, targets, params) => {
			console.log("DEBUG submitQueries")
      dispatch(actions.service.serviceSubmitQueries(action, fieldMap, targets, params));
    }
  }
}

const styles = theme => ({
  root: {
    display: 'flex',
    alignItems: 'center',
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
  mapDispatchToProps
)(withStyles(styles, {withTheme: true})(FormDialog));
