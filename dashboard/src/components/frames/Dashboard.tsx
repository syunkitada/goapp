import * as React from "react";
import { connect } from "react-redux";

import { Theme } from "@material-ui/core/styles/createMuiTheme";
import createStyles from "@material-ui/core/styles/createStyles";
import withStyles, {
  StyleRules,
  WithStyles
} from "@material-ui/core/styles/withStyles";

import AppBar from "@material-ui/core/AppBar";
import Avatar from "@material-ui/core/Avatar";
import ClickAwayListener from "@material-ui/core/ClickAwayListener";
import CssBaseline from "@material-ui/core/CssBaseline";
import Drawer from "@material-ui/core/Drawer";
import Hidden from "@material-ui/core/Hidden";
import IconButton from "@material-ui/core/IconButton";
import MenuItem from "@material-ui/core/MenuItem";
import MenuList from "@material-ui/core/MenuList";
import Popover from "@material-ui/core/Popover";
import Toolbar from "@material-ui/core/Toolbar";
import Typography from "@material-ui/core/Typography";

import { fade } from "@material-ui/core/styles/colorManipulator";

import ExitToAppIcon from "@material-ui/icons/ExitToApp";
import MenuIcon from "@material-ui/icons/Menu";

import LeftSidebar from "./LeftSidebar";

import actions from "../../actions";

const drawerWidth = 240;

const styles = (theme: Theme): StyleRules =>
  createStyles({
    appBar: {
      marginLeft: drawerWidth,
      [theme.breakpoints.up("md")]: {
        width: `calc(100% - ${drawerWidth}px)`
      }
    },
    appBarShift: {
      marginLeft: drawerWidth,
      transition: theme.transitions.create(["width", "margin"], {
        duration: theme.transitions.duration.enteringScreen,
        easing: theme.transitions.easing.sharp
      }),
      width: `calc(100% - ${drawerWidth}px)`
    },
    appBarSpacer: theme.mixins.toolbar,
    chartContainer: {
      marginLeft: -22
    },
    content: {
      flexGrow: 1,
      padding: theme.spacing(1),
      width: "100%",
      [theme.breakpoints.up("md")]: {
        marginLeft: drawerWidth
      }
    },
    drawerPaper: {
      width: drawerWidth
    },
    grow: {
      flexGrow: 1
    },
    inputInput: {
      paddingBottom: theme.spacing(1),
      paddingLeft: theme.spacing(10),
      paddingRight: theme.spacing(1),
      paddingTop: theme.spacing(1),
      transition: theme.transitions.create("width"),
      width: "100%",
      [theme.breakpoints.up("md")]: {
        width: 200
      }
    },
    inputRoot: {
      color: "inherit",
      width: "100%"
    },
    menuButton: {
      marginLeft: 12,
      marginRight: 36
    },
    menuButtonHidden: {
      display: "none"
    },
    navIconHide: {
      [theme.breakpoints.up("md")]: {
        display: "none"
      }
    },
    root: {
      display: "flex"
    },
    search: {
      "&:hover": {
        backgroundColor: fade(theme.palette.common.white, 0.25)
      },
      backgroundColor: fade(theme.palette.common.white, 0.15),
      borderRadius: theme.shape.borderRadius,
      marginLeft: 0,
      marginRight: theme.spacing(2),
      position: "relative",
      width: "100%",
      [theme.breakpoints.up("sm")]: {
        marginLeft: theme.spacing(3),
        width: "auto"
      }
    },
    searchIcon: {
      alignItems: "center",
      display: "flex",
      height: "100%",
      justifyContent: "center",
      pointerEvents: "none",
      position: "absolute",
      width: theme.spacing(9)
    },
    sectionDesktop: {
      display: "none",
      [theme.breakpoints.up("md")]: {
        display: "flex"
      }
    },
    sectionMobile: {
      display: "flex",
      [theme.breakpoints.up("md")]: {
        display: "none"
      }
    },
    tableContainer: {
      height: 320
    },
    title: {
      flexGrow: 1
    },
    toolbar: {
      paddingRight: 24 // keep right padding when drawer closed
    },
    toolbarIcon: {
      alignItems: "center",
      display: "flex",
      justifyContent: "flex-end",
      padding: "0 8px",
      ...theme.mixins.toolbar
    }
  });

interface IDashboard extends WithStyles<typeof styles> {
  children;
  auth;
  services;
  selectedServiceIndex;
  onClickLogout;
  projectService;
  match;
}

class Dashboard extends React.Component<IDashboard> {
  public state = {
    anchorEl: null,
    mobileOpen: false,
    open: true
  };

  public render() {
    const { anchorEl } = this.state;
    const {
      classes,
      children,
      auth,
      onClickLogout,
      services,
      selectedServiceIndex,
      projectService,
      match
    } = this.props;
    const isMenuOpen = Boolean(anchorEl);

    let userName: any;
    let userPhotoUrl: any;
    let title: any;
    let drawer: any;
    if (services) {
      const service = services[selectedServiceIndex];
      title = service.Name;
      userName = auth.user.displayName;
      userPhotoUrl = auth.user.photoURL;

      drawer = (
        <div>
          <div className={classes.toolbar} />
          <LeftSidebar
            services={services}
            selectedServiceIndex={selectedServiceIndex}
          />
        </div>
      );
    } else {
      title = match.url;
      userName = auth.user.Name;
      drawer = (
        <div>
          <div className={classes.toolbar} />
          <LeftSidebar
            projectService={projectService}
            match={match}
            history={history}
          />
        </div>
      );
    }

    return (
      <React.Fragment>
        <CssBaseline />
        <div className={classes.root}>
          <AppBar className={classes.appBar} position="fixed">
            <Toolbar>
              <IconButton
                color="inherit"
                aria-label="Open drawer"
                onClick={this.handleDrawerToggle}
                className={classes.navIconHide}
              >
                <MenuIcon />
              </IconButton>
              <Typography variant="subtitle1" color="inherit" noWrap={true}>
                {title}
              </Typography>

              <div className={classes.grow} />

              <IconButton
                aria-owns={isMenuOpen ? "menu-list-grow" : ""}
                aria-haspopup="true"
                color="inherit"
                onClick={this.handleMenuOpen}
              >
                <Avatar src={userPhotoUrl} />
                <span>&nbsp;&nbsp;</span>
                {userName}
              </IconButton>
              <Popover
                open={isMenuOpen}
                anchorEl={anchorEl}
                anchorOrigin={{
                  horizontal: "right",
                  vertical: "bottom"
                }}
                transformOrigin={{
                  horizontal: "right",
                  vertical: "top"
                }}
              >
                <ClickAwayListener onClickAway={this.handleMenuClose}>
                  <MenuList>
                    <MenuItem onClick={onClickLogout}>
                      <ExitToAppIcon /> Logout
                    </MenuItem>
                  </MenuList>
                </ClickAwayListener>
              </Popover>
            </Toolbar>
          </AppBar>

          <Hidden mdUp={true}>
            <Drawer
              variant="temporary"
              anchor={"left"}
              open={this.state.mobileOpen}
              onClose={this.handleDrawerToggle}
              classes={{
                paper: classes.drawerPaper
              }}
              ModalProps={{
                keepMounted: true // Better open performance on mobile.
              }}
            >
              {drawer}
            </Drawer>
          </Hidden>
          <Hidden smDown={true} implementation="css">
            <Drawer
              variant="permanent"
              open={true}
              classes={{
                paper: classes.drawerPaper
              }}
            >
              {drawer}
            </Drawer>
          </Hidden>

          <main className={classes.content}>
            <div className={classes.appBarSpacer} />
            {children}
          </main>
        </div>
      </React.Fragment>
    );
  }

  private handleDrawerToggle = () => {
    this.setState(state => ({ mobileOpen: !this.state.mobileOpen }));
  };

  private handleMenuOpen = event => {
    this.setState({ anchorEl: event.currentTarget });
  };

  private handleMenuClose = () => {
    this.setState({ anchorEl: null });
  };
}

function mapStateToProps(state, ownProps) {
  const auth = state.auth;

  return { auth };
}

function mapDispatchToProps(dispatch, ownProps) {
  return {
    onClickLogout: () => dispatch(actions.auth.authLogout())
  };
}

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(withStyles(styles)(Dashboard));
