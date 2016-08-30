package assets

//go:generate go-bindata --debug -tags !ndebug -pkg assets static/ static/pure-release-0.6.0/ templates/
//go:generate go-bindata -o bindata_ndebug.go -tags ndebug -pkg assets static/ static/pure-release-0.6.0/ templates/
