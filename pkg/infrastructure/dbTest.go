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

var (
	testdb   *gorm.DB
	resource *dockertest.Resource
	pool     *dockertest.Pool
)

func CreateContainer() (*dockertest.Resource, *dockertest.Pool) {
	// Dockerとの接続
	pool, err := dockertest.NewPool("")
	pool.MaxWait = time.Minute * 1
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

func CloseContainer(resource *dockertest.Resource, pool *dockertest.Pool) {
	// コンテナの終了
	fmt.Printf("close container")
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
}

func ConnectDB(pool *dockertest.Pool) *gorm.DB {
	if err := pool.Retry(func() error {
		fmt.Println("waiting to db start up....")
		time.Sleep(time.Second * 3)

		var err error
		config := fmt.Sprintf(dsn, user, password, port, dbName)
		testdb, err = gorm.Open("postgres", config)
		if err != nil {
			return err
		}
		return testdb.DB().Ping()
	}); err != nil {
		log.Fatalf("接続エラー: %s", err)
	}
	autoMigrate(testdb)
	execSeeds(testdb)
	fmt.Println("[INFO]TestDB setup done!")
	print("testdb: %v", testdb)
	return testdb
}

func GetTestDB() *gorm.DB {
	return testdb
}
