package models

type Skill struct {
	Name string `json:"name"`
	SkillType string `json:"skillType"`
	Link string `json:"link"`
	Comment string `json:"comment"`
}

func NewSkill(name string, skillType string, link string, comment string) Skill {
	return Skill {
		Name: name,
		SkillType: skillType,
		Link: link,
		Comment: comment,
	}
}