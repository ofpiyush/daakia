# Daakia

High performance, Simple messaging framework. 

Describe your Methods -> Generate code for client and server -> Code!

## Note
As is evident from the commits, daakia is not being actively maintained.

My current job commitments leave little room for anything else. I wish to resume work on it someday, but that day does not look near.

If you'd like to work on it, open an issue :)
## Installation

```bash
$ go get -u github.com/daakia/daakia
$ go install github.com/daakia/daakia/cmd/daakia
```


## Usage

### Create a toml file

```toml
# daakia.toml
namespace = "daakia"
[[services]]
	name = "TwoWayPublish"
	[[services.server]]
		name = "Publish"
		in = true
	[[services.client]]
		name = "Publish"
		in = true
```

### Generate code
```
$ daakia --go --js -i ./daakia.toml -o ./out/path/here
```

## Why Daakia?

#### Traditional messaging protocols are rigid.
__Possible Solution:__

Write code around/on top of the messaging software!
A Sub-protocol on application messages sound familiar? (Text message/buttons/audio/system messages etc...)
__Cons:__

* It needs developers across different teams to implement the same set of features and use documentation to co-ordinate.
* Forwards and backwards compatibility can be very tricky given the different speed of releases between mobile native apps, websites, and servers.


#### Traditional messaging servers/brokers are rigid.
__Possible Solution:__

Write the plugins/extensions for your feature!

__Cons:__
* Often, the software is in an unfamiliar language.
* A plugin is limited to the code accessible. The part you want to tweak might just need hacking the software.
* Your plugin might need changes every now and then depending on the release.

#### Bonus Problems
* One WebSocket/TCP connection for messaging, random number of expensive http calls to services for request/response type of messages.
* Serialize and deserialize in various formats (often with allocations) between the client and various services.

### The Daakia solution.
You write your own protocol and export in your language. Daakia works on any transport protocol in any language. 
 
A single connection handles everything from request response to streaming use cases.

Our recommended way to use Daakia is with flatbuffers.

You can get cross language 0 allocation serialization/deserialization with flatbuffers.
