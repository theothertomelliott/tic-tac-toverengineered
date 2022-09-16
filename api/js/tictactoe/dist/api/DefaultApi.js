"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports["default"] = void 0;

var _ApiClient = _interopRequireDefault(require("../ApiClient"));

var _Error = _interopRequireDefault(require("../model/Error"));

var _Grid = _interopRequireDefault(require("../model/Grid"));

var _Match = _interopRequireDefault(require("../model/Match"));

var _MatchPair = _interopRequireDefault(require("../model/MatchPair"));

var _MatchPending = _interopRequireDefault(require("../model/MatchPending"));

var _Winner = _interopRequireDefault(require("../model/Winner"));

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { "default": obj }; }

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }

function _defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } }

function _createClass(Constructor, protoProps, staticProps) { if (protoProps) _defineProperties(Constructor.prototype, protoProps); if (staticProps) _defineProperties(Constructor, staticProps); Object.defineProperty(Constructor, "prototype", { writable: false }); return Constructor; }

/**
* Default service.
* @module api/DefaultApi
* @version 1.0.0
*/
var DefaultApi = /*#__PURE__*/function () {
  /**
  * Constructs a new DefaultApi. 
  * @alias module:api/DefaultApi
  * @class
  * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
  * default to {@link module:ApiClient#instance} if unspecified.
  */
  function DefaultApi(apiClient) {
    _classCallCheck(this, DefaultApi);

    this.apiClient = apiClient || _ApiClient["default"].instance;
  }
  /**
   * Callback function to receive the result of the currentPlayer operation.
   * @callback module:api/DefaultApi~currentPlayerCallback
   * @param {String} error Error message, if any.
   * @param {String} data The data returned by the service call.
   * @param {String} response The complete HTTP response.
   */

  /**
   * Return the current player for a game
   * @param {String} game ID of game
   * @param {module:api/DefaultApi~currentPlayerCallback} callback The callback function, accepting three arguments: error, data, response
   * data is of type: {@link String}
   */


  _createClass(DefaultApi, [{
    key: "currentPlayer",
    value: function currentPlayer(game, callback) {
      var postBody = null; // verify the required parameter 'game' is set

      if (game === undefined || game === null) {
        throw new _Error["default"]("Missing the required parameter 'game' when calling currentPlayer");
      }

      var pathParams = {
        'game': game
      };
      var queryParams = {};
      var headerParams = {};
      var formParams = {};
      var authNames = [];
      var contentTypes = [];
      var accepts = ['application/json'];
      var returnType = 'String';
      return this.apiClient.callApi('/{game}/player/current', 'GET', pathParams, queryParams, headerParams, formParams, postBody, authNames, contentTypes, accepts, returnType, null, callback);
    }
    /**
     * Callback function to receive the result of the gameGrid operation.
     * @callback module:api/DefaultApi~gameGridCallback
     * @param {String} error Error message, if any.
     * @param {module:model/Grid} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Return the grid state for a game
     * @param {String} game ID of game
     * @param {module:api/DefaultApi~gameGridCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/Grid}
     */

  }, {
    key: "gameGrid",
    value: function gameGrid(game, callback) {
      var postBody = null; // verify the required parameter 'game' is set

      if (game === undefined || game === null) {
        throw new _Error["default"]("Missing the required parameter 'game' when calling gameGrid");
      }

      var pathParams = {
        'game': game
      };
      var queryParams = {};
      var headerParams = {};
      var formParams = {};
      var authNames = [];
      var contentTypes = [];
      var accepts = ['application/json'];
      var returnType = _Grid["default"];
      return this.apiClient.callApi('/{game}/grid', 'GET', pathParams, queryParams, headerParams, formParams, postBody, authNames, contentTypes, accepts, returnType, null, callback);
    }
    /**
     * Callback function to receive the result of the index operation.
     * @callback module:api/DefaultApi~indexCallback
     * @param {String} error Error message, if any.
     * @param {Array.<String>} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Returns a list of game IDs
     * @param {Object} opts Optional parameters
     * @param {Number} opts.offset starting offset for results
     * @param {Number} opts.max maximum number of results to return
     * @param {module:api/DefaultApi~indexCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link Array.<String>}
     */

  }, {
    key: "index",
    value: function index(opts, callback) {
      opts = opts || {};
      var postBody = null;
      var pathParams = {};
      var queryParams = {
        'offset': opts['offset'],
        'max': opts['max']
      };
      var headerParams = {};
      var formParams = {};
      var authNames = [];
      var contentTypes = [];
      var accepts = ['application/json'];
      var returnType = ['String'];
      return this.apiClient.callApi('/', 'GET', pathParams, queryParams, headerParams, formParams, postBody, authNames, contentTypes, accepts, returnType, null, callback);
    }
    /**
     * Callback function to receive the result of the matchStatus operation.
     * @callback module:api/DefaultApi~matchStatusCallback
     * @param {String} error Error message, if any.
     * @param {module:model/Match} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get status of a match request
     * @param {String} requestID ID of match request to be checked
     * @param {module:api/DefaultApi~matchStatusCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/Match}
     */

  }, {
    key: "matchStatus",
    value: function matchStatus(requestID, callback) {
      var postBody = null; // verify the required parameter 'requestID' is set

      if (requestID === undefined || requestID === null) {
        throw new _Error["default"]("Missing the required parameter 'requestID' when calling matchStatus");
      }

      var pathParams = {};
      var queryParams = {
        'requestID': requestID
      };
      var headerParams = {};
      var formParams = {};
      var authNames = [];
      var contentTypes = [];
      var accepts = ['application/json'];
      var returnType = _Match["default"];
      return this.apiClient.callApi('/match', 'GET', pathParams, queryParams, headerParams, formParams, postBody, authNames, contentTypes, accepts, returnType, null, callback);
    }
    /**
     * Callback function to receive the result of the play operation.
     * @callback module:api/DefaultApi~playCallback
     * @param {String} error Error message, if any.
     * @param {String} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Make a move in a game
     * @param {String} game ID of game
     * @param {String} token token of player making move
     * @param {Number} i column in grid
     * @param {Number} j row in grid
     * @param {module:api/DefaultApi~playCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link String}
     */

  }, {
    key: "play",
    value: function play(game, token, i, j, callback) {
      var postBody = null; // verify the required parameter 'game' is set

      if (game === undefined || game === null) {
        throw new _Error["default"]("Missing the required parameter 'game' when calling play");
      } // verify the required parameter 'token' is set


      if (token === undefined || token === null) {
        throw new _Error["default"]("Missing the required parameter 'token' when calling play");
      } // verify the required parameter 'i' is set


      if (i === undefined || i === null) {
        throw new _Error["default"]("Missing the required parameter 'i' when calling play");
      } // verify the required parameter 'j' is set


      if (j === undefined || j === null) {
        throw new _Error["default"]("Missing the required parameter 'j' when calling play");
      }

      var pathParams = {
        'game': game
      };
      var queryParams = {
        'token': token,
        'i': i,
        'j': j
      };
      var headerParams = {};
      var formParams = {};
      var authNames = [];
      var contentTypes = [];
      var accepts = ['application/json'];
      var returnType = 'String';
      return this.apiClient.callApi('/{game}/play', 'POST', pathParams, queryParams, headerParams, formParams, postBody, authNames, contentTypes, accepts, returnType, null, callback);
    }
    /**
     * Callback function to receive the result of the requestMatch operation.
     * @callback module:api/DefaultApi~requestMatchCallback
     * @param {String} error Error message, if any.
     * @param {module:model/MatchPending} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Request a new /match
     * @param {module:api/DefaultApi~requestMatchCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/MatchPending}
     */

  }, {
    key: "requestMatch",
    value: function requestMatch(callback) {
      var postBody = null;
      var pathParams = {};
      var queryParams = {};
      var headerParams = {};
      var formParams = {};
      var authNames = [];
      var contentTypes = [];
      var accepts = ['application/json'];
      var returnType = _MatchPending["default"];
      return this.apiClient.callApi('/match', 'POST', pathParams, queryParams, headerParams, formParams, postBody, authNames, contentTypes, accepts, returnType, null, callback);
    }
    /**
     * Callback function to receive the result of the requestMatchPair operation.
     * @callback module:api/DefaultApi~requestMatchPairCallback
     * @param {String} error Error message, if any.
     * @param {module:model/MatchPair} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Request matches for both players in a game, to be used for one-player games.
     * @param {module:api/DefaultApi~requestMatchPairCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/MatchPair}
     */

  }, {
    key: "requestMatchPair",
    value: function requestMatchPair(callback) {
      var postBody = null;
      var pathParams = {};
      var queryParams = {};
      var headerParams = {};
      var formParams = {};
      var authNames = [];
      var contentTypes = [];
      var accepts = ['application/json'];
      var returnType = _MatchPair["default"];
      return this.apiClient.callApi('/match/pair', 'POST', pathParams, queryParams, headerParams, formParams, postBody, authNames, contentTypes, accepts, returnType, null, callback);
    }
    /**
     * Callback function to receive the result of the winner operation.
     * @callback module:api/DefaultApi~winnerCallback
     * @param {String} error Error message, if any.
     * @param {module:model/Winner} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Return the winner, if any for a game
     * @param {String} game ID of game
     * @param {module:api/DefaultApi~winnerCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/Winner}
     */

  }, {
    key: "winner",
    value: function winner(game, callback) {
      var postBody = null; // verify the required parameter 'game' is set

      if (game === undefined || game === null) {
        throw new _Error["default"]("Missing the required parameter 'game' when calling winner");
      }

      var pathParams = {
        'game': game
      };
      var queryParams = {};
      var headerParams = {};
      var formParams = {};
      var authNames = [];
      var contentTypes = [];
      var accepts = ['application/json'];
      var returnType = _Winner["default"];
      return this.apiClient.callApi('/{game}/winner', 'GET', pathParams, queryParams, headerParams, formParams, postBody, authNames, contentTypes, accepts, returnType, null, callback);
    }
  }]);

  return DefaultApi;
}();

exports["default"] = DefaultApi;