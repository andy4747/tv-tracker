import React, { FC } from 'react';
import { BrowserRouter as Router, Route } from 'react-router-dom';
import './App.css';
import { Home } from './Components/Auth/Home';

export const App: FC = () => {
	return (
		<>
			<Router>
				<Route path='/' exact component={Home} />
				<Route path='/signup' exact component={Home} />
				<Route path='/login' exact component={Home} />
				{/* <Route path='/' component={Home} /> */}
			</Router>
		</>
	);
};
