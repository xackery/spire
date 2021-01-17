package models

import (
	"github.com/volatiletech/null/v8"
)

type Spawn2 struct {
	ID                     int         `json:"id" gorm:"Column:id"`
	SpawngroupID           int         `json:"spawngroup_id" gorm:"Column:spawngroupID"`
	Zone                   null.String `json:"zone" gorm:"Column:zone"`
	Version                int16       `json:"version" gorm:"Column:version"`
	X                      float32     `json:"x" gorm:"Column:x"`
	Y                      float32     `json:"y" gorm:"Column:y"`
	Z                      float32     `json:"z" gorm:"Column:z"`
	Heading                float32     `json:"heading" gorm:"Column:heading"`
	Respawntime            int         `json:"respawntime" gorm:"Column:respawntime"`
	Variance               int         `json:"variance" gorm:"Column:variance"`
	Pathgrid               int         `json:"pathgrid" gorm:"Column:pathgrid"`
	Condition              uint32      `json:"_condition" gorm:"Column:_condition"`
	CondValue              int32       `json:"cond_value" gorm:"Column:cond_value"`
	Enabled                uint8       `json:"enabled" gorm:"Column:enabled"`
	Animation              uint8       `json:"animation" gorm:"Column:animation"`
	MinExpansion           uint8       `json:"min_expansion" gorm:"Column:min_expansion"`
	MaxExpansion           uint8       `json:"max_expansion" gorm:"Column:max_expansion"`
	ContentFlags           null.String `json:"content_flags" gorm:"Column:content_flags"`
	ContentFlagsDisabled   null.String `json:"content_flags_disabled" gorm:"Column:content_flags_disabled"`
}

func (Spawn2) TableName() string {
    return "spawn2"
}

func (Spawn2) Relationships() []string {
    return []string{}
}

func (Spawn2) Connection() string {
    return "eqemu_content"
}
