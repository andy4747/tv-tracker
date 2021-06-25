import React, { FC } from 'react';
import { Navbar } from '../Navbar';

export const Home: FC = () => {
	return (
		<>
			<Navbar auth={false} user={{}} />
		</>
	);
};
