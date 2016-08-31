package zkb

import (
	"github.com/pkg/errors"
	"github.com/random-j-farmer/bobstore"
	"github.com/random-j-farmer/jq"
)

// Killmail contains the data from a eve json killmail
type Killmail struct {
	KillID          uint64
	KillTime        string // "2016.08.28 18:10:28"
	SolarSystemName string
	SolarSystemID   uint64

	WarID uint64

	Victim Victim

	AttackerCount int
	Attackers     []Attacker

	ZKBTotalValue float64
	ZKBPoints     int
}

// Attacker - an attacker
type Attacker struct {
	CharID          uint64
	CharName        string
	CorporationID   uint64
	CorporationName string
	AllianceID      uint64
	AllianceName    string

	SecStatus  float64
	DamageDone float64

	ShipID     uint64
	ShipName   string
	WeaponID   uint64
	WeaponName string
}

// Victim - the victim
type Victim struct {
	CharID          uint64
	CharName        string
	CorporationID   uint64
	CorporationName string
	AllianceID      uint64
	AllianceName    string

	DamageTaken float64

	ShipID   uint64
	ShipName string

	Position Position

	Items []Item
}

// Position - coordinates in the solarsystem of the kill
type Position struct {
	X float64
	Y float64
	Z float64
}

// Item - an item that dropped or was destroyed
type Item struct {
	ItemID            uint64
	ItemName          string
	QuantityDropped   int
	QuantityDestroyed int
}

// KillmailWithRef - a killmail with associated ref
type KillmailWithRef struct {
	*Killmail
	bobstore.Ref
}

// Parse a json killmail
func Parse(b []byte) (*Killmail, error) {
	q, err := jq.Unmarshal(b)
	if err != nil {
		return nil, errors.Wrap(err, "jq.New")
	}

	kq := jq.New(q.Map("package", "killmail"))
	vq := jq.New(kq.Map("victim"))

	km := &Killmail{
		KillID:          kq.UInt64("killID"),
		KillTime:        kq.String("killTime"),
		SolarSystemID:   kq.UInt64("solarSystem", "id"),
		SolarSystemName: kq.String("solarSystem", "name"),

		WarID: kq.UInt64("warID"),

		ZKBTotalValue: q.Float("zkb", "totalValue"),
		ZKBPoints:     q.Int("zkb", "points"),

		AttackerCount: kq.Int("attackerCount"),

		Victim: Victim{
			CharID:          vq.UInt64("character", "id"),
			CharName:        vq.String("character", "name"),
			CorporationID:   vq.UInt64("corporation", "id"),
			CorporationName: vq.String("corporation", "name"),
			AllianceID:      vq.UInt64("alliance", "id"),
			AllianceName:    vq.String("alliance", "name"),

			DamageTaken: vq.Float("damageDone"),

			ShipID:   vq.UInt64("shipType", "id"),
			ShipName: vq.String("shipType", "name"),

			Position: Position{
				X: vq.Float("position", "x"),
				Y: vq.Float("position", "y"),
				Z: vq.Float("position", "z"),
			},
		},
	}

	items := vq.Slice("items")
	km.Victim.Items = make([]Item, len(items))
	for i, jitem := range items {
		iq := jq.New(jitem)
		item := &km.Victim.Items[i]
		item.ItemID = iq.UInt64("itemType", "id")
		item.ItemName = iq.String("itemType", "name")
		item.QuantityDestroyed = iq.Int("quantityDestroyed")
		item.QuantityDropped = iq.Int("quantityDropped")
	}

	attackers := kq.Slice("attackers")
	km.Attackers = make([]Attacker, len(attackers))
	for i, jatt := range attackers {
		aq := jq.New(jatt)
		a := &km.Attackers[i]
		a.CharID = aq.UInt64("character", "id")
		a.CharName = aq.String("character", "name")
		a.CorporationID = aq.UInt64("corporation", "id")
		a.CorporationName = aq.String("corporation", "name")
		a.AllianceID = aq.UInt64("alliance", "id")
		a.AllianceName = aq.String("alliance", "name")

		a.SecStatus = aq.Float("securityStatus")
		a.DamageDone = aq.Float("damageDone")

		a.ShipID = aq.UInt64("shipType", "id")
		a.ShipName = aq.String("shipType", "name")
		a.WeaponID = aq.UInt64("weaponType", "id")
		a.WeaponName = aq.String("weaponType", "name")
	}

	return km, nil
}
