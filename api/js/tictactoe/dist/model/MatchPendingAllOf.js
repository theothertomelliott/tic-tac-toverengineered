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
 * The MatchPendingAllOf model module.
 * @module model/MatchPendingAllOf
 * @version 1.0.0
 */
var MatchPendingAllOf = /*#__PURE__*/function () {
  /**
   * Constructs a new <code>MatchPendingAllOf</code>.
   * @alias module:model/MatchPendingAllOf
   * @param requestID {String} 
   */
  function MatchPendingAllOf(requestID) {
    _classCallCheck(this, MatchPendingAllOf);

    MatchPendingAllOf.initialize(this, requestID);
  }
  /**
   * Initializes the fields of this object.
   * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
   * Only for internal use.
   */


  _createClass(MatchPendingAllOf, null, [{
    key: "initialize",
    value: function initialize(obj, requestID) {
      obj['requestID'] = requestID;
    }
    /**
     * Constructs a <code>MatchPendingAllOf</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/MatchPendingAllOf} obj Optional instance to populate.
     * @return {module:model/MatchPendingAllOf} The populated <code>MatchPendingAllOf</code> instance.
     */

  }, {
    key: "constructFromObject",
    value: function constructFromObject(data, obj) {
      if (data) {
        obj = obj || new MatchPendingAllOf();

        if (data.hasOwnProperty('requestID')) {
          obj['requestID'] = _ApiClient["default"].convertToType(data['requestID'], 'String');
        }
      }

      return obj;
    }
  }]);

  return MatchPendingAllOf;
}();
/**
 * @member {String} requestID
 */


MatchPendingAllOf.prototype['requestID'] = undefined;
var _default = MatchPendingAllOf;
exports["default"] = _default;