package run

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func Attack() bool {
	User.Rebuild()
	Monster.Rebuild()
	Clear()
	// fmt.Printf("[=========================== %d ===========================]\n", User.Level)
	User.ShowStatus()
	Monster.ShowStatus()
	time.Sleep(333 * time.Millisecond)
	User_round := 0
	monster_round := 0
	for {
		// Clear()
		// fmt.Printf("[=========================== %d ===========================]\n", User.Level)
		User.ShowStatus()
		Monster.ShowStatus()
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
			fmt.Println("you die !")
			return false
		}
		if !Monster.Alive() {
			fmt.Println("you won,next monster!")
			Monster.BuildNew(5, 25)
			// GetWantInput()
			return true
		}
		time.Sleep(100 * time.Millisecond)

	}
}

func Clear() {
	cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func InitTarget() {
	User = &TargetInfoExt{
		Level:   1,
		Name:    "User",
		AttInfo: new(TargetInfoString),
		NorAtt: AttributesBasicExt{
			HP:          BuildRandAtt(500, 80, 100),
			Def:         BuildRandAtt(5, 0, 100),
			Attack:      BuildRandAtt(40, 50, 100),
			AttackSpeed: BuildRandAtt(100, 50, 80),
		},
		Tmp: AttributesBasicExt{},
	}
	Monster = &TargetInfoExt{
		Name:    "monster",
		Level:   1,
		AttInfo: new(TargetInfoString),
		NorAtt: AttributesBasicExt{
			HP:          BuildRandAtt(600, 80, 100),
			Def:         BuildRandAtt(3, 0, 100),
			Attack:      BuildRandAtt(20, 50, 100),
			AttackSpeed: BuildRandAtt(100, 50, 80),
		},
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
