# Daakia

Daakia helps you hack a messaging protocol / build your own.

Code and API are extremely unstable. The project is at an experimental stage right now. Use at your own risk


## Why Daakia?

### The conflict
* Traditional messaging servers and protocols are treated as rigid black boxes.
* Businesses change quickly and need flexibility. 
* Often, it boils down to features that warrant a slightly different protocol/protocol that could be extended.
* Different types of messaging often need us to work with different set of tools/ open multiple connections of various lifetimes with the backend.

### Current possible resolutions

#### Write code around/on top of the messaging layer.
This is the most common approach taken. A Sub-protocol on application messages sound familiar? (Text message/buttons/audio/system messages etc...)

This approach can open gates to a host of bugs.
* It needs developers across different teams to implement the same set of features.
* Given the different speeds of releases between mobile native apps, websites and servers, this can be very tricky in terms of forwards and backwards compatibility.

#### Use the plugins/extensions
If the messaging server has a good extensions/plugins system, write extensions to support! This by far is the best solution out there. It still doesn't solve the problems with sub-protocols.

#### Hack the messaging server.
These hacks are often put in places which might make it incompatible with future updates of the software.

### The Daakia way.
Daakia borrows the best ideas from all types of messaging systems and (to my knowledge) throws in some of it's own.

* The protocol and message types should be extensible, forwards and backwards compatible.
* Messages should be serialized in a binary format that allows random access through a byte buffer across languages and platforms without additional parsing/ allocation. 
* Shared objects(messages), serialization/deserialization and routing code should be generated from a source file written in a common language (daak lang?).
* Multiple underlying transports should be supported and each should use the strengths of the transport.
* Write libraries and tutorials for most common synchronous/asynchronous messaging use cases.
* Marshalling and Unmarshalling to popular formats should be easy at the cost of a few allocations.
* Allow various types of messaging mechanisms to co-exist on exactly one persistent connection.
