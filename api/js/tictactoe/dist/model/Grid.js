"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports["default"] = void 0;

var _ApiClient = _interopRequireDefault(require("../ApiClient"));

var _GridAllOf = _interopRequireDefault(require("./GridAllOf"));

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { "default": obj }; }

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }

function _defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } }

function _createClass(Constructor, protoProps, staticProps) { if (protoProps) _defineProperties(Constructor.prototype, protoProps); if (staticProps) _defineProperties(Constructor, staticProps); Object.defineProperty(Constructor, "prototype", { writable: false }); return Constructor; }

/**
 * The Grid model module.
 * @module model/Grid
 * @version 1.0.0
 */
var Grid = /*#__PURE__*/function () {
  /**
   * Constructs a new <code>Grid</code>.
   * @alias module:model/Grid
   * @implements module:model/GridAllOf
   * @param grid {Array.<Array.<String>>} 
   */
  function Grid(grid) {
    _classCallCheck(this, Grid);

    _GridAllOf["default"].initialize(this, grid);

    Grid.initialize(this, grid);
  }
  /**
   * Initializes the fields of this object.
   * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
   * Only for internal use.
   */


  _createClass(Grid, null, [{
    key: "initialize",
    value: function initialize(obj, grid) {
      obj['grid'] = grid;
    }
    /**
     * Constructs a <code>Grid</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/Grid} obj Optional instance to populate.
     * @return {module:model/Grid} The populated <code>Grid</code> instance.
     */

  }, {
    key: "constructFromObject",
    value: function constructFromObject(data, obj) {
      if (data) {
        obj = obj || new Grid();

        _GridAllOf["default"].constructFromObject(data, obj);

        if (data.hasOwnProperty('grid')) {
          obj['grid'] = _ApiClient["default"].convertToType(data['grid'], [['String']]);
        }
      }

      return obj;
    }
  }]);

  return Grid;
}();
/**
 * @member {Array.<Array.<String>>} grid
 */


Grid.prototype['grid'] = undefined; // Implement GridAllOf interface:

/**
 * @member {Array.<Array.<String>>} grid
 */

_GridAllOf["default"].prototype['grid'] = undefined;
var _default = Grid;
exports["default"] = _default;