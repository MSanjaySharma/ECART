import React from "react";
import { BrowserRouter, Route, Switch } from "react-router-dom";

import Header from "./components/layout/Header";
import StickyFooter from "./components/layout/Footer";
import { PrivateRoute } from "./utils/components/PrivateRoute";
import { Home, Signin, Signup } from "./pages";

function App() {
  return (
    <div className="App">
      <BrowserRouter>
        <Header />
        <div style={{ height: "64px" }} />

        <Switch>
          <PrivateRoute path="/" component={Home} exact={true} />
          <Route path="/signin" component={Signin} exact={true} />
          <Route path="/signup" component={Signup} exact={true} />
        </Switch>
        <StickyFooter />
      </BrowserRouter>
    </div>
  );
}

export default App;
