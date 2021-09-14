package models

type Skill struct {
	Name string `json:"name"`
	SkillType string `json:"skillType"`
	Link string `json:"link"`
}

func NewSkill(name string, skillType string, link string) Skill {
	return Skill {
		Name: name,
		SkillType: skillType,
		Link: link,
	}
}