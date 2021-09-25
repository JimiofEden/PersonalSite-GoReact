package models

type SkillType struct {
	Id int `json:"id"`
	SkillTypeId int `json:"skillTypeId"`
	SkillTypeName string `json:"skillTypeName"`
	Sequence int `json:"sequence"`
}

func NewSkillType(skillTypeId int, skillTypeName string, sequence int) SkillType {
	return SkillType {
		SkillTypeId: skillTypeId,
		SkillTypeName: skillTypeName,
		Sequence: sequence,
	}
}