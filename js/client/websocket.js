/**
 * A Websocket Listener
 * @constructor
 * @param {string} url
 */
function WebSocketTransport(url) {
  if(!this instanceof WebSocketTransport) {
    throw Error("Websocket Transport called without the new keyword")
  }
  this._url = url;
  this._ping = new Uint8Array([1]);
  this._connect();
  this._attempts = 0;
}
/**
 * WebSocket object
 * @type {WebSocket}
 * @private
 */
WebSocketTransport.prototype._ws = null;
WebSocketTransport.prototype._on_connect_func = null;

WebSocketTransport.prototype._connect = function () {
  this._ws = null;
  this._ws = new WebSocket(this._url);
  this._ws.binaryType = "arraybuffer";
  this._ws.onclose = this._on_close;
  this._ws.onerror = this._on_close;
  this._ws.onopen = this._on_connect;
  this._ws.onmessage = this._on_message;
  if(this._attempts <8) {
      this._attempts++;
  }
};

WebSocketTransport.prototype._on_close = function (e) {
  clearInterval(this._ping_interval);
  var delay = (Math.pow(2, this._attempts) - 1)*1000;
  setTimeout(this._connect,delay);
};

WebSocketTransport.prototype._on_connect = function() {
  this._on_connect_func = this._on_connect_func || function(){};
  var self = this;
  this._ping_interval = setInterval(function() {
    if(self._ws && self._ws.readyState == 1) {
      self._ws.send(self._ping);
    }
  },1000);
  this._on_connect_func();
};


WebSocketTransport.prototype.OnMessage = function(on_msg_func){
  this._on_message = on_msg_func;
  this._ws.onmessage = on_msg_func;
};

WebSocketTransport.prototype.OnConnect = function(on_connect_func) {
  this._on_connect_func = on_connect_func;
};

WebSocketTransport.prototype.Send = function(message) {
  this._ws.send(message);
};

this.WebSocketTransport = WebSocketTransport;