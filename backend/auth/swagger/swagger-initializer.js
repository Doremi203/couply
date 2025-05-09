window.onload = function() {
  window.ui = SwaggerUIBundle({
    // Можно указать несколько JSON-файлов и дать им названия
    urls: [
      { url: "api/registration/registration.swagger.json", name: "Registration API" },
      { url: "api/login/login.swagger.json", name: "Login API" },
      { url: "api/phone-confirm/confirm.swagger.json", name: "Phone Confirm API" },
    ],
    dom_id: '#swagger-ui',
    deepLinking: true,
    presets: [
      SwaggerUIBundle.presets.apis,
      SwaggerUIStandalonePreset
    ],
    layout: "StandaloneLayout"
  });
};