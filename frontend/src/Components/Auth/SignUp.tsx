import Avatar from '@material-ui/core/Avatar';
import Box from '@material-ui/core/Box';
import Button from '@material-ui/core/Button';
import Container from '@material-ui/core/Container';
import CssBaseline from '@material-ui/core/CssBaseline';
import Grid from '@material-ui/core/Grid';
import Link from '@material-ui/core/Link';
import { makeStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Typography from '@material-ui/core/Typography';
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
import React, { ChangeEvent, FC, SyntheticEvent, useState } from 'react';
import { RouterProps } from 'react-router';
import { API_PROXY } from '../../Utils/Constants';
import { Copyright } from '../Copyright';

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
		marginTop: theme.spacing(3),
	},
	submit: {
		margin: theme.spacing(3, 0, 2),
	},
}));

export const SignUp: FC<RouterProps> = ({ history }: RouterProps) => {
	const classes = useStyles();
	const [username, setUsername] = useState<string>('');
	const [email, setEmail] = useState<string>('');
	const [password, setPassword] = useState<string>('');
	const [confirm_password, setConfirmPassword] = useState<string>('');

	const signupHandler = async (e: SyntheticEvent) => {
		e.preventDefault();
		await fetch(`${API_PROXY}/api/auth/register`, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
			},
			body: JSON.stringify({
				username: username,
				email: email,
				password: password,
				confirm_password: confirm_password,
			}),
		});
		history.push('/login');
	};

	return (
		<Container component='main' maxWidth='xs'>
			<CssBaseline />
			<div className={classes.paper}>
				<Avatar className={classes.avatar}>
					<LockOutlinedIcon />
				</Avatar>
				<Typography component='h1' variant='h5'>
					Sign up
				</Typography>
				<form className={classes.form} onSubmit={signupHandler} noValidate>
					<Grid container spacing={2}>
						<Grid item xs={12}>
							<TextField
								variant='outlined'
								required
								fullWidth
								id='username'
								label='Username'
								name='username'
								autoComplete='username'
								value={username}
								onChange={(e: ChangeEvent<HTMLInputElement>) => {
									setUsername(e.target.value);
								}}
							/>
						</Grid>
						<Grid item xs={12}>
							<TextField
								variant='outlined'
								required
								fullWidth
								id='email'
								label='Email Address'
								name='email'
								autoComplete='email'
								value={email}
								onChange={(e: ChangeEvent<HTMLInputElement>) => {
									setEmail(e.target.value);
								}}
							/>
						</Grid>
						<Grid item xs={12}>
							<TextField
								variant='outlined'
								required
								fullWidth
								name='password'
								label='Password'
								type='password'
								id='password'
								value={password}
								onChange={(e: ChangeEvent<HTMLInputElement>) => {
									setPassword(e.target.value);
								}}
							/>
						</Grid>
						<Grid item xs={12}>
							<TextField
								variant='outlined'
								required
								fullWidth
								name='confirm_password'
								label='Confirm Password'
								type='password'
								id='confirm_password'
								value={confirm_password}
								onChange={(e: ChangeEvent<HTMLInputElement>) => {
									setConfirmPassword(e.target.value);
								}}
							/>
						</Grid>
					</Grid>
					<Button
						type='submit'
						fullWidth
						variant='contained'
						color='primary'
						className={classes.submit}>
						Sign Up
					</Button>
					<Grid container justify='flex-end'>
						<Grid item>
							<Link href='/login' variant='body2'>
								Already have an account? Sign in
							</Link>
						</Grid>
					</Grid>
				</form>
			</div>
			<Box mt={5}>
				<Copyright />
			</Box>
		</Container>
	);
};
