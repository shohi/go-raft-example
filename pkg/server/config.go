package server

type RaftConfig struct {
	RaftPort int
	JoinAddr string
	NodeID   string
}

type Config struct {
	RaftConfig

	HTTPPort int
}
