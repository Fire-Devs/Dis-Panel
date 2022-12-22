import "./navbar.css";

import AppBar from "@mui/material/AppBar";
import Box from "@mui/material/Box";
import Toolbar from "@mui/material/Toolbar";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import { createMuiTheme } from "@mui/material/styles";

const theme = createMuiTheme({
  palette: {
    primary: {
      main: "#23272a",
    },
    secondary: {
      main: "#ffffff",
    },
  },
});

export default function Navbar() {
  return (
    <Box sx={{ flexGrow: 1 }}>
      <AppBar component="nav" theme={theme} >
        <Toolbar>
          <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
            DisPanel
          </Typography>
          <Button variant="outlined" color="inherit">
            Login
          </Button>
        </Toolbar>
      </AppBar>
    </Box>
  );
}
