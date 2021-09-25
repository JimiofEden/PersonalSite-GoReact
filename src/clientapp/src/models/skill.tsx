import { SkillType } from './skillType';

export class Skill {
	id: number = 0;
	skillName: string = "";
	skillType: SkillType = new SkillType();
	url: string = "";
	comment: string = "";
	sequence: number = 0;
}