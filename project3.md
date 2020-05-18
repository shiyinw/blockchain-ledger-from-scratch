# Project 3: Blockchain Ledger
**Shiyin Wang, Yixiang Zhang, Rui Lu**  (2020 Spring Tsinghua Operating Systems Group Project)

We start to work on some problems related to transactions on Distributed Systems. In project 3, we will first enable durability across crash failures on a single machine.

In this project, we build a key-value ledger (ledger is a database that stores monetary transactions) that commits data to the disk and can recover after a crash failure by reading from the disk. Specifically, the ledger we will build here will not only store the final states of each account, but all the transactions in the ledger, in a log format. Sometimes, people refer to these transaction logs as a ``block chain’’ (of course, block chains have more security guarantees, as we will see from the next project). In this project, we only use a single server implementation, assuming the disk files are durable (i.e. disks never fail).

### 1. (20%, 50 lines) RPC calls to support the five basic ledger transactions


### 2. (20%, 20 lines) Persistence through server crash

### 3. (15%, 20 lines) Integrity constrains
This task involves two actions in the system: Withdraw and Transfer. Whenever we want to reduce the amount of money from a user, we need to check his/her balance first. The implementation is simple.
```go
func (s *server) Withdraw(ctx context.Context, in *pb.Request) (*pb.BooleanResponse, error) {
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