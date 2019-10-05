package conf

const (
	// frigg自己的网络信息，LISTEN_IP和SERVICE_IP区别，是因为腾讯云服务器分内网ip和外网ip，需要监听内网ip，外网访问外网ip会转发到监听的内网地址
	LISTEN_IP = "172.27.0.15"
	SERVICE_IP = "139.155.46.62"
	//LISTEN_IP = "127.0.0.1"
	//SERVICE_IP = "127.0.0.1"
	SERVICE_PORT  = 42222
	SERVICE_NAME = "frigg_db_service"
	TTL = 5

	// mimiron's name
	ALGORITHM_SERVICE_NAME = "mimiron_algorithm_service"
)
