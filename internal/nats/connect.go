package nats

import "github.com/nats-io/stan.go"

type Conf struct {
	Cluster string
	Client  string
}

// Init подключение к nats-streaming
func Init(conf Conf) (*stan.Conn, error) {
	sc, err := stan.Connect(conf.Cluster, conf.Client)
	if err != nil {
		return nil, err
	}

	return &sc, nil
}

func Close(conn stan.Conn) error {

	err := conn.Close()
	return err
}
