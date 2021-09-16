import React, {useState, useEffect} from 'react';
import {Skill} from '../models/skill';
import {fetchSkills} from '../util/ApiHandler';
import {PulseLoader} from 'react-spinners'

export const Skills = (props: any) => {

	const [loading, setLoading] = useState<boolean>(true);
	const [skills, setSkills] = useState<Skill[]>([]);

	useEffect(() => {
		// TODO - Load in skills here
		const goGetSkills = () => fetchSkills().then((results: any) => results.data);

		let subscribed = true;
		setLoading(true);

		goGetSkills().then(
			(results: any) => {
				if (subscribed) {
					console.log(results.data);
					setLoading(false);
					setSkills(results.data);
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

	const [skillTypes, setSkillTypes] = useState<string[]>([]);

	useEffect(() => {
		const newSkillTypes = (
			skills.map((skill: any) => {
				return skill.skillType;
			}));
		setSkillTypes(newSkillTypes.filter((v, i, a) => {return a.indexOf(v) === i}));
	}, [skills]);

	return (
		loading
		? <PulseLoader/>
		: <div>
			<p>For a quick glance, I am familiar with the following technologies:</p>
			<div className="skills">
			{
				skillTypes.map((skillType: string, i: number) => {
					return (
						<div className="skill-category" key={skillType}>
						<h4>{skillType}</h4>
						{
							skills.filter((skill: Skill) => { return skill.skillType === skillType})
							.map((skill: Skill, i: number) => {
								return (
									<p key={skill.name}>
										{skill.name}
									</p>
								)
							})
						}
						</div>
					)
				})
			}
			</div>
		</div>
	)
}