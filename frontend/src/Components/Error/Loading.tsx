import { Grid } from '@material-ui/core';
import CircularProgress from '@material-ui/core/CircularProgress';
import { makeStyles } from '@material-ui/core/styles';
import React, { FC } from 'react';
import { LoadingSize } from '../../Utils/Types';

const useStyles = makeStyles((theme) => ({
	root: {
		display: 'flex',
		'& > * + *': {
			marginLeft: theme.spacing(2),
		},
	},
}));

interface Props {
	size: LoadingSize;
}

export const Loading: FC<Props> = ({ size }: Props) => {
	const classes = useStyles();

	return (
		<Grid
			container
			spacing={0}
			direction='column'
			alignItems='center'
			justify='center'
			style={{ minHeight: '100vh' }}>
			<Grid item xs={3}>
				<div className={classes.root}>
					<CircularProgress color='primary' size={size} />
				</div>
			</Grid>
		</Grid>
	);
};
