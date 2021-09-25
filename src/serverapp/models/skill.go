package models

type Skill struct {
	Id int `json:"id"`
	SkillName string `json:"skillName"`
	SkillTypeId int `json:"skillType"`
	Url string `json:"url"`
	Comment string `json:"comment"`
	Sequence int `json:"sequence"`
}

func NewSkill(skillName string, skillTypeId int, url string, comment string, sequence int) Skill {
	return Skill {
		SkillName: skillName,
		SkillTypeId: skillTypeId,
		Url: url,
		Comment: comment,
		Sequence: sequence,
	}
}