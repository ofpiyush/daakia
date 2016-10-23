/**
 * Daakia
 * @constructor
 * @param {string} url
 * @param {Parser} Parser
 */

function Daakia (url, Protocol, Server, Transport) {
  if(!this instanceof Daakia) {
    throw Error("Daakia called without the new keyword")
  }
  var tp = Transport || WebSocketTransport;

  this.transport = new tp(url);
  this.transport.OnMessage(Protocol(Server));
}

Daakia.prototype.Send = function(message) {
  this.transport.Send(message);
};


this.Daakia = Daakia;