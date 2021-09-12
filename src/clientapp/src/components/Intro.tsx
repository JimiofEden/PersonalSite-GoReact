import React from 'react';
import './Intro.css';

function Header(props: any) {
	return <h3 className="header-text">Hi, I'm Adam! Welcome to my website!</h3>;
}

function AboutMe(props: any) {
	return <div>
		<p>I'm a 10+ years professional Software Engineer who likes games, music, and <a href={props.twitterLink} target="_blank" rel="noreferrer">cooking</a> (especially anything spicy)</p>
	</div>;
}

function Resume(props: any) {
	return <div>
		<p>There's not much here at the moment, since I prefer websites that load very quickly with minimal content, but you can check out my <a href={props.resumeLink} target="_blank" rel="noreferrer">resume</a></p>
		<p>Also feel free to check out my <a href={props.githubLink} target="_blank" rel="noreferrer">github</a></p>
	</div>
}

function Skills(props: any) {
	return <div>
		<p>For a quick glance, I am familiar with the following technologies:</p>
		<div className="skills">
			<div className="skill-category">
				<h4>Backend</h4>
				<p>C#</p>
				<p>.NET</p>
				<p>Go (This site served by Go!)</p>
				<p>Node</p>
				<p>Python</p>
				<p>Ruby</p>
			</div>
			<div className="skill-category">
				<h4>Frontend</h4>
				<p>React (This frontend is being rendered with React!)</p>
				<p>Typescript (This frontend was built with Typescript!)</p>
				<p>Angular</p>
				<p>Knockout</p>
			</div>
			<div className="skill-category">
				<h4>Database</h4>
				<p>SQL Server</p>
				<p>MySQL</p>
				<p>MongoDB</p>
				<p>Oracle</p>
			</div>
			<div className="skill-category">
				<h4>Serverside</h4>
				<p>Azure</p>
				<p>IIS</p>
				<p>Apache</p>
			</div>
			<div className="skill-category">
				<h4>Misc. Skills</h4>
				<p>Continuous Integration (Azure Devops, TeamCity, Octopus Deploy)</p>
				<p>Version Control (Git, SVN)</p>
				<p>Unit Testing (Jest, NUnit)</p>
				<p>Project Managing (Agile, Scrum)</p>
				<p>Translating Business Requests into Requirements</p>
			</div>
		</div>
	</div>
}

function Contact(props: any) {
	return <div>
		<p>Feel free to message me if you want to talk about development or if you want a cat pic!</p>
	</div>
}

export function Intro() {
	return <div className="site-container">
		<Header/>
		<AboutMe
			twitterLink={"https://twitter.com/JimiofEden"}
		/>
		<Resume
			resumeLink={"./AH-Resume_0721-linkedin.pdf"}
			githubLink={"https://github.com/jimiofeden"}
		/>
		<Skills/>
		<Contact/>
	</div>;
}