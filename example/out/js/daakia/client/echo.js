
var daakia = this.daakia || {};

daakia.Echo = function(Client,Transport) {
  if (!this instanceof daakia.Echo) {
    throw Error("daakia.Echo called without the new keyword");
  }
  this.transport = Transport;
  this.client = new Client(this);
  this.transport.OnConnect(this.client._Connected.bind(this.client));
  
  this.yoloHeader =  new Uint8Array([daakia.Echo.Methods.Server_yolo])
  
  this.transport.Next(this.Route.bind(this));
  this.transport.Connect();
};
daakia.Echo.Methods = { 
	Server_yolo: 11,
};
daakia.Echo.prototype.Route = function (message) {
	var route = new Uint8Array(message,1,1);
    var buffer = new Uint8Array(message,1);
    switch (route[0]) {
    
    }
};

daakia.Echo.prototype.yolo = function() {
  this.transport.Send(this.yoloHeader);
};


daakia.Echo.prototype.close = function() {
	this.transport.close();
}
this.daakia = daakia;