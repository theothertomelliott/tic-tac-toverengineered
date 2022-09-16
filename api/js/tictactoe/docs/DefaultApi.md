# TicTacToe.DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**currentPlayer**](DefaultApi.md#currentPlayer) | **GET** /{game}/player/current | 
[**gameGrid**](DefaultApi.md#gameGrid) | **GET** /{game}/grid | 
[**index**](DefaultApi.md#index) | **GET** / | 
[**matchStatus**](DefaultApi.md#matchStatus) | **GET** /match | 
[**play**](DefaultApi.md#play) | **POST** /{game}/play | 
[**requestMatch**](DefaultApi.md#requestMatch) | **POST** /match | 
[**requestMatchPair**](DefaultApi.md#requestMatchPair) | **POST** /match/pair | 
[**winner**](DefaultApi.md#winner) | **GET** /{game}/winner | 



## currentPlayer

> String currentPlayer(game)



Return the current player for a game

### Example

```javascript
import TicTacToe from 'tic_tac_toe';

let apiInstance = new TicTacToe.DefaultApi();
let game = "game_example"; // String | ID of game
apiInstance.currentPlayer(game, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **game** | **String**| ID of game | 

### Return type

**String**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## gameGrid

> Grid gameGrid(game)



Return the grid state for a game

### Example

```javascript
import TicTacToe from 'tic_tac_toe';

let apiInstance = new TicTacToe.DefaultApi();
let game = "game_example"; // String | ID of game
apiInstance.gameGrid(game, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **game** | **String**| ID of game | 

### Return type

[**Grid**](Grid.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## index

> [String] index(opts)



Returns a list of game IDs

### Example

```javascript
import TicTacToe from 'tic_tac_toe';

let apiInstance = new TicTacToe.DefaultApi();
let opts = {
  'offset': 789, // Number | starting offset for results
  'max': 789 // Number | maximum number of results to return
};
apiInstance.index(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **Number**| starting offset for results | [optional] 
 **max** | **Number**| maximum number of results to return | [optional] 

### Return type

**[String]**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## matchStatus

> Match matchStatus(requestID)



Get status of a match request

### Example

```javascript
import TicTacToe from 'tic_tac_toe';

let apiInstance = new TicTacToe.DefaultApi();
let requestID = "requestID_example"; // String | ID of match request to be checked
apiInstance.matchStatus(requestID, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestID** | **String**| ID of match request to be checked | 

### Return type

[**Match**](Match.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## play

> String play(game, token, i, j)



Make a move in a game

### Example

```javascript
import TicTacToe from 'tic_tac_toe';

let apiInstance = new TicTacToe.DefaultApi();
let game = "game_example"; // String | ID of game
let token = "token_example"; // String | token of player making move
let i = 56; // Number | column in grid
let j = 56; // Number | row in grid
apiInstance.play(game, token, i, j, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **game** | **String**| ID of game | 
 **token** | **String**| token of player making move | 
 **i** | **Number**| column in grid | 
 **j** | **Number**| row in grid | 

### Return type

**String**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## requestMatch

> MatchPending requestMatch()



Request a new /match

### Example

```javascript
import TicTacToe from 'tic_tac_toe';

let apiInstance = new TicTacToe.DefaultApi();
apiInstance.requestMatch((error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**MatchPending**](MatchPending.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## requestMatchPair

> MatchPair requestMatchPair()



Request matches for both players in a game, to be used for one-player games.

### Example

```javascript
import TicTacToe from 'tic_tac_toe';

let apiInstance = new TicTacToe.DefaultApi();
apiInstance.requestMatchPair((error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**MatchPair**](MatchPair.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## winner

> Winner winner(game)



Return the winner, if any for a game

### Example

```javascript
import TicTacToe from 'tic_tac_toe';

let apiInstance = new TicTacToe.DefaultApi();
let game = "game_example"; // String | ID of game
apiInstance.winner(game, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **game** | **String**| ID of game | 

### Return type

[**Winner**](Winner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

