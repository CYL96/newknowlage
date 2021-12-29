package gui

import (
	"os"
	"os/exec"
	"time"

	"fyne.io/fyne/v2/data/binding"
)

func Attack() bool {
	User.Rebuild()
	Monster.Rebuild()
	Clear()
	// fmt.Printf("[=========================== %d ===========================]\n", User.Level)
	// User.ShowStatus()
	// Monster.ShowStatus()
	time.Sleep(10 * time.Millisecond)
	User_round := 0
	monster_round := 0
	for {
		// Clear()
		// fmt.Printf("[=========================== %d ===========================]\n", User.Level)
		// User.ShowStatus()
		// Monster.ShowStatus()
		// fmt.Println("User_round:", User_round, "monster_round:", monster_round)
		t := GetAttackTarget()
		switch t {
		case 1:
			User.Attack(Monster)
			Monster.FlushStr()
			User_round++
		case 2:
			Monster.Attack(User)
			User.FlushStr()
			monster_round++

		}

		if !User.Alive() {
			return false
		}
		if !Monster.Alive() {
			Monster.BuildNew(5, 25)
			// GetWantInput()
			return true
		}
		time.Sleep(40 * time.Millisecond)

	}
}

func Clear() {
	cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func InitTarget() {
	// User init
	User = &TargetInfoExt{
		Level: 1,
		Name:  "User",
		// AttInfo: new(TargetInfoString),
		NorAtt: AttributesBasicExt{
			HP:          BuildRandAtt(500, 80, 100),
			LeftPer:     1,
			Def:         BuildRandAtt(5, 0, 100),
			Attack:      BuildRandAtt(40, 50, 100),
			AttackSpeed: BuildRandAtt(100, 50, 80),
		},
		Tmp: AttributesBasicExt{},
	}
	User.Tmp.LeftPerBd = binding.BindFloat(&User.Tmp.LeftPer)
	User.Tmp.HPStrBd = binding.BindString(&User.Tmp.HPStr)
	User.Tmp.DefStrBd = binding.BindString(&User.Tmp.DefStr)
	User.Tmp.AttackStrBd = binding.BindString(&User.Tmp.AttackStr)
	User.Tmp.AttackSpeedStrBd = binding.BindString(&User.Tmp.AttackSpeedStr)
	User.LevelBd = binding.BindInt(&User.Level)
	User.StrBd = binding.BindString(&User.Str)
	for i := range User.statusStr {
		User.status = append(User.status, binding.BindString(&User.statusStr[i]))
	}

	// monster init
	Monster = &TargetInfoExt{
		Name:  "monster",
		Level: 1,
		// AttInfo: new(TargetInfoString),
		NorAtt: AttributesBasicExt{
			HP:          BuildRandAtt(600, 80, 100),
			LeftPer:     1,
			Def:         BuildRandAtt(3, 0, 100),
			Attack:      BuildRandAtt(20, 50, 100),
			AttackSpeed: BuildRandAtt(100, 50, 80),
		},
	}
	Monster.Tmp.LeftPerBd = binding.BindFloat(&Monster.Tmp.LeftPer)
	Monster.Tmp.HPStrBd = binding.BindString(&Monster.Tmp.HPStr)
	Monster.Tmp.DefStrBd = binding.BindString(&Monster.Tmp.DefStr)
	Monster.Tmp.AttackStrBd = binding.BindString(&Monster.Tmp.AttackStr)
	Monster.Tmp.AttackSpeedStrBd = binding.BindString(&Monster.Tmp.AttackSpeedStr)
	Monster.LevelBd = binding.BindInt(&Monster.Level)
	Monster.StrBd = binding.BindString(&Monster.Str)
	for i := range User.statusStr {
		User.status = append(User.status, binding.BindString(&User.statusStr[i]))
	}

}

var User *TargetInfoExt
var Monster *TargetInfoExt

func GetAttackTarget() int {
	if User.Tmp.RunAttackSpeed > Monster.Tmp.RunAttackSpeed {
		User.Tmp.RunAttackSpeed = Decimal(User.Tmp.RunAttackSpeed-Monster.Tmp.RunAttackSpeed, 1)
		if Monster.Tmp.RunAttackSpeed < Monster.Tmp.AttackSpeed {
			Monster.Tmp.RunAttackSpeed += Monster.Tmp.AttackSpeed
		}
		return 1
	} else {
		Monster.Tmp.RunAttackSpeed = Decimal(Monster.Tmp.RunAttackSpeed-User.Tmp.RunAttackSpeed, 1)
		if User.Tmp.RunAttackSpeed < User.Tmp.AttackSpeed {
			User.Tmp.RunAttackSpeed += User.Tmp.AttackSpeed
		}
		return 2
	}
}

func init() {
	InitTarget()
}
