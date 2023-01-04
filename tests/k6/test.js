import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  stages: [
    { duration: '30s', target: 20 },
    { duration: '1m30s', target: 10 },
    { duration: '20s', target: 0 },
  ],
};

export default function () {
    const host = 'http://localhost:8081';

    const pairRes = http.post(host + '/match/pair');
    let pair = JSON.parse(pairRes.body);
    let gameId = pair["o"]["gameID"];
    let oToken = pair["o"]["token"];
    let xToken = pair["x"]["token"];

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

    const winRes = http.get(host + '/' + gameId + '/winner');
    const winner = JSON.parse(winRes.body);
    check(winner, {
        'not a draw': (r) => !r["draw"],
        'X won': (r) => r["winner"] == 'X',
     });
}