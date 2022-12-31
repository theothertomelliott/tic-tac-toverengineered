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

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD.
    define(['expect.js', process.cwd()+'/src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require(process.cwd()+'/src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.TicTacToe);
  }
}(this, function(expect, TicTacToe) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new TicTacToe.MatchPendingAllOf();
  });

  var getProperty = function(object, getter, property) {
    // Use getter method if present; otherwise, get the property directly.
    if (typeof object[getter] === 'function')
      return object[getter]();
    else
      return object[property];
  }

  var setProperty = function(object, setter, property, value) {
    // Use setter method if present; otherwise, set the property directly.
    if (typeof object[setter] === 'function')
      object[setter](value);
    else
      object[property] = value;
  }

  describe('MatchPendingAllOf', function() {
    it('should create an instance of MatchPendingAllOf', function() {
      // uncomment below and update the code to test MatchPendingAllOf
      //var instance = new TicTacToe.MatchPendingAllOf();
      //expect(instance).to.be.a(TicTacToe.MatchPendingAllOf);
    });

    it('should have the property requestID (base name: "requestID")', function() {
      // uncomment below and update the code to test the property requestID
      //var instance = new TicTacToe.MatchPendingAllOf();
      //expect(instance).to.be();
    });

  });

}));