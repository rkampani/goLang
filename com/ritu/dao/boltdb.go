package dao

import (
	"context"
	"fmt"
	"strconv"

	"github.com/rituK/com/ritu/utils"

	"encoding/json"

	"github.com/boltdb/bolt"
	"github.com/rituK/com/ritu/modal"
)

type IBoltDB interface {
	OpenBoltDB()
	Seed()
	QueryUsers(ctx context.Context) ([]modal.Person, error)
	QueryUser(ctx context.Context, userId string) (modal.Person, error)
	SaveOrUpdateUser(ctx context.Context, jsonUserBytes []byte, key string) error
}

// BoltClient is the real implementation
type BoltClient struct {
	boltDB *bolt.DB
}

func (bc *BoltClient) OpenBoltDB() {

	var _error error
	bc.boltDB, _error = bolt.Open("Person.db", 0600, nil)
	if _error != nil {
		panic(_error)
	}
	utils.LogTracingPrint("Connection Open Succefully")
}

// Seed can be used to seed some sample accounts
func (bc *BoltClient) Seed() {

	utils.LogTracingPrint("seeds --- starts")
	bc.initializerBucket()
	bc.setUpUsers()

	utils.LogTracingPrint("seed --- ends")

}

// Creates an "PersonBucket" in our BoltDB. It will overwrite any existing bucket of the same name.
func (bc *BoltClient) initializerBucket() {
	bc.boltDB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("PersonBucket"))
		if err != nil {
			return fmt.Errorf("create bucket failed: %s", err)
		}
		return nil
	})

}

func (bc *BoltClient) setUpUsers() {
	utils.LogTracingPrint("user setUp --- starts")
	rangecount := 105
	for i := 0; i < rangecount; i++ {

		// Create an instance of our Person & Address struct
		key := strconv.Itoa(i)
		temp2 := modal.Address{

			AddressLine1: "500 West Madison",
			City:         "Chicago",
			State:        "IL",
			ZipCode:      "60661",
		}
		temp1 := modal.Person{
			Fname:   "Adam" + key,
			Lname:   "Smith" + key,
			ID:      i,
			Address: &temp2,
		}

		//Serialize the struct and save in DB

		jsonBytes, _ := json.Marshal(temp1)

		bc.save(key, jsonBytes)
	}

	utils.LogTracingPrint("user Setup --- Ends")
}

func (bc *BoltClient) save(key string, jsonBytes []byte) {
	bc.boltDB.Update(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte("PersonBucket"))
		err := b.Put([]byte(key), jsonBytes)
		return err
	})
}

func (bc *BoltClient) QueryUsers(ctx context.Context) ([]modal.Person, error) {

	utils.LogTracingPrint("Query users --- starts")

	var tmp []modal.Person

	person := modal.Person{}
	// Read an object from the bucket using boltDB.View
	err := bc.boltDB.View(func(tx *bolt.Tx) error {
		// Read the bucket from the DB
		b := tx.Bucket([]byte("PersonBucket"))

		// Read the value identified by our accountId supplied as []byte
		b.ForEach(func(k, v []byte) error {
			fmt.Printf("key=%s, value=%s\n", k, v)
			personBytes := v
			// Unmarshal the returned bytes into the account struct we created at
			// the top of the function
			json.Unmarshal(personBytes, &person)
			tmp = append(tmp, person)
			return nil

		})
		return nil
	})

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	utils.LogTracingPrint("Query users --- ends")
	// Return the Account struct and nil as error.
	return tmp, nil

}
func (bc *BoltClient) QueryUser(ctx context.Context, userId string) (modal.Person, error) {

	utils.LogTracingPrint("GetUser --starts" + userId)
	person := modal.Person{}

	// Read an object from the bucket using boltDB.View
	err := bc.boltDB.View(func(tx *bolt.Tx) error {
		// Read the bucket from the DB
		b := tx.Bucket([]byte("PersonBucket"))

		// Read the value identified by our accountId supplied as []byte
		personBytes := b.Get([]byte(userId))
		if personBytes == nil {
			return fmt.Errorf("No account found for " + userId)
		}
		// Unmarshal the returned bytes into the account struct we created at
		// the top of the function
		json.Unmarshal(personBytes, &person)

		// Return nil to indicate nothing went wrong, e.g no error
		return nil
	})
	if err != nil {
		panic(err)
	}
	utils.LogTracingPrint("GetUser --- ends")
	return person, nil
}

func (bc *BoltClient) SaveOrUpdateUser(ctx context.Context, jsonUserBytes []byte, key string) error {

	bc.boltDB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("PersonBucket"))
		err := b.Put([]byte(key), jsonUserBytes)
		return err
	})

	return nil
}
