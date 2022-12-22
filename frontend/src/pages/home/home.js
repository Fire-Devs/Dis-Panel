import "./home.css";
import Button from "@mui/material/Button";
// import { createMuiTheme } from "@mui/material/styles";
// const theme = createMuiTheme({
//   palette: {
//     primary: {
//       main: "#23272a",
//     },
//     secondary: {
//       main: "#ffffff",
//     },
//   },
// });

export default function Home() {
  return (
    <section>
      <div className="home">
        <h1 className="home_heading">DisPanel</h1>
        <p className="home_para">
          Proably the best Panel you can get to host your Discord Bot!
        </p>
        <div className="flex_buttons">
          <Button
            size="large"
            variant="contained"
            color="primary"
            href="/admin"
          >
            Login
          </Button>
          <Button
            size="large"
            variant="contained"
            color="primary"
            href="/login"
          >
            Support
          </Button>
        </div>
      </div>
    </section>
  );
}
