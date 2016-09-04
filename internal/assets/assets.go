package assets

//go:generate go-bindata -o bindata_dev.go --dev -tags dev -pkg assets static/ static/pure-release-0.6.0/ templates/
//go:generate go-bindata -tags !dev -pkg assets static/ static/pure-release-0.6.0/ templates/
