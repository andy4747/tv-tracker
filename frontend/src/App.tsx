import React, { FC } from 'react';
import { BrowserRouter as Router, Switch } from 'react-router-dom';
import './App.css';
import { Home } from './Components/Auth/Home';
import Login from './Components/Auth/Login';
import { SignUp } from './Components/Auth/SignUp';
import { Dashboard } from './Components/Private/Dashboard';
import { PrivateRoute } from './Utils/PrivateRouter';
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
