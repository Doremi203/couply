window.onload = function() {
  window.ui = SwaggerUIBundle({
    // Можно указать несколько JSON-файлов и дать им названия
    urls: [
      { url: "api/payment-service/v1/payment-service.swagger.json", name: "Payment Service API" },
      { url: "api/subscription-service/v1/subscription-service.swagger.json", name: "Subscription Service API" },
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