"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports["default"] = void 0;

var _ApiClient = _interopRequireDefault(require("../ApiClient"));

var _WinnerAllOf = _interopRequireDefault(require("./WinnerAllOf"));

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { "default": obj }; }

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }

function _defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } }

function _createClass(Constructor, protoProps, staticProps) { if (protoProps) _defineProperties(Constructor.prototype, protoProps); if (staticProps) _defineProperties(Constructor, staticProps); Object.defineProperty(Constructor, "prototype", { writable: false }); return Constructor; }

/**
 * The Winner model module.
 * @module model/Winner
 * @version 1.0.0
 */
var Winner = /*#__PURE__*/function () {
  /**
   * Constructs a new <code>Winner</code>.
   * @alias module:model/Winner
   * @implements module:model/WinnerAllOf
   */
  function Winner() {
    _classCallCheck(this, Winner);

    _WinnerAllOf["default"].initialize(this);

    Winner.initialize(this);
  }
  /**
   * Initializes the fields of this object.
   * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
   * Only for internal use.
   */


  _createClass(Winner, null, [{
    key: "initialize",
    value: function initialize(obj) {}
    /**
     * Constructs a <code>Winner</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/Winner} obj Optional instance to populate.
     * @return {module:model/Winner} The populated <code>Winner</code> instance.
     */

  }, {
    key: "constructFromObject",
    value: function constructFromObject(data, obj) {
      if (data) {
        obj = obj || new Winner();

        _WinnerAllOf["default"].constructFromObject(data, obj);

        if (data.hasOwnProperty('winner')) {
          obj['winner'] = _ApiClient["default"].convertToType(data['winner'], 'String');
        }

        if (data.hasOwnProperty('draw')) {
          obj['draw'] = _ApiClient["default"].convertToType(data['draw'], 'Boolean');
        }
      }

      return obj;
    }
  }]);

  return Winner;
}();
/**
 * @member {String} winner
 */


Winner.prototype['winner'] = undefined;
/**
 * @member {Boolean} draw
 */

Winner.prototype['draw'] = undefined; // Implement WinnerAllOf interface:

/**
 * @member {String} winner
 */

_WinnerAllOf["default"].prototype['winner'] = undefined;
/**
 * @member {Boolean} draw
 */

_WinnerAllOf["default"].prototype['draw'] = undefined;
var _default = Winner;
exports["default"] = _default;