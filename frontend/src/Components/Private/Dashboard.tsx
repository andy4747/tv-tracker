import React, { FC } from 'react';
import { User } from '../../Utils/Types';
import { Navbar } from '../Navbar';

interface Props {
	user: User;
}

export const Dashboard: FC<Props> = ({ user }: Props) => {
	return (
		<>
			<Navbar auth={true} user={user} />
			<h1>Dashboard</h1>
			<p>{user.email}</p>
			<p>{user.username}</p>
		</>
	);
};
