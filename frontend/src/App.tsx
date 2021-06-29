import React, { FC } from 'react';
import { BrowserRouter as Router, Switch } from 'react-router-dom';
import './App.css';
import Login from './Components/Auth/Login';
import { SignUp } from './Components/Auth/SignUp';
import { Home } from './Components/Home/Home';
import { Dashboard } from './Components/Private/Dashboard';
import { PrivateRoute } from './Utils/PrivateRoute';
import { PublicRoute } from './Utils/PublicRoute';

export const App: FC = () => {
	return (
		<>
			<Router>
				<Switch>
					<PublicRoute path='/' exact component={Home} />
					<PublicRoute path='/signup' exact component={SignUp} />
					<PublicRoute path='/login' exact component={Login} />
					<PrivateRoute path='/dashboard' exact component={Dashboard} />
				</Switch>
			</Router>
		</>
	);
};
