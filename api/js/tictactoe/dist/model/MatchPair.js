"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports["default"] = void 0;

var _ApiClient = _interopRequireDefault(require("../ApiClient"));

var _Match = _interopRequireDefault(require("./Match"));

var _MatchPairAllOf = _interopRequireDefault(require("./MatchPairAllOf"));

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { "default": obj }; }

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }

function _defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } }

function _createClass(Constructor, protoProps, staticProps) { if (protoProps) _defineProperties(Constructor.prototype, protoProps); if (staticProps) _defineProperties(Constructor, staticProps); Object.defineProperty(Constructor, "prototype", { writable: false }); return Constructor; }

/**
 * The MatchPair model module.
 * @module model/MatchPair
 * @version 1.0.0
 */
var MatchPair = /*#__PURE__*/function () {
  /**
   * Constructs a new <code>MatchPair</code>.
   * @alias module:model/MatchPair
   * @implements module:model/MatchPairAllOf
   * @param x {module:model/Match} 
   * @param o {module:model/Match} 
   */
  function MatchPair(x, o) {
    _classCallCheck(this, MatchPair);

    _MatchPairAllOf["default"].initialize(this, x, o);

    MatchPair.initialize(this, x, o);
  }
  /**
   * Initializes the fields of this object.
   * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
   * Only for internal use.
   */


  _createClass(MatchPair, null, [{
    key: "initialize",
    value: function initialize(obj, x, o) {
      obj['x'] = x;
      obj['o'] = o;
    }
    /**
     * Constructs a <code>MatchPair</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/MatchPair} obj Optional instance to populate.
     * @return {module:model/MatchPair} The populated <code>MatchPair</code> instance.
     */

  }, {
    key: "constructFromObject",
    value: function constructFromObject(data, obj) {
      if (data) {
        obj = obj || new MatchPair();

        _MatchPairAllOf["default"].constructFromObject(data, obj);

        if (data.hasOwnProperty('x')) {
          obj['x'] = _Match["default"].constructFromObject(data['x']);
        }

        if (data.hasOwnProperty('o')) {
          obj['o'] = _Match["default"].constructFromObject(data['o']);
        }
      }

      return obj;
    }
  }]);

  return MatchPair;
}();
/**
 * @member {module:model/Match} x
 */


MatchPair.prototype['x'] = undefined;
/**
 * @member {module:model/Match} o
 */

MatchPair.prototype['o'] = undefined; // Implement MatchPairAllOf interface:

/**
 * @member {module:model/Match} x
 */

_MatchPairAllOf["default"].prototype['x'] = undefined;
/**
 * @member {module:model/Match} o
 */

_MatchPairAllOf["default"].prototype['o'] = undefined;
var _default = MatchPair;
exports["default"] = _default;