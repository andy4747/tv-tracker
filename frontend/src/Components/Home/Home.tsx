import React, { FC } from 'react';
import { RouterProps } from 'react-router';
import { Navbar } from './Navbar';

export const Home: FC<RouterProps> = () => {
	return (
		<>
			<Navbar auth={false} user={{}} />
		</>
	);
};
