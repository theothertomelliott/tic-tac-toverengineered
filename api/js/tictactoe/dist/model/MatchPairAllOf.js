"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports["default"] = void 0;

var _ApiClient = _interopRequireDefault(require("../ApiClient"));

var _Match = _interopRequireDefault(require("./Match"));

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { "default": obj }; }

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }

function _defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } }

function _createClass(Constructor, protoProps, staticProps) { if (protoProps) _defineProperties(Constructor.prototype, protoProps); if (staticProps) _defineProperties(Constructor, staticProps); Object.defineProperty(Constructor, "prototype", { writable: false }); return Constructor; }

/**
 * The MatchPairAllOf model module.
 * @module model/MatchPairAllOf
 * @version 1.0.0
 */
var MatchPairAllOf = /*#__PURE__*/function () {
  /**
   * Constructs a new <code>MatchPairAllOf</code>.
   * @alias module:model/MatchPairAllOf
   * @param x {module:model/Match} 
   * @param o {module:model/Match} 
   */
  function MatchPairAllOf(x, o) {
    _classCallCheck(this, MatchPairAllOf);

    MatchPairAllOf.initialize(this, x, o);
  }
  /**
   * Initializes the fields of this object.
   * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
   * Only for internal use.
   */


  _createClass(MatchPairAllOf, null, [{
    key: "initialize",
    value: function initialize(obj, x, o) {
      obj['x'] = x;
      obj['o'] = o;
    }
    /**
     * Constructs a <code>MatchPairAllOf</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/MatchPairAllOf} obj Optional instance to populate.
     * @return {module:model/MatchPairAllOf} The populated <code>MatchPairAllOf</code> instance.
     */

  }, {
    key: "constructFromObject",
    value: function constructFromObject(data, obj) {
      if (data) {
        obj = obj || new MatchPairAllOf();

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

  return MatchPairAllOf;
}();
/**
 * @member {module:model/Match} x
 */


MatchPairAllOf.prototype['x'] = undefined;
/**
 * @member {module:model/Match} o
 */

MatchPairAllOf.prototype['o'] = undefined;
var _default = MatchPairAllOf;
exports["default"] = _default;