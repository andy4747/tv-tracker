import AppBar from '@material-ui/core/AppBar';
import IconButton from '@material-ui/core/IconButton';
import { makeStyles } from '@material-ui/core/styles';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import TvIcon from '@material-ui/icons/Tv';
import React, { FC } from 'react';
import { User } from '../../Utils/Types';
import { AuthorizedNav } from '../Private/AuthorizedNav';
import { PublicNav } from './PublicNav';

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
							<AuthorizedNav user={user} />
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
							<PublicNav />
						</Toolbar>
					</AppBar>
				</div>
			);
		default:
			return null;
	}
};
