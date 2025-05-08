# Mobile Screenshot Tests for Couply Frontend

This project uses Playwright and Storybook to create mobile screenshot tests for all components. Screenshot tests help ensure that the visual appearance of components doesn't change unexpectedly on mobile devices.

## Setup

The screenshot tests are set up to capture screenshots of components rendered in Storybook. Each component has a Storybook story file (`.stories.tsx`) that defines how the component should be rendered.

## Running the Tests

To run the mobile screenshot tests, follow these steps:

1. Start the Storybook server:
   ```
   npm run storybook
   ```

2. In a separate terminal, run the mobile screenshot tests:
   ```
   npm run test:screenshots:mobile
   ```

The first time you run the tests, they will fail because there are no reference screenshots to compare against. To create the reference screenshots, run:

```
npm run test:screenshots:mobile:update
```

This will create reference screenshots for all components in mobile view (iPhone 8 dimensions: 375x667). After that, you can run `npm run test:screenshots:mobile` to compare against these references.

## Screenshot Directories

The mobile screenshots are organized in the following directories:

- `screenshots-mobile/reference/` - Reference screenshots (the expected appearance on mobile)
- `screenshots-mobile/actual/` - Actual screenshots taken during the test run
- `screenshots-mobile/diff/` - Difference images showing what changed (if any)

## Components Covered by Screenshot Tests

The following components are covered by screenshot tests:

1. **Shared Components**:
   - NavBar
   - CustomButton (Default and Disabled states)
   - CustomInput (Text, Password, and Email types)
   - ToggleButtons (Default, Three Options, and No Selection states)

2. **Feature Components**:
   - ProfileSlider

3. **Page Components**:
   - HomePage
   - EnterInfoPage
   - AuthPage

## Adding New Tests

To add a screenshot test for a new component:

1. Create a Storybook story for the component in the same directory as the component
2. Add the component to the `components` array in `screenshot-tests.js`
3. Run `npm run test:screenshots:update` to create the reference screenshots

## Implementation Details

The mobile screenshot tests are implemented using Playwright, which provides a way to take screenshots of web pages and compare them. The tests are run using a custom script (`screenshot-tests.js`) that:

1. Launches a headless Chrome browser with mobile device emulation (iPhone 8, 375x667)
2. Navigates to each component in Storybook
3. Takes screenshots of the components as they appear on mobile
4. Compares the screenshots with reference images

### Mobile Device Emulation

The tests use Playwright's device emulation capabilities to simulate an iPhone 8 with the following characteristics:
- Viewport: 375x667 pixels
- Device scale factor: 2
- Mobile user agent
- Touch enabled

## Troubleshooting

If you encounter issues with the mobile screenshot tests:

- Make sure Storybook is running on port 6006
- Check that the component is properly rendered in Storybook in mobile view
- Verify that the URL in the screenshot test script matches the Storybook URL for the component
- Try updating the reference screenshots with `npm run test:screenshots:mobile:update`
- If components don't look right on mobile, check their responsive design implementation