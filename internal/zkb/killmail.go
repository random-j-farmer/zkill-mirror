package zkb

import (
	"regexp"

	"github.com/pkg/errors"
	"github.com/random-j-farmer/bobstore"
	"github.com/random-j-farmer/jq"
)

var dateNormalizer = regexp.MustCompile(`[.: ]`)

// Killmail contains the data from a eve json killmail
type Killmail struct {
	KillID int
	// "2016.08.28 18:10:28"
	KillTime        string
	SolarSystemName string
	SolarSystemID   int

	WarID int

	Victim Victim

	AttackerCount int
	Attackers     []Attacker

	ZKBTotalValue float64
	ZKBPoints     int
}

// Attacker - an attacker
type Attacker struct {
	CharID          int
	CharName        string
	CorporationID   int
	CorporationName string
	AllianceID      int
	AllianceName    string

	SecStatus  float32
	DamageDone float32

	ShipID     int
	ShipName   string
	WeaponID   int
	WeaponName string
}

// Victim - the victim
type Victim struct {
	CharID          int
	CharName        string
	CorporationID   int
	CorporationName string
	AllianceID      int
	AllianceName    string

	DamageTaken float32

	ShipID   int
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
	ItemID            int
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
	q, err := jq.New(b)
	if err != nil {
		return nil, errors.Wrap(err, "jq.New")
	}

	kq := jq.NewFromInterface(q.Map("package", "killmail"))
	vq := jq.NewFromInterface(kq.Map("victim"))

	killtime := dateNormalizer.ReplaceAllLiteralString(kq.String("killTime"), "")

	km := &Killmail{
		KillID:          kq.Int("killID"),
		KillTime:        killtime,
		SolarSystemID:   kq.Int("solarSystem", "id"),
		SolarSystemName: kq.String("solarSystem", "name"),

		WarID: kq.Int("warID"),

		ZKBTotalValue: q.Float("zkb", "totalValue"),
		ZKBPoints:     q.Int("zkb", "points"),

		AttackerCount: kq.Int("attackerCount"),

		Victim: Victim{
			CharID:          vq.Int("character", "id"),
			CharName:        vq.String("character", "name"),
			CorporationID:   vq.Int("corporation", "id"),
			CorporationName: vq.String("corporation", "name"),
			AllianceID:      vq.Int("alliance", "id"),
			AllianceName:    vq.String("alliance", "name"),

			DamageTaken: float32(vq.Float("damageDone")),

			ShipID:   vq.Int("shipType", "id"),
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
		iq := jq.NewFromInterface(jitem)
		item := &km.Victim.Items[i]
		item.ItemID = iq.Int("itemType", "id")
		item.ItemName = iq.String("itemType", "name")
		item.QuantityDestroyed = iq.Int("quantityDestroyed")
		item.QuantityDropped = iq.Int("quantityDropped")
	}

	attackers := kq.Slice("attackers")
	km.Attackers = make([]Attacker, len(attackers))
	for i, jatt := range attackers {
		aq := jq.NewFromInterface(jatt)
		a := &km.Attackers[i]
		a.CharID = aq.Int("character", "id")
		a.CharName = aq.String("character", "name")
		a.CorporationID = aq.Int("corporation", "id")
		a.CorporationName = aq.String("corporation", "name")
		a.AllianceID = aq.Int("alliance", "id")
		a.AllianceName = aq.String("alliance", "name")

		a.SecStatus = float32(aq.Float("securityStatus"))
		a.DamageDone = float32(aq.Float("damageDone"))

		a.ShipID = aq.Int("shipType", "id")
		a.ShipName = aq.String("shipType", "name")
		a.WeaponID = aq.Int("weaponType", "id")
		a.WeaponName = aq.String("weaponType", "name")
	}

	return km, nil
}
