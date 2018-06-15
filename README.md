# RPC Compare
A quick project to test drive the various available RPC libraries:

# Table Of Contents
- [Overview](#overview)
- [Conclusion](#conclusion)

# Overview
The 2 libraries I decided to test were:

- [GRPC](https://grpc.io)
- [Cap'N Proto](https://capnproto.org/)

With both of these libraries I attempted to make a simple `Calculator` service 
with an `Add` method. This method took a list of integers and returned the 
summed result.

# Conclusion
After testing both libraries I decided that GRPC is the best fit for me.  

## GRPC
GRPC allows for a much simpler flow.  

You can use built in language types and parameters are passed directly to 
functions. Additionally much less boilerplate code is needed to get it up and 
running.  

I was able to learn how to use GRPC and finish a quick "hello world" in 10-15 
minutes.  

Implementing a Service method:  

```go
type CalcServer struct{}

// RPC parameters are passed directly to functions
// Results are returned by functions
func (c *CalcServer) Add(ctx context.Context, numbers *calc.Numbers) (
	*calc.Result, error) {

	// ...

	return result, nil;
}
```

Calling a Service method:  

```go
client := calc.NewCalculatorClient(conn)

numbers := &calc.Numbers{
	Numbers: []int32{1, 2, 3, 4, 5},
}

// Calling a RPC service method is very similar to calling a normal method
res, err := client.Add(ctx, numbers)
```

## Cap'N Proto
I found Cap'N Proto to be a much more complex and heavy library.  

Much boilerplate code was needed to build a service, implement methods, and 
call methods via a client. Native types could not be used, and I had to dig 
through the documentation to figure out how to use the special Cap'N Proto 
Go types.  

In fact I ended up not finishing my Cap'N Proto demo after running into a bug 
with the special Go Cap'N Proto Int32List type. It became clear that too much 
extra code was needed to use Cap'N Proto.  

Implementing a Service method:  

```go
// RPC call parameters are wrapped in a special call structure
func (c CalcServer) Add(call calc.Calculator_add) error {
	// Extra code is needed to extract RCP call parameters
	numbers, err := call.Params.Numbers()
	if err != nil {
		return fmt.Errorf("error reading data: %s", err.Error())
	}

	// ...

	// Results can not be returned like a normal function, they must be 
	// set via a special API
	call.Results.SetResult(accum)
}
```

Calling a Service method:  

```go
// Calling an RPC Service method is slightly unintuative and verbose
res, err := calcClient.Add(ctx, func(params calc.Calculator_add_Params) error {
	// ...

	// RPC service method parameters must be set via a custom API
	return params.SetNumbers(numbers)
}).Struct()
```

## Wrap Up
Overall I believe GRPC is the best fit for my needs.  

It has more documentation and a much simpler usage flow. Using GRPC is almost 
the same as writing native Go code which does not use RPC.  

Using Cap'N Proto felt like I was writing code specifically for RPC. And was 
not smooth or intuitive like GRPC was.  

I believe that Cap'N Proto has much potential. It offers an very interesting 
feature set that I believe many people can benefit from. However at this 
moment it is not the best fit for my needs.
