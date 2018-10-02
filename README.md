# blockchain-with-go
referring to [it](https://hackernoon.com/learn-blockchains-by-building-one-117428612f46)

## usage

### dependency
```
make deps 
```

### run 
```
make run
```
or 

```
go run main.go -p <port>
```

## endpoint

### **GET**```/mine```  : current transaction to block (mining with proof of work by node )
### **POST** ```/transactions/new``` : make transaction 
example

```
{
	"sender": "6fdf241f5f8f47af80936b247898afb2",
	"recipient": "03418427b36a4ad38de40b9e16af7509",
	"amount": 5
}
```
### **GET**```/chain```  : get fullchain   
### **GET**```/amount?node=<node>```  : get amount of node  
### **GET**```/nodes```  : get all nodes that is registered  
### **POST**```/nodes/register``` : register nodes
```
{
	"nodes": ["http://localhost:5001"]
}
```   
### **GET**```/nodes/resolve``` : audit all registered node  
If there is a node holding a longer chain than the chain held by the registered node, update the chain




