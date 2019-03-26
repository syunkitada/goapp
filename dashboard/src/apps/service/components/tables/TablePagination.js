import React, {Component} from 'react';
import { connect } from 'react-redux';

import OriginTablePagination from '@material-ui/core/TablePagination';

import TablePaginationActions from './TablePaginationActions'

class TablePagination extends Component {
  render() {
    const { count, rowsPerPage, page, onChangePage, onChangeRowsPerPage } = this.props

    return (
      <OriginTablePagination
        rowsPerPageOptions={[5, 10, 25]}
        component="span"
        count={count}
        rowsPerPage={rowsPerPage}
        page={page}
        backIconButtonProps={{
          'aria-label': 'Previous Page',
        }}
        nextIconButtonProps={{
          'aria-label': 'Next Page',
        }}
        onChangePage={onChangePage}
        onChangeRowsPerPage={onChangeRowsPerPage}
        ActionsComponent={TablePaginationActions}
      />
    );
  }
}

function mapStateToProps(state, ownProps) {
  return {}
}

function mapDispatchToProps(dispatch, ownProps) {
  return {}
}

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(TablePagination);
