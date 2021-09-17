import React, { useState, useEffect } from 'react';
import { Header } from './Header';
import { AboutMe } from './AboutMe';
import { Resume } from './Resume';
import { Skills } from './Skills';
import { Contact } from './Contact';
import { StoredLink } from '../models/storedLink';
import { fetchLinks } from '../util/ApiHandler';
import './Site.css';

export const Site = () => {

	const [loading, setLoading] = useState<boolean>(true);
	const [links, setLinks] = useState<StoredLink[]>([]);

	useEffect(() => {
		// TODO - Load in skills here
		const goGetLinks = () => fetchLinks().then((results: any) => results.data);

		let subscribed = true;
		setLoading(true);

		goGetLinks().then(
			(results: any) => {
				if (subscribed) {
					console.log(results.data);
					setLoading(false);
					setLinks(results.data);
				}
			},
			(err: any) => {
				if (subscribed) {
					setLoading(false);
					console.log(err);
				}
			}
		);
		return () => {
			subscribed = false;
		};
	}, []);

	return (
		<div className="site-container">
		<Header/>
		<AboutMe
			// TODO - link name should be in an appsetting somewhere
			twitterLink={links.filter((x) => {return x.name === "twitter"})[0].url}
		/>
		<Resume
			resumeLink={links.filter((x) => {return x.name === "resume"})[0].url}
			githubLink={links.filter((x) => {return x.name === "github"})[0].url}
		/>
		<Skills/>
		<Contact
			contactEmail={links.filter((x) => {return x.name === "email"})[0].url}
		/>
	</div>
	);
}