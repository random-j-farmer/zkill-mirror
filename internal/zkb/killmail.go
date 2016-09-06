package zkb

import (
	"github.com/pkg/errors"
	"github.com/random-j-farmer/bobstore"
	"github.com/random-j-farmer/jq"
	"github.com/random-j-farmer/zkill-mirror/internal/mapdata"
)

// Killmail contains the data from a eve json killmail
// the layout of these aims to map 1:1 to the json
// returned by zkillboard queries (which is not what is pulled,
// that is eve crest api json).
// Exceptions:
// * shipTypeName, solarSystemName, regionID and regionName are returned,
//   even if they are not in zkillboard data.
type Killmail struct {
	KillID          uint64 `json:"killID"`
	KillTime        string `json:"killTime"` // "2016.08.28 18:10:28"
	SolarSystemName string `json:"solarSystemName"`
	SolarSystemID   uint64 `json:"solarSystemID"`

	WarID uint64 `json:"warID"`

	Victim Victim `json:"victim"`

	AttackerCount int        `json:"attackerCount"`
	Attackers     []Attacker `json:"attackers"`

	ZKB ZKB `json:"zkb"`

	Position Position `json:"position"`

	// actually not present in the killmail
	// added by the indexing step

	RegionID   uint64       `json:"regionID"`
	RegionName string       `json:"regionName"`
	Ref        bobstore.Ref `json:"-"`
}

// Attacker - an attacker
type Attacker struct {
	CharacterID     uint64 `json:"characterID"`
	CharacterName   string `json:"characterName"`
	CorporationID   uint64 `json:"corporationID"`
	CorporationName string `json:"corporationName"`
	AllianceID      uint64 `json:"allianceID"`
	AllianceName    string `json:"allianceName"`
	FactionID       uint64 `json:"factionID"`
	FactionName     string `json:"factionName"`

	SecStatus  float64 `json:"secStatus"`
	DamageDone float64 `json:"damageDone"`
	FinalBlow  int     `json:"finalBlow"`

	ShipTypeID     uint64 `json:"shipTypeID"`
	ShipTypeName   string `json:"shipTypeName"`
	WeaponTypeID   uint64 `json:"weaponTypeID"`
	WeaponTypeName string `json:"-"`
}

// Victim - the victim
type Victim struct {
	CharacterID     uint64 `json:"characterID"`
	CharacterName   string `json:"characterName"`
	CorporationID   uint64 `json:"corporationID"`
	CorporationName string `json:"corporationName"`
	AllianceID      uint64 `json:"allianceID"`
	AllianceName    string `json:"allianceName"`
	FactionID       uint64 `json:"factionID"`
	FactionName     string `json:"factionName"`

	DamageTaken float64 `json:"damageTaken"`

	ShipTypeID   uint64 `json:"shipTypeID"`
	ShipTypeName string `json:"shipTypeName"`

	Items []Item `json:"items"`
}

// Position - coordinates in the solarsystem of the kill
type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

// Item - an item that dropped or was destroyed
type Item struct {
	ItemTypeID        uint64 `json:"itemTypeID"`
	ItemTypeName      string `json:"itemTypeName"`
	QuantityDropped   int    `json:"quantityDropped"`
	QuantityDestroyed int    `json:"quantityDestroyed"`
	Singleton         int    `json:"singleton"`
	Flag              int    `json:"flag"`
}

// ZKB zkillboard data
type ZKB struct {
	TotalValue float64 `json:"totalValue"`
	LocationID uint64  `json:"locationID"`
	Hash       string  `json:"hash"`
	Points     int     `json:"points"`
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

		ZKB: ZKB{
			TotalValue: q.Float("package", "zkb", "totalValue"),
			Points:     q.Int("package", "zkb", "points"),
			Hash:       q.String("package", "zkb", "hash"),
			LocationID: q.UInt64("package", "zkb", "locationID"),
		},

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
			DamageTaken:     vq.Float("damageTaken"),
			ShipTypeID:      vq.UInt64("shipType", "id"),
			ShipTypeName:    vq.String("shipType", "name"),
		},

		Ref: ref,
	}

	km.RegionID, km.RegionName = mapdata.RegionBySolarSystem(km.SolarSystemID)

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
		a.FinalBlow = aq.Int("finalBlow")
		a.ShipTypeID = aq.UInt64("shipType", "id")
		a.ShipTypeName = aq.String("shipType", "name")
		a.WeaponTypeID = aq.UInt64("weaponType", "id")
		a.WeaponTypeName = aq.String("weaponType", "name")
	}

	return km, nil
}
