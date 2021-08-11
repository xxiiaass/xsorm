package xsorm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestStruct struct {
	UserID    int64  `gorm:"column:user_id;PRIMARY_KEY"    json:"user_id"`
	ShortUUID string `gorm:"column:short_uuid"              json:"short_uuid"`
	Salt      string `gorm:"column:salt"                    json:"salt"`
	Password  string `gorm:"column:password"               json:"password"`
}

func (TestStruct) TableName() string {
	return "user"
}
func (u *TestStruct) AfterFind() (err error) {
	u.Password = "xxxxxx"
	return
}

func TestAfterHook(t *testing.T) {
	ts := new(TestStruct)
	NewBuild(ts).Where("user_id", ">", 0).First()
	assert.Equal(t, "xxxxxx", ts.Password)

	tss := make([]TestStruct, 0)
	NewBuild(&tss).Where("user_id", ">", 0).Limit(3).Get()
	assert.Equal(t, 3, len(tss))
	assert.Equal(t, "xxxxxx", tss[0].Password)
	assert.Equal(t, "xxxxxx", tss[1].Password)
	assert.Equal(t, "xxxxxx", tss[2].Password)
}
