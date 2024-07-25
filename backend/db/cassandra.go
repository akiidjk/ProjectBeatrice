package db

import (
	"backend/logger"
	"fmt"
	"time"

	"backend/config" // Add this import statement

	"github.com/gocql/gocql"
)

var Session *gocql.Session

func Init() {
	var err error
	keyspace_name := "ollama"
	for i := 0; i < 5; i++ {
		cluster := gocql.NewCluster(config.CASSANDRA)
		cluster.Keyspace = ""
		Session, err = cluster.CreateSession()
		if err == nil {
			err = ensureKeyspace(Session, keyspace_name)
			if err != nil {
				logger.Error("Failed to create or verify keyspace: " + err.Error())
				Session.Close()
				continue
			}
			cluster.Keyspace = keyspace_name
			Session, err = cluster.CreateSession()
			if err == nil {
				break
			}
		}
		logger.Warning("Failed to connect to Cassandra. Retrying...")
		time.Sleep(5 * time.Second)
	}
	if err != nil {
		panic(err)
	}
	logger.Success("Cassandra well initialized")
}

func GetSession() *gocql.Session {
	return Session
}

func ensureKeyspace(session *gocql.Session, keyspace string) error {
	query := fmt.Sprintf(`SELECT keyspace_name FROM system_schema.keyspaces WHERE keyspace_name = '%s'`, keyspace)
	var ksName string
	if err := session.Query(query).Consistency(gocql.One).Scan(&ksName); err != nil {
		if err == gocql.ErrNotFound {
			createKeyspaceQuery := fmt.Sprintf(`
				CREATE KEYSPACE %s
				WITH replication = {
					'class': 'SimpleStrategy',
					'replication_factor': 1
				}`, keyspace)
			if err := session.Query(createKeyspaceQuery).Exec(); err != nil {
				return fmt.Errorf("failed to create keyspace %s: %w", keyspace, err)
			}
			return nil
		}
		return fmt.Errorf("failed to check keyspace existence: %w", err)
	}
	return nil
}
