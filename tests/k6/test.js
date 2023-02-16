import { playGame } from './game.js';

export const options = {
  scenarios: {
    default: {
      executor: "ramping-vus",
      env: { TEST_HOSTNAME: 'localhost:8081' },
      stages: [
        { duration: '30s', target: 2 },
        { duration: '30s', target: 40 },
        { duration: '30s', target: 2 },
      ],
    }
  },
  thresholds: {
    http_req_failed: ['rate<0.01'], // http errors should be less than 1%
    http_req_duration: ['p(90)<400'], // 90% of requests should be below 400ms
  },
};

export default playGame;