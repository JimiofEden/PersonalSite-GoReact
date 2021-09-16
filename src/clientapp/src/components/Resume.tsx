export const Resume = (props: any) => {
	return <div>
		<p>There's not much here at the moment, since I prefer websites that load very quickly with minimal content, but you can check out my <a href={props.resumeLink} target="_blank" rel="noreferrer">resume</a></p>
		<p>Also feel free to check out my <a href={props.githubLink} target="_blank" rel="noreferrer">github</a></p>
	</div>
}