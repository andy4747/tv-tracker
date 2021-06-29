import Button from '@material-ui/core/Button';
import Menu from '@material-ui/core/Menu';
import MenuItem from '@material-ui/core/MenuItem';
import { makeStyles } from '@material-ui/core/styles';
import AddCircleIcon from '@material-ui/icons/AddCircle';
import React, { FC } from 'react';
import { MovieModal } from './MovieModal';
import { TvModal } from './TvModal';

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

export const CreateMenu: FC = () => {
	const classes = useStyles();
	const [anchorEl, setAnchorEl] =
		React.useState<(EventTarget & HTMLElement) | null>(null);
	const handleClose = () => {
		setAnchorEl(null);
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
				<AddCircleIcon />
			</Button>
			<Menu
				id='simple-menu'
				anchorEl={anchorEl}
				keepMounted
				open={Boolean(anchorEl)}
				onClose={handleClose}>
				<MenuItem onClick={handleClose}>
					<TvModal />
				</MenuItem>
				<MenuItem onClick={handleClose}>
					<MovieModal />
				</MenuItem>
			</Menu>
		</div>
	);
};
