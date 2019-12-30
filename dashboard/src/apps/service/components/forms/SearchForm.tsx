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

interface ISearchForm extends WithStyles<typeof styles> {
  targets;
  routes;
  data;
  index;
  searchQueries: any;
}

class SearchForm extends React.Component<ISearchForm> {
  public state = {
    searchQueries: {},
  };

  public render() {
    const {classes, index, selected, searchQueries} = this.props;
    logger.info('SearchForm', 'render', index, selected);

    const inputs: any[] = [];
    for (let i = 0, len = index.Inputs.length; i < len; i++) {
      const input = index.Inputs[i];
      let defaultValue = searchQueries[input.Name];
      if (!defaultValue) {
        defaultValue = input.Default;
      }
      switch (input.Type) {
        case 'Selector':
          let selectorData: any;
          if (input.DataKey) {
            selectorData = data[input.DataKey];
          } else if (input.Data) {
            selectorData = input.Data;
          }
          const options: any = [];
          if (!selectorData) {
            continue;
          }
          for (let j = 0, lenj = selectorData.length; j < lenj; j++) {
            options.push(selectorData[j]);
          }

          const selectorProps = {
            getOptionLabel: option => option,
            options,
          };

          inputs.push(
            <Grid item={true} key={input.Name}>
              <Autocomplete
                {...selectorProps}
                multiple={input.Multiple}
                disableCloseOnSelect={true}
                defaultValue={defaultValue}
                onChange={(event, values) =>
                  this.handleSelectorChange(event, input.Name, values)
                }
                renderInput={params => (
                  <TextField
                    {...params}
                    size="small"
                    label={input.Name}
                    variant="outlined"
                  />
                )}
              />
            </Grid>,
          );
          break;

        case 'Text':
          inputs.push(
            <Grid item={true} key={input.Name}>
              <TextField
                label={input.Name}
                defaultValue={defaultValue}
                variant="outlined"
                size="small"
                name={input.Name}
                onChange={this.handleInputChange}
              />
            </Grid>,
          );
          break;

        case 'DateTime':
          inputs.push(
            <Grid item={true} key={input.Name}>
              <MuiPickersUtilsProvider utils={DateFnsUtils}>
                <DateTimePicker
                  label={input.Name}
                  inputVariant="outlined"
                  size="small"
                  value={defaultValue}
                  format="yyyy/MM/dd HH:mm"
                  showTodayButton={true}
                  onChange={(date: Date | null) =>
                    this.handleDateChange(input.Name, date)
                  }
                />
              </MuiPickersUtilsProvider>
            </Grid>,
          );
          break;
      }
    }

    let inputsForm = (
      <form onSubmit={this.handleInputSubmit}>
        <Grid container={true} direction="row" spacing={1}>
          {inputs}
          <Grid item={true}>
            <Button
              variant="outlined"
              type="submit"
              size="medium"
              color="primary"
              startIcon={<SearchIcon />}>
              Submit
            </Button>
          </Grid>
        </Grid>
      </form>
    );
  }

  private handleSelectorChange = (event, name, values) => {
    const {searchQueries} = this.state;
    searchQueries[name] = values;
    this.setState({searchQueries});
  };

  private handleInputSubmit = event => {
    event.preventDefault();
    const {routes} = this.props;
    const {searchQueries} = this.state;
    const route = routes[routes.length - 1];
    const location = route.location;
    const queryStr = encodeURIComponent(JSON.stringify(searchQueries));
    location.search = 'q=' + queryStr;
    route.history.push(location);
  };

  private handleInputChange = event => {
    const {searchQueries} = this.state;
    searchQueries[event.target.name] = event.target.value;
    this.setState({searchQueries});
  };

  private handleDateChange = (name: string, date: Date | null) => {
    const {searchQueries} = this.state;
    searchQueries[name] = date;
    this.setState({searchQueries});
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
  const {queryKind, dataKind} = ownProps;
  return {
    submitQueries: (index, items, fieldMap, params) => {
      dispatch(
        actions.service.serviceSubmitQueries({
          action: index,
          dataKind,
          fieldMap,
          items,
          params,
          queryKind,
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
      margin: theme.spacing(1),
      position: 'relative',
    },
  });

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(withStyles(styles, {withTheme: true})(SearchForm));
