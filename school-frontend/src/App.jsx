import { BrowserRouter } from "react-router-dom";
import AppRouter from "./router/AppRouter";
import React from "react";
import Login from "./pages/Login";


const App = () => (
  <BrowserRouter>
    <AppRouter />
    {/* <Login/> */}
  </BrowserRouter>
);

export default App;
