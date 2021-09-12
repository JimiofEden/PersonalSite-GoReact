import React from 'react';
import Header from './Header';
import AboutMe from './AboutMe';
import Resume from './Resume';
import Skills from './Skills';
import Contact from './Contact';
import './Site.css';

export function Site() {
	return <div className="site-container">
		<Header/>
		{/*TODO - Links should come from backend store*/}
		<AboutMe
			twitterLink={"https://twitter.com/JimiofEden"}
		/>
		<Resume
			resumeLink={"./AH-Resume_0721-linkedin.pdf"}
			githubLink={"https://github.com/jimiofeden"}
		/>
		<Skills/>
		<Contact
			contactEmail={"mailto:jimiofeden@gmail.com"}
		/>
	</div>;
}