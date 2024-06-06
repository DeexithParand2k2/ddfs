package p2p

/*
Interfaces over strcuts:

Flexibility: Structs with methods allow you to define behavior specific to a single type, while interfaces allow you to define behavior shared by multiple types.
Implicit Implementation: Types must explicitly declare that they implement an interface by providing implementations for all the methods declared in the interface. Struct methods, on the other hand, are defined directly within the struct and are automatically associated with instances of that type.
Type-specific Methods: Struct methods are tied to a specific type and operate directly on instances of that type. Interfaces provide polymorphic behavior, allowing different types to be treated uniformly based on shared behavior defined by the interface.
*/

// method signatures that represent the actions/process of remote nodes
type Peer interface{}

// method signatures that handles the communication between  remote nodes in a network.
// ... can be TCP, UDP, WebSockets etc.
type Transport interface{}
