// module.exports = {
//     gridUrl: "local",
//     baseUrl: "http://localhost",
//     pageLoadTimeout: 0,
//     httpTimeout: 60000,
//     testTimeout: 90000,
//     resetCursor: false,
//     sets: {
//         desktop: {
//             files: [
//                 "testplane-tests/**/*.testplane.(t|j)s"
//             ],
//             browsers: [
//                 "chrome"
//             ]
//         }
//     },
//     browsers: {
//         chrome: {
//             headless: false, // Changed to false to show the browser UI
//             windowSize: {
//                 width: 1920,
//                 height: 1080
//             },
//             desiredCapabilities: {
//                 browserName: "chrome"
//             }
//         }
//     },
//     plugins: {
//         "html-reporter/testplane": {
//             // https://github.com/gemini-testing/html-reporter
//             enabled: true,
//             path: "testplane-report",
//             defaultView: "all",
//             diffMode: "3-up-scaled",
//             saveFormat: 'sqlite', // Use sqlite format as required
//             pluginsEnabled: true, // Enable plugins
//             saveErrorDetails: true, // Save error details
//             metaInfoBaseUrls: {
//                 file: 'file://'
//             }
//         },
//         "storybook/testplane": {
//             enabled: true,
//             storybookUrl: "http://localhost:6006",
//             ignoreElements: [],
//             screenshotDelay: 1000,
//             allowViewportOverflow: true
//         }
//     },
//     system: {
//         mochaOpts: {
//             timeout: 60000 // Increase timeout for tests
//         },
//         debug: true // Enable debug mode for more detailed logs
//     },
//     screenshotMode: "viewport",
//     windowSize: {
//         width: 1920,
//         height: 1080
//     }
// };