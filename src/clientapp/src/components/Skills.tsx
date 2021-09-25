import React, { useState, useEffect } from 'react';
import { Skill } from '../models/skill';
import { SkillType } from '../models/skillType';
import { fetchSkills } from '../util/ApiHandler';
import { PulseLoader } from 'react-spinners'
import "./Skills.css";

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
					//console.log(results.data);
					setLoading(false);
					setSkills(results.data.skills);
				}
			},
			(err: any) => {
				if (subscribed) {
					setLoading(false);
					//console.log(err);
				}
			}
		);
		return () => {
			subscribed = false;
		};
	}, []);

	const [skillTypes, setSkillTypes] = useState<SkillType[]>([]);

	useEffect(() => {
		const newSkillTypes = (
			skills.map((skill: Skill) => {
				return skill.skillType;
			}));
		var uniqueSkillTypes = [];
		for(var i = 0; i < skills.length; i++){
			if (uniqueSkillTypes.filter((x) => {return x.skillTypeName === skills[i].skillType.skillTypeName}).length == 0) {
				uniqueSkillTypes.push(skills[i].skillType);
			}

		}
		setSkillTypes(uniqueSkillTypes);
	}, [skills]);

	return (
		loading
		? <PulseLoader/>
		: <div>
			<p>For a quick glance, I am familiar with the following technologies:</p>
			<div className="skills">
			{
				skillTypes
				.sort((a, b) => {return a.sequence - b.sequence})
				.map((skillType: SkillType, i: number) => {
					return (
						<div className="skill-category" key={i}>
						<h4>{skillType.skillTypeName}</h4>
						{
							skills
							.sort((a, b) => {return a.sequence - b.sequence})
							.filter((skill: Skill) => { return skill.skillType.skillTypeName === skillType.skillTypeName})
							.map((skill: Skill, j: number) => {
								return (
									<p key={j}>
										{skill.skillName}{skill.comment.trim() !== ""
											? (skill.url.trim() !== ""
												? <span> - <a href={skill.url} target="_blank" rel="noreferrer">{skill.comment}</a></span>
												: <span> - {skill.comment}</span>
												)
											: ("")
										}
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