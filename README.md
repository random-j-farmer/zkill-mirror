ZKillboard Mirror
=================

Development
===========

Additional packages:

    # for live reloading while developing
    go get github.com/codegangsta/gin

    # to compile assets (static files, templates) into the appropriate
    get github.com/jteeuwen/go-bindata/...

    # boltdb
    go get -u github.com/boltdb/bolt/...


To run the server with gin:

    ZKM_PORT=8081 ZKM_VERBOSE=true ZKM_CACHE_TEMPLATES=false gin -p 8080 -a 8081 run

Or create a custom config file in ~/.ZKILL-MIRROR/, this overrides the
example config in the distribution.

  time http http://localhost:3000/api/ | jq -e '.[].package.killmail.solarSystem.name' | wc -l
  1000

  real	0m0.489s
  user	0m0.437s
  sys	0m0.067s

Embedding templates and static files
------------------------------------

Run `go generate` in
internal/assets.  For development, use `go generate -tags debug` -
this will generate special embeddings that point back to the original file.
In combination with ZKM_CACHE_TEMPLATES=false, changed files will be
served immediately.

Before making a production release, run go generate without the debug tag.
This will really embed the files in internal/assets, only the executable
is needed on the production server.
