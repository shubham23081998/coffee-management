package dlt

import (
	"backend/modals"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

//  type FabricDlt interface{
// 	GetAllCoffee()(string,error)
// 	CreateCoffee()(id,color,size,owner string)
//  }

var contract *gateway.Contract

func init() {
	// Perform one-time initialization here
	log.Println("Initializing Fabric service...")

	// Connect to the gateway and get the contract
	// You can customize this based on your initialization requirements

	err := os.Setenv("DISCOVERY_AS_LOCALHOST", "true")
	if err != nil {
		log.Fatalf("Error setting DISCOVERY_AS_LOCALHOST environemnt variable: %v", err)
	}

	wallet, err := gateway.NewFileSystemWallet("wallet")
	if err != nil {
		log.Fatalf("Failed to create wallet: %v", err)
	}

	if !wallet.Exists("appUser3") {
		err = populateWallet(wallet)
		if err != nil {
			log.Fatalf("Failed to populate wallet contents: %v", err)
		}
	}

	ccpPath := filepath.Join(
		"..",
		"fabric-samples",
		"test-network",
		"organizations",
		"peerOrganizations",
		"org1.example.com",
		"connection-org1.yaml",
	)
	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(filepath.Clean(ccpPath))),
		gateway.WithIdentity(wallet, "appUser3"),
	)
	if err != nil {
		log.Fatalf("Failed to connect to gateway: %v", err)
	}
	defer gw.Close()

	network, err := gw.GetNetwork("mychannel")
	if err != nil {
		log.Fatalf("Failed to get network: %v", err)
	}

	contract = network.GetContract("basic")

	log.Println("Fabric service initialized.")
}

// FabricService interface and other functions as before...

func GetAllCoffee(ctx context.Context) ([]modals.DltResponse, error) {
	log.Println("--> Evaluate Transaction: GetAllCoffee, function returns all the current assets on the ledger")
	result, err := contract.EvaluateTransaction("GetAllCoffee")
	if err != nil {
		log.Fatalf("Failed to evaluate transaction: %v", err)
	}
	var dltResponse []modals.DltResponse
	err = json.Unmarshal(result, &dltResponse)
	if err != nil {
		return nil, err
	}
	log.Println(string(result))
	return dltResponse, err
}

func GetCoffee(ctx context.Context, getCoffee modals.DltResponse) (modals.DltResponse, error) {
	log.Println("--> Evaluate Transaction: GetAllCoffee, function returns all the current assets on the ledger")
	result, err := contract.EvaluateTransaction("QueryCoffeeBatch", getCoffee.BatchID)
	if err != nil {
		log.Fatalf("Failed to evaluate transaction: %v", err)
	}

	log.Println("Transaction successful, result: ", string(result))
	var dltResponse modals.DltResponse
	err = json.Unmarshal(result, &dltResponse)
	if err != nil {
		return dltResponse, err
	}
	log.Println(string(result))
	return dltResponse, err
}
func CreateCoffeeBatch(ctx context.Context, createCoffee modals.DltResponse) error {
	log.Println("--> Submit Transaction: CreateAsset, creates new asset with ID, color, owner, size, and appraisedValue arguments")
	result, err := contract.SubmitTransaction("CreateCoffeeBatch", createCoffee.BatchID, createCoffee.Variety, createCoffee.Quantity, createCoffee.Origin, createCoffee.Price, createCoffee.ProducedBy)
	if err != nil {
		fmt.Errorf("Failed to Submit transaction: %v", err)
		return err
	}
	log.Println(string(result))
	return nil
}

func UpdateCoffeeBatch(ctx context.Context, updateCoffee modals.DltResponse) error {
	log.Println("--> Submit Transaction: CreateAsset, creates new asset with ID, color, owner, size, and appraisedValue arguments")
	// log.Println(updateCoffee)
	result, err := contract.SubmitTransaction("UpdateBatchStatus", updateCoffee.BatchID, updateCoffee.Status, updateCoffee.Price)
	if err != nil {
		fmt.Errorf("Failed to update transaction: %v", err)
		return err
	}
	log.Println(string(result))
	return nil
}

func TransferCoffeeOwnerShip(ctx context.Context, transferCoffee modals.DltResponse) error {
	log.Println("--> Submit Transaction: CreateAsset, creates new asset with ID, color, owner, size, and appraisedValue arguments")
	result, err := contract.SubmitTransaction("TransferOwnership", transferCoffee.BatchID, transferCoffee.CurentOwner, transferCoffee.NewOwner)
	if err != nil {
		fmt.Errorf("Failed to Submit transaction: %v", err)
		return err
	}
	log.Println(string(result))
	return nil
}

// func main() {
// 	log.Println("============ application-golang starts ============")

// 	fmt.Println("contract ==>", contract)
// 	log.Println("--> Submit Transaction: InitLedger, function creates the initial set of assets on the ledger")
// 	result, err := contract.SubmitTransaction("InitLedger")
// 	fmt.Println("err==>", err)
// 	fmt.Println()
// 	if err != nil {
// 		log.Fatalf("Failed to Submit transaction: %v", err)
// 	}
// 	log.Println(string(result))

// 	log.Println("--> Evaluate Transaction: GetAllCoffee, function returns all the current assets on the ledger")
// 	result, err = contract.EvaluateTransaction("GetAllCoffee")
// 	if err != nil {
// 		log.Fatalf("Failed to evaluate transaction: %v", err)
// 	}
// 	log.Println(string(result))

// 	log.Println("--> Submit Transaction: CreateAsset, creates new asset with ID, color, owner, size, and appraisedValue arguments")
// 	result, err = contract.SubmitTransaction("CreateCoffeeBatch","3","123","american","100","Euthopia","120","farmer")
// 	if err != nil {
// 		log.Fatalf("Failed to Submit transaction: %v", err)
// 	}
// 	log.Println(string(result))

// 	// log.Println("--> Evaluate Transaction: ReadAsset, function returns an asset with a given assetID")
// 	// result, err = contract.EvaluateTransaction("ReadAsset", "asset13")
// 	// if err != nil {
// 	// 	log.Fatalf("Failed to evaluate transaction: %v\n", err)
// 	// }
// 	// log.Println(string(result))

// 	// log.Println("--> Evaluate Transaction: AssetExists, function returns 'true' if an asset with given assetID exist")
// 	// result, err = contract.EvaluateTransaction("AssetExists", "asset1")
// 	// if err != nil {
// 	// 	log.Fatalf("Failed to evaluate transaction: %v\n", err)
// 	// }
// 	// log.Println(string(result))

// 	// log.Println("--> Submit Transaction: TransferAsset asset1, transfer to new owner of Tom")
// 	// _, err = contract.SubmitTransaction("TransferAsset", "asset1", "Tom")
// 	// if err != nil {
// 	// 	log.Fatalf("Failed to Submit transaction: %v", err)
// 	// }

// 	// log.Println("--> Evaluate Transaction: ReadAsset, function returns 'asset1' attributes")
// 	// result, err = contract.EvaluateTransaction("ReadAsset", "asset1")
// 	// if err != nil {
// 	// 	log.Fatalf("Failed to evaluate transaction: %v", err)
// 	// }
// 	// log.Println(string(result))
// 	// log.Println("============ application-golang ends ============")
// }

func populateWallet(wallet *gateway.Wallet) error {
	log.Println("============ Populating wallet ============")
	credPath := filepath.Join(
		"..",
		"fabric-samples",
		"test-network",
		"organizations",
		"peerOrganizations",
		"org1.example.com",
		"users",
		"User1@org1.example.com",
		"msp",
	)

	certPath := filepath.Join(credPath, "signcerts", "User1@org1.example.com-cert.pem")
	// read the certificate pem
	cert, err := ioutil.ReadFile(filepath.Clean(certPath))
	if err != nil {
		return err
	}

	keyDir := filepath.Join(credPath, "keystore")
	// there's a single file in this dir containing the private key
	files, err := ioutil.ReadDir(keyDir)
	if err != nil {
		return err
	}
	if len(files) != 1 {
		return fmt.Errorf("keystore folder should have contain one file")
	}
	keyPath := filepath.Join(keyDir, files[0].Name())
	key, err := ioutil.ReadFile(filepath.Clean(keyPath))
	if err != nil {
		return err
	}

	identity := gateway.NewX509Identity("Org1MSP", string(cert), string(key))

	return wallet.Put("appUser3", identity)
}
