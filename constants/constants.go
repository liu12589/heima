package constants

const (
	UserName = "root"
	Password = "123456"
	Host     = "10.233.12.236"
	Port     = "3306"
	Dbname   = "heima"
	Timeout  = "10s"
)

const (
	Login = "login"
)

const (
	KubeConfig = "kubeconfig"
)

var Ports = []int{8001, 8002, 8003, 8004, 8005, 8006, 8007}

func GetPort() int32 {
	var port int
	if len(Ports) != 0 {
		port = Ports[0]
	}
	Ports = Ports[1:]
	return int32(port)
}
