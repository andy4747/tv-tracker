import Button from '@material-ui/core/Button';
import Menu from '@material-ui/core/Menu';
import MenuItem from '@material-ui/core/MenuItem';
import { makeStyles } from '@material-ui/core/styles';
import AccountCircleIcon from '@material-ui/icons/AccountCircle';
import ExitToAppIcon from '@material-ui/icons/ExitToApp';
import PersonIcon from '@material-ui/icons/Person';
import React, { FC, SyntheticEvent } from 'react';
import { useHistory } from 'react-router';
import { API_PROXY } from '../../Utils/Constants';
import { User } from '../../Utils/Types';

const useStyles = makeStyles((theme) => ({
	menuButton: {
		marginRight: theme.spacing(2),
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
	user: User;
}

export const AccountsMenu: FC<Props> = ({ user }: Props) => {
	const classes = useStyles();
	const [anchorEl, setAnchorEl] =
		React.useState<(EventTarget & HTMLElement) | null>(null);
	const handleClose = () => {
		setAnchorEl(null);
	};
	const history = useHistory();

	const logoutHandler = async (e: SyntheticEvent) => {
		e.preventDefault();
		await fetch(`${API_PROXY}/api/auth/logout`, {
			headers: { 'Content-Type': 'application/json' },
			credentials: 'include',
		});
		history.push('/');
		handleClose();
	};

	return (
		<div>
			<Button
				aria-controls='simple-menu'
				aria-haspopup='true'
				className={classes.menuButton}
				color='inherit'
				onClick={(e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
					setAnchorEl(e.currentTarget);
				}}>
				<AccountCircleIcon />
			</Button>
			<Menu
				id='simple-menu'
				anchorEl={anchorEl}
				keepMounted
				open={Boolean(anchorEl)}
				onClose={handleClose}>
				<MenuItem onClick={handleClose}>
					<PersonIcon /> {user.username}
				</MenuItem>
				<MenuItem onClick={logoutHandler}>
					<ExitToAppIcon /> Logout
				</MenuItem>
			</Menu>
		</div>
	);
};
