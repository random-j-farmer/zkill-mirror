#! /bin/bash
# victim stats - zkillboard returns
jq '.[] | {"killid": .killID, "killTime": .killTime, "system": .solarSystemName, "region": .regionName, "pilot": .victim.characterName, "id": .victim.characterID, "corp":  .victim.corporationName, "corpid": .victim.corporationID, "alliance": .victim.allianceName , "allid": .victim.allianceID, "ship": .victim.shipTypeName}' 
