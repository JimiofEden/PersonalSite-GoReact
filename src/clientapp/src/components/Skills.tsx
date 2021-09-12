export default function Skills(props: any) {
	return <div>
		<p>For a quick glance, I am familiar with the following technologies:</p>
		<div className="skills">
			<div className="skill-category">
				{/*TODO - These should come from the backend store*/}
				<h4>Backend</h4>
				<p>C#</p>
				<p>.NET</p>
				<p>Go (This site served by Go!)</p>
				<p>Node</p>
				{/*TODO - Link to github PyMeth*/}
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
				<p>Oracle</p>
			</div>
			<div className="skill-category">
				<h4>Serverside</h4>
				<p>Docker</p>
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