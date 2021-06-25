import React, { FC, useEffect, useState } from 'react';
import { Redirect, Route, RouteProps } from 'react-router-dom';
import { Loading } from '../Components/Error/Loading';
import { API_PROXY } from './Constants';
import { LoadingSize } from './Types';

interface Props {
	path: RouteProps['path'];
	exact: RouteProps['exact'];
	component: React.ElementType;
	location?: RouteProps['location'];
}

export const PublicRoute: FC<Props> = ({
	component: Component,
	...rest
}: Props) => {
	const [status, setStatus] = useState(400);
	const [loading, setLoading] = useState(true);

	const authHandler = async () => {
		const response = await fetch(`${API_PROXY}/api/auth/user`, {
			headers: { 'Content-Type': 'application/json' },
			credentials: 'include',
		});
		const status = await response.status;
		setStatus(status);
		setLoading(false);
	};

	useEffect(() => {
		authHandler();
		return () => {
			setLoading(true);
			setStatus(400);
		};
	}, []);

	return (
		<>
			<Route
				{...rest}
				render={(props) => {
					if (status === 200 && loading === false) {
						return (
							<Redirect
								to={{
									pathname: '/dashboard',
									state: {
										from: props.location,
									},
								}}
							/>
						);
					} else if (loading === true) {
						return (
							<>
								<Loading size={LoadingSize.xl} />
							</>
						);
					} else {
						return <Component {...props} />;
					}
				}}
			/>
		</>
	);
};
