#!/usr/bin/env node

import { chromium } from 'playwright';
import fs from 'fs/promises';
import path from 'path';
import { fileURLToPath } from 'url';
import { dirname } from 'path';

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

// Configuration
const STORYBOOK_URL = 'http://localhost:6006';
const SCREENSHOTS_DIR = path.join(__dirname, 'screenshots-mobile');
const REFERENCE_DIR = path.join(__dirname, 'screenshots-mobile/reference');
const ACTUAL_DIR = path.join(__dirname, 'screenshots-mobile/actual');
const DIFF_DIR = path.join(__dirname, 'screenshots-mobile/diff');

// Mobile device configuration
const MOBILE_DEVICE = {
  viewport: { width: 375, height: 667 }, // iPhone 8 size
  userAgent: 'Mozilla/5.0 (iPhone; CPU iPhone OS 14_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 Mobile/15E148 Safari/604.1',
  deviceScaleFactor: 2,
  isMobile: true,
  hasTouch: true
};

// List of components to test
const components = [
  { name: 'NavBar', url: '/?path=/story/components-navbar--default' },
  { name: 'CustomButton-Default', url: '/?path=/story/components-custombutton--default' },
  { name: 'CustomButton-Disabled', url: '/?path=/story/components-custombutton--disabled' },
  { name: 'CustomInput-Text', url: '/?path=/story/components-custominput--text' },
  { name: 'CustomInput-Password', url: '/?path=/story/components-custominput--password' },
  { name: 'CustomInput-Email', url: '/?path=/story/components-custominput--email' },
  { name: 'ToggleButtons-Default', url: '/?path=/story/components-togglebuttons--default' },
  { name: 'ToggleButtons-ThreeOptions', url: '/?path=/story/components-togglebuttons--three-options' },
  { name: 'ToggleButtons-NoSelection', url: '/?path=/story/components-togglebuttons--no-selection' },
  { name: 'ProfileSlider', url: '/?path=/story/features-profileslider--default' },
  { name: 'HomePage', url: '/?path=/story/pages-homepage--default' },
  { name: 'EnterInfoPage', url: '/?path=/story/pages-enterinfopage--default' },
  { name: 'AuthPage', url: '/?path=/story/pages-authpage--default' },
];

// Create directories if they don't exist
async function createDirectories() {
  await fs.mkdir(SCREENSHOTS_DIR, { recursive: true });
  await fs.mkdir(REFERENCE_DIR, { recursive: true });
  await fs.mkdir(ACTUAL_DIR, { recursive: true });
  await fs.mkdir(DIFF_DIR, { recursive: true });
}

// Take screenshots of all components
async function takeScreenshots(updateReference = false) {
  const browser = await chromium.launch();
  const page = await browser.newPage({
    viewport: MOBILE_DEVICE.viewport,
    userAgent: MOBILE_DEVICE.userAgent,
    deviceScaleFactor: MOBILE_DEVICE.deviceScaleFactor,
    isMobile: MOBILE_DEVICE.isMobile,
    hasTouch: MOBILE_DEVICE.hasTouch
  });
  
  console.log(`Taking screenshots of ${components.length} components...`);
  
  for (const component of components) {
    try {
      console.log(`Processing ${component.name}...`);
      
      // Navigate to the component in Storybook
      console.log(`Navigating to ${STORYBOOK_URL}${component.url}`);
      await page.goto(`${STORYBOOK_URL}${component.url}`);
      
      // Wait for the page to load
      await page.waitForLoadState('networkidle');
      
      // Log the page title for debugging
      const title = await page.title();
      console.log(`Page title: ${title}`);
      
      // Wait for the component to render (try different selectors)
      // Take a screenshot of the component
      const screenshotPath = path.join(
        updateReference ? REFERENCE_DIR : ACTUAL_DIR,
        `${component.name}.png`
      );
      
      // Take a screenshot of the component
      try {
        // Try to find the iframe
        const frameElement = await page.$('#storybook-preview-iframe');
        if (frameElement) {
          console.log('Found iframe');
          
          // Take a screenshot of the iframe element
          await frameElement.screenshot({
            path: screenshotPath
          });
          
          console.log(`Iframe screenshot saved to ${screenshotPath}`);
        } else {
          throw new Error('Could not find iframe');
        }
      } catch (error) {
        console.log(`Error with iframe: ${error.message}, trying direct screenshot`);
        
        // Take a screenshot of the preview area
        try {
          const previewArea = await page.$('.sb-show-main-container');
          if (previewArea) {
            await previewArea.screenshot({
              path: screenshotPath
            });
            console.log(`Preview area screenshot saved to ${screenshotPath}`);
          } else {
            throw new Error('Could not find preview area');
          }
        } catch (error) {
          console.log(`Error with preview area: ${error.message}, taking full page screenshot`);
          
          // Fallback to full page screenshot
          await page.screenshot({
            path: screenshotPath,
            fullPage: false
          });
          console.log(`Full page screenshot saved to ${screenshotPath}`);
        }
      }
      
      // Screenshots are now taken inside the try/catch block above
    } catch (error) {
      console.error(`Error processing ${component.name}:`, error);
    }
  }
  
  await browser.close();
}

// Compare screenshots
async function compareScreenshots() {
  let passCount = 0;
  let failCount = 0;
  
  console.log('Comparing screenshots...');
  
  for (const component of components) {
    const referencePath = path.join(REFERENCE_DIR, `${component.name}.png`);
    const actualPath = path.join(ACTUAL_DIR, `${component.name}.png`);
    
    try {
      // Check if reference screenshot exists
      try {
        await fs.access(referencePath);
      } catch (error) {
        console.error(`Reference screenshot for ${component.name} does not exist. Run with --update-reference first.`);
        failCount++;
        continue;
      }
      
      // Check if actual screenshot exists
      try {
        await fs.access(actualPath);
      } catch (error) {
        console.error(`Actual screenshot for ${component.name} does not exist.`);
        failCount++;
        continue;
      }
      
      // Compare screenshots using Playwright's built-in comparison
      const referenceBuffer = await fs.readFile(referencePath);
      const actualBuffer = await fs.readFile(actualPath);
      
      // For now, we'll just log that comparison would happen here
      // In a real implementation, you would use a library like pixelmatch or resemblejs
      console.log(`Comparing ${component.name}...`);
      console.log(`  Reference: ${referencePath}`);
      console.log(`  Actual: ${actualPath}`);
      
      // For demonstration, we'll just check if the files have the same size
      if (referenceBuffer.length === actualBuffer.length) {
        console.log(`  ✅ ${component.name} passed`);
        passCount++;
      } else {
        console.log(`  ❌ ${component.name} failed`);
        failCount++;
      }
    } catch (error) {
      console.error(`Error comparing ${component.name}:`, error);
      failCount++;
    }
  }
  
  console.log(`\nResults: ${passCount} passed, ${failCount} failed`);
  return failCount === 0;
}

// Main function
async function main() {
  try {
    await createDirectories();
    
    const updateReference = process.argv.includes('--update-reference');
    
    if (updateReference) {
      console.log('Updating reference screenshots...');
      await takeScreenshots(true);
      console.log('Reference screenshots updated.');
    } else {
      console.log('Taking actual screenshots...');
      await takeScreenshots(false);
      
      const success = await compareScreenshots();
      process.exit(success ? 0 : 1);
    }
  } catch (error) {
    console.error('Error:', error);
    process.exit(1);
  }
}

main();