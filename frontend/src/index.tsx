import { ThemeProvider } from '@material-ui/core/styles';
import React from 'react';
import ReactDOM from 'react-dom';
import { App } from './App';
import { theme } from './Theme';

ReactDOM.render(
	// <React.StrictMode>
	<ThemeProvider theme={theme}>
		<App />
	</ThemeProvider>,
	// </React.StrictMode>,
	document.getElementById('root')
);
