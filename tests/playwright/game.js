const expect = require('@playwright/test').expect

module.exports = { playGame };

async function playGame(page) {
  // Go to http://localhost:8080/
  await page.goto('http://localhost:8080/');

  // Click text=New Game
  await page.click('text=New Game');

  // Click button
  await page.click('[id="space:0,0"]');
  await expect(page.locator('[id="space:0,0"]')).toHaveText('X', { timeout: 10000 } );

  // Click [id="space:0,1"]
  await page.click('[id="space:0,1"]');
  await expect(page.locator('[id="space:0,1"]')).toHaveText('O', { timeout: 10000 } );

  // Click [id="space:0,2"]
  await page.click('[id="space:0,2"]');
  await expect(page.locator('[id="space:0,2"]')).toHaveText('X', { timeout: 10000 } );

  // Click [id="space:1,0"]
  await page.click('[id="space:1,0"]');
  await expect(page.locator('[id="space:1,0"]')).toHaveText('O', { timeout: 10000 } );

  // Click [id="space:1,1"]
  await page.click('[id="space:1,1"]');
  await expect(page.locator('[id="space:1,1"]')).toHaveText('X', { timeout: 10000 } );

  // Click [id="space:1,2"]
  await page.click('[id="space:1,2"]');
  await expect(page.locator('[id="space:1,2"]')).toHaveText('O', { timeout: 10000 } );

  // Click [id="space:2,0"]
  await page.click('[id="space:2,0"]');
  await expect(page.locator('[id="space:2,0"]')).toHaveText('X', { timeout: 10000 } );

  // X should have won
  await expect(page.getByText('Winner: X')).toBeVisible();
}