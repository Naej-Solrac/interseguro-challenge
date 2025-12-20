// Configuración de Swagger/OpenAPI
import swaggerJsdoc from 'swagger-jsdoc';

const options: swaggerJsdoc.Options = {
  definition: {
    openapi: '3.0.0',
    info: {
      title: 'Node Statistics API',
      version: '1.0.0',
      description: 'API REST para calcular estadísticas de matrices Q y R provenientes de factorización QR',
      contact: {
        name: 'Interseguro Challenge',
      },
    },
    servers: [
      {
        url: 'http://localhost:3000',
        description: 'Servidor de desarrollo',
      },
      {
        url: 'http://node-api:3000',
        description: 'Servidor en Docker',
      },
    ],
    components: {
      schemas: {
        Matrix: {
          type: 'array',
          items: {
            type: 'array',
            items: {
              type: 'number',
            },
          },
          example: [[1.5, 2.3], [3.1, 4.7]],
        },
        MatrixStatistics: {
          type: 'object',
          properties: {
            avg: {
              type: 'number',
              description: 'Promedio de todos los elementos',
              example: 2.9,
            },
            min: {
              type: 'number',
              description: 'Valor mínimo',
              example: 1.5,
            },
            max: {
              type: 'number',
              description: 'Valor máximo',
              example: 4.7,
            },
            sum: {
              type: 'number',
              description: 'Suma de todos los elementos',
              example: 11.6,
            },
            isDiagonal: {
              type: 'boolean',
              description: 'Indica si la matriz es diagonal',
              example: false,
            },
          },
        },
        ErrorResponse: {
          type: 'object',
          properties: {
            error: {
              type: 'string',
              example: 'Validación fallida',
            },
            details: {
              type: 'array',
              items: {
                type: 'string',
              },
              example: ['Q es obligatoria', 'R debe ser una matriz (array de arrays)'],
            },
          },
        },
      },
    },
  },
  apis: ['./src/infrastructure/*.ts'], // Rutas donde buscar anotaciones JSDoc
};

export const swaggerSpec = swaggerJsdoc(options);
