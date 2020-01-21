import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import AndroidIcon from '@material-ui/icons/Android';

const useStyles = makeStyles(theme => ({
  menuIconPadRight: {
    marginRight: theme.spacing(2),
  },
}));

function App() {
  const classes = useStyles();
  return (
    <div className="App">
      <AppBar color='primary' position='static'>
        <Toolbar>
          <AndroidIcon className={classes.menuIconPadRight} />
          <Typography>MARVIN</Typography>
        </Toolbar>
      </AppBar>
    </div>
  );
}

export default App;
