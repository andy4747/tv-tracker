import Link from '@material-ui/core/Link';
import Typography from '@material-ui/core/Typography';
import React, { FC } from 'react';

export const Copyright: FC = () => {
	return (
		<Typography variant='body2' color='textSecondary' align='center'>
			{'Copyright © '}
			<Link color='inherit' href='/'>
				My Shows Tracker
			</Link>{' '}
			{new Date().getFullYear()}
			{'.'}
		</Typography>
	);
};
