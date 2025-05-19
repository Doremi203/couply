window.onload = function() {
  window.ui = SwaggerUIBundle({
    // Можно указать несколько JSON-файлов и дать им названия
    urls: [
      { url: "api/push/push.swagger.json", name: "Push API" },
      { url: "api/push/admin.swagger.json", name: "Push Admin API" },
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