/**
 * Tic Tac Toe
 * An API for games of Tic Tac Toe
 *
 * The version of the OpenAPI document: 1.0.0
 * Contact: tom.w.elliott@gmail.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';

/**
 * The WinnerAllOf model module.
 * @module model/WinnerAllOf
 * @version 1.0.0
 */
class WinnerAllOf {
    /**
     * Constructs a new <code>WinnerAllOf</code>.
     * @alias module:model/WinnerAllOf
     */
    constructor() { 
        
        WinnerAllOf.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>WinnerAllOf</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/WinnerAllOf} obj Optional instance to populate.
     * @return {module:model/WinnerAllOf} The populated <code>WinnerAllOf</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new WinnerAllOf();

            if (data.hasOwnProperty('winner')) {
                obj['winner'] = ApiClient.convertToType(data['winner'], 'String');
            }
            if (data.hasOwnProperty('draw')) {
                obj['draw'] = ApiClient.convertToType(data['draw'], 'Boolean');
            }
        }
        return obj;
    }


}

/**
 * @member {String} winner
 */
WinnerAllOf.prototype['winner'] = undefined;

/**
 * @member {Boolean} draw
 */
WinnerAllOf.prototype['draw'] = undefined;






export default WinnerAllOf;
