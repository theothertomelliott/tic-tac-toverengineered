import http from 'k6/http';
import { group, check, sleep } from 'k6';

export const options = {
  thresholds: {
    http_req_failed: ['rate<0.01'], // http errors should be less than 1%
    http_req_duration: ['p(90)<400'], // 90% of requests should be below 400ms
  },
  stages: [
    { duration: '60s', target: 500 },
  ],
};

export default function () {
    const host = 'http://localhost:8079';

    let pairRes = {}
    group('create game', function () {
      pairRes = http.post(host + '/match/pair');
    });

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
}