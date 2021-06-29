import Button from '@material-ui/core/Button';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogTitle from '@material-ui/core/DialogTitle';
import FormControl from '@material-ui/core/FormControl';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select';
import { makeStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import MovieIcon from '@material-ui/icons/Movie';
import React, { FC } from 'react';

const useStyles = makeStyles((theme) => ({
	paper: {
		marginTop: theme.spacing(8),
		display: 'flex',
		flexDirection: 'column',
		alignItems: 'center',
	},
	avatar: {
		margin: theme.spacing(1),
		backgroundColor: theme.palette.primary.main,
	},
	form: {
		width: '100%', // Fix IE 11 issue.
		marginTop: theme.spacing(1),
	},
	submit: {
		margin: theme.spacing(3, 0, 2),
	},
	formControl: {
		marginTop: theme.spacing(1),
		minWidth: '100%',
	},
	statusLabel: {
		marginLeft: theme.spacing(1),
	},
}));

export const MovieModal: FC = () => {
	const classes = useStyles();
	const [open, setOpen] = React.useState(false);
	const [watchingStatus, setWatchingStatus] = React.useState(3);
	const [openSelect, setOpenSelect] = React.useState(false);

	const handleSelectClose = () => {
		setOpenSelect(false);
	};

	const handleSelectOpen = () => {
		setOpenSelect(true);
	};

	// const handleOpen = () => {
	// 	setOpen(true);
	// };

	const handleClickOpen = () => {
		setOpen(true);
	};

	const handleClose = () => {
		setOpen(false);
	};

	return (
		<div>
			<Button onClick={handleClickOpen}>
				<MovieIcon />
				Movies
			</Button>
			<Dialog
				open={open}
				onClose={handleClose}
				aria-labelledby='form-dialog-title'>
				<DialogTitle id='form-dialog-title'>Subscribe</DialogTitle>
				<DialogContent>
					<DialogContentText>
						To subscribe to this website, please enter your email address here.
						We will send updates occasionally.
					</DialogContentText>
					<form method='post' className={classes.form}>
						{/* Movie Name */}
						<TextField
							variant='outlined'
							margin='normal'
							required
							fullWidth
							id='movieName'
							label='Movie'
							name='Movie'
							autoFocus
						/>
						{/* Movie Status */}
						<FormControl className={classes.formControl}>
							<InputLabel
								id='demo-controlled-open-select-label'
								className={classes.statusLabel}>
								Status
							</InputLabel>
							<Select
								labelId='demo-controlled-open-select-label'
								id='demo-controlled-open-select'
								variant='outlined'
								fullWidth
								open={openSelect}
								onClose={handleSelectClose}
								onOpen={handleSelectOpen}
								value={watchingStatus}
								onChange={(e) => {
									setWatchingStatus(e.target.value as number);
								}}>
								<MenuItem value={1}>Planned</MenuItem>
								<MenuItem value={2}>Completed</MenuItem>
								<MenuItem value={3}>Watching</MenuItem>
							</Select>
						</FormControl>
						{watchingStatus === 3 ? (
							<TextField
								variant='outlined'
								margin='normal'
								required
								fullWidth
								id='currentLength'
								label='Current Length'
								name='minutes'
								autoComplete='movie'
							/>
						) : (
							<></>
						)}

						{/* Movie Language */}
						<TextField
							variant='outlined'
							margin='normal'
							required
							fullWidth
							id='movieLang'
							label='Language'
							name='language'
						/>
					</form>
				</DialogContent>
				<DialogActions>
					<Button onClick={handleClose} color='primary'>
						Cancel
					</Button>
					<Button onClick={handleClose} type='submit' color='primary'>
						Create
					</Button>
				</DialogActions>
			</Dialog>
		</div>
	);
};
