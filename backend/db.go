package anyNoteX

import (
	"log"
	"os"

	"github.com/Xonline-Tech/AnyNoteX/backend/platform"
	"github.com/nutsdb/nutsdb"
)

var db *nutsdb.DB

// DataStructure represents the data structure we have already supported
type DataStructure = uint16

const (
	// DataStructureSet represents the data structure set flag
	DataStructureSet DataStructure = 0

	// DataStructureSortedSet represents the data structure sorted set flag
	DataStructureSortedSet DataStructure = 1

	// DataStructureBTree represents the data structure b tree flag
	DataStructureBTree DataStructure = 2

	// DataStructureList represents the data structure list flag
	DataStructureList DataStructure = 3
)

// InitializeDb initializes db
func InitializeDb() {
	var err error
	opt := nutsdb.DefaultOptions

	// define database parameters
	opt.Dir = platform.GetAppDir() + string(os.PathSeparator) + "db"

	// open database
	if db, err = nutsdb.Open(opt); err != nil {
		log.Panic(err)
	}
}

// CloseDb .
func CloseDb() {
	db.Close()
}

// NewBucketIfNotExists 当Bucket不存在时创建桶
func NewBucketIfNotExists(bucketName string) {
	db.Update(func(tx *nutsdb.Tx) error {

		var err error

		// check whether the bucket exists
		if isExist := tx.ExistBucket(DataStructureSet, bucketName); isExist == true {
			log.Println("bucket \"", bucketName, "\" already exists, skip creating...")
			return err
		}

		if err := tx.NewKVBucket(bucketName); err != nil {
			log.Println(err)
		}
		return err
	})
}
