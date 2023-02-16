import { sleep, group } from 'k6'
import http from 'k6/http'

export const options = {
  thresholds: {},
  scenarios: {
    Scenario_1: {
      executor: 'ramping-vus',
      gracefulStop: '30s',
      stages: [
        { target: 20, duration: '1m' },
        { target: 20, duration: '3m30s' },
        { target: 0, duration: '1m' },
      ],
      gracefulRampDown: '30s',
      exec: 'scenario_1',
    },
  },
}

export function scenario_1() {
  let response

  group('page_2 - http://localhost:8080/new', function () {
    response = http.get('http://localhost:8080/new', {
      headers: {
        'upgrade-insecure-requests': '1',
        'sec-ch-ua': '"Not_A Brand";v="99", "Google Chrome";v="109", "Chromium";v="109"',
        'sec-ch-ua-mobile': '?0',
        'sec-ch-ua-platform': '"macOS"',
      },
    })
    response = http.get(
      'http://localhost:8081/0e5d63cb-bad8-4286-8f95-f2d8b1623727/player/current',
      {
        headers: {
          accept: 'application/json',
          'sec-ch-ua': '"Not_A Brand";v="99", "Google Chrome";v="109", "Chromium";v="109"',
          'sec-ch-ua-mobile': '?0',
          'sec-ch-ua-platform': '"macOS"',
        },
      }
    )
    response = http.get('http://localhost:8081/0e5d63cb-bad8-4286-8f95-f2d8b1623727/winner', {
      headers: {
        accept: 'application/json',
        'sec-ch-ua': '"Not_A Brand";v="99", "Google Chrome";v="109", "Chromium";v="109"',
        'sec-ch-ua-mobile': '?0',
        'sec-ch-ua-platform': '"macOS"',
      },
    })
    response = http.get('http://localhost:8081/0e5d63cb-bad8-4286-8f95-f2d8b1623727/grid', {
      headers: {
        accept: 'application/json',
        'sec-ch-ua': '"Not_A Brand";v="99", "Google Chrome";v="109", "Chromium";v="109"',
        'sec-ch-ua-mobile': '?0',
        'sec-ch-ua-platform': '"macOS"',
      },
    })
    sleep(0.9)
    response = http.post(
      'http://localhost:8081/0e5d63cb-bad8-4286-8f95-f2d8b1623727/play?token=%7B%22Game%22%3A%220e5d63cb-bad8-4286-8f95-f2d8b1623727%22%2C%22Player%22%3A%22X%22%7D&i=0&j=0',
      null,
      {
        headers: {
          accept: 'application/json',
          'sec-ch-ua': '"Not_A Brand";v="99", "Google Chrome";v="109", "Chromium";v="109"',
          'sec-ch-ua-mobile': '?0',
          'sec-ch-ua-platform': '"macOS"',
        },
      }
    )
    response = http.get(
      'http://localhost:8081/0e5d63cb-bad8-4286-8f95-f2d8b1623727/player/current',
      {
        headers: {
          accept: 'application/json',
          'sec-ch-ua': '"Not_A Brand";v="99", "Google Chrome";v="109", "Chromium";v="109"',
          'sec-ch-ua-mobile': '?0',
          'sec-ch-ua-platform': '"macOS"',
        },
      }
    )
    response = http.get('http://localhost:8081/0e5d63cb-bad8-4286-8f95-f2d8b1623727/winner', {
      headers: {
        accept: 'application/json',
        'sec-ch-ua': '"Not_A Brand";v="99", "Google Chrome";v="109", "Chromium";v="109"',
        'sec-ch-ua-mobile': '?0',
        'sec-ch-ua-platform': '"macOS"',
      },
    })
    response = http.get('http://localhost:8081/0e5d63cb-bad8-4286-8f95-f2d8b1623727/grid', {
      headers: {
        accept: 'application/json',
        'sec-ch-ua': '"Not_A Brand";v="99", "Google Chrome";v="109", "Chromium";v="109"',
        'sec-ch-ua-mobile': '?0',
        'sec-ch-ua-platform': '"macOS"',
      },
    })
    sleep(0.6)
    response = http.post(
      'http://localhost:8081/0e5d63cb-bad8-4286-8f95-f2d8b1623727/play?token=%7B%22Game%22%3A%220e5d63cb-bad8-4286-8f95-f2d8b1623727%22%2C%22Player%22%3A%22O%22%7D&i=0&j=1',
      null,
      {
        headers: {
          accept: 'application/json',
          'sec-ch-ua': '"Not_A Brand";v="99", "Google Chrome";v="109", "Chromium";v="109"',
          'sec-ch-ua-mobile': '?0',
          'sec-ch-ua-platform': '"macOS"',
        },
      }
    )
    response = http.get(
      'http://localhost:8081/0e5d63cb-bad8-4286-8f95-f2d8b1623727/player/current',
      {
        headers: {
          accept: 'application/json',
          'sec-ch-ua': '"Not_A Brand";v="99", "Google Chrome";v="109", "Chromium";v="109"',
          'sec-ch-ua-mobile': '?0',
          'sec-ch-ua-platform': '"macOS"',
        },
      }
    )
    response = http.get('http://localhost:8081/0e5d63cb-bad8-4286-8f95-f2d8b1623727/winner', {
      headers: {
        accept: 'application/json',
        'sec-ch-ua': '"Not_A Brand";v="99", "Google Chrome";v="109", "Chromium";v="109"',
        'sec-ch-ua-mobile': '?0',
        'sec-ch-ua-platform': '"macOS"',
      },
    })
    response = http.get('http://localhost:8081/0e5d63cb-bad8-4286-8f95-f2d8b1623727/grid', {
      headers: {
        accept: 'application/json',
        'sec-ch-ua': '"Not_A Brand";v="99", "Google Chrome";v="109", "Chromium";v="109"',
        'sec-ch-ua-mobile': '?0',
        'sec-ch-ua-platform': '"macOS"',
      },
    })
    sleep(0.6)
    response = http.post(
      'http://localhost:8081/0e5d63cb-bad8-4286-8f95-f2d8b1623727/play?token=%7B%22Game%22%3A%220e5d63cb-bad8-4286-8f95-f2d8b1623727%22%2C%22Player%22%3A%22X%22%7D&i=0&j=2',
      null,
      {
        headers: {
          accept: 'application/json',
          'sec-ch-ua': '"Not_A Brand";v="99", "Google Chrome";v="109", "Chromium";v="109"',
          'sec-ch-ua-mobile': '?0',
          'sec-ch-ua-platform': '"macOS"',
        },
      }
    )
    response = http.get(
      'http://localhost:8081/0e5d63cb-bad8-4286-8f95-f2d8b1623727/player/current',
      {
        headers: {
          accept: 'application/json',
          'sec-ch-ua': '"Not_A Brand";v="99", "Google Chrome";v="109", "Chromium";v="109"',
          'sec-ch-ua-mobile': '?0',
          'sec-ch-ua-platform': '"macOS"',
        },
      }
    )
    response = http.get('http://localhost:8081/0e5d63cb-bad8-4286-8f95-f2d8b1623727/winner', {
      headers: {
        accept: 'application/json',
        'sec-ch-ua': '"Not_A Brand";v="99", "Google Chrome";v="109", "Chromium";v="109"',
        'sec-ch-ua-mobile': '?0',
        'sec-ch-ua-platform': '"macOS"',
      },
    })
    response = http.get('http://localhost:8081/0e5d63cb-bad8-4286-8f95-f2d8b1623727/grid', {
      headers: {
        accept: 'application/json',
        'sec-ch-ua': '"Not_A Brand";v="99", "Google Chrome";v="109", "Chromium";v="109"',
        'sec-ch-ua-mobile': '?0',
        'sec-ch-ua-platform': '"macOS"',
      },
    })
    sleep(0.6)
    response = http.post(
      'http://localhost:8081/0e5d63cb-bad8-4286-8f95-f2d8b1623727/play?token=%7B%22Game%22%3A%220e5d63cb-bad8-4286-8f95-f2d8b1623727%22%2C%22Player%22%3A%22O%22%7D&i=1&j=0',
      null,
      {
        headers: {
          accept: 'application/json',
          'sec-ch-ua': '"Not_A Brand";v="99", "Google Chrome";v="109", "Chromium";v="109"',
          'sec-ch-ua-mobile': '?0',
          'sec-ch-ua-platform': '"macOS"',
        },
      }
    )
    response = http.get(
      'http://localhost:8081/0e5d63cb-bad8-4286-8f95-f2d8b1623727/player/current',
      {
        headers: {
          accept: 'application/json',
          'sec-ch-ua': '"Not_A Brand";v="99", "Google Chrome";v="109", "Chromium";v="109"',
          'sec-ch-ua-mobile': '?0',
          'sec-ch-ua-platform': '"macOS"',
        },
      }
    )
    response = http.get('http://localhost:8081/0e5d63cb-bad8-4286-8f95-f2d8b1623727/winner', {
      headers: {
        accept: 'application/json',
        'sec-ch-ua': '"Not_A Brand";v="99", "Google Chrome";v="109", "Chromium";v="109"',
        'sec-ch-ua-mobile': '?0',
        'sec-ch-ua-platform': '"macOS"',
      },
    })
    response = http.get('http://localhost:8081/0e5d63cb-bad8-4286-8f95-f2d8b1623727/grid', {
      headers: {
        accept: 'application/json',
        'sec-ch-ua': '"Not_A Brand";v="99", "Google Chrome";v="109", "Chromium";v="109"',
        'sec-ch-ua-mobile': '?0',
        'sec-ch-ua-platform': '"macOS"',
      },
    })
    sleep(0.5)
    response = http.post(
      'http://localhost:8081/0e5d63cb-bad8-4286-8f95-f2d8b1623727/play?token=%7B%22Game%22%3A%220e5d63cb-bad8-4286-8f95-f2d8b1623727%22%2C%22Player%22%3A%22X%22%7D&i=1&j=1',
      null,
      {
        headers: {
          accept: 'application/json',
          'sec-ch-ua': '"Not_A Brand";v="99", "Google Chrome";v="109", "Chromium";v="109"',
          'sec-ch-ua-mobile': '?0',
          'sec-ch-ua-platform': '"macOS"',
        },
      }
    )
    response = http.get(
      'http://localhost:8081/0e5d63cb-bad8-4286-8f95-f2d8b1623727/player/current',
      {
        headers: {
          accept: 'application/json',
          'sec-ch-ua': '"Not_A Brand";v="99", "Google Chrome";v="109", "Chromium";v="109"',
          'sec-ch-ua-mobile': '?0',
          'sec-ch-ua-platform': '"macOS"',
        },
      }
    )
    response = http.get('http://localhost:8081/0e5d63cb-bad8-4286-8f95-f2d8b1623727/winner', {
      headers: {
        accept: 'application/json',
        'sec-ch-ua': '"Not_A Brand";v="99", "Google Chrome";v="109", "Chromium";v="109"',
        'sec-ch-ua-mobile': '?0',
        'sec-ch-ua-platform': '"macOS"',
      },
    })
    response = http.get('http://localhost:8081/0e5d63cb-bad8-4286-8f95-f2d8b1623727/grid', {
      headers: {
        accept: 'application/json',
        'sec-ch-ua': '"Not_A Brand";v="99", "Google Chrome";v="109", "Chromium";v="109"',
        'sec-ch-ua-mobile': '?0',
        'sec-ch-ua-platform': '"macOS"',
      },
    })
    sleep(0.5)
    response = http.post(
      'http://localhost:8081/0e5d63cb-bad8-4286-8f95-f2d8b1623727/play?token=%7B%22Game%22%3A%220e5d63cb-bad8-4286-8f95-f2d8b1623727%22%2C%22Player%22%3A%22O%22%7D&i=1&j=2',
      null,
      {
        headers: {
          accept: 'application/json',
          'sec-ch-ua': '"Not_A Brand";v="99", "Google Chrome";v="109", "Chromium";v="109"',
          'sec-ch-ua-mobile': '?0',
          'sec-ch-ua-platform': '"macOS"',
        },
      }
    )
    response = http.get(
      'http://localhost:8081/0e5d63cb-bad8-4286-8f95-f2d8b1623727/player/current',
      {
        headers: {
          accept: 'application/json',
          'sec-ch-ua': '"Not_A Brand";v="99", "Google Chrome";v="109", "Chromium";v="109"',
          'sec-ch-ua-mobile': '?0',
          'sec-ch-ua-platform': '"macOS"',
        },
      }
    )
    response = http.get('http://localhost:8081/0e5d63cb-bad8-4286-8f95-f2d8b1623727/winner', {
      headers: {
        accept: 'application/json',
        'sec-ch-ua': '"Not_A Brand";v="99", "Google Chrome";v="109", "Chromium";v="109"',
        'sec-ch-ua-mobile': '?0',
        'sec-ch-ua-platform': '"macOS"',
      },
    })
    response = http.get('http://localhost:8081/0e5d63cb-bad8-4286-8f95-f2d8b1623727/grid', {
      headers: {
        accept: 'application/json',
        'sec-ch-ua': '"Not_A Brand";v="99", "Google Chrome";v="109", "Chromium";v="109"',
        'sec-ch-ua-mobile': '?0',
        'sec-ch-ua-platform': '"macOS"',
      },
    })
    sleep(0.6)
    response = http.post(
      'http://localhost:8081/0e5d63cb-bad8-4286-8f95-f2d8b1623727/play?token=%7B%22Game%22%3A%220e5d63cb-bad8-4286-8f95-f2d8b1623727%22%2C%22Player%22%3A%22X%22%7D&i=2&j=0',
      null,
      {
        headers: {
          accept: 'application/json',
          'sec-ch-ua': '"Not_A Brand";v="99", "Google Chrome";v="109", "Chromium";v="109"',
          'sec-ch-ua-mobile': '?0',
          'sec-ch-ua-platform': '"macOS"',
        },
      }
    )
    response = http.get(
      'http://localhost:8081/0e5d63cb-bad8-4286-8f95-f2d8b1623727/player/current',
      {
        headers: {
          accept: 'application/json',
          'sec-ch-ua': '"Not_A Brand";v="99", "Google Chrome";v="109", "Chromium";v="109"',
          'sec-ch-ua-mobile': '?0',
          'sec-ch-ua-platform': '"macOS"',
        },
      }
    )
    response = http.get('http://localhost:8081/0e5d63cb-bad8-4286-8f95-f2d8b1623727/winner', {
      headers: {
        accept: 'application/json',
        'sec-ch-ua': '"Not_A Brand";v="99", "Google Chrome";v="109", "Chromium";v="109"',
        'sec-ch-ua-mobile': '?0',
        'sec-ch-ua-platform': '"macOS"',
      },
    })
    response = http.get('http://localhost:8081/0e5d63cb-bad8-4286-8f95-f2d8b1623727/grid', {
      headers: {
        accept: 'application/json',
        'sec-ch-ua': '"Not_A Brand";v="99", "Google Chrome";v="109", "Chromium";v="109"',
        'sec-ch-ua-mobile': '?0',
        'sec-ch-ua-platform': '"macOS"',
      },
    })
  })
}
