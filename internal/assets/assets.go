package assets

//go:generate go-bindata -o bindata_dev.go --dev -tags dev -pkg assets static static/datatables.net static/datatables.net/js static/datatables.net-dt static/datatables.net-dt/css static/datatables.net-dt/images static/jquery static/jquery/dist static/pure templates
//go:generate go-bindata -tags !dev -pkg assets static static/datatables.net static/datatables.net/js static/datatables.net-dt static/datatables.net-dt/css static/datatables.net-dt/images static/jquery static/jquery/dist static/pure templates
