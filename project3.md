# Project 3: Blockchain Ledger
**Shiyin Wang, Yixiang Zhang, Rui Lu**  (2020 Spring Tsinghua Operating Systems Group Project)

[GitHub](https://github.com/shiyinw/blockchain-ledger-from-scratch/) https://github.com/shiyinw/blockchain-ledger-from-scratch/

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
func (s *server) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetResponse, error) {
	return &pb.GetResponse{Value: data[in.UserID]}, nil
}
```

```go
func (s *server) Put(ctx context.Context, in *pb.Request) (*pb.BooleanResponse, error) {
	loglen++
	data[in.UserID] = in.Value
	// Json I/O
	entry := Dictionary{"Type":"PUT", "UserID":in.UserID, "Value":in.Value}
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
        loglen = 0
		fileidx++
		file.BlockID = fileidx
		file.Transactions = []Dictionary{}
	}
	return &pb.BooleanResponse{Success: true}, nil
}
```

```go
func (s *server) Deposit(ctx context.Context, in *pb.Request) (*pb.BooleanResponse, error) {
	loglen++
	data[in.UserID] += in.Value
	// Json I/O
	entry := Dictionary{"Type": "DEPOSIT", "UserID": in.UserID, "Value": in.Value}
	file.Transactions = append(file.Transactions, entry)
	data, err1 := json.Marshal(file)
	if err1 != nil {
		return &pb.BooleanResponse{Success: false}, err1
	}
	err2 := ioutil.WriteFile(dataDir+strconv.FormatInt(fileidx, 10)+".json", data, 0644)
	if err2 != nil {
		return &pb.BooleanResponse{Success: false}, err2
	}
	if loglen % blockSize == 0 {
        loglen = 0
		fileidx++
		file.BlockID = fileidx
		file.Transactions = []Dictionary{}
	}
	return &pb.BooleanResponse{Success: true}, nil
}
```

```go
func (s *server) Withdraw(ctx context.Context, in *pb.Request) (*pb.BooleanResponse, error) {
	got := false
	for true {
		if got = mutex.TryLock(in.UserID); got {
			break
		}
	}
	if got {
		defer mutex.Unlock(in.UserID)
		// Integrity constrain
		if data[in.UserID] >= in.Value {
			loglen++
			data[in.UserID] -= in.Value
			// Json I/O
			entry := Dictionary{"Type": "WITHDRAW", "UserID": in.UserID, "Value": in.Value}
			file.Transactions = append(file.Transactions, entry)
			data, err1 := json.Marshal(file)
			if err1 != nil {
				return &pb.BooleanResponse{Success: false}, err1
			}
			err2 := ioutil.WriteFile(dataDir+strconv.FormatInt(fileidx, 10)+".json", data, 0644)
			if err2 != nil {
				return &pb.BooleanResponse{Success: false}, err2
			}
			if loglen%blockSize == 0 {
                loglen = 0
				fileidx++
				file.BlockID = fileidx
				file.Transactions = []Dictionary{}
			}
			return &pb.BooleanResponse{Success: true}, nil
		} else {
			return &pb.BooleanResponse{Success: false}, nil
		}
	}
	return &pb.BooleanResponse{Success: false}, nil
}
```

```go
func (s *server) Transfer(ctx context.Context, in *pb.TransferRequest) (*pb.BooleanResponse, error) {
	got := false
	for true {
		if got = mutex.TryLock(in.FromID); got {
			break
		}
	}
	if got {
		defer mutex.Unlock(in.FromID)
		// Integrity constrain
		var userid, err3 = uuid.NewUUID()
		if err3 != nil {
			return &pb.BooleanResponse{Success: false}, err3
		}
		if data[in.FromID] >= in.Value {
			loglen++
			data[in.FromID] -= in.Value
			data[in.ToID] += in.Value
			// Json I/O
			entry := Dictionary{"Type": "TRANSFER", "FromID": in.FromID, "ToID": in.ToID, "Value": in.Value}
			file.Transactions = append(file.Transactions, entry)
			data, err1 := json.Marshal(file)
			if err1 != nil {
				return &pb.BooleanResponse{Success: false}, err1
			}
			err2 := ioutil.WriteFile(dataDir+strconv.FormatInt(fileidx, 10)+".json", data, 0644)
			if err2 != nil {
				return &pb.BooleanResponse{Success: false}, err2
			}
			if loglen%blockSize == 0 {
                loglen = 0
				fileidx++
				file.BlockID = fileidx
				file.Transactions = []Dictionary{}
			}
			return &pb.BooleanResponse{Success: true}, nil
		} else {
			log.Printf("Transaction %d failed with: %s", in.Value, userid)
			return &pb.BooleanResponse{Success: false}, nil
		}
	}
	return &pb.BooleanResponse{Success: false}, nil
}
```
Let's take `TRANSFER` for example. We assign a lock to each user to constrain that only one thread can decrease the balance of a certain user. And we use the `defer` syntax to unlock. We assign a unique Transaction ID to each transaction by package `github.com/google/uuid`. After that, we use a `if` statement to constrain integrity(negative balance is illegal) and write the transaction log into the `$ID.json` blocks. Finally, if the number of logs in the current block reaches the limit, we initialize a new block.

### 2. (20%, 20 lines) Persistence through server crash

We insert the recovery script in `main.go` after reading the config. The `dataDir` in `config.json` define the directory to store blocks. We first check whether this directory exists.

```go
// Unused variable
dataDir = outputDir
if _, err := os.Stat(dataDir); os.IsNotExist(err) {
  os.Mkdir(outputDir, 0777)
}
```

Then we recursively load the content of each block to process the transactions. Since the professor and TAs have confirmed in the course group chat that we assume the blocks are all correct, we don't need to bother the constrains here.

```go
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
```

After loading the transactions, we check whether number of logs in the current block reaches the limit, which is the same as the previous task.

```go
if loglen > 0 && loglen % blockSize == 0{
  loglen = 0
  fileidx++
  file.BlockID = fileidx
  file.Transactions = []Dictionary{}
}
log.Print(loglen)
```

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

We also need to avoid contradictory transactions. We have observed that the problem happens on the senders' side. Therefore, we assign a lock to each user to constrain that **only one thread can decrease the balance of a certain user**. And we use the `defer` syntax to unlock. Here `mutex` is an instance of `github.com/EagleChen/mapmutex`.

```go
func (s *server) Withdraw(ctx context.Context, in *pb.Request) (*pb.BooleanResponse, error) {
	got := false
	for true {
		if got = mutex.TryLock(in.UserID); got {
			break
		}
	}
	if got {
    // code here
    return &pb.BooleanResponse{Success: true}, nil
	}
  return &pb.BooleanResponse{Success: false}, nil
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

If you want to make a clean start(ignore all the previous blocks), then you need to delete all the blocks in the `dataDir`:

```bash
cd tmp
rm -rf *
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