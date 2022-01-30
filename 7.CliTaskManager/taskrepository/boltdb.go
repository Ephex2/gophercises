package taskrepository

import (
	"time"

	bolt "go.etcd.io/bbolt"
)

var path = "tasks.db"
var rootBucketName = []byte("TaskList")
var boltRepo *boltRepository
var loadError error

type boltRepository struct {
	db *bolt.DB
	//rootBucket *bolt.Bucket
}

func newBoltRepository() (*boltRepository, error) {
	repo := &boltRepository{}
	err := repo.Open()
	if err != nil {
		return nil, err
	}
	defer repo.db.Close()

	err = repo.db.Update(func(tx *bolt.Tx) (err error) {
		// Root bucket for this project is TaskList, creating it in init function.
		_, err = tx.CreateBucketIfNotExists(rootBucketName)
		if err != nil {
			return err
		}

		return err
	})

	return repo, err
}

func (br *boltRepository) Open() (err error) {
	br.db, err = bolt.Open(path, 0600, &bolt.Options{Timeout: 1 * time.Second})
	return err
}

// This is a command line application, no need to keep DB open
// Opening and defer closing db in each function. Seems inelegant but works. Do not expect to ever deal with concurrency issues.
func init() {
	boltRepo, loadError = newBoltRepository()
	if loadError != nil {
		panic(loadError)
	}
}

func (br *boltRepository) readBolt() (tasks []string, err error) {
	br.Open()
	defer br.db.Close()

	err = br.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(rootBucketName)
		return b.ForEach(func(k, v []byte) error {
			// Get each task from store. Currently k == v for each entry.
			tasks = append(tasks, string(v))
			return nil
		})
	})

	return tasks, err
}

func (br *boltRepository) updateBolt(taskName string) error {
	br.Open()
	defer br.db.Close()

	return br.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(rootBucketName)

		// Persist task to bucket. May consider using sequential IDs as keys, but this makes tasks harder to retrieve (user must know id?).
		// Will error out if task name is too long.
		return b.Put([]byte(taskName), []byte(taskName))
	})
}

func (br *boltRepository) deleteBolt(taskName string) error {
	br.Open()
	defer br.db.Close()

	return br.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(rootBucketName)
		return b.Delete([]byte(taskName))
	})
}

func (br *boltRepository) clearBolt() error {
	br.Open()
	defer br.db.Close()

	return br.db.Update(func(tx *bolt.Tx) (innerErr error) {
		if innerErr = tx.DeleteBucket(rootBucketName); innerErr != nil {
			return innerErr
		}

		// In current version, no task calls multiple repository functions, but this may change and this next line will be important.
		// Currently, bucket is only created when this package is initialized.
		_, innerErr = tx.CreateBucket(rootBucketName)

		// Whether innerErr is nil or no, return
		return
	})
}
