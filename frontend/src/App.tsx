import React, { FC } from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import './App.css';
import { Home } from './Components/Auth/Home';
import Login from './Components/Auth/Login';
import { SignUp } from './Components/Auth/SignUp';
import { Dashboard } from './Components/Private/Dashboard';
import { PrivateRoute } from './Utils/PrivateRouter';

export const App: FC = () => {
	return (
		<>
			<Router>
				<Switch>
					<Route path='/' exact component={Home} />
					<Route path='/signup' exact component={SignUp} />
					<Route path='/login' exact component={Login} />
					<PrivateRoute path='/dashboard' exact component={Dashboard} />
				</Switch>
			</Router>
		</>
	);
};
