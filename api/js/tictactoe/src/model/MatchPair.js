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
import Match from './Match';
import MatchPairAllOf from './MatchPairAllOf';

/**
 * The MatchPair model module.
 * @module model/MatchPair
 * @version 1.0.0
 */
class MatchPair {
    /**
     * Constructs a new <code>MatchPair</code>.
     * @alias module:model/MatchPair
     * @implements module:model/MatchPairAllOf
     * @param x {module:model/Match} 
     * @param o {module:model/Match} 
     */
    constructor(x, o) { 
        MatchPairAllOf.initialize(this, x, o);
        MatchPair.initialize(this, x, o);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, x, o) { 
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
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new MatchPair();
            MatchPairAllOf.constructFromObject(data, obj);

            if (data.hasOwnProperty('x')) {
                obj['x'] = Match.constructFromObject(data['x']);
            }
            if (data.hasOwnProperty('o')) {
                obj['o'] = Match.constructFromObject(data['o']);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/Match} x
 */
MatchPair.prototype['x'] = undefined;

/**
 * @member {module:model/Match} o
 */
MatchPair.prototype['o'] = undefined;


// Implement MatchPairAllOf interface:
/**
 * @member {module:model/Match} x
 */
MatchPairAllOf.prototype['x'] = undefined;
/**
 * @member {module:model/Match} o
 */
MatchPairAllOf.prototype['o'] = undefined;




export default MatchPair;
