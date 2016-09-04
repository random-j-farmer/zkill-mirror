ZKillboard Mirror
=================

Mirrors zkillboard data by using the pull api.

Implements limited zkillboard-like json api (see example scripts).

Installation
------------

You need Go installed.  I use 1.7, it used to work with 1.6.

    # install & compile
    go get -v github.com/random-j-farmer/zkill-mirror

    cd DIR_FOR_YOUDATE  # around 300MB per week
    zkill-mirror serve  # i do this in screen, or in a docker container

Updating
--------

After updating to a new version, you usually have to reindex the stored killmails.
On my server, this takes about one minute per weeks worth of killmails.
This is done like so:

    # install & compile
    go get -v -u github.com/random-j-farmer/zkill-mirror
    # kill old process, update to new binary
    kill WHATEVER_THE_PID_IS

    rm -f zkill-mirror.bolt  # delete the bolt db
    zkill-mirror reindex     # create new bolt db with current index scheme
    zkill-mirror serve       # profit

This might become automated at some point, or the indexing scheme might become
more stable.  In the near future, my #1 goal is to reduce the db/index size,
so it is not going to be stable.

Development
===========

Additional packages:

    # for live reloading while developing
    # go get github.com/codegangsta/gin
    # my fork has a patch merged in for build tags support
    go get github.com/random-j-farmer/gin

    # to compile assets (static files, templates) into the appropriate
    get github.com/jteeuwen/go-bindata/...

    # boltdb
    go get -u github.com/boltdb/bolt/...


To run the server with gin:

    ZKM_PORT=8081 ZKM_VERBOSE=true ZKM_CACHE_TEMPLATES=false gin --tags dev -p 8080 -a 8081 run

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


Reindexing
----------

* Naive approach: Reindexing 1.7MB of gzipped input took 15 seconds (db_nosync=true).
  bobstore gzip on the same input took 0.5 seconds. Doing it in batches of 100 (workers=4)
  turned that into 0.5 seconds for the same input.  With db_nosync=true, reindex_workers=8,
  reindex_batch_size=1000 turned it into 0.293 seconds.  Same settings but db_nosync_false:
  0.297 seconds.  So: db_nosync not needed anymore
* Naive approach for 100MB was 11(? i think) minutes (db_nosync=true). Batching the day after:
  no sync, batch size 100, 4 workers: 38 seconds.  with 8 workers and batch_size 1000: 32 seconds.
  bobstore gzip takes 1min 12 seconds (although decompress/compress is done in only 1 thread).
  cp of the input file: 0.5 seconds.
