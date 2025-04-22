window.onload = function() {
  window.ui = SwaggerUIBundle({
    // Можно указать несколько JSON-файлов и дать им названия
    urls: [
      { url: "api/matching-service/v1/matching_service.swagger.json", name: "Matching Service API" },
      { url: "api/search-service/v1/search_service.swagger.json", name: "Search Service API" },
      { url: "api/user-service/v1/user_service.swagger.json", name: "User Service API" },
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