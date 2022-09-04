"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports["default"] = void 0;

var _ApiClient = _interopRequireDefault(require("../ApiClient"));

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { "default": obj }; }

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }

function _defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } }

function _createClass(Constructor, protoProps, staticProps) { if (protoProps) _defineProperties(Constructor.prototype, protoProps); if (staticProps) _defineProperties(Constructor, staticProps); Object.defineProperty(Constructor, "prototype", { writable: false }); return Constructor; }

/**
 * The MatchAllOf model module.
 * @module model/MatchAllOf
 * @version 1.0.0
 */
var MatchAllOf = /*#__PURE__*/function () {
  /**
   * Constructs a new <code>MatchAllOf</code>.
   * @alias module:model/MatchAllOf
   * @param gameID {String} 
   * @param mark {String} 
   * @param token {String} 
   */
  function MatchAllOf(gameID, mark, token) {
    _classCallCheck(this, MatchAllOf);

    MatchAllOf.initialize(this, gameID, mark, token);
  }
  /**
   * Initializes the fields of this object.
   * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
   * Only for internal use.
   */


  _createClass(MatchAllOf, null, [{
    key: "initialize",
    value: function initialize(obj, gameID, mark, token) {
      obj['gameID'] = gameID;
      obj['mark'] = mark;
      obj['token'] = token;
    }
    /**
     * Constructs a <code>MatchAllOf</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/MatchAllOf} obj Optional instance to populate.
     * @return {module:model/MatchAllOf} The populated <code>MatchAllOf</code> instance.
     */

  }, {
    key: "constructFromObject",
    value: function constructFromObject(data, obj) {
      if (data) {
        obj = obj || new MatchAllOf();

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

  return MatchAllOf;
}();
/**
 * @member {String} gameID
 */


MatchAllOf.prototype['gameID'] = undefined;
/**
 * @member {String} mark
 */

MatchAllOf.prototype['mark'] = undefined;
/**
 * @member {String} token
 */

MatchAllOf.prototype['token'] = undefined;
var _default = MatchAllOf;
exports["default"] = _default;