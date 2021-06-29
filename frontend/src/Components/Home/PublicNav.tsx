import { ButtonGroup } from '@material-ui/core';
import Button from '@material-ui/core/Button';
import { makeStyles } from '@material-ui/core/styles';
import React, { FC } from 'react';

const useStyles = makeStyles(() => ({
	btnGroup: {
		marginRight: 50,
	},
	leftBtn: {
		marginRight: 20,
		border: 'none',
	},
	rightBtn: {
		border: 'none',
	},
}));

export const PublicNav: FC = () => {
	const classes = useStyles();

	return (
		<ButtonGroup className={classes.btnGroup}>
			<Button color='inherit' className={classes.leftBtn} href='/signup'>
				Signup
			</Button>
			<Button color='inherit' className={classes.rightBtn} href='/login'>
				Login
			</Button>
		</ButtonGroup>
	);
};
