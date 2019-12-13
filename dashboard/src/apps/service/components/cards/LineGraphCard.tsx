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

import * as moment from 'moment';

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
    const {classes, data} = this.props;
    const colorPaletts = [
      '#8884d8',
      '#82ca9d',
      '#82ca9d',
      '#82ca9d',
      '#82ca9d',
      '#82ca9d',
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
            {data.Name}
          </Typography>
          <div style={{height: 300, padding: '0px 0px 60px 0px'}}>
            <ResponsiveContainer>
              <LineChart
                data={data.Values}
                syncId="system"
                margin={{
                  bottom: 0,
                  left: 0,
                  right: 30,
                  top: 10,
                }}>
                <CartesianGrid strokeDasharray="3 3" />
                <XAxis
                  dataKey="time"
                  type="number"
                  domain={['dataMin', 'dataMax']}
                  tickFormatter={t => moment(new Date(t)).format('MM/DD hh:mm')}
                />
                <YAxis
                  tickFormatter={y => {
                    if (y >= 1000000000) {
                      return y / 1000000000 + 'G';
                    } else if (y >= 1000000) {
                      return y / 1000000 + 'M';
                    } else if (y >= 1000) {
                      return y / 1000 + 'K';
                    }
                    return y;
                  }}
                />
                <Tooltip
                  labelFormatter={t =>
                    moment(new Date(t)).format('YYYY/MM/DD hh:mm')
                  }
                />
                <Legend />
                {data.Keys.map((key, index) => (
                  <Line
                    key={key}
                    type="monotone"
                    dataKey={key}
                    stroke={colorPaletts[index]}
                  />
                ))}
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
