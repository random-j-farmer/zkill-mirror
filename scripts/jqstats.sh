#! /bin/bash
# victim stats
jq '.[].package.killmail | {"killid": .killID, "killTime": .killTime, "system": .solarSystem.name, "pilot": .victim.character.name, "id": .victim.character.id, "corp":  .victim.corporation.name, "corpid": .victim.corporation.id, "alliance": .victim.alliance.name , "allid": .victim.alliance.id, "ship": .victim.shipType.name}' 
