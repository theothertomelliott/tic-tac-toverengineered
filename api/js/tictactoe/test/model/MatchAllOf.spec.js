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
    instance = new TicTacToe.MatchAllOf();
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

  describe('MatchAllOf', function() {
    it('should create an instance of MatchAllOf', function() {
      // uncomment below and update the code to test MatchAllOf
      //var instance = new TicTacToe.MatchAllOf();
      //expect(instance).to.be.a(TicTacToe.MatchAllOf);
    });

    it('should have the property gameID (base name: "gameID")', function() {
      // uncomment below and update the code to test the property gameID
      //var instance = new TicTacToe.MatchAllOf();
      //expect(instance).to.be();
    });

    it('should have the property mark (base name: "mark")', function() {
      // uncomment below and update the code to test the property mark
      //var instance = new TicTacToe.MatchAllOf();
      //expect(instance).to.be();
    });

    it('should have the property token (base name: "token")', function() {
      // uncomment below and update the code to test the property token
      //var instance = new TicTacToe.MatchAllOf();
      //expect(instance).to.be();
    });

  });

}));