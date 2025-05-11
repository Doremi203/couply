window.onload = function() {
  window.ui = SwaggerUIBundle({
    // Можно указать несколько JSON-файлов и дать им названия
    urls: [
      { url: "api/blocker-service/v1/blocker-service.swagger.json", name: "Blocker Service API" },
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