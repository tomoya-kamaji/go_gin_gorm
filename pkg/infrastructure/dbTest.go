package infrastructure

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
)

var (
	user     = "postgres"
	password = "secret"
	dbName   = "unittest"
	port     = "5433"
	dialect  = "postgres"
	dsn      = "postgres://%s:%s@localhost:%s/%s?sslmode=disable"
)

func createContainer() (*dockertest.Resource, *dockertest.Pool) {
	// Dockerとの接続
	pool, err := dockertest.NewPool("")
	pool.MaxWait = time.Minute * 2
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// Dockerコンテナ起動時の細かいオプションを指定する
	runOptions := &dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "12.3",
		Env: []string{
			"POSTGRES_USER=" + user,
			"POSTGRES_PASSWORD=" + password,
			"POSTGRES_DB=" + dbName,
		},
		ExposedPorts: []string{"5432"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5432": {
				{HostIP: "0.0.0.0", HostPort: port},
			},
		},
	}

	// コンテナを起動
	resource, err := pool.RunWithOptions(runOptions)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	return resource, pool
}

func closeContainer(resource *dockertest.Resource, pool *dockertest.Pool) {
	// コンテナの終了
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
}

func connectDB(pool *dockertest.Pool) *gorm.DB {
	var db *gorm.DB
	if err := pool.Retry(func() error {
		fmt.Println("waiting to db start up....")
		time.Sleep(time.Second * 3)

		var err error
		config := fmt.Sprintf(dsn, user, password, port, dbName)
		db, err = gorm.Open("postgres", config)
		if err != nil {
			return err
		}
		return db.DB().Ping()
	}); err != nil {
		log.Fatalf("接続エラー: %s", err)
	}
	autoMigrate(db)
	execSeeds(db)
	return db
}

func InitTest() *gorm.DB {
	resource, pool := createContainer()
	defer closeContainer(resource, pool)
	db := connectDB(pool)
	defer db.Close()
	return db
}

func GetTestDB() *gorm.DB {
	return db
}
