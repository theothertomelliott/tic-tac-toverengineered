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
 * The WinnerAllOf model module.
 * @module model/WinnerAllOf
 * @version 1.0.0
 */
var WinnerAllOf = /*#__PURE__*/function () {
  /**
   * Constructs a new <code>WinnerAllOf</code>.
   * @alias module:model/WinnerAllOf
   */
  function WinnerAllOf() {
    _classCallCheck(this, WinnerAllOf);

    WinnerAllOf.initialize(this);
  }
  /**
   * Initializes the fields of this object.
   * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
   * Only for internal use.
   */


  _createClass(WinnerAllOf, null, [{
    key: "initialize",
    value: function initialize(obj) {}
    /**
     * Constructs a <code>WinnerAllOf</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/WinnerAllOf} obj Optional instance to populate.
     * @return {module:model/WinnerAllOf} The populated <code>WinnerAllOf</code> instance.
     */

  }, {
    key: "constructFromObject",
    value: function constructFromObject(data, obj) {
      if (data) {
        obj = obj || new WinnerAllOf();

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

  return WinnerAllOf;
}();
/**
 * @member {String} winner
 */


WinnerAllOf.prototype['winner'] = undefined;
/**
 * @member {Boolean} draw
 */

WinnerAllOf.prototype['draw'] = undefined;
var _default = WinnerAllOf;
exports["default"] = _default;