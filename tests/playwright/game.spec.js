const test = require('@playwright/test').test
const playGame = require ('./game').playGame

test('test', async ({ page }) => {
  await playGame(page);
});