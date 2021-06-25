import { ButtonGroup } from '@material-ui/core';
import AppBar from '@material-ui/core/AppBar';
import Button from '@material-ui/core/Button';
import IconButton from '@material-ui/core/IconButton';
import { makeStyles } from '@material-ui/core/styles';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import TvIcon from '@material-ui/icons/Tv';
import React, { FC, SyntheticEvent } from 'react';
import { useHistory } from 'react-router';
import { API_PROXY } from '../Utils/Constants';
import { User } from '../Utils/Types';

const useStyles = makeStyles((theme) => ({
	root: {
		flexGrow: 1,
	},
	menuButton: {
		marginRight: theme.spacing(2),
	},
	title: {
		flexGrow: 1,
	},
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

interface Props {
	auth: boolean;
	user: User;
}

export const Navbar: FC<Props> = ({ auth, user }: Props) => {
	const classes = useStyles();
	const history = useHistory();
	const logoutHandler = async (e: SyntheticEvent) => {
		e.preventDefault();
		await fetch(`${API_PROXY}/api/auth/logout`, {
			headers: { 'Content-Type': 'application/json' },
			credentials: 'include',
		});
		history.push('/');
	};
	switch (auth) {
		case true:
			return (
				<div className={classes.root}>
					<AppBar position='static' color='primary'>
						<Toolbar>
							<IconButton
								edge='start'
								className={classes.menuButton}
								color='inherit'
								aria-label='menu'>
								<TvIcon />
							</IconButton>
							<Typography variant='h6' className={classes.title}>
								My Shows Tracker
							</Typography>
							<ButtonGroup className={classes.btnGroup}>
								<Button color='inherit' className={classes.leftBtn}>
									{user.username}
								</Button>
								<Button
									color='inherit'
									onClick={logoutHandler}
									className={classes.rightBtn}>
									Logout
								</Button>
							</ButtonGroup>
						</Toolbar>
					</AppBar>
				</div>
			);
		case false:
			return (
				<div className={classes.root}>
					<AppBar position='static' color='primary'>
						<Toolbar>
							<IconButton
								edge='start'
								className={classes.menuButton}
								color='inherit'
								aria-label='menu'>
								<TvIcon />
							</IconButton>
							<Typography variant='h6' className={classes.title}>
								My Shows Tracker
							</Typography>
							<ButtonGroup className={classes.btnGroup}>
								<Button
									color='inherit'
									className={classes.leftBtn}
									href='/signup'>
									Signup
								</Button>
								<Button
									color='inherit'
									className={classes.rightBtn}
									href='/login'>
									Login
								</Button>
							</ButtonGroup>
						</Toolbar>
					</AppBar>
				</div>
			);
		default:
			return null;
	}
};
