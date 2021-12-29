package run

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type (
	AttributesBasicExt struct {
		HP             float64
		HPStr          string
		Def            float64
		DefStr         string
		Attack         float64
		AttackStr      string
		AttackSpeed    float64
		AttackSpeedStr string
		RunAttackSpeed float64
	}
	TargetInfoExt struct {
		Level   int
		Name    string
		AttInfo *TargetInfoString
		NorAtt  AttributesBasicExt
		Tmp     AttributesBasicExt
	}
)
type TargetInfoString struct {
	AttInfo bytes.Buffer
	binding.DataItem
}

// 	Get() (string, error)
//	Set(string) error
func (this *TargetInfoString) Get() (string, error) {
	return this.AttInfo.String(), nil
}
func (this *TargetInfoString) Set(str string) error {
	this.AttInfo.Reset()
	this.AttInfo.WriteString(str)
	return nil
}
func (this *TargetInfoExt) BindUi() fyne.CanvasObject {
	vb := container.NewVBox()
	hb := container.NewHBox()

	name := widget.NewLabel(this.Name)
	vb.Add(name)

	bar := widget.NewProgressBarWithData(binding.BindFloat(&this.Tmp.HP))
	vb.Add(bar)

	hp := widget.NewLabelWithData(binding.BindString(&this.Tmp.HPStr))
	vb.Add(hp)
	att := widget.NewLabelWithData(binding.BindString(&this.Tmp.AttackStr))
	vb.Add(att)
	def := widget.NewLabelWithData(binding.BindString(&this.Tmp.DefStr))
	vb.Add(def)

	attsp := widget.NewLabelWithData(binding.BindString(&this.Tmp.AttackSpeedStr))
	vb.Add(attsp)

	a := widget.NewLabel("A")
	hb.Add(a)
	b := widget.NewLabel("B")
	hb.Add(b)
	c := widget.NewLabel("C")
	hb.Add(c)

	vb.Add(hb)
	return vb
}

func (this *TargetInfoExt) Alive() bool {
	return this.Tmp.HP > 0
}

func (this *TargetInfoExt) Attack(tar *TargetInfoExt) {
	dm := this.Tmp.Attack * (1 - tar.GetReduce())
	tar.Tmp.HP -= dm
	st := fmt.Sprintf("%s attack %s,%s hp-%0.1f", this.Name, tar.Name, tar.Name, dm, " -", Decimal(dm, 1))

	this.AttInfo.AttInfo.WriteString(st)
	this.AttInfo.AttInfo.WriteString("\n")
	tar.AttInfo.AttInfo.WriteString(st)
	tar.AttInfo.AttInfo.WriteString("\n")
	// fmt.Println(tar.Name, " -", Decimal(dm, 1))
}
func (this *TargetInfoExt) Rebuild() {
	this.Tmp = this.NorAtt
	this.AttInfo.AttInfo.Reset()
	this.Tmp.RunAttackSpeed = this.Tmp.AttackSpeed
	this.Tmp.HPStr = fmt.Sprintf("HP: %.1f / %.1f", this.Tmp.HP, this.NorAtt.HP)
	this.Tmp.DefStr = fmt.Sprintf("ATT: [ %.1f ]", this.Tmp.Attack)
	this.Tmp.AttackStr = fmt.Sprintf("DEF: %.1f [%.1f%%]", this.Tmp.Def, this.GetReduce()*100)
	this.Tmp.AttackSpeedStr = fmt.Sprintf("ATTSPEED: [%.1f]", this.Tmp.AttackSpeed)
}
func (this *TargetInfoExt) FlushStr() {
	this.Tmp.HPStr = fmt.Sprintf("HP: %.1f / %.1f", this.Tmp.HP, this.NorAtt.HP)
	this.Tmp.DefStr = fmt.Sprintf("ATT: [ %.1f ]", this.Tmp.Attack)
	this.Tmp.AttackStr = fmt.Sprintf("DEF: %.1f [%.1f%%]", this.Tmp.Def, this.GetReduce()*100)
	this.Tmp.AttackSpeedStr = fmt.Sprintf("ATTSPEED: [%.1f]", this.Tmp.AttackSpeed)
}

func (this *TargetInfoExt) ShowStatus() {
	// this.Tmp.HPStr = fmt.Sprintf("HP: %.1f / %.1f", this.Tmp.HP, this.NorAtt.HP)
	fmt.Printf("%s HP:[ %.1f / %.1f ] ATT: [ %.1f ] DEF: %.1f [%.1f%%] ATTSPEED: [%.1f]\n", this.Name, this.Tmp.HP, this.NorAtt.HP, this.Tmp.Attack, this.Tmp.Def, this.GetReduce()*100, this.Tmp.AttackSpeed)
	// fmt.Println(this.Name, " hp:", this.Tmp.HP, " attack:", this.Tmp.Attack, " def:", this.Tmp.Def, fmt.Sprintf("[%g%%]", this.GetReduce()*100))
}
func (this *TargetInfoExt) GetReduce() (percent float64) {
	return Decimal((0.052*this.Tmp.Def)/(0.9+0.048*this.Tmp.Def), 2)
}
func (this *TargetInfoExt) Upgrade() fyne.CanvasObject {
	this.Level++
	Clear()
	User.Rebuild()
	Monster.Rebuild()
	fmt.Println("Upgrade：", "LV:", this.Level)
	var Att_1 []AttributesBasicExt
	Att_1 = make([]AttributesBasicExt, 5)
	// Attack>hp=def
	Att_1[0].HP = BuildRandAtt(200, 10, 55)
	Att_1[0].Def = BuildRandAtt(5, 5, 45)
	Att_1[0].Attack = BuildRandAtt(15, 40, 100)
	Att_1[0].AttackSpeed = BuildRandAtt(20, 5, 30)
	// hp>Attack=def
	Att_1[1].HP = BuildRandAtt(250, 40, 100)
	Att_1[1].Def = BuildRandAtt(5, 5, 45)
	Att_1[1].Attack = BuildRandAtt(10, 10, 55)
	Att_1[1].AttackSpeed = BuildRandAtt(20, 5, 30)
	// Attack=hp=def= attsp
	Att_1[2].HP = BuildRandAtt(200, 30, 90)
	Att_1[2].Def = BuildRandAtt(5, 20, 65)
	Att_1[2].Attack = BuildRandAtt(12, 30, 100)
	Att_1[2].AttackSpeed = BuildRandAtt(30, 5, 30)
	// def>att = hp
	Att_1[3].HP = BuildRandAtt(200, 20, 55)
	Att_1[3].Def = BuildRandAtt(6, 40, 100)
	Att_1[3].Attack = BuildRandAtt(10, 10, 55)
	Att_1[3].AttackSpeed = BuildRandAtt(30, 10, 40)
	// def>att = hp
	Att_1[4].HP = BuildRandAtt(150, 10, 55)
	Att_1[4].Def = BuildRandAtt(3, 5, 45)
	Att_1[4].Attack = BuildRandAtt(12, 15, 85)
	Att_1[4].AttackSpeed = BuildRandAtt(40, 5, 100)
	User.ShowStatus()
	Monster.ShowStatus()
	fmt.Println("==========================================================================")
	str1 := fmt.Sprintf("1: attack: +", Att_1[0].Attack, " hp: +", Att_1[0].HP, " def: +", Att_1[0].Def, " AttackSpeed: +", Att_1[0].AttackSpeed)
	str2 := fmt.Sprintf("2: attack: +", Att_1[1].Attack, " hp: +", Att_1[1].HP, " def: +", Att_1[1].Def, " AttackSpeed: +", Att_1[1].AttackSpeed)
	str3 := fmt.Sprintf("3: attack: +", Att_1[2].Attack, " hp: +", Att_1[2].HP, " def: +", Att_1[2].Def, " AttackSpeed: +", Att_1[2].AttackSpeed)
	str4 := fmt.Sprintf("4: attack: +", Att_1[3].Attack, " hp: +", Att_1[3].HP, " def: +", Att_1[3].Def, " AttackSpeed: +", Att_1[3].AttackSpeed)
	str5 := fmt.Sprintf("5: attack: +", Att_1[4].Attack, " hp: +", Att_1[4].HP, " def: +", Att_1[4].Def, " AttackSpeed: +", Att_1[4].AttackSpeed)
	fmt.Println("==========================================================================")
	fmt.Println("input choose：")
	bt1 := widget.NewButton(str1, func() {
		key := 0
		this.NorAtt.HP = Decimal(this.NorAtt.HP+Att_1[key].HP, 1)
		this.NorAtt.Attack = Decimal(this.NorAtt.Attack+Att_1[key].Attack, 1)
		this.NorAtt.Def = Decimal(this.NorAtt.Def+Att_1[key].Def, 1)
		this.NorAtt.AttackSpeed = Decimal(this.NorAtt.AttackSpeed+Att_1[key].AttackSpeed, 1)
	})
	bt2 := widget.NewButton(str2, func() {
		key := 1
		this.NorAtt.HP = Decimal(this.NorAtt.HP+Att_1[key].HP, 1)
		this.NorAtt.Attack = Decimal(this.NorAtt.Attack+Att_1[key].Attack, 1)
		this.NorAtt.Def = Decimal(this.NorAtt.Def+Att_1[key].Def, 1)
		this.NorAtt.AttackSpeed = Decimal(this.NorAtt.AttackSpeed+Att_1[key].AttackSpeed, 1)
	})
	bt3 := widget.NewButton(str3, func() {
		key := 2
		this.NorAtt.HP = Decimal(this.NorAtt.HP+Att_1[key].HP, 1)
		this.NorAtt.Attack = Decimal(this.NorAtt.Attack+Att_1[key].Attack, 1)
		this.NorAtt.Def = Decimal(this.NorAtt.Def+Att_1[key].Def, 1)
		this.NorAtt.AttackSpeed = Decimal(this.NorAtt.AttackSpeed+Att_1[key].AttackSpeed, 1)
	})
	bt4 := widget.NewButton(str4, func() {
		key := 3
		this.NorAtt.HP = Decimal(this.NorAtt.HP+Att_1[key].HP, 1)
		this.NorAtt.Attack = Decimal(this.NorAtt.Attack+Att_1[key].Attack, 1)
		this.NorAtt.Def = Decimal(this.NorAtt.Def+Att_1[key].Def, 1)
		this.NorAtt.AttackSpeed = Decimal(this.NorAtt.AttackSpeed+Att_1[key].AttackSpeed, 1)
	})
	bt5 := widget.NewButton(str5, func() {
		key := 4
		this.NorAtt.HP = Decimal(this.NorAtt.HP+Att_1[key].HP, 1)
		this.NorAtt.Attack = Decimal(this.NorAtt.Attack+Att_1[key].Attack, 1)
		this.NorAtt.Def = Decimal(this.NorAtt.Def+Att_1[key].Def, 1)
		this.NorAtt.AttackSpeed = Decimal(this.NorAtt.AttackSpeed+Att_1[key].AttackSpeed, 1)
	})
	box := container.NewHBox(bt1, bt2, bt3, bt4, bt5)
	// res := GetWantInput(1, 2, 3, 4, 5)
	// this.NorAtt.HP = Decimal(this.NorAtt.HP+Att_1[res-1].HP, 1)
	// this.NorAtt.Attack = Decimal(this.NorAtt.Attack+Att_1[res-1].Attack, 1)
	// this.NorAtt.Def = Decimal(this.NorAtt.Def+Att_1[res-1].Def, 1)
	// this.NorAtt.AttackSpeed = Decimal(this.NorAtt.AttackSpeed+Att_1[res-1].AttackSpeed, 1)
	return box
}

func BuildRandAtt(att, min, max float64) float64 {
	rand.Seed(time.Now().UnixNano())
	r := (rand.Float64()*(max-min) + min) / 100
	time.Sleep(time.Duration(rand.Int63n(100000)) * time.Nanosecond)
	return Decimal(att*r, 1)
}

func (this *TargetInfoExt) BuildNew(min, max float64) {
	// rand.Seed(time.Now().UnixNano())
	// r := (rand.Float64()*(max-min) + min) / 100
	this.NorAtt.HP = Decimal(this.NorAtt.HP+BuildRandAtt(200, 1, 100), 1)
	this.NorAtt.Attack = Decimal(this.NorAtt.Attack+BuildRandAtt(10, 1, 100), 1)
	this.NorAtt.Def = Decimal(this.NorAtt.Def+BuildRandAtt(3, 1, 100), 1)
	this.NorAtt.AttackSpeed = Decimal(this.NorAtt.AttackSpeed+BuildRandAtt(25, 0, 100), 1)
	this.Tmp = this.NorAtt
	return
}
