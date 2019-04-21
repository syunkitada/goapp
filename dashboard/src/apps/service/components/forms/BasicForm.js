import React, {Component} from 'react';
import { connect } from 'react-redux';
import { Link } from 'react-router-dom';

import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableRow from '@material-ui/core/TableRow';
import Checkbox from '@material-ui/core/Checkbox';
import Button from '@material-ui/core/Button';
import Menu from '@material-ui/core/Menu';
import MenuItem from '@material-ui/core/MenuItem';
import KeyboardArrowDownIcon from '@material-ui/icons/KeyboardArrowDown';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import TextField from '@material-ui/core/TextField';

import sort_utils from '../../../../modules/sort_utils'
import icon_utils from '../../../../modules/icon_utils'

class BasicForm extends Component {
  state = {
    order: 'asc',
    orderBy: 0,
    selected: [],
    data: [],
    page: 0,
    rowsPerPage: 5,
    searchRegExp: null,
    anchorEl: null,
		actionTarget: null,
		actionName: null,
    fieldMap: {},
  };

  handleTestFieldChange = (event, field) => {
    const { fieldMap } = this.state;
    let text = event.target.value
    let error = ""
    let re = new RegExp(field.RegExp)
    let len = text.length

    if (len < field.Min) {
      error += `Please enter ${field.Min} or more charactors. `
    } else if (len > field.Max) {
      error += `Please enter ${field.Max} or less charactors. `
    }
    if (!re.test(text)) {
      if (field.RegExpMsg) {
        error += field.RegExpMsg + " "
      } else {
        error += "Invalid characters. "
      }
    }

    fieldMap[field.Name] = {value: event.target.value, error: error, type: field.Type}

    this.setState({fieldMap: fieldMap})
  };

  handleSelectFieldChange = (event, field) => {
    const { fieldMap } = this.state;
    fieldMap[field.Name] = {value: event.target.value, error: null, type: field.Type}
    this.setState({fieldMap: fieldMap})
  };

  render() {
    const { routes, classes, index, data} = this.props
    const { fieldMap, selected, anchorEl, order, orderBy, rowsPerPage, page, searchRegExp, actionName, actionTarget } = this.state;

    let rawData = data[index.DataKey]
    if (!rawData) {
      return null
    }

    let fields = []
    for (let i = 0, len = index.Fields.length; i < len; i++) {
      let field = index.Fields[i]
      let fieldState = fieldMap[field.Name]
      let isError = false
      let helperText = ""
      if (fieldState) {
        if (fieldState.error !== "") {
          isError = true
          helperText = fieldState.error
        }
      }

      let autoFocus = false
      if (i === 0) {
        autoFocus = true
      }


      let f = fieldMap[field.Name]
      let value = ""
      if (f) {
        value = f.value
      } else {
        value = rawData[field.Name]
      }

      switch (field.Type) {
      case "text":
        fields.push(
          <TextField id={field.Name} key={field.Name}
            label={field.Name}
            autoFocus={autoFocus} margin="dense" type={field.Type} fullWidth
            onChange={event => {this.handleTestFieldChange(event, field)}}
            helperText={helperText}
            value={value}
            error={isError}
          />
        )
        break
      case "select":
        let options = field.Options
        if (!options) {
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
        if (!value || value === "") {
          value = options[0]
        }

        fields.push(
          <TextField
            select
            key={field.Name}
            label={field.Name}
            className={classes.textField}
            value={value}
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
      <div className={classes.root}>
        {fields}
      </div>
    );
  }
}

const styles = theme => ({
  root: {
    // margin: theme.spacing.unit * 2,
    width: '100%',
  },
  table: {
    width: '100%',
  },
  tableWrapper: {
    overflowX: 'auto',
  },
  margin: {
    // margin: theme.spacing.unit,
  },
  spacer: {
    flex: '1 1 100%',
  },
  actions: {
    color: theme.palette.text.secondary,
  },
  title: {
    flex: '0 0 auto',
  },
});

BasicForm.propTypes = {
  classes: PropTypes.object.isRequired,
};

function mapStateToProps(state, ownProps) {
  const auth = state.auth

  return {
    auth: auth,
  }
}

function mapDispatchToProps(dispatch, ownProps) {
  return {}
}

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(withStyles(styles)(BasicForm));
