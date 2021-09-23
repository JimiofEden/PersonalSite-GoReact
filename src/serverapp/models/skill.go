package models

type Skill struct {
	Id int `json:"id"`
	Name string `json:"name"`
	SkillTypeId int `json:"skillType"`
	Link string `json:"link"`
	Comment string `json:"comment"`
}

func NewSkill(name string, skillTypeId int, link string, comment string) Skill {
	return Skill {
		Name: name,
		SkillTypeId: skillTypeId,
		Link: link,
		Comment: comment,
	}
}