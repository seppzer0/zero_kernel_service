package config

type Config struct {
	Server struct {
		ListenAddress string
	}
	Worker struct {
		Location    string
		Concurrency string
	}
	DeviceList struct {
		Codename []string
	}
}
