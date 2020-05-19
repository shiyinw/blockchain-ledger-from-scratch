package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"

	pb "github.com/shiyinw/blockchain-ledger-from-scratch/protobuf/go"

	"github.com/EagleChen/mapmutex"
	"github.com/google/uuid"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const blockSize = 50
type Dictionary map[string]interface{}

var data = make(map[string]int32)
var loglen int32
var fileidx int64 = 1 // store the ledger in [fileidx].json
var dataDir string // store the blocks

// no two contradictory transactions
var mutex = mapmutex.NewMapMutex()

// json I/O
type fileio struct{
	BlockID int64
	PrevHash string `default:"00000000"`
	Transactions []Dictionary
	Nonce string `default:"00000000"`
}
var file fileio = fileio{fileidx, "00000000", []Dictionary{}, "00000000"}

type server struct{}
// Database Interface 
func (s *server) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetResponse, error) {
	return &pb.GetResponse{Value: data[in.UserID]}, nil
}
func (s *server) Put(ctx context.Context, in *pb.Request) (*pb.BooleanResponse, error) {
	mutex.TryLock(in.UserID)
	defer mutex.Unlock(in.UserID)
	loglen++
	data[in.UserID] = in.Value
	// Json I/O
	var userid, err3 = uuid.NewUUID()
	if err3!=nil{
		return &pb.BooleanResponse{Success: false}, err3
	}
	entry := Dictionary{"Type":"PUT", "UserID":in.UserID, "Value":in.Value, "TransactionID":userid}
	file.Transactions = append(file.Transactions, entry)
	data, err1 := json.Marshal(file)
	if err1!=nil{
		return &pb.BooleanResponse{Success: false}, err1
	}
	err2 := ioutil.WriteFile(dataDir+strconv.FormatInt(fileidx, 10)+".json", data, 0644)
	if err2!=nil{
		return &pb.BooleanResponse{Success: false}, err2
	}
	if loglen % blockSize == 0{
		fileidx++
		file.BlockID = fileidx
		file.Transactions = []Dictionary{}
	}
	return &pb.BooleanResponse{Success: true}, nil
}
func (s *server) Deposit(ctx context.Context, in *pb.Request) (*pb.BooleanResponse, error) {
	mutex.TryLock(in.UserID)
	defer mutex.Unlock(in.UserID)
	loglen++
	data[in.UserID] += in.Value
	// Json I/O
	var userid, err3 = uuid.NewUUID()
	if err3!=nil{
		return &pb.BooleanResponse{Success: false}, err3
	}
	entry := Dictionary{"Type":"DEPOSIT", "UserID":in.UserID, "Value":in.Value, "TransactionID":userid}
	file.Transactions = append(file.Transactions, entry)
	data, err1 := json.Marshal(file)
	if err1!=nil{
		return &pb.BooleanResponse{Success: false}, err1
	}
	err2 := ioutil.WriteFile(dataDir+strconv.FormatInt(fileidx, 10)+".json", data, 0644)
	if err2!=nil{
		return &pb.BooleanResponse{Success: false}, err2
	}
	if loglen % blockSize == 0{
		fileidx++
		file.BlockID = fileidx
		file.Transactions = []Dictionary{}
	}
	return &pb.BooleanResponse{Success: true}, nil
}
func (s *server) Withdraw(ctx context.Context, in *pb.Request) (*pb.BooleanResponse, error) {
	mutex.TryLock(in.UserID)
	defer mutex.Unlock(in.UserID)
	// Proj 3-3 integrity constrains
	var userid, err3 = uuid.NewUUID()
	if err3!=nil{
		return &pb.BooleanResponse{Success: false}, err3
	}
	if data[in.UserID]>=in.Value{
		loglen++
		data[in.UserID] -= in.Value
		// Json I/O
		entry := Dictionary{"Type":"WITHDRAW", "UserID":in.UserID, "Value":in.Value, "TransactionID":userid}
		file.Transactions = append(file.Transactions, entry)
		data, err1 := json.Marshal(file)
		if err1!=nil{
			return &pb.BooleanResponse{Success: false}, err1
		}
		err2 := ioutil.WriteFile(dataDir+strconv.FormatInt(fileidx, 10)+".json", data, 0644)
		if err2!=nil{
			return &pb.BooleanResponse{Success: false}, err2
		}
		if loglen % blockSize == 0{
			fileidx++
			file.BlockID = fileidx
			file.Transactions = []Dictionary{}
		}
		return &pb.BooleanResponse{Success: true}, nil
	}else{
		return &pb.BooleanResponse{Success: false}, errors.New("Transaction ["+userid.String()+"] "+strconv.FormatInt(int64(in.Value), 10)+" failed with: insufficient balance")
	}
}
func (s *server) Transfer(ctx context.Context, in *pb.TransferRequest) (*pb.BooleanResponse, error) {
	mutex.TryLock(in.FromID)
	mutex.TryLock(in.ToID)
	defer mutex.Unlock(in.FromID)
	defer mutex.Unlock(in.ToID)
	// Proj 3-3 integrity constrains
	var userid, err3 = uuid.NewUUID()
	if err3!=nil{
		return &pb.BooleanResponse{Success: false}, err3
	}
	if data[in.FromID]>=in.Value{
		loglen++
		data[in.FromID] -= in.Value
		data[in.ToID] += in.Value
		// Json I/O
		entry := Dictionary{"Type":"TRANSFER", "FromID":in.FromID, "ToID":in.ToID, "Value":in.Value, "TransactionID":userid}
		file.Transactions = append(file.Transactions, entry)
		data, err1 := json.Marshal(file)
		if err1!=nil{
			return &pb.BooleanResponse{Success: false}, err1
		}
		err2 := ioutil.WriteFile(dataDir+strconv.FormatInt(fileidx, 10)+".json", data, 0644)
		if err2!=nil{
			return &pb.BooleanResponse{Success: false}, err2
		}
		if loglen % blockSize == 0{
			fileidx++
			file.BlockID = fileidx
			file.Transactions = []Dictionary{}
		}
		return &pb.BooleanResponse{Success: true}, nil
	}else{
		return &pb.BooleanResponse{Success: false}, errors.New("Transaction ["+userid.String()+"] "+strconv.FormatInt(int64(in.Value), 10)+" failed with: insufficient balance")
	}
}
// Interface with test grader
func (s *server) LogLength(ctx context.Context, in *pb.Null) (*pb.GetResponse, error) {
	return &pb.GetResponse{Value: loglen}, nil
}

// Main function, RPC server initialization
func main() {
	// Read config
	address, outputDir := func() (string, string) {
		conf, err := ioutil.ReadFile("config.json")
		if err != nil {
			panic(err)
		}
		var dat map[string]interface{}
		err = json.Unmarshal(conf, &dat)
		if err != nil {
			panic(err)
		}
		dat = dat["1"].(map[string]interface{}) // should be dat[myNum] in the future
		return fmt.Sprintf("%s:%s", dat["ip"], dat["port"]), fmt.Sprintf("%s", dat["dataDir"])
	}()
	// Unused variable
	dataDir = outputDir
	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		os.Mkdir(outputDir, 0777)
	}
	// Recover
	log.Print("Retrieving data from blocks......")
	files, err := ioutil.ReadDir(dataDir)
	if err != nil {
		log.Fatal(err)
	}
	for i := 1; i <= len(files); i++ {
		fileidx = int64(i)
		var cached_string, _ = ioutil.ReadFile(dataDir + strconv.Itoa(i) + ".json")
		json.Unmarshal(cached_string, &file)
		for _, tran := range file.Transactions {
			loglen++
			switch tran["Type"] {
			default:
				log.Fatal("Unknown operation.")
			case "PUT":
				data[tran["UserID"].(string)] = int32(tran["Value"].(float64))
			case "DEPOSIT":
				data[tran["UserID"].(string)] += int32(tran["Value"].(float64))
			case "WITHDRAW":
				data[tran["UserID"].(string)] -= int32(tran["Value"].(float64))
			case "TRANSFER":
				data[tran["FromID"].(string)] -= int32(tran["Value"].(float64))
				data[tran["ToID"].(string)] += int32(tran["Value"].(float64))
			}
		}
	}
	if loglen % blockSize == 0{
		fileidx++
		file.BlockID = fileidx
		file.Transactions = []Dictionary{}
	}
	log.Print(loglen)

	// Bind to port
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Listening: %s ...", address)

	// Create gRPC server
	s := grpc.NewServer()
	pb.RegisterBlockDatabaseServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)

	// Start server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}