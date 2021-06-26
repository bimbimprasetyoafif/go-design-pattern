package main

import "fmt"

type MinumanType string

const (
	AirPutih  MinumanType = "airputih"
	KopiHitam MinumanType = "kopihitam"
	LemonTea  MinumanType = "lemontea"
)

type IMinuman interface {
	tambahkanAir()
	tambahkanGula()
	tambahkanPerisaBubukLain()
	hidangkan() Minuman
}

func pesanMinuman(minumanType MinumanType) IMinuman {
	switch minumanType {
	case AirPutih:
		return &MinumanAirPutih{}
	case KopiHitam:
		return &MinumanKopi{}
	case LemonTea:
		return &MinumanlemonTea{}
	}
	return nil
}

type Minuman struct {
	kadarAir     int
	kadarGula    int
	campuranLain string
}

type MinumanAirPutih struct {
	Minuman
}

func (a *MinumanAirPutih) tambahkanAir() {
	a.kadarAir = 100
}

func (a *MinumanAirPutih) tambahkanGula() {
	a.kadarGula = 0
}

func (a *MinumanAirPutih) tambahkanPerisaBubukLain() {
	a.campuranLain = "tidak ada"
}

func (a *MinumanAirPutih) hidangkan() Minuman {
	return Minuman{
		kadarAir:     a.kadarAir,
		kadarGula:    a.kadarGula,
		campuranLain: a.campuranLain,
	}
}

type MinumanKopi struct {
	Minuman
}

func (a *MinumanKopi) tambahkanAir() {
	a.kadarAir = 50
}

func (a *MinumanKopi) tambahkanGula() {
	a.kadarGula = 10
}

func (a *MinumanKopi) tambahkanPerisaBubukLain() {
	a.campuranLain = "kopi kapal api"
}

func (a *MinumanKopi) hidangkan() Minuman {
	return Minuman{
		kadarAir:     a.kadarAir,
		kadarGula:    a.kadarGula,
		campuranLain: a.campuranLain,
	}
}

type MinumanlemonTea struct {
	Minuman
}

func (a *MinumanlemonTea) tambahkanAir() {
	a.kadarAir = 50
}

func (a *MinumanlemonTea) tambahkanGula() {
	a.kadarGula = 50
}

func (a *MinumanlemonTea) tambahkanPerisaBubukLain() {
	a.campuranLain = "marimas jeruk"
}

func (a *MinumanlemonTea) hidangkan() Minuman {
	return Minuman{
		kadarAir:     a.kadarAir,
		kadarGula:    a.kadarGula,
		campuranLain: a.campuranLain,
	}
}

type Kedai struct {
	Bartender IMinuman
}

func panggilPelayan(m IMinuman) *Kedai {
	return &Kedai{
		Bartender: m,
	}
}

func (k *Kedai) gantiPelayan(m IMinuman) {
	k.Bartender = m
}

func (k *Kedai) sajikanMinuman() Minuman {
	k.Bartender.tambahkanAir()
	k.Bartender.tambahkanGula()
	k.Bartender.tambahkanPerisaBubukLain()
	return k.Bartender.hidangkan()
}

func main() {
	airPutih := pesanMinuman(AirPutih)
	kopi := pesanMinuman(KopiHitam)
	lemonTea := pesanMinuman(LemonTea)

	bartender := panggilPelayan(airPutih)
	minumanSiap := bartender.sajikanMinuman()
	fmt.Println(minumanSiap.campuranLain)

	bartender.gantiPelayan(kopi)
	minumanSiap = bartender.sajikanMinuman()
	fmt.Println(minumanSiap.campuranLain)

	bartender.gantiPelayan(lemonTea)
	minumanSiap = bartender.sajikanMinuman()
	fmt.Println(minumanSiap.campuranLain)
}
