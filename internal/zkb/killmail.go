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

	Position Position

	// actually not present in the killmail
	// added by the indexing step

	RegionID   uint64
	RegionName string
	Ref        bobstore.Ref
}

// Attacker - an attacker
type Attacker struct {
	CharacterID     uint64
	CharacterName   string
	CorporationID   uint64
	CorporationName string
	AllianceID      uint64
	AllianceName    string
	FactionID       uint64
	FactionName     string

	SecStatus  float64
	DamageDone float64
	FinalBlow  bool

	ShipTypeID     uint64
	ShipTypeName   string
	WeaponTypeID   uint64
	WeaponTypeName string
}

// Victim - the victim
type Victim struct {
	CharacterID     uint64
	CharacterName   string
	CorporationID   uint64
	CorporationName string
	AllianceID      uint64
	AllianceName    string
	FactionID       uint64
	FactionName     string

	DamageTaken float64

	ShipTypeID   uint64
	ShipTypeName string

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
	ItemTypeID        uint64
	ItemTypeName      string
	QuantityDropped   int
	QuantityDestroyed int
	Singleton         int
	Flag              int
}

// Parse a json killmail
func Parse(b []byte, ref bobstore.Ref) (*Killmail, error) {
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

		Position: Position{
			X: vq.Float("position", "x"),
			Y: vq.Float("position", "y"),
			Z: vq.Float("position", "z"),
		},

		Victim: Victim{
			CharacterID:     vq.UInt64("character", "id"),
			CharacterName:   vq.String("character", "name"),
			CorporationID:   vq.UInt64("corporation", "id"),
			CorporationName: vq.String("corporation", "name"),
			AllianceID:      vq.UInt64("alliance", "id"),
			AllianceName:    vq.String("alliance", "name"),
			FactionID:       vq.UInt64("faction", "id"),
			FactionName:     vq.String("faction", "name"),

			DamageTaken: vq.Float("damageDone"),

			ShipTypeID:   vq.UInt64("shipType", "id"),
			ShipTypeName: vq.String("shipType", "name"),
		},

		Ref: ref,
	}

	km.RegionID, km.RegionName = RegionBySolarSystem(km.SolarSystemID)

	items := vq.Slice("items")
	km.Victim.Items = make([]Item, len(items))
	for i, jitem := range items {
		iq := jq.New(jitem)
		item := &km.Victim.Items[i]
		item.ItemTypeID = iq.UInt64("itemType", "id")
		item.ItemTypeName = iq.String("itemType", "name")
		item.QuantityDestroyed = iq.Int("quantityDestroyed")
		item.QuantityDropped = iq.Int("quantityDropped")
		item.Flag = iq.Int("flag")
		item.Singleton = iq.Int("singleton")
	}

	attackers := kq.Slice("attackers")
	km.Attackers = make([]Attacker, len(attackers))
	for i, jatt := range attackers {
		aq := jq.New(jatt)
		a := &km.Attackers[i]
		a.CharacterID = aq.UInt64("character", "id")
		a.CharacterName = aq.String("character", "name")
		a.CorporationID = aq.UInt64("corporation", "id")
		a.CorporationName = aq.String("corporation", "name")
		a.AllianceID = aq.UInt64("alliance", "id")
		a.AllianceName = aq.String("alliance", "name")
		a.FactionID = aq.UInt64("faction", "id")
		a.FactionName = aq.String("faction", "name")

		a.SecStatus = aq.Float("securityStatus")
		a.DamageDone = aq.Float("damageDone")

		a.ShipTypeID = aq.UInt64("shipType", "id")
		a.ShipTypeName = aq.String("shipType", "name")
		a.WeaponTypeID = aq.UInt64("weaponType", "id")
		a.WeaponTypeName = aq.String("weaponType", "name")
	}

	return km, nil
}
