package conf

const (
	DriveName = "mysql"
)

type DbConf struct {
	Host   string
	Port   string
	User   string
	Pwd    string
	DbName string
}

var MasterDbConf DbConf = DbConf{
	Host:   "127.0.01",
	Port:   "3306",
	User:   "root",
	Pwd:    "123456",
	DbName: "supperstar",
}

var SlaveDbConf DbConf = DbConf{
	Host:   "127.0.01",
	Port:   "3306",
	User:   "root",
	Pwd:    "root",
	DbName: "supperstar",
}
