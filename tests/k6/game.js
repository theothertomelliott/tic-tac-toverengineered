import http from 'k6/http';
import { group, check, sleep } from 'k6';

export function playGame() {
    const host = `http://${__ENV.TEST_HOSTNAME}`;

    let pairRes = {}
    let succeeded = false;
    group('create game', function () {
      pairRes = http.post(host + '/match/pair');
      succeeded = check(pairRes, { 'matching succeeded': (r) => r.status == 200 });
    });

    // Do not attempt to proceed if the game couldn't be created
    if(!succeeded) {
      sleep(1);
      return;
    }

    let pair = JSON.parse(pairRes.body);
    let gameId = pair["o"]["gameID"];
    let oToken = pair["o"]["token"];
    let xToken = pair["x"]["token"];

    group('moves', function() {
      let playRes = http.post(host + '/' + gameId + '/play?game=' + gameId + '&token=' + xToken + '&i=0&j=0');
      check(playRes, { 'first move succeeded': (r) => r.status == 200 });

      playRes = http.post(host + '/' + gameId + '/play?game=' + gameId + '&token=' + oToken + '&i=0&j=1');
      check(playRes, { 'second move succeeded': (r) => r.status == 200 });

      playRes = http.post(host + '/' + gameId + '/play?game=' + gameId + '&token=' + xToken + '&i=0&j=2');
      check(playRes, { 'third move succeeded': (r) => r.status == 200 });

      playRes = http.post(host + '/' + gameId + '/play?game=' + gameId + '&token=' + oToken + '&i=1&j=0');
      check(playRes, { 'fourth move succeeded': (r) => r.status == 200 });

      playRes = http.post(host + '/' + gameId + '/play?game=' + gameId + '&token=' + xToken + '&i=1&j=1');
      check(playRes, { 'fifth move succeeded': (r) => r.status == 200 });

      playRes = http.post(host + '/' + gameId + '/play?game=' + gameId + '&token=' + oToken + '&i=1&j=2');
      check(playRes, { 'sixth move succeeded': (r) => r.status == 200 });

      playRes = http.post(host + '/' + gameId + '/play?game=' + gameId + '&token=' + xToken + '&i=2&j=0');
      check(playRes, { 'seventh move succeeded': (r) => r.status == 200 });
    });

    group('winner', function() {
      const winRes = http.get(host + '/' + gameId + '/winner');
      const winner = JSON.parse(winRes.body);
      check(winner, {
          'not a draw': (r) => !r["draw"],
          'X won': (r) => r["winner"] == 'X',
      });
    });

    // Wait before playing another game
    sleep(1);
};