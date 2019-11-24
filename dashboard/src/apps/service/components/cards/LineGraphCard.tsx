import * as React from 'react';
import {connect} from 'react-redux';

import {Theme} from '@material-ui/core/styles/createMuiTheme';
import createStyles from '@material-ui/core/styles/createStyles';
import withStyles, {
  StyleRules,
  WithStyles,
} from '@material-ui/core/styles/withStyles';

import Button from '@material-ui/core/Button';
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import Typography from '@material-ui/core/Typography';

import {
  CartesianGrid,
  Legend,
  Line,
  LineChart,
  ResponsiveContainer,
  Tooltip,
  XAxis,
  YAxis,
} from 'recharts';

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

class LineGraphCard extends React.Component<IBasicView> {
  public render() {
    const {classes} = this.props;

    const data = [
      {
        amt: 2400,
        name: 'Page A',
        pv: 2400,
        uv: 4000,
      },
      {
        amt: 2210,
        name: 'Page B',
        pv: 1398,
        uv: 3000,
      },
      {
        amt: 2290,
        name: 'Page C',
        pv: 9800,
        uv: 2000,
      },
      {
        amt: 2000,
        name: 'Page D',
        pv: 3908,
        uv: 2780,
      },
      {
        amt: 2181,
        name: 'Page E',
        pv: 4800,
        uv: 1890,
      },
      {
        amt: 2500,
        name: 'Page F',
        pv: 3800,
        uv: 2390,
      },
      {
        amt: 2100,
        name: 'Page G',
        pv: 4300,
        uv: 3490,
      },
    ];

    return (
      <Card className={classes.card}>
        <CardContent>
          <Typography
            variant="subtitle1"
            color="inherit"
            noWrap={true}
            style={{
              borderBottom: '1px solid rgba(0, 0, 0, .125)',
              marginBottom: 10,
            }}>
            CPU Usage
          </Typography>
          <div style={{height: 300, padding: '0px 0px 60px 0px'}}>
            <ResponsiveContainer>
              <LineChart
                data={data}
                syncId="system"
                margin={{
                  bottom: 0,
                  left: 0,
                  right: 30,
                  top: 10,
                }}>
                <CartesianGrid strokeDasharray="3 3" />
                <XAxis dataKey="name" />
                <YAxis />
                <Tooltip />
                <Legend />
                <Line type="monotone" dataKey="uv" stroke="#8884d8" />
              </LineChart>
            </ResponsiveContainer>
          </div>
        </CardContent>
        <CardActions>
          <Button size="small">Learn More</Button>
        </CardActions>
      </Card>
    );
  }
}

function mapStateToProps(state, ownProps) {
  return {};
}

function mapDispatchToProps(dispatch, ownProps) {
  return {};
}

const styles = (theme: Theme): StyleRules =>
  createStyles({
    card: {
      height: 300,
      minWidth: 275,
    },
  });

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(withStyles(styles, {withTheme: true})(LineGraphCard));
