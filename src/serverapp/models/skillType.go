package models

type SkillType struct {
	Id int `json:"id"`
	SkillTypeId int `json:"skillTypeName"`
	SkillTypeName string `json:"skillTypeName"`
}

func NewSkillType(skillTypeId int, skillTypeName string) SkillType {
	return SkillType {
		SkillTypeId: skillTypeId,
		SkillTypeName: skillTypeName,
	}
}