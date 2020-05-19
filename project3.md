# Project 3: Blockchain Ledger
**Shiyin Wang, Yixiang Zhang, Rui Lu**  (2020 Spring Tsinghua Operating Systems Group Project)

We start to work on some problems related to transactions on Distributed Systems. In project 3, we will first enable durability across crash failures on a single machine.

In this project, we build a key-value ledger (ledger is a database that stores monetary transactions) that commits data to the disk and can recover after a crash failure by reading from the disk. Specifically, the ledger we will build here will not only store the final states of each account, but all the transactions in the ledger, in a log format. Sometimes, people refer to these transaction logs as a ``block chain’’ (of course, block chains have more security guarantees, as we will see from the next project). In this project, we only use a single server implementation, assuming the disk files are durable (i.e. disks never fail).

We use external packages, so please run the following command if you don't have them installed:
```bash
go get github.com/EagleChen/mapmutex
go get github.com/google/uuid
```
### 1. (20%, 50 lines) RPC calls to support the five basic ledger transactions
The `main.go` has defined the basic framework for RPC calls. Therefore, we simply need to add the I/O commands inside corresponding functions.
```go
func (s *server) Transfer(ctx context.Context, in *pb.TransferRequest) (*pb.BooleanResponse, error) {
  mutex.TryLock(in.FromID)
	mutex.TryLock(in.ToID)
	defer mutex.Unlock(in.FromID)
	defer mutex.Unlock(in.ToID)
	
  // generate Transaction ID
	var userid, err3 = uuid.NewUUID()
	if err3!=nil{
		return &pb.BooleanResponse{Success: false}, err3
	}
  
  // integrity constrains
	if data[in.FromID]>=in.Value{
		loglen++
		data[in.FromID] -= in.Value
		data[in.ToID] += in.Value
    
		// Json I/O
		entry := Dictionary{"Type":"TRANSFER", "FromID":in.FromID, "ToID":in.ToID, "Value":in.Value, "TransactionID":userid}
		file.Transactions = append(file.Transactions, entry)
		data, err1 := json.Marshal(file)
		if err1!=nil{return &pb.BooleanResponse{Success: false}, err1}
		err2 := ioutil.WriteFile(dataDir+strconv.FormatInt(fileidx, 10)+".json", data, 0644)
		if err2!=nil{return &pb.BooleanResponse{Success: false}, err2}
    
    // Start a new block if the number of logs in the current block >= 50
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
```
Let's take `TRANSFER` for example. We assign a lock to each user to constrain that only one thread can edit the balance of a certain user. And we use the `defer` syntax to unlock. We assign a unique Transaction ID to each transaction by package `github.com/google/uuid`. After that, we use a `if` statement to constrain integrity(negative balance is illegal) and write the transaction log into the `$ID.json` blocks. Finally, if the number of logs in the current block reaches the limit, we initialize a new block.


### 2. (20%, 20 lines) Persistence through server crash

### 3. (15%, 20 lines) Integrity constrains
This task involves two actions in the system: Withdraw and Transfer. Whenever we want to reduce the amount of money from a user, we need to check his/her balance first. The implementation is simple.
```go
func (s *server) Withdraw(ctx context.Context, in *pb.Request) (*pb.BooleanResponse, error) {
  ......
	// Proj 3-3 integrity constrains
	if data[in.UserID]>=in.Value{
		loglen++
		data[in.UserID] -= in.Value
		return &pb.BooleanResponse{Success: true}, nil
	}else{
		return &pb.BooleanResponse{Success: false}, errors.New("Transaction "+string(in.Value)+" failed with: insufficient balance")
	}
}
```
```go
func (s *server) Transfer(ctx context.Context, in *pb.TransferRequest) (*pb.BooleanResponse, error) {
  ......
	if data[in.FromID]>=in.Value{
		loglen++
		data[in.FromID] -= in.Value
		data[in.ToID] += in.Value
		return &pb.BooleanResponse{Success: true}, nil
	}else{
		return &pb.BooleanResponse{Success: false}, errors.New("Transaction "+string(in.Value)+" failed with: insufficient balance")
	}
}
```

### 4. (15%, 30 lines) The non-GET transactions need to be saved as JSON blocks

```go
// Json I/O
entry := Dictionary{"Type":"TRANSFER", "FromID":in.FromID, "ToID":in.ToID, "Value":in.Value, "TransactionID":userid}
file.Transactions = append(file.Transactions, entry)
data, err1 := json.Marshal(file)
if err1!=nil{return &pb.BooleanResponse{Success: false}, err1}
err2 := ioutil.WriteFile(dataDir+strconv.FormatInt(fileidx, 10)+".json", data, 0644)
if err2!=nil{return &pb.BooleanResponse{Success: false}, err2}
```

We use the `encoding/json` package to handle the json file I/O with our ledger. The format is defined in the project handbook.

### 5. (10%, 2 lines) Startup script of your program (start.sh), and a compile script to compile your program (compile.sh).

You can simply use the `compile.sh` and `start.sh`, which are provided in the skeleton, to run our blockchain ledger. Since the compiling is contained in `start.sh`, so you don't need to run both of them.

```bash
./compile.sh
./start.sh
```

### 6. (20%, you decide how long it is.) Test cases.

Our test cases are included in the `our_test.sh` script. We repeat all the tasks in the provided skeleton `test_run.sh` file.

```bash
cd example
./our_test.sh
```

We also test the integrity constraints on contradictory transactions. And we rerun our script for several times to test the persistence through server crashes.

```shell
echo "Step 6: Try conflict transaction"
go run ./example/test_client.go -T=TRANSFER -from=TEST--1 -to=TEST---2  -value=5
go run ./example/test_client.go -T=TRANSFER -from=TEST--1 -to=TEST---3  -value=5
go run ./example/test_client.go -T=TRANSFER -from=TEST--1 -to=TEST---4  -value=5
go run ./example/test_client.go -T=DEPOSIT -user=TEST---1 -value=10
echo "Check value: expecting value=5"
go run ./example/test_client.go -T=GET -user=TEST---2
```