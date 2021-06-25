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
import { RouterProps, withRouter } from 'react-router';
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
		marginTop: theme.spacing(1),
	},
	submit: {
		margin: theme.spacing(3, 0, 2),
	},
}));

const Login: FC<RouterProps> = ({ history }: RouterProps) => {
	const [email, setEmail] = useState<string>('');
	const [password, setPassword] = useState<string>('');
	const classes = useStyles();

	const loginHandler = async (e: SyntheticEvent) => {
		e.preventDefault();
		await fetch(`${API_PROXY}/api/auth/login`, {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			credentials: 'include',
			body: JSON.stringify({
				email: email,
				password: password,
			}),
		});
		history.push('/dashboard');
	};

	return (
		<Container component='main' maxWidth='xs'>
			<CssBaseline />
			<div className={classes.paper}>
				<Avatar className={classes.avatar} color='primary'>
					<LockOutlinedIcon />
				</Avatar>
				<Typography component='h1' variant='h5'>
					Login
				</Typography>
				<form className={classes.form} onSubmit={loginHandler} noValidate>
					<TextField
						variant='outlined'
						margin='normal'
						required
						fullWidth
						id='email'
						label='Email Address'
						name='email'
						autoComplete='email'
						value={email}
						onChange={(e: ChangeEvent<HTMLTextAreaElement>) =>
							setEmail(e.target.value)
						}
						autoFocus
					/>
					<TextField
						variant='outlined'
						margin='normal'
						required
						fullWidth
						name='password'
						label='Password'
						type='password'
						id='password'
						value={password}
						onChange={(e: ChangeEvent<HTMLTextAreaElement>) =>
							setPassword(e.target.value)
						}
						autoComplete='current-password'
					/>
					<Button
						type='submit'
						fullWidth
						variant='contained'
						color='primary'
						className={classes.submit}>
						Login
					</Button>
					<Grid container>
						<Grid item xs>
							<Link href='#' variant='body2'>
								Forgot password?
							</Link>
						</Grid>
						<Grid item>
							<Link href='/signup' variant='body2'>
								{"Don't have an account? Sign Up"}
							</Link>
						</Grid>
					</Grid>
				</form>
			</div>
			<Box mt={8}>
				<Copyright />
			</Box>
		</Container>
	);
};

export default withRouter(Login);
