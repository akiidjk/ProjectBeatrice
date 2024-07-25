CREATE KEYSPACE IF NOT EXISTS ollama
WITH REPLICATION = {
  'class': 'SimpleStrategy',
  'replication_factor': 1
};