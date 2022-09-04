"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports["default"] = void 0;

var _ApiClient = _interopRequireDefault(require("../ApiClient"));

var _MatchAllOf = _interopRequireDefault(require("./MatchAllOf"));

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { "default": obj }; }

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }

function _defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } }

function _createClass(Constructor, protoProps, staticProps) { if (protoProps) _defineProperties(Constructor.prototype, protoProps); if (staticProps) _defineProperties(Constructor, staticProps); Object.defineProperty(Constructor, "prototype", { writable: false }); return Constructor; }

/**
 * The Match model module.
 * @module model/Match
 * @version 1.0.0
 */
var Match = /*#__PURE__*/function () {
  /**
   * Constructs a new <code>Match</code>.
   * @alias module:model/Match
   * @implements module:model/MatchAllOf
   * @param gameID {String} 
   * @param mark {String} 
   * @param token {String} 
   */
  function Match(gameID, mark, token) {
    _classCallCheck(this, Match);

    _MatchAllOf["default"].initialize(this, gameID, mark, token);

    Match.initialize(this, gameID, mark, token);
  }
  /**
   * Initializes the fields of this object.
   * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
   * Only for internal use.
   */


  _createClass(Match, null, [{
    key: "initialize",
    value: function initialize(obj, gameID, mark, token) {
      obj['gameID'] = gameID;
      obj['mark'] = mark;
      obj['token'] = token;
    }
    /**
     * Constructs a <code>Match</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/Match} obj Optional instance to populate.
     * @return {module:model/Match} The populated <code>Match</code> instance.
     */

  }, {
    key: "constructFromObject",
    value: function constructFromObject(data, obj) {
      if (data) {
        obj = obj || new Match();

        _MatchAllOf["default"].constructFromObject(data, obj);

        if (data.hasOwnProperty('gameID')) {
          obj['gameID'] = _ApiClient["default"].convertToType(data['gameID'], 'String');
        }

        if (data.hasOwnProperty('mark')) {
          obj['mark'] = _ApiClient["default"].convertToType(data['mark'], 'String');
        }

        if (data.hasOwnProperty('token')) {
          obj['token'] = _ApiClient["default"].convertToType(data['token'], 'String');
        }
      }

      return obj;
    }
  }]);

  return Match;
}();
/**
 * @member {String} gameID
 */


Match.prototype['gameID'] = undefined;
/**
 * @member {String} mark
 */

Match.prototype['mark'] = undefined;
/**
 * @member {String} token
 */

Match.prototype['token'] = undefined; // Implement MatchAllOf interface:

/**
 * @member {String} gameID
 */

_MatchAllOf["default"].prototype['gameID'] = undefined;
/**
 * @member {String} mark
 */

_MatchAllOf["default"].prototype['mark'] = undefined;
/**
 * @member {String} token
 */

_MatchAllOf["default"].prototype['token'] = undefined;
var _default = Match;
exports["default"] = _default;