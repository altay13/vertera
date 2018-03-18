package cassandra

import (
	"github.com/altay13/vertera/eventHandler"
	"github.com/gocql/gocql"
)

type Cassandra struct {
	Config
	cluster *gocql.ClusterConfig
	session *gocql.Session
}

func NewCassandra(conf *Config) *Cassandra {
	c := &Cassandra{
		Config: *conf,
	}

	return c
}

func (c *Cassandra) newCluster() {
	c.cluster = gocql.NewCluster(c.Host...)
	c.session, _ = c.cluster.CreateSession()
	c.Consistency = c.Consistency
}

func (c *Cassandra) Disconnect() {
	c.session.Close()
}

func (c *Cassandra) Set(event *eventHandler.Event) *eventHandler.Event {
	return nil
}

func (c *Cassandra) Get(event *eventHandler.Event) *eventHandler.Event {
	return nil
}

func (c *Cassandra) GetName() string {
	return eventHandler.CASSANDRA
}
