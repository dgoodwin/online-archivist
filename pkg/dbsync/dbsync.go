package dbsync

import (
	"database/sql"
	"fmt"

	"github.com/openshift/online-archivist/pkg/config"

	"github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
)

type DBSyncer struct {
	cfg    config.ArchivistConfig
	logger *logrus.Entry
	db     *sql.DB
}

func NewDBSyncer(cfg config.ArchivistConfig) *DBSyncer {
	logger := logrus.WithField("component", "dbsync")
	return &DBSyncer{cfg: cfg, logger: logger}
}

func (dbs *DBSyncer) Run() {
	dbs.logger.Infof("connecting to MySQL database: %s@%s",
		dbs.cfg.Database.Username, dbs.cfg.Database.Address)
	_, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)", dbs.cfg.Database.Username,
		dbs.cfg.Database.Password, dbs.cfg.Database.Address))
	if err != nil {
		dbs.logger.Fatal("error connecting to database:", err)
	}
	dbs.logger.Info("db connection successful")
	// TODO: connect to correct database once we're sure it exists
}
