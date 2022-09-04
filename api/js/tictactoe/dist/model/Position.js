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
 * The Position model module.
 * @module model/Position
 * @version 1.0.0
 */
var Position = /*#__PURE__*/function () {
  /**
   * Constructs a new <code>Position</code>.
   * @alias module:model/Position
   * @param i {Number} 
   * @param j {Number} 
   */
  function Position(i, j) {
    _classCallCheck(this, Position);

    Position.initialize(this, i, j);
  }
  /**
   * Initializes the fields of this object.
   * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
   * Only for internal use.
   */


  _createClass(Position, null, [{
    key: "initialize",
    value: function initialize(obj, i, j) {
      obj['i'] = i;
      obj['j'] = j;
    }
    /**
     * Constructs a <code>Position</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/Position} obj Optional instance to populate.
     * @return {module:model/Position} The populated <code>Position</code> instance.
     */

  }, {
    key: "constructFromObject",
    value: function constructFromObject(data, obj) {
      if (data) {
        obj = obj || new Position();

        if (data.hasOwnProperty('i')) {
          obj['i'] = _ApiClient["default"].convertToType(data['i'], 'Number');
        }

        if (data.hasOwnProperty('j')) {
          obj['j'] = _ApiClient["default"].convertToType(data['j'], 'Number');
        }
      }

      return obj;
    }
  }]);

  return Position;
}();
/**
 * @member {Number} i
 */


Position.prototype['i'] = undefined;
/**
 * @member {Number} j
 */

Position.prototype['j'] = undefined;
var _default = Position;
exports["default"] = _default;