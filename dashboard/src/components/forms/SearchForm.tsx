import * as React from "react";
import { connect } from "react-redux";

import { Theme } from "@material-ui/core/styles/createMuiTheme";
import createStyles from "@material-ui/core/styles/createStyles";
import withStyles, {
  StyleRules,
  WithStyles
} from "@material-ui/core/styles/withStyles";

import Button from "@material-ui/core/Button";
import Grid from "@material-ui/core/Grid";
import TextField from "@material-ui/core/TextField";

import green from "@material-ui/core/colors/green";
import red from "@material-ui/core/colors/red";

import actions from "../../actions";
import form_utils from "../../lib/form_utils";
import logger from "../../lib/logger";

import Autocomplete from "@material-ui/lab/Autocomplete";

import { DateTimePicker } from "@material-ui/pickers";
import { MuiPickersUtilsProvider } from "@material-ui/pickers";

import DateFnsUtils from "@date-io/date-fns";

import SearchIcon from "@material-ui/icons/Search";

// import Icon from '../../../../components/icons/Icon';

interface ISearchForm extends WithStyles<typeof styles> {
  targets;
  routes;
  data;
  selected;
  index;
  isSubmitting;
  searchQueries: any;
  submitQueries;
  onSubmit;
  onChange;
}

class SearchForm extends React.Component<ISearchForm> {
  public state = {
    fieldMap: {},
    searchQueries: {}
  };

  public componentWillMount() {
    const searchQueries = form_utils.getSearchQueries();
    console.log("DEBUG searchQueries", searchQueries);
    this.setState({ searchQueries });
  }

  public render() {
    const { data, index, selected, isSubmitting } = this.props;
    const { searchQueries } = this.state;
    logger.info("SearchForm", "render", index, selected, isSubmitting);
    if (!index.Inputs) {
      logger.warning("SearchForm", "Not found index.Inputs");
      return null;
    }

    const inputs: any[] = [];
    for (let i = 0, len = index.Inputs.length; i < len; i++) {
      const input = index.Inputs[i];
      let defaultValue: any;
      console.log("TODO DEBUG SearchQueries", searchQueries);
      if (searchQueries && input.Name in searchQueries) {
        defaultValue = searchQueries[input.Name];
      } else {
        defaultValue = input.Default;
      }
      switch (input.Type) {
        case "Select":
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
            options
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
            </Grid>
          );
          break;
        case "Text":
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
            </Grid>
          );
          break;
        case "DateTime":
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
            </Grid>
          );
          break;
      }
    }

    return (
      <form onSubmit={this.handleInputSubmit}>
        <Grid container={true} direction="row" spacing={1}>
          {inputs}
          <Grid item={true}>
            <Button
              variant="outlined"
              type="submit"
              size="medium"
              color="primary"
              startIcon={<SearchIcon />}
            >
              Search
            </Button>
          </Grid>
        </Grid>
      </form>
    );
  }

  private handleSelectorChange = (event, name, values) => {
    const { searchQueries } = this.state;
    searchQueries[name] = values;
    this.setState({ searchQueries });
  };

  private handleInputSubmit = event => {
    event.preventDefault();
    const { routes, index, onSubmit } = this.props;
    const { searchQueries } = this.state;
    const route = routes[routes.length - 1];
    const location = route.location;
    const queryStr = encodeURIComponent(JSON.stringify(searchQueries));
    location.search = "q=" + queryStr;
    route.history.push(location);
    if (onSubmit) {
      onSubmit(event, index, searchQueries);
    }
  };

  private handleInputChange = event => {
    const { searchQueries } = this.state;
    searchQueries[event.target.name] = event.target.value;
    this.setState({ searchQueries });
  };

  private handleDateChange = (name: string, date: Date | null) => {
    const { searchQueries } = this.state;
    searchQueries[name] = date;
    this.setState({ searchQueries });
  };
}

function mapStateToProps(state, ownProps) {
  const { auth, service, isSubmitting, isSubmitSuccess } = state.service;
  return {
    auth,
    isSubmitSuccess,
    isSubmitting,
    service
  };
}

function mapDispatchToProps(dispatch, ownProps) {
  const { index } = ownProps;
  return {
    submitQueries: (items, fieldMap, route) => {
      dispatch(
        actions.service.serviceSubmitQueries({
          fieldMap,
          index,
          items,
          route
        })
      );
    }
  };
}

const styles = (theme: Theme): StyleRules =>
  createStyles({
    button: {
      margin: theme.spacing(1)
    },
    buttonFailed: {
      "&:hover": {
        backgroundColor: red[700]
      },
      backgroundColor: red[500]
    },
    buttonProgress: {
      color: green[500],
      left: "50%",
      marginLeft: -12,
      marginTop: -12,
      position: "absolute",
      top: "50%"
    },
    buttonSuccess: {
      "&:hover": {
        backgroundColor: green[700]
      },
      backgroundColor: green[500]
    },
    fabProgress: {
      color: green[500],
      left: -6,
      position: "absolute",
      top: -6,
      zIndex: 1
    },
    root: {
      width: "100%"
    },
    wrapper: {
      margin: theme.spacing(1),
      position: "relative"
    }
  });

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(withStyles(styles, { withTheme: true })(SearchForm));
