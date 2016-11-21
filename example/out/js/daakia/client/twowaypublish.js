
var daakia = this.daakia || {};

daakia.TwoWayPublish = function(Client,Transport) {
  if (!this instanceof daakia.TwoWayPublish) {
    throw Error("daakia.TwoWayPublish called without the new keyword");
  }
  this.transport = Transport;
  this.client = new Client(this);
  this.transport.OnConnect(this.client._Connected.bind(this.client));
  
  this.PublishHeader =  new Uint8Array([daakia.TwoWayPublish.Methods.Server_Publish])
  
  this.PubAckHeader =  new Uint8Array([daakia.TwoWayPublish.Methods.Server_PubAck])
  
  this.transport.Next(this.Route.bind(this));
  this.transport.Connect();
};
daakia.TwoWayPublish.Methods = { 
	Server_Publish: 11,
	Server_PubAck: 12,
	Client_Publish: 13,
	Client_PubAck: 14,
};
daakia.TwoWayPublish.prototype.Route = function (message) {
	var route = new Uint8Array(message,1,1);
    var buffer = new Uint8Array(message,1);
    switch (route[0]) {
    
      case daakia.TwoWayPublish.Methods.Client_Publish:
      	this.client.Publish(buffer);
        break;
    
      case daakia.TwoWayPublish.Methods.Client_PubAck:
      	this.client.PubAck(buffer);
        break;
    
    }
};

daakia.TwoWayPublish.prototype.Publish = function(payload) {
  this.transport.Send(this.PublishHeader, payload);
};

daakia.TwoWayPublish.prototype.PubAck = function(payload) {
  this.transport.Send(this.PubAckHeader, payload);
};


daakia.TwoWayPublish.prototype.close = function() {
	this.transport.close();
}
this.daakia = daakia;