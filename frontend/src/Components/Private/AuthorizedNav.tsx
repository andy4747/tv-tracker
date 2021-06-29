import { ButtonGroup } from '@material-ui/core';
import { makeStyles } from '@material-ui/core/styles';
import React, { FC } from 'react';
import { User } from '../../Utils/Types';
import { AccountsMenu } from './AccountsMenu';
import { CreateMenu } from './CreateMenu';

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

interface Props {
	user: User;
}

export const AuthorizedNav: FC<Props> = ({ user }: Props) => {
	const classes = useStyles();

	return (
		<ButtonGroup className={classes.btnGroup}>
			<CreateMenu />
			<AccountsMenu user={user} />
		</ButtonGroup>
	);
};
